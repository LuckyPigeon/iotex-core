// Code generated by MockGen. DO NOT EDIT.
// Source: ./api/coreservice.go

// Package mock_apicoreservice is a generated GoMock package.
package mock_apicoreservice

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	hash "github.com/iotexproject/go-pkgs/hash"
	address "github.com/iotexproject/iotex-address/address"
	action "github.com/iotexproject/iotex-core/action"
	logfilter "github.com/iotexproject/iotex-core/api/logfilter"
	apitypes "github.com/iotexproject/iotex-core/api/types"
	block "github.com/iotexproject/iotex-core/blockchain/block"
	iotexapi "github.com/iotexproject/iotex-proto/golang/iotexapi"
	iotextypes "github.com/iotexproject/iotex-proto/golang/iotextypes"
)

// MockCoreService is a mock of CoreService interface.
type MockCoreService struct {
	ctrl     *gomock.Controller
	recorder *MockCoreServiceMockRecorder
}

// MockCoreServiceMockRecorder is the mock recorder for MockCoreService.
type MockCoreServiceMockRecorder struct {
	mock *MockCoreService
}

// NewMockCoreService creates a new mock instance.
func NewMockCoreService(ctrl *gomock.Controller) *MockCoreService {
	mock := &MockCoreService{ctrl: ctrl}
	mock.recorder = &MockCoreServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCoreService) EXPECT() *MockCoreServiceMockRecorder {
	return m.recorder
}

// Account mocks base method.
func (m *MockCoreService) Account(addr address.Address) (*iotextypes.AccountMeta, *iotextypes.BlockIdentifier, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Account", addr)
	ret0, _ := ret[0].(*iotextypes.AccountMeta)
	ret1, _ := ret[1].(*iotextypes.BlockIdentifier)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Account indicates an expected call of Account.
func (mr *MockCoreServiceMockRecorder) Account(addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Account", reflect.TypeOf((*MockCoreService)(nil).Account), addr)
}

// ActPoolActions mocks base method.
func (m *MockCoreService) ActPoolActions(actHashes []string) ([]*iotextypes.Action, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ActPoolActions", actHashes)
	ret0, _ := ret[0].([]*iotextypes.Action)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ActPoolActions indicates an expected call of ActPoolActions.
func (mr *MockCoreServiceMockRecorder) ActPoolActions(actHashes interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActPoolActions", reflect.TypeOf((*MockCoreService)(nil).ActPoolActions), actHashes)
}

// Action mocks base method.
func (m *MockCoreService) Action(actionHash string, checkPending bool) (*iotexapi.ActionInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Action", actionHash, checkPending)
	ret0, _ := ret[0].(*iotexapi.ActionInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Action indicates an expected call of Action.
func (mr *MockCoreServiceMockRecorder) Action(actionHash, checkPending interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Action", reflect.TypeOf((*MockCoreService)(nil).Action), actionHash, checkPending)
}

// ActionByActionHash mocks base method.
func (m *MockCoreService) ActionByActionHash(h hash.Hash256) (action.SealedEnvelope, hash.Hash256, uint64, uint32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ActionByActionHash", h)
	ret0, _ := ret[0].(action.SealedEnvelope)
	ret1, _ := ret[1].(hash.Hash256)
	ret2, _ := ret[2].(uint64)
	ret3, _ := ret[3].(uint32)
	ret4, _ := ret[4].(error)
	return ret0, ret1, ret2, ret3, ret4
}

// ActionByActionHash indicates an expected call of ActionByActionHash.
func (mr *MockCoreServiceMockRecorder) ActionByActionHash(h interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActionByActionHash", reflect.TypeOf((*MockCoreService)(nil).ActionByActionHash), h)
}

// Actions mocks base method.
func (m *MockCoreService) Actions(start, count uint64) ([]*iotexapi.ActionInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Actions", start, count)
	ret0, _ := ret[0].([]*iotexapi.ActionInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Actions indicates an expected call of Actions.
func (mr *MockCoreServiceMockRecorder) Actions(start, count interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Actions", reflect.TypeOf((*MockCoreService)(nil).Actions), start, count)
}

// ActionsByAddress mocks base method.
func (m *MockCoreService) ActionsByAddress(addr address.Address, start, count uint64) ([]*iotexapi.ActionInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ActionsByAddress", addr, start, count)
	ret0, _ := ret[0].([]*iotexapi.ActionInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ActionsByAddress indicates an expected call of ActionsByAddress.
func (mr *MockCoreServiceMockRecorder) ActionsByAddress(addr, start, count interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActionsByAddress", reflect.TypeOf((*MockCoreService)(nil).ActionsByAddress), addr, start, count)
}

// ActionsByBlock mocks base method.
func (m *MockCoreService) ActionsByBlock(blkHash string, start, count uint64) ([]*iotexapi.ActionInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ActionsByBlock", blkHash, start, count)
	ret0, _ := ret[0].([]*iotexapi.ActionInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ActionsByBlock indicates an expected call of ActionsByBlock.
func (mr *MockCoreServiceMockRecorder) ActionsByBlock(blkHash, start, count interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActionsByBlock", reflect.TypeOf((*MockCoreService)(nil).ActionsByBlock), blkHash, start, count)
}

// ActionsInBlockByHash mocks base method.
func (m *MockCoreService) ActionsInBlockByHash(arg0 string) ([]action.SealedEnvelope, []*action.Receipt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ActionsInBlockByHash", arg0)
	ret0, _ := ret[0].([]action.SealedEnvelope)
	ret1, _ := ret[1].([]*action.Receipt)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ActionsInBlockByHash indicates an expected call of ActionsInBlockByHash.
func (mr *MockCoreServiceMockRecorder) ActionsInBlockByHash(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActionsInBlockByHash", reflect.TypeOf((*MockCoreService)(nil).ActionsInBlockByHash), arg0)
}

// BlockHashByBlockHeight mocks base method.
func (m *MockCoreService) BlockHashByBlockHeight(blkHeight uint64) (hash.Hash256, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockHashByBlockHeight", blkHeight)
	ret0, _ := ret[0].(hash.Hash256)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BlockHashByBlockHeight indicates an expected call of BlockHashByBlockHeight.
func (mr *MockCoreServiceMockRecorder) BlockHashByBlockHeight(blkHeight interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockHashByBlockHeight", reflect.TypeOf((*MockCoreService)(nil).BlockHashByBlockHeight), blkHeight)
}

// BlockMetaByHash mocks base method.
func (m *MockCoreService) BlockMetaByHash(blkHash string) (*iotextypes.BlockMeta, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockMetaByHash", blkHash)
	ret0, _ := ret[0].(*iotextypes.BlockMeta)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BlockMetaByHash indicates an expected call of BlockMetaByHash.
func (mr *MockCoreServiceMockRecorder) BlockMetaByHash(blkHash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockMetaByHash", reflect.TypeOf((*MockCoreService)(nil).BlockMetaByHash), blkHash)
}

// BlockMetas mocks base method.
func (m *MockCoreService) BlockMetas(start, count uint64) ([]*iotextypes.BlockMeta, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockMetas", start, count)
	ret0, _ := ret[0].([]*iotextypes.BlockMeta)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BlockMetas indicates an expected call of BlockMetas.
func (mr *MockCoreServiceMockRecorder) BlockMetas(start, count interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockMetas", reflect.TypeOf((*MockCoreService)(nil).BlockMetas), start, count)
}

// ChainID mocks base method.
func (m *MockCoreService) ChainID() uint32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChainID")
	ret0, _ := ret[0].(uint32)
	return ret0
}

// ChainID indicates an expected call of ChainID.
func (mr *MockCoreServiceMockRecorder) ChainID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChainID", reflect.TypeOf((*MockCoreService)(nil).ChainID))
}

// ChainListener mocks base method.
func (m *MockCoreService) ChainListener() apitypes.Listener {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChainListener")
	ret0, _ := ret[0].(apitypes.Listener)
	return ret0
}

// ChainListener indicates an expected call of ChainListener.
func (mr *MockCoreServiceMockRecorder) ChainListener() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChainListener", reflect.TypeOf((*MockCoreService)(nil).ChainListener))
}

// ChainMeta mocks base method.
func (m *MockCoreService) ChainMeta() (*iotextypes.ChainMeta, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChainMeta")
	ret0, _ := ret[0].(*iotextypes.ChainMeta)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ChainMeta indicates an expected call of ChainMeta.
func (mr *MockCoreServiceMockRecorder) ChainMeta() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChainMeta", reflect.TypeOf((*MockCoreService)(nil).ChainMeta))
}

// EVMNetworkID mocks base method.
func (m *MockCoreService) EVMNetworkID() uint32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EVMNetworkID")
	ret0, _ := ret[0].(uint32)
	return ret0
}

// EVMNetworkID indicates an expected call of EVMNetworkID.
func (mr *MockCoreServiceMockRecorder) EVMNetworkID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EVMNetworkID", reflect.TypeOf((*MockCoreService)(nil).EVMNetworkID))
}

// ElectionBuckets mocks base method.
func (m *MockCoreService) ElectionBuckets(epochNum uint64) ([]*iotextypes.ElectionBucket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ElectionBuckets", epochNum)
	ret0, _ := ret[0].([]*iotextypes.ElectionBucket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ElectionBuckets indicates an expected call of ElectionBuckets.
func (mr *MockCoreServiceMockRecorder) ElectionBuckets(epochNum interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ElectionBuckets", reflect.TypeOf((*MockCoreService)(nil).ElectionBuckets), epochNum)
}

// EpochMeta mocks base method.
func (m *MockCoreService) EpochMeta(epochNum uint64) (*iotextypes.EpochData, uint64, []*iotexapi.BlockProducerInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EpochMeta", epochNum)
	ret0, _ := ret[0].(*iotextypes.EpochData)
	ret1, _ := ret[1].(uint64)
	ret2, _ := ret[2].([]*iotexapi.BlockProducerInfo)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// EpochMeta indicates an expected call of EpochMeta.
func (mr *MockCoreServiceMockRecorder) EpochMeta(epochNum interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EpochMeta", reflect.TypeOf((*MockCoreService)(nil).EpochMeta), epochNum)
}

// EstimateExecutionGasConsumption mocks base method.
func (m *MockCoreService) EstimateExecutionGasConsumption(ctx context.Context, sc *action.Execution, callerAddr address.Address) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EstimateExecutionGasConsumption", ctx, sc, callerAddr)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EstimateExecutionGasConsumption indicates an expected call of EstimateExecutionGasConsumption.
func (mr *MockCoreServiceMockRecorder) EstimateExecutionGasConsumption(ctx, sc, callerAddr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EstimateExecutionGasConsumption", reflect.TypeOf((*MockCoreService)(nil).EstimateExecutionGasConsumption), ctx, sc, callerAddr)
}

// EstimateGasForAction mocks base method.
func (m *MockCoreService) EstimateGasForAction(ctx context.Context, in *iotextypes.Action) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EstimateGasForAction", ctx, in)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EstimateGasForAction indicates an expected call of EstimateGasForAction.
func (mr *MockCoreServiceMockRecorder) EstimateGasForAction(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EstimateGasForAction", reflect.TypeOf((*MockCoreService)(nil).EstimateGasForAction), ctx, in)
}

// EstimateGasForNonExecution mocks base method.
func (m *MockCoreService) EstimateGasForNonExecution(arg0 action.Action) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EstimateGasForNonExecution", arg0)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EstimateGasForNonExecution indicates an expected call of EstimateGasForNonExecution.
func (mr *MockCoreServiceMockRecorder) EstimateGasForNonExecution(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EstimateGasForNonExecution", reflect.TypeOf((*MockCoreService)(nil).EstimateGasForNonExecution), arg0)
}

// LogsInBlockByHash mocks base method.
func (m *MockCoreService) LogsInBlockByHash(filter *logfilter.LogFilter, blockHash hash.Hash256) ([]*action.Log, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LogsInBlockByHash", filter, blockHash)
	ret0, _ := ret[0].([]*action.Log)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LogsInBlockByHash indicates an expected call of LogsInBlockByHash.
func (mr *MockCoreServiceMockRecorder) LogsInBlockByHash(filter, blockHash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogsInBlockByHash", reflect.TypeOf((*MockCoreService)(nil).LogsInBlockByHash), filter, blockHash)
}

// LogsInRange mocks base method.
func (m *MockCoreService) LogsInRange(filter *logfilter.LogFilter, start, end, paginationSize uint64) ([]*action.Log, []hash.Hash256, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LogsInRange", filter, start, end, paginationSize)
	ret0, _ := ret[0].([]*action.Log)
	ret1, _ := ret[1].([]hash.Hash256)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// LogsInRange indicates an expected call of LogsInRange.
func (mr *MockCoreServiceMockRecorder) LogsInRange(filter, start, end, paginationSize interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogsInRange", reflect.TypeOf((*MockCoreService)(nil).LogsInRange), filter, start, end, paginationSize)
}

// PendingNonce mocks base method.
func (m *MockCoreService) PendingNonce(arg0 address.Address) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PendingNonce", arg0)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PendingNonce indicates an expected call of PendingNonce.
func (mr *MockCoreServiceMockRecorder) PendingNonce(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PendingNonce", reflect.TypeOf((*MockCoreService)(nil).PendingNonce), arg0)
}

// RawBlocks mocks base method.
func (m *MockCoreService) RawBlocks(startHeight, count uint64, withReceipts, withTransactionLogs bool) ([]*iotexapi.BlockInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RawBlocks", startHeight, count, withReceipts, withTransactionLogs)
	ret0, _ := ret[0].([]*iotexapi.BlockInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RawBlocks indicates an expected call of RawBlocks.
func (mr *MockCoreServiceMockRecorder) RawBlocks(startHeight, count, withReceipts, withTransactionLogs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RawBlocks", reflect.TypeOf((*MockCoreService)(nil).RawBlocks), startHeight, count, withReceipts, withTransactionLogs)
}

// ReadContract mocks base method.
func (m *MockCoreService) ReadContract(ctx context.Context, callerAddr address.Address, sc *action.Execution) (string, *iotextypes.Receipt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadContract", ctx, callerAddr, sc)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(*iotextypes.Receipt)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ReadContract indicates an expected call of ReadContract.
func (mr *MockCoreServiceMockRecorder) ReadContract(ctx, callerAddr, sc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadContract", reflect.TypeOf((*MockCoreService)(nil).ReadContract), ctx, callerAddr, sc)
}

// ReadContractStorage mocks base method.
func (m *MockCoreService) ReadContractStorage(ctx context.Context, addr address.Address, key []byte) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadContractStorage", ctx, addr, key)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadContractStorage indicates an expected call of ReadContractStorage.
func (mr *MockCoreServiceMockRecorder) ReadContractStorage(ctx, addr, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadContractStorage", reflect.TypeOf((*MockCoreService)(nil).ReadContractStorage), ctx, addr, key)
}

// ReadState mocks base method.
func (m *MockCoreService) ReadState(protocolID, height string, methodName []byte, arguments [][]byte) (*iotexapi.ReadStateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadState", protocolID, height, methodName, arguments)
	ret0, _ := ret[0].(*iotexapi.ReadStateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadState indicates an expected call of ReadState.
func (mr *MockCoreServiceMockRecorder) ReadState(protocolID, height, methodName, arguments interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadState", reflect.TypeOf((*MockCoreService)(nil).ReadState), protocolID, height, methodName, arguments)
}

// ReceiptByActionHash mocks base method.
func (m *MockCoreService) ReceiptByActionHash(h hash.Hash256) (*action.Receipt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReceiptByActionHash", h)
	ret0, _ := ret[0].(*action.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReceiptByActionHash indicates an expected call of ReceiptByActionHash.
func (mr *MockCoreServiceMockRecorder) ReceiptByActionHash(h interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReceiptByActionHash", reflect.TypeOf((*MockCoreService)(nil).ReceiptByActionHash), h)
}

// ReceiveBlock mocks base method.
func (m *MockCoreService) ReceiveBlock(blk *block.Block) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReceiveBlock", blk)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReceiveBlock indicates an expected call of ReceiveBlock.
func (mr *MockCoreServiceMockRecorder) ReceiveBlock(blk interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReceiveBlock", reflect.TypeOf((*MockCoreService)(nil).ReceiveBlock), blk)
}

// SendAction mocks base method.
func (m *MockCoreService) SendAction(ctx context.Context, in *iotextypes.Action) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendAction", ctx, in)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendAction indicates an expected call of SendAction.
func (mr *MockCoreServiceMockRecorder) SendAction(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendAction", reflect.TypeOf((*MockCoreService)(nil).SendAction), ctx, in)
}

// ServerMeta mocks base method.
func (m *MockCoreService) ServerMeta() (string, string, string, string, string) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServerMeta")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(string)
	ret3, _ := ret[3].(string)
	ret4, _ := ret[4].(string)
	return ret0, ret1, ret2, ret3, ret4
}

// ServerMeta indicates an expected call of ServerMeta.
func (mr *MockCoreServiceMockRecorder) ServerMeta() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerMeta", reflect.TypeOf((*MockCoreService)(nil).ServerMeta))
}

// SimulateExecution mocks base method.
func (m *MockCoreService) SimulateExecution(arg0 context.Context, arg1 address.Address, arg2 *action.Execution) ([]byte, *action.Receipt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SimulateExecution", arg0, arg1, arg2)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(*action.Receipt)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// SimulateExecution indicates an expected call of SimulateExecution.
func (mr *MockCoreServiceMockRecorder) SimulateExecution(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SimulateExecution", reflect.TypeOf((*MockCoreService)(nil).SimulateExecution), arg0, arg1, arg2)
}

// Start mocks base method.
func (m *MockCoreService) Start(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start.
func (mr *MockCoreServiceMockRecorder) Start(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockCoreService)(nil).Start), ctx)
}

// Stop mocks base method.
func (m *MockCoreService) Stop(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop.
func (mr *MockCoreServiceMockRecorder) Stop(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockCoreService)(nil).Stop), ctx)
}

// SuggestGasPrice mocks base method.
func (m *MockCoreService) SuggestGasPrice() (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SuggestGasPrice")
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SuggestGasPrice indicates an expected call of SuggestGasPrice.
func (mr *MockCoreServiceMockRecorder) SuggestGasPrice() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SuggestGasPrice", reflect.TypeOf((*MockCoreService)(nil).SuggestGasPrice))
}

// SyncingProgress mocks base method.
func (m *MockCoreService) SyncingProgress() (uint64, uint64, uint64) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncingProgress")
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(uint64)
	ret2, _ := ret[2].(uint64)
	return ret0, ret1, ret2
}

// SyncingProgress indicates an expected call of SyncingProgress.
func (mr *MockCoreServiceMockRecorder) SyncingProgress() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncingProgress", reflect.TypeOf((*MockCoreService)(nil).SyncingProgress))
}

// TipHeight mocks base method.
func (m *MockCoreService) TipHeight() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TipHeight")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// TipHeight indicates an expected call of TipHeight.
func (mr *MockCoreServiceMockRecorder) TipHeight() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TipHeight", reflect.TypeOf((*MockCoreService)(nil).TipHeight))
}

// TransactionLogByActionHash mocks base method.
func (m *MockCoreService) TransactionLogByActionHash(actHash string) (*iotextypes.TransactionLog, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransactionLogByActionHash", actHash)
	ret0, _ := ret[0].(*iotextypes.TransactionLog)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TransactionLogByActionHash indicates an expected call of TransactionLogByActionHash.
func (mr *MockCoreServiceMockRecorder) TransactionLogByActionHash(actHash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransactionLogByActionHash", reflect.TypeOf((*MockCoreService)(nil).TransactionLogByActionHash), actHash)
}

// TransactionLogByBlockHeight mocks base method.
func (m *MockCoreService) TransactionLogByBlockHeight(blockHeight uint64) (*iotextypes.BlockIdentifier, *iotextypes.TransactionLogs, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransactionLogByBlockHeight", blockHeight)
	ret0, _ := ret[0].(*iotextypes.BlockIdentifier)
	ret1, _ := ret[1].(*iotextypes.TransactionLogs)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// TransactionLogByBlockHeight indicates an expected call of TransactionLogByBlockHeight.
func (mr *MockCoreServiceMockRecorder) TransactionLogByBlockHeight(blockHeight interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransactionLogByBlockHeight", reflect.TypeOf((*MockCoreService)(nil).TransactionLogByBlockHeight), blockHeight)
}

// UnconfirmedActionsByAddress mocks base method.
func (m *MockCoreService) UnconfirmedActionsByAddress(address string, start, count uint64) ([]*iotexapi.ActionInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnconfirmedActionsByAddress", address, start, count)
	ret0, _ := ret[0].([]*iotexapi.ActionInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnconfirmedActionsByAddress indicates an expected call of UnconfirmedActionsByAddress.
func (mr *MockCoreServiceMockRecorder) UnconfirmedActionsByAddress(address, start, count interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnconfirmedActionsByAddress", reflect.TypeOf((*MockCoreService)(nil).UnconfirmedActionsByAddress), address, start, count)
}

// MockintrinsicGasCalculator is a mock of intrinsicGasCalculator interface.
type MockintrinsicGasCalculator struct {
	ctrl     *gomock.Controller
	recorder *MockintrinsicGasCalculatorMockRecorder
}

// MockintrinsicGasCalculatorMockRecorder is the mock recorder for MockintrinsicGasCalculator.
type MockintrinsicGasCalculatorMockRecorder struct {
	mock *MockintrinsicGasCalculator
}

// NewMockintrinsicGasCalculator creates a new mock instance.
func NewMockintrinsicGasCalculator(ctrl *gomock.Controller) *MockintrinsicGasCalculator {
	mock := &MockintrinsicGasCalculator{ctrl: ctrl}
	mock.recorder = &MockintrinsicGasCalculatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockintrinsicGasCalculator) EXPECT() *MockintrinsicGasCalculatorMockRecorder {
	return m.recorder
}

// IntrinsicGas mocks base method.
func (m *MockintrinsicGasCalculator) IntrinsicGas() (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IntrinsicGas")
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IntrinsicGas indicates an expected call of IntrinsicGas.
func (mr *MockintrinsicGasCalculatorMockRecorder) IntrinsicGas() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IntrinsicGas", reflect.TypeOf((*MockintrinsicGasCalculator)(nil).IntrinsicGas))
}
