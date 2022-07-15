// Code generated by MockGen. DO NOT EDIT.
// Source: registry/client.go

// Package mock_registry is a generated GoMock package.
package mock_registry

import (
	big "math/big"
	reflect "reflect"

	registry "github.com/forta-network/forta-core-go/domain/registry"
	registry0 "github.com/forta-network/forta-core-go/registry"
	gomock "github.com/golang/mock/gomock"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// DisableScanner mocks base method.
func (m *MockClient) DisableScanner(ScannerPermission registry0.ScannerPermission, scannerAddress string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DisableScanner", ScannerPermission, scannerAddress)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DisableScanner indicates an expected call of DisableScanner.
func (mr *MockClientMockRecorder) DisableScanner(ScannerPermission, scannerAddress interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisableScanner", reflect.TypeOf((*MockClient)(nil).DisableScanner), ScannerPermission, scannerAddress)
}

// EnableScanner mocks base method.
func (m *MockClient) EnableScanner(ScannerPermission registry0.ScannerPermission, scannerAddress string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnableScanner", ScannerPermission, scannerAddress)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnableScanner indicates an expected call of EnableScanner.
func (mr *MockClientMockRecorder) EnableScanner(ScannerPermission, scannerAddress interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableScanner", reflect.TypeOf((*MockClient)(nil).EnableScanner), ScannerPermission, scannerAddress)
}

// ForEachAgent mocks base method.
func (m *MockClient) ForEachAgent(handler func(*registry0.Agent) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForEachAgent", handler)
	ret0, _ := ret[0].(error)
	return ret0
}

// ForEachAgent indicates an expected call of ForEachAgent.
func (mr *MockClientMockRecorder) ForEachAgent(handler interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForEachAgent", reflect.TypeOf((*MockClient)(nil).ForEachAgent), handler)
}

// ForEachAgentID mocks base method.
func (m *MockClient) ForEachAgentID(handler func(string) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForEachAgentID", handler)
	ret0, _ := ret[0].(error)
	return ret0
}

// ForEachAgentID indicates an expected call of ForEachAgentID.
func (mr *MockClientMockRecorder) ForEachAgentID(handler interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForEachAgentID", reflect.TypeOf((*MockClient)(nil).ForEachAgentID), handler)
}

// ForEachAssignedAgent mocks base method.
func (m *MockClient) ForEachAssignedAgent(scannerID string, handler func(*registry0.Agent) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForEachAssignedAgent", scannerID, handler)
	ret0, _ := ret[0].(error)
	return ret0
}

// ForEachAssignedAgent indicates an expected call of ForEachAssignedAgent.
func (mr *MockClientMockRecorder) ForEachAssignedAgent(scannerID, handler interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForEachAssignedAgent", reflect.TypeOf((*MockClient)(nil).ForEachAssignedAgent), scannerID, handler)
}

// ForEachAssignedScanner mocks base method.
func (m *MockClient) ForEachAssignedScanner(agentID string, handler func(*registry0.Scanner) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForEachAssignedScanner", agentID, handler)
	ret0, _ := ret[0].(error)
	return ret0
}

// ForEachAssignedScanner indicates an expected call of ForEachAssignedScanner.
func (mr *MockClientMockRecorder) ForEachAssignedScanner(agentID, handler interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForEachAssignedScanner", reflect.TypeOf((*MockClient)(nil).ForEachAssignedScanner), agentID, handler)
}

// ForEachChainAgent mocks base method.
func (m *MockClient) ForEachChainAgent(chainID int64, handler func(*registry0.Agent) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForEachChainAgent", chainID, handler)
	ret0, _ := ret[0].(error)
	return ret0
}

// ForEachChainAgent indicates an expected call of ForEachChainAgent.
func (mr *MockClientMockRecorder) ForEachChainAgent(chainID, handler interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForEachChainAgent", reflect.TypeOf((*MockClient)(nil).ForEachChainAgent), chainID, handler)
}

// ForEachScanner mocks base method.
func (m *MockClient) ForEachScanner(handler func(*registry0.Scanner) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForEachScanner", handler)
	ret0, _ := ret[0].(error)
	return ret0
}

// ForEachScanner indicates an expected call of ForEachScanner.
func (mr *MockClientMockRecorder) ForEachScanner(handler interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForEachScanner", reflect.TypeOf((*MockClient)(nil).ForEachScanner), handler)
}

// GetAgent mocks base method.
func (m *MockClient) GetAgent(agentID string) (*registry0.Agent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAgent", agentID)
	ret0, _ := ret[0].(*registry0.Agent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAgent indicates an expected call of GetAgent.
func (mr *MockClientMockRecorder) GetAgent(agentID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAgent", reflect.TypeOf((*MockClient)(nil).GetAgent), agentID)
}

// GetAssignmentHash mocks base method.
func (m *MockClient) GetAssignmentHash(scannerID string) (*registry0.AssignmentHash, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAssignmentHash", scannerID)
	ret0, _ := ret[0].(*registry0.AssignmentHash)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAssignmentHash indicates an expected call of GetAssignmentHash.
func (mr *MockClientMockRecorder) GetAssignmentHash(scannerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAssignmentHash", reflect.TypeOf((*MockClient)(nil).GetAssignmentHash), scannerID)
}

// GetScanner mocks base method.
func (m *MockClient) GetScanner(scannerID string) (*registry0.Scanner, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetScanner", scannerID)
	ret0, _ := ret[0].(*registry0.Scanner)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetScanner indicates an expected call of GetScanner.
func (mr *MockClientMockRecorder) GetScanner(scannerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetScanner", reflect.TypeOf((*MockClient)(nil).GetScanner), scannerID)
}

// GetScannerNodePrereleaseVersion mocks base method.
func (m *MockClient) GetScannerNodePrereleaseVersion() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetScannerNodePrereleaseVersion")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetScannerNodePrereleaseVersion indicates an expected call of GetScannerNodePrereleaseVersion.
func (mr *MockClientMockRecorder) GetScannerNodePrereleaseVersion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetScannerNodePrereleaseVersion", reflect.TypeOf((*MockClient)(nil).GetScannerNodePrereleaseVersion))
}

// GetScannerNodeVersion mocks base method.
func (m *MockClient) GetScannerNodeVersion() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetScannerNodeVersion")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetScannerNodeVersion indicates an expected call of GetScannerNodeVersion.
func (mr *MockClientMockRecorder) GetScannerNodeVersion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetScannerNodeVersion", reflect.TypeOf((*MockClient)(nil).GetScannerNodeVersion))
}

// GetStakingThreshold mocks base method.
func (m *MockClient) GetStakingThreshold(scannerID string) (*registry0.StakingThreshold, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStakingThreshold", scannerID)
	ret0, _ := ret[0].(*registry0.StakingThreshold)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStakingThreshold indicates an expected call of GetStakingThreshold.
func (mr *MockClientMockRecorder) GetStakingThreshold(scannerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStakingThreshold", reflect.TypeOf((*MockClient)(nil).GetStakingThreshold), scannerID)
}

// IsAssigned mocks base method.
func (m *MockClient) IsAssigned(scannerID, agentID string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsAssigned", scannerID, agentID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsAssigned indicates an expected call of IsAssigned.
func (mr *MockClientMockRecorder) IsAssigned(scannerID, agentID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsAssigned", reflect.TypeOf((*MockClient)(nil).IsAssigned), scannerID, agentID)
}

// IsEnabledScanner mocks base method.
func (m *MockClient) IsEnabledScanner(scannerID string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsEnabledScanner", scannerID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsEnabledScanner indicates an expected call of IsEnabledScanner.
func (mr *MockClientMockRecorder) IsEnabledScanner(scannerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsEnabledScanner", reflect.TypeOf((*MockClient)(nil).IsEnabledScanner), scannerID)
}

// PegBlock mocks base method.
func (m *MockClient) PegBlock(blockNum *big.Int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PegBlock", blockNum)
}

// PegBlock indicates an expected call of PegBlock.
func (mr *MockClientMockRecorder) PegBlock(blockNum interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PegBlock", reflect.TypeOf((*MockClient)(nil).PegBlock), blockNum)
}

// PegLatestBlock mocks base method.
func (m *MockClient) PegLatestBlock() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PegLatestBlock")
	ret0, _ := ret[0].(error)
	return ret0
}

// PegLatestBlock indicates an expected call of PegLatestBlock.
func (mr *MockClientMockRecorder) PegLatestBlock() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PegLatestBlock", reflect.TypeOf((*MockClient)(nil).PegLatestBlock))
}

// RegisterScanner mocks base method.
func (m *MockClient) RegisterScanner(ownerAddress string, chainID int64, metadata string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterScanner", ownerAddress, chainID, metadata)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterScanner indicates an expected call of RegisterScanner.
func (mr *MockClientMockRecorder) RegisterScanner(ownerAddress, chainID, metadata interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterScanner", reflect.TypeOf((*MockClient)(nil).RegisterScanner), ownerAddress, chainID, metadata)
}

// RegistryContracts mocks base method.
func (m *MockClient) RegistryContracts() *registry.RegistryContracts {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegistryContracts")
	ret0, _ := ret[0].(*registry.RegistryContracts)
	return ret0
}

// RegistryContracts indicates an expected call of RegistryContracts.
func (mr *MockClientMockRecorder) RegistryContracts() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegistryContracts", reflect.TypeOf((*MockClient)(nil).RegistryContracts))
}

// ResetOpts mocks base method.
func (m *MockClient) ResetOpts() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ResetOpts")
}

// ResetOpts indicates an expected call of ResetOpts.
func (mr *MockClientMockRecorder) ResetOpts() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetOpts", reflect.TypeOf((*MockClient)(nil).ResetOpts))
}