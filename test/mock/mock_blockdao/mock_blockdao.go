// Code generated by MockGen. DO NOT EDIT.
// Source: ./blockchain/blockdao/blockdao.go

// Package mock_blockdao is a generated GoMock package.
package mock_blockdao

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	hash "github.com/iotexproject/go-pkgs/hash"
	action "github.com/iotexproject/iotex-core/action"
	block "github.com/iotexproject/iotex-core/blockchain/block"
	iotextypes "github.com/iotexproject/iotex-proto/golang/iotextypes"
)

// MockBlockDAO is a mock of BlockDAO interface.
type MockBlockDAO struct {
	ctrl     *gomock.Controller
	recorder *MockBlockDAOMockRecorder
}

// MockBlockDAOMockRecorder is the mock recorder for MockBlockDAO.
type MockBlockDAOMockRecorder struct {
	mock *MockBlockDAO
}

// NewMockBlockDAO creates a new mock instance.
func NewMockBlockDAO(ctrl *gomock.Controller) *MockBlockDAO {
	mock := &MockBlockDAO{ctrl: ctrl}
	mock.recorder = &MockBlockDAOMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBlockDAO) EXPECT() *MockBlockDAOMockRecorder {
	return m.recorder
}

// ContainsTransactionLog mocks base method.
func (m *MockBlockDAO) ContainsTransactionLog() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContainsTransactionLog")
	ret0, _ := ret[0].(bool)
	return ret0
}

// ContainsTransactionLog indicates an expected call of ContainsTransactionLog.
func (mr *MockBlockDAOMockRecorder) ContainsTransactionLog() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainsTransactionLog", reflect.TypeOf((*MockBlockDAO)(nil).ContainsTransactionLog))
}

// DeleteBlockToTarget mocks base method.
func (m *MockBlockDAO) DeleteBlockToTarget(arg0 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBlockToTarget", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteBlockToTarget indicates an expected call of DeleteBlockToTarget.
func (mr *MockBlockDAOMockRecorder) DeleteBlockToTarget(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBlockToTarget", reflect.TypeOf((*MockBlockDAO)(nil).DeleteBlockToTarget), arg0)
}

// DeleteTipBlock mocks base method.
func (m *MockBlockDAO) DeleteTipBlock() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTipBlock")
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTipBlock indicates an expected call of DeleteTipBlock.
func (mr *MockBlockDAOMockRecorder) DeleteTipBlock() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTipBlock", reflect.TypeOf((*MockBlockDAO)(nil).DeleteTipBlock))
}

// FooterByHeight mocks base method.
func (m *MockBlockDAO) FooterByHeight(arg0 uint64) (*block.Footer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FooterByHeight", arg0)
	ret0, _ := ret[0].(*block.Footer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FooterByHeight indicates an expected call of FooterByHeight.
func (mr *MockBlockDAOMockRecorder) FooterByHeight(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FooterByHeight", reflect.TypeOf((*MockBlockDAO)(nil).FooterByHeight), arg0)
}

// GetActionByActionHash mocks base method.
func (m *MockBlockDAO) GetActionByActionHash(arg0 hash.Hash256, arg1 uint64) (action.SealedEnvelope, uint32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActionByActionHash", arg0, arg1)
	ret0, _ := ret[0].(action.SealedEnvelope)
	ret1, _ := ret[1].(uint32)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetActionByActionHash indicates an expected call of GetActionByActionHash.
func (mr *MockBlockDAOMockRecorder) GetActionByActionHash(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActionByActionHash", reflect.TypeOf((*MockBlockDAO)(nil).GetActionByActionHash), arg0, arg1)
}

// GetBlock mocks base method.
func (m *MockBlockDAO) GetBlock(arg0 hash.Hash256) (*block.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlock", arg0)
	ret0, _ := ret[0].(*block.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlock indicates an expected call of GetBlock.
func (mr *MockBlockDAOMockRecorder) GetBlock(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlock", reflect.TypeOf((*MockBlockDAO)(nil).GetBlock), arg0)
}

// GetBlockByHeight mocks base method.
func (m *MockBlockDAO) GetBlockByHeight(arg0 uint64) (*block.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlockByHeight", arg0)
	ret0, _ := ret[0].(*block.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlockByHeight indicates an expected call of GetBlockByHeight.
func (mr *MockBlockDAOMockRecorder) GetBlockByHeight(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockByHeight", reflect.TypeOf((*MockBlockDAO)(nil).GetBlockByHeight), arg0)
}

// GetBlockHash mocks base method.
func (m *MockBlockDAO) GetBlockHash(arg0 uint64) (hash.Hash256, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlockHash", arg0)
	ret0, _ := ret[0].(hash.Hash256)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlockHash indicates an expected call of GetBlockHash.
func (mr *MockBlockDAOMockRecorder) GetBlockHash(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockHash", reflect.TypeOf((*MockBlockDAO)(nil).GetBlockHash), arg0)
}

// GetBlockHeight mocks base method.
func (m *MockBlockDAO) GetBlockHeight(arg0 hash.Hash256) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlockHeight", arg0)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlockHeight indicates an expected call of GetBlockHeight.
func (mr *MockBlockDAOMockRecorder) GetBlockHeight(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockHeight", reflect.TypeOf((*MockBlockDAO)(nil).GetBlockHeight), arg0)
}

// GetReceiptByActionHash mocks base method.
func (m *MockBlockDAO) GetReceiptByActionHash(arg0 hash.Hash256, arg1 uint64) (*action.Receipt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReceiptByActionHash", arg0, arg1)
	ret0, _ := ret[0].(*action.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReceiptByActionHash indicates an expected call of GetReceiptByActionHash.
func (mr *MockBlockDAOMockRecorder) GetReceiptByActionHash(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReceiptByActionHash", reflect.TypeOf((*MockBlockDAO)(nil).GetReceiptByActionHash), arg0, arg1)
}

// GetReceipts mocks base method.
func (m *MockBlockDAO) GetReceipts(arg0 uint64) ([]*action.Receipt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReceipts", arg0)
	ret0, _ := ret[0].([]*action.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReceipts indicates an expected call of GetReceipts.
func (mr *MockBlockDAOMockRecorder) GetReceipts(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReceipts", reflect.TypeOf((*MockBlockDAO)(nil).GetReceipts), arg0)
}

// Header mocks base method.
func (m *MockBlockDAO) Header(arg0 hash.Hash256) (*block.Header, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header", arg0)
	ret0, _ := ret[0].(*block.Header)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header.
func (mr *MockBlockDAOMockRecorder) Header(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockBlockDAO)(nil).Header), arg0)
}

// HeaderByHeight mocks base method.
func (m *MockBlockDAO) HeaderByHeight(arg0 uint64) (*block.Header, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HeaderByHeight", arg0)
	ret0, _ := ret[0].(*block.Header)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HeaderByHeight indicates an expected call of HeaderByHeight.
func (mr *MockBlockDAOMockRecorder) HeaderByHeight(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HeaderByHeight", reflect.TypeOf((*MockBlockDAO)(nil).HeaderByHeight), arg0)
}

// Height mocks base method.
func (m *MockBlockDAO) Height() (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Height")
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Height indicates an expected call of Height.
func (mr *MockBlockDAOMockRecorder) Height() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Height", reflect.TypeOf((*MockBlockDAO)(nil).Height))
}

// PutBlock mocks base method.
func (m *MockBlockDAO) PutBlock(arg0 context.Context, arg1 *block.Block) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutBlock", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// PutBlock indicates an expected call of PutBlock.
func (mr *MockBlockDAOMockRecorder) PutBlock(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutBlock", reflect.TypeOf((*MockBlockDAO)(nil).PutBlock), arg0, arg1)
}

// Start mocks base method.
func (m *MockBlockDAO) Start(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start.
func (mr *MockBlockDAOMockRecorder) Start(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockBlockDAO)(nil).Start), ctx)
}

// Stop mocks base method.
func (m *MockBlockDAO) Stop(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop.
func (mr *MockBlockDAOMockRecorder) Stop(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockBlockDAO)(nil).Stop), ctx)
}

// TransactionLogs mocks base method.
func (m *MockBlockDAO) TransactionLogs(arg0 uint64) (*iotextypes.TransactionLogs, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransactionLogs", arg0)
	ret0, _ := ret[0].(*iotextypes.TransactionLogs)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TransactionLogs indicates an expected call of TransactionLogs.
func (mr *MockBlockDAOMockRecorder) TransactionLogs(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransactionLogs", reflect.TypeOf((*MockBlockDAO)(nil).TransactionLogs), arg0)
}
