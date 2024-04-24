// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/confluentinc/cli/v3/pkg/flink/types (interfaces: StoreInterface)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"
	types "github.com/confluentinc/cli/v3/pkg/flink/types"
	gomock "go.uber.org/mock/gomock"
)

// MockStoreInterface is a mock of StoreInterface interface.
type MockStoreInterface struct {
	ctrl     *gomock.Controller
	recorder *MockStoreInterfaceMockRecorder
}

// MockStoreInterfaceMockRecorder is the mock recorder for MockStoreInterface.
type MockStoreInterfaceMockRecorder struct {
	mock *MockStoreInterface
}

// NewMockStoreInterface creates a new mock instance.
func NewMockStoreInterface(ctrl *gomock.Controller) *MockStoreInterface {
	mock := &MockStoreInterface{ctrl: ctrl}
	mock.recorder = &MockStoreInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStoreInterface) EXPECT() *MockStoreInterfaceMockRecorder {
	return m.recorder
}

// DeleteStatement mocks base method.
func (m *MockStoreInterface) DeleteStatement(arg0 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteStatement", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// DeleteStatement indicates an expected call of DeleteStatement.
func (mr *MockStoreInterfaceMockRecorder) DeleteStatement(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteStatement", reflect.TypeOf((*MockStoreInterface)(nil).DeleteStatement), arg0)
}

// FetchStatementResults mocks base method.
func (m *MockStoreInterface) FetchStatementResults(arg0 types.ProcessedStatement) (*types.ProcessedStatement, *types.StatementError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchStatementResults", arg0)
	ret0, _ := ret[0].(*types.ProcessedStatement)
	ret1, _ := ret[1].(*types.StatementError)
	return ret0, ret1
}

// FetchStatementResults indicates an expected call of FetchStatementResults.
func (mr *MockStoreInterfaceMockRecorder) FetchStatementResults(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchStatementResults", reflect.TypeOf((*MockStoreInterface)(nil).FetchStatementResults), arg0)
}

// GetAuthToken mocks base method.
func (m *MockStoreInterface) GetAuthToken() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAuthToken")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetAuthToken indicates an expected call of GetAuthToken.
func (mr *MockStoreInterfaceMockRecorder) GetAuthToken() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAuthToken", reflect.TypeOf((*MockStoreInterface)(nil).GetAuthToken))
}

// GetComputePool mocks base method.
func (m *MockStoreInterface) GetComputePool() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetComputePool")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetComputePool indicates an expected call of GetComputePool.
func (mr *MockStoreInterfaceMockRecorder) GetComputePool() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetComputePool", reflect.TypeOf((*MockStoreInterface)(nil).GetComputePool))
}

// GetEnvironmentId mocks base method.
func (m *MockStoreInterface) GetEnvironmentId() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEnvironmentId")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetEnvironmentId indicates an expected call of GetEnvironmentId.
func (mr *MockStoreInterfaceMockRecorder) GetEnvironmentId() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEnvironmentId", reflect.TypeOf((*MockStoreInterface)(nil).GetEnvironmentId))
}

// GetOrganizationId mocks base method.
func (m *MockStoreInterface) GetOrganizationId() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrganizationId")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetOrganizationId indicates an expected call of GetOrganizationId.
func (mr *MockStoreInterfaceMockRecorder) GetOrganizationId() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrganizationId", reflect.TypeOf((*MockStoreInterface)(nil).GetOrganizationId))
}

// ProcessStatement mocks base method.
func (m *MockStoreInterface) ProcessStatement(arg0 string) (*types.ProcessedStatement, *types.StatementError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessStatement", arg0)
	ret0, _ := ret[0].(*types.ProcessedStatement)
	ret1, _ := ret[1].(*types.StatementError)
	return ret0, ret1
}

// ProcessStatement indicates an expected call of ProcessStatement.
func (mr *MockStoreInterfaceMockRecorder) ProcessStatement(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessStatement", reflect.TypeOf((*MockStoreInterface)(nil).ProcessStatement), arg0)
}

// StopStatement mocks base method.
func (m *MockStoreInterface) StopStatement(arg0 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StopStatement", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// StopStatement indicates an expected call of StopStatement.
func (mr *MockStoreInterfaceMockRecorder) StopStatement(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopStatement", reflect.TypeOf((*MockStoreInterface)(nil).StopStatement), arg0)
}

// WaitForTerminalStatementState mocks base method.
func (m *MockStoreInterface) WaitForTerminalStatementState(arg0 context.Context, arg1 types.ProcessedStatement) (*types.ProcessedStatement, *types.StatementError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitForTerminalStatementState", arg0, arg1)
	ret0, _ := ret[0].(*types.ProcessedStatement)
	ret1, _ := ret[1].(*types.StatementError)
	return ret0, ret1
}

// WaitForTerminalStatementState indicates an expected call of WaitForTerminalStatementState.
func (mr *MockStoreInterfaceMockRecorder) WaitForTerminalStatementState(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForTerminalStatementState", reflect.TypeOf((*MockStoreInterface)(nil).WaitForTerminalStatementState), arg0, arg1)
}

// WaitPendingStatement mocks base method.
func (m *MockStoreInterface) WaitPendingStatement(arg0 context.Context, arg1 types.ProcessedStatement) (*types.ProcessedStatement, *types.StatementError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitPendingStatement", arg0, arg1)
	ret0, _ := ret[0].(*types.ProcessedStatement)
	ret1, _ := ret[1].(*types.StatementError)
	return ret0, ret1
}

// WaitPendingStatement indicates an expected call of WaitPendingStatement.
func (mr *MockStoreInterfaceMockRecorder) WaitPendingStatement(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitPendingStatement", reflect.TypeOf((*MockStoreInterface)(nil).WaitPendingStatement), arg0, arg1)
}
