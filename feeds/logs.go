package feeds

import (
	"context"
	"fmt"
	"math/big"

	"github.com/forta-network/forta-core-go/domain"

	"github.com/forta-network/forta-core-go/utils"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"

	eth "github.com/forta-network/forta-core-go/ethereum"
)

type logFeed struct {
	ctx        context.Context
	startBlock *big.Int
	endBlock   *big.Int
	topics     [][]string
	addresses  []string
	client     eth.Client
	offset     int
}

func (l *logFeed) GetLogsForRange(startBlock *big.Int, endBlock *big.Int) ([]types.Log, error) {
	addrs := make([]common.Address, 0, len(l.addresses))
	for _, addr := range l.addresses {
		addrs = append(addrs, common.HexToAddress(addr))
	}

	var topics [][]common.Hash
	for _, topicSet := range l.topics {
		var topicPosition []common.Hash
		for _, topic := range topicSet {
			topicHash := common.HexToHash(topic)
			topicPosition = append(topicPosition, topicHash)
		}
		topics = append(topics, topicPosition)
	}

	q := ethereum.FilterQuery{
		FromBlock: startBlock,
		ToBlock:   endBlock,
		Addresses: addrs,
		Topics:    topics,
	}

	return l.client.GetLogs(l.ctx, q)
}

func (l *logFeed) GetLogsForLastBlocks(blocksAgo int64) ([]types.Log, error) {

	blk, err := l.client.BlockByNumber(l.ctx, nil)
	if err != nil {
		return nil, err
	}

	endBlock, err := utils.HexToBigInt(blk.Number)
	if err != nil {
		return nil, err
	}

	startBlock := big.NewInt(endBlock.Int64() - blocksAgo)

	return l.GetLogsForRange(startBlock, endBlock)
}

func (l *logFeed) ForEachLog(handler func(blk *domain.Block, logEntry types.Log) error, finishBlockHandler func(blk *domain.Block) error) error {
	eg, ctx := errgroup.WithContext(l.ctx)

	addrs := make([]common.Address, 0, len(l.addresses))
	for _, addr := range l.addresses {
		addrs = append(addrs, common.HexToAddress(addr))
	}

	var topics [][]common.Hash
	for _, topicSet := range l.topics {
		var topicPosition []common.Hash
		for _, topic := range topicSet {
			topicHash := common.HexToHash(topic)
			topicPosition = append(topicPosition, topicHash)
		}
		topics = append(topics, topicPosition)
	}

	currentBlock := l.startBlock
	increment := big.NewInt(1)
	eg.Go(func() error {
		for {
			if ctx.Err() != nil {
				return ctx.Err()
			}
			if currentBlock != nil && l.endBlock != nil {
				if currentBlock.Cmp(l.endBlock) > 0 {
					log.Infof("completed processing logs (endBlock reached)")
					return nil
				}
			}

			blk, err := l.client.BlockByNumber(l.ctx, currentBlock)
			if err != nil {
				log.Error("error while getting latest block number:", err)
				return err
			}

			// initialize current block if nil
			if currentBlock == nil {
				currentBlock, err = utils.HexToBigInt(blk.Number)
				if err != nil {
					log.Errorf("error while converting latest block number: %s, %s", blk.Number, err)
					return err
				}
			}

			blockToRetrieve := big.NewInt(currentBlock.Int64() - int64(l.offset))

			// if offset is set, get previous block instead
			if l.offset > 0 {
				pastBlock, err := l.client.BlockByNumber(l.ctx, blockToRetrieve)
				if err != nil {
					log.WithError(err).Error("error while getting past block")
					return err
				}
				blk = pastBlock
			}

			q := ethereum.FilterQuery{
				FromBlock: blockToRetrieve,
				ToBlock:   blockToRetrieve,
				Addresses: addrs,
				Topics:    topics,
			}
			logs, err := l.client.GetLogs(l.ctx, q)
			if err != nil {
				return err
			}

			for _, lg := range logs {
				if err := handler(blk, lg); err != nil {
					log.Error("handler returned error, exiting log subscription:", err)
					return err
				}
			}

			currentBlock = currentBlock.Add(currentBlock, increment)
			if err := finishBlockHandler(blk); err != nil {
				return err
			}
		}
	})
	log.Infof("subscribed to logs: address=%v, topics=%v, startBlock=%s, endBlock=%s", l.addresses, l.topics, l.startBlock, l.endBlock)
	defer func() {
		log.Info("log subscription closed")
	}()
	return eg.Wait()
}

type LogFeedConfig struct {
	Topics     [][]string
	Addresses  []string
	StartBlock *big.Int
	EndBlock   *big.Int
	Offset     int
}

func NewLogFeed(ctx context.Context, client eth.Client, cfg LogFeedConfig) (*logFeed, error) {
	if cfg.Offset < 0 {
		return nil, fmt.Errorf("offset cannot be below zero: offset=%d", cfg.Offset)
	}
	return &logFeed{
		ctx:        ctx,
		client:     client,
		topics:     cfg.Topics,
		addresses:  cfg.Addresses,
		startBlock: cfg.StartBlock,
		endBlock:   cfg.EndBlock,
		offset:     cfg.Offset,
	}, nil
}
