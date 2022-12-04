// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/lasthyphen/dijetsnodego/vms/platformvm/txs (interfaces: UnsignedTx)

// Package txs is a generated GoMock package.
package txs

import (
	ids "github.com/lasthyphen/dijetsnodego/ids"
	snow "github.com/lasthyphen/dijetsnodego/snow"
	djtx "github.com/lasthyphen/dijetsnodego/vms/components/djtx"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockUnsignedTx is a mock of UnsignedTx interface
type MockUnsignedTx struct {
	ctrl     *gomock.Controller
	recorder *MockUnsignedTxMockRecorder
}

// MockUnsignedTxMockRecorder is the mock recorder for MockUnsignedTx
type MockUnsignedTxMockRecorder struct {
	mock *MockUnsignedTx
}

// NewMockUnsignedTx creates a new mock instance
func NewMockUnsignedTx(ctrl *gomock.Controller) *MockUnsignedTx {
	mock := &MockUnsignedTx{ctrl: ctrl}
	mock.recorder = &MockUnsignedTxMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUnsignedTx) EXPECT() *MockUnsignedTxMockRecorder {
	return m.recorder
}

// Bytes mocks base method
func (m *MockUnsignedTx) Bytes() []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Bytes")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// Bytes indicates an expected call of Bytes
func (mr *MockUnsignedTxMockRecorder) Bytes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bytes", reflect.TypeOf((*MockUnsignedTx)(nil).Bytes))
}

// InitCtx mocks base method
func (m *MockUnsignedTx) InitCtx(arg0 *snow.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "InitCtx", arg0)
}

// InitCtx indicates an expected call of InitCtx
func (mr *MockUnsignedTxMockRecorder) InitCtx(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitCtx", reflect.TypeOf((*MockUnsignedTx)(nil).InitCtx), arg0)
}

// Initialize mocks base method
func (m *MockUnsignedTx) Initialize(arg0 []byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Initialize", arg0)
}

// Initialize indicates an expected call of Initialize
func (mr *MockUnsignedTxMockRecorder) Initialize(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Initialize", reflect.TypeOf((*MockUnsignedTx)(nil).Initialize), arg0)
}

// InputIDs mocks base method
func (m *MockUnsignedTx) InputIDs() ids.Set {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InputIDs")
	ret0, _ := ret[0].(ids.Set)
	return ret0
}

// InputIDs indicates an expected call of InputIDs
func (mr *MockUnsignedTxMockRecorder) InputIDs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InputIDs", reflect.TypeOf((*MockUnsignedTx)(nil).InputIDs))
}

// Outputs mocks base method
func (m *MockUnsignedTx) Outputs() []*djtx.TransferableOutput {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Outputs")
	ret0, _ := ret[0].([]*djtx.TransferableOutput)
	return ret0
}

// Outputs indicates an expected call of Outputs
func (mr *MockUnsignedTxMockRecorder) Outputs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Outputs", reflect.TypeOf((*MockUnsignedTx)(nil).Outputs))
}

// SyntacticVerify mocks base method
func (m *MockUnsignedTx) SyntacticVerify(arg0 *snow.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyntacticVerify", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SyntacticVerify indicates an expected call of SyntacticVerify
func (mr *MockUnsignedTxMockRecorder) SyntacticVerify(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyntacticVerify", reflect.TypeOf((*MockUnsignedTx)(nil).SyntacticVerify), arg0)
}

// Visit mocks base method
func (m *MockUnsignedTx) Visit(arg0 Visitor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Visit", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Visit indicates an expected call of Visit
func (mr *MockUnsignedTxMockRecorder) Visit(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Visit", reflect.TypeOf((*MockUnsignedTx)(nil).Visit), arg0)
}
