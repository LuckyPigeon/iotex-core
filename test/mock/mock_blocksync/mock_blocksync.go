// Code generated by MockGen. DO NOT EDIT.
// Source: ./blocksync/blocksync.go

// Package mock_blocksync is a generated GoMock package.
package mock_blocksync

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	block "github.com/iotexproject/iotex-core/blockchain/block"
)

// MockBlockSync is a mock of BlockSync interface.
type MockBlockSync struct {
	ctrl     *gomock.Controller
	recorder *MockBlockSyncMockRecorder
}

// MockBlockSyncMockRecorder is the mock recorder for MockBlockSync.
type MockBlockSyncMockRecorder struct {
	mock *MockBlockSync
}

// NewMockBlockSync creates a new mock instance.
func NewMockBlockSync(ctrl *gomock.Controller) *MockBlockSync {
	mock := &MockBlockSync{ctrl: ctrl}
	mock.recorder = &MockBlockSyncMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBlockSync) EXPECT() *MockBlockSyncMockRecorder {
	return m.recorder
}

// ProcessBlock mocks base method.
func (m *MockBlockSync) ProcessBlock(arg0 context.Context, arg1 string, arg2 *block.Block) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessBlock", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProcessBlock indicates an expected call of ProcessBlock.
func (mr *MockBlockSyncMockRecorder) ProcessBlock(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessBlock", reflect.TypeOf((*MockBlockSync)(nil).ProcessBlock), arg0, arg1, arg2)
}

// ProcessSyncRequest mocks base method.
func (m *MockBlockSync) ProcessSyncRequest(arg0 context.Context, arg1, arg2 uint64, arg3 func(context.Context, *block.Block) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessSyncRequest", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProcessSyncRequest indicates an expected call of ProcessSyncRequest.
func (mr *MockBlockSyncMockRecorder) ProcessSyncRequest(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessSyncRequest", reflect.TypeOf((*MockBlockSync)(nil).ProcessSyncRequest), arg0, arg1, arg2, arg3)
}

// Start mocks base method.
func (m *MockBlockSync) Start(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start.
func (mr *MockBlockSyncMockRecorder) Start(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockBlockSync)(nil).Start), arg0)
}

// Stop mocks base method.
func (m *MockBlockSync) Stop(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop.
func (mr *MockBlockSyncMockRecorder) Stop(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockBlockSync)(nil).Stop), arg0)
}

// SyncStatus mocks base method.
func (m *MockBlockSync) SyncStatus() (uint64, uint64, uint64, string) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncStatus")
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(uint64)
	ret2, _ := ret[2].(uint64)
	ret3, _ := ret[3].(string)
	return ret0, ret1, ret2, ret3
}

// SyncStatus indicates an expected call of SyncStatus.
func (mr *MockBlockSyncMockRecorder) SyncStatus() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncStatus", reflect.TypeOf((*MockBlockSync)(nil).SyncStatus))
}

// TargetHeight mocks base method.
func (m *MockBlockSync) TargetHeight() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TargetHeight")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// TargetHeight indicates an expected call of TargetHeight.
func (mr *MockBlockSyncMockRecorder) TargetHeight() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TargetHeight", reflect.TypeOf((*MockBlockSync)(nil).TargetHeight))
}
