// Code generated by MockGen. DO NOT EDIT.
// Source: api/server/server.go

// Package server is a generated GoMock package.
package server

import (
	io "io"
	reflect "reflect"
	sync "sync"
	time "time"

	ids "github.com/lasthyphen/dijetsnodego/ids"
	snow "github.com/lasthyphen/dijetsnodego/snow"
	common "github.com/lasthyphen/dijetsnodego/snow/engine/common"
	logging "github.com/lasthyphen/dijetsnodego/utils/logging"
	gomock "github.com/golang/mock/gomock"
)

// MockPathAdder is a mock of PathAdder interface.
type MockPathAdder struct {
	ctrl     *gomock.Controller
	recorder *MockPathAdderMockRecorder
}

// MockPathAdderMockRecorder is the mock recorder for MockPathAdder.
type MockPathAdderMockRecorder struct {
	mock *MockPathAdder
}

// NewMockPathAdder creates a new mock instance.
func NewMockPathAdder(ctrl *gomock.Controller) *MockPathAdder {
	mock := &MockPathAdder{ctrl: ctrl}
	mock.recorder = &MockPathAdderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPathAdder) EXPECT() *MockPathAdderMockRecorder {
	return m.recorder
}

// AddAliases mocks base method.
func (m *MockPathAdder) AddAliases(endpoint string, aliases ...string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{endpoint}
	for _, a := range aliases {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddAliases", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddAliases indicates an expected call of AddAliases.
func (mr *MockPathAdderMockRecorder) AddAliases(endpoint interface{}, aliases ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{endpoint}, aliases...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAliases", reflect.TypeOf((*MockPathAdder)(nil).AddAliases), varargs...)
}

// AddRoute mocks base method.
func (m *MockPathAdder) AddRoute(handler *common.HTTPHandler, lock *sync.RWMutex, base, endpoint string, loggingWriter io.Writer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRoute", handler, lock, base, endpoint, loggingWriter)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddRoute indicates an expected call of AddRoute.
func (mr *MockPathAdderMockRecorder) AddRoute(handler, lock, base, endpoint, loggingWriter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRoute", reflect.TypeOf((*MockPathAdder)(nil).AddRoute), handler, lock, base, endpoint, loggingWriter)
}

// MockPathAdderWithReadLock is a mock of PathAdderWithReadLock interface.
type MockPathAdderWithReadLock struct {
	ctrl     *gomock.Controller
	recorder *MockPathAdderWithReadLockMockRecorder
}

// MockPathAdderWithReadLockMockRecorder is the mock recorder for MockPathAdderWithReadLock.
type MockPathAdderWithReadLockMockRecorder struct {
	mock *MockPathAdderWithReadLock
}

// NewMockPathAdderWithReadLock creates a new mock instance.
func NewMockPathAdderWithReadLock(ctrl *gomock.Controller) *MockPathAdderWithReadLock {
	mock := &MockPathAdderWithReadLock{ctrl: ctrl}
	mock.recorder = &MockPathAdderWithReadLockMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPathAdderWithReadLock) EXPECT() *MockPathAdderWithReadLockMockRecorder {
	return m.recorder
}

// AddAliasesWithReadLock mocks base method.
func (m *MockPathAdderWithReadLock) AddAliasesWithReadLock(endpoint string, aliases ...string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{endpoint}
	for _, a := range aliases {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddAliasesWithReadLock", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddAliasesWithReadLock indicates an expected call of AddAliasesWithReadLock.
func (mr *MockPathAdderWithReadLockMockRecorder) AddAliasesWithReadLock(endpoint interface{}, aliases ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{endpoint}, aliases...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAliasesWithReadLock", reflect.TypeOf((*MockPathAdderWithReadLock)(nil).AddAliasesWithReadLock), varargs...)
}

// AddRouteWithReadLock mocks base method.
func (m *MockPathAdderWithReadLock) AddRouteWithReadLock(handler *common.HTTPHandler, lock *sync.RWMutex, base, endpoint string, loggingWriter io.Writer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRouteWithReadLock", handler, lock, base, endpoint, loggingWriter)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddRouteWithReadLock indicates an expected call of AddRouteWithReadLock.
func (mr *MockPathAdderWithReadLockMockRecorder) AddRouteWithReadLock(handler, lock, base, endpoint, loggingWriter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRouteWithReadLock", reflect.TypeOf((*MockPathAdderWithReadLock)(nil).AddRouteWithReadLock), handler, lock, base, endpoint, loggingWriter)
}

// MockServer is a mock of Server interface.
type MockServer struct {
	ctrl     *gomock.Controller
	recorder *MockServerMockRecorder
}

// MockServerMockRecorder is the mock recorder for MockServer.
type MockServerMockRecorder struct {
	mock *MockServer
}

// NewMockServer creates a new mock instance.
func NewMockServer(ctrl *gomock.Controller) *MockServer {
	mock := &MockServer{ctrl: ctrl}
	mock.recorder = &MockServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServer) EXPECT() *MockServerMockRecorder {
	return m.recorder
}

// AddAliases mocks base method.
func (m *MockServer) AddAliases(endpoint string, aliases ...string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{endpoint}
	for _, a := range aliases {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddAliases", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddAliases indicates an expected call of AddAliases.
func (mr *MockServerMockRecorder) AddAliases(endpoint interface{}, aliases ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{endpoint}, aliases...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAliases", reflect.TypeOf((*MockServer)(nil).AddAliases), varargs...)
}

// AddAliasesWithReadLock mocks base method.
func (m *MockServer) AddAliasesWithReadLock(endpoint string, aliases ...string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{endpoint}
	for _, a := range aliases {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddAliasesWithReadLock", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddAliasesWithReadLock indicates an expected call of AddAliasesWithReadLock.
func (mr *MockServerMockRecorder) AddAliasesWithReadLock(endpoint interface{}, aliases ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{endpoint}, aliases...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAliasesWithReadLock", reflect.TypeOf((*MockServer)(nil).AddAliasesWithReadLock), varargs...)
}

// AddChainRoute mocks base method.
func (m *MockServer) AddChainRoute(handler *common.HTTPHandler, ctx *snow.ConsensusContext, base, endpoint string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddChainRoute", handler, ctx, base, endpoint)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddChainRoute indicates an expected call of AddChainRoute.
func (mr *MockServerMockRecorder) AddChainRoute(handler, ctx, base, endpoint interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddChainRoute", reflect.TypeOf((*MockServer)(nil).AddChainRoute), handler, ctx, base, endpoint)
}

// AddRoute mocks base method.
func (m *MockServer) AddRoute(handler *common.HTTPHandler, lock *sync.RWMutex, base, endpoint string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRoute", handler, lock, base, endpoint)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddRoute indicates an expected call of AddRoute.
func (mr *MockServerMockRecorder) AddRoute(handler, lock, base, endpoint interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRoute", reflect.TypeOf((*MockServer)(nil).AddRoute), handler, lock, base, endpoint)
}

// AddRouteWithReadLock mocks base method.
func (m *MockServer) AddRouteWithReadLock(handler *common.HTTPHandler, lock *sync.RWMutex, base, endpoint string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRouteWithReadLock", handler, lock, base, endpoint)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddRouteWithReadLock indicates an expected call of AddRouteWithReadLock.
func (mr *MockServerMockRecorder) AddRouteWithReadLock(handler, lock, base, endpoint interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRouteWithReadLock", reflect.TypeOf((*MockServer)(nil).AddRouteWithReadLock), handler, lock, base, endpoint)
}

// Dispatch mocks base method.
func (m *MockServer) Dispatch() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Dispatch")
	ret0, _ := ret[0].(error)
	return ret0
}

// Dispatch indicates an expected call of Dispatch.
func (mr *MockServerMockRecorder) Dispatch() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Dispatch", reflect.TypeOf((*MockServer)(nil).Dispatch))
}

// DispatchTLS mocks base method.
func (m *MockServer) DispatchTLS(certBytes, keyBytes []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DispatchTLS", certBytes, keyBytes)
	ret0, _ := ret[0].(error)
	return ret0
}

// DispatchTLS indicates an expected call of DispatchTLS.
func (mr *MockServerMockRecorder) DispatchTLS(certBytes, keyBytes interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DispatchTLS", reflect.TypeOf((*MockServer)(nil).DispatchTLS), certBytes, keyBytes)
}

// Initialize mocks base method.
func (m *MockServer) Initialize(log logging.Logger, factory logging.Factory, host string, port uint16, allowedOrigins []string, shutdownTimeout time.Duration, nodeID ids.NodeID, wrappers ...Wrapper) {
	m.ctrl.T.Helper()
	varargs := []interface{}{log, factory, host, port, allowedOrigins, shutdownTimeout, nodeID}
	for _, a := range wrappers {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Initialize", varargs...)
}

// Initialize indicates an expected call of Initialize.
func (mr *MockServerMockRecorder) Initialize(log, factory, host, port, allowedOrigins, shutdownTimeout, nodeID interface{}, wrappers ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{log, factory, host, port, allowedOrigins, shutdownTimeout, nodeID}, wrappers...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Initialize", reflect.TypeOf((*MockServer)(nil).Initialize), varargs...)
}

// RegisterChain mocks base method.
func (m *MockServer) RegisterChain(chainName string, engine common.Engine) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterChain", chainName, engine)
}

// RegisterChain indicates an expected call of RegisterChain.
func (mr *MockServerMockRecorder) RegisterChain(chainName, engine interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterChain", reflect.TypeOf((*MockServer)(nil).RegisterChain), chainName, engine)
}

// Shutdown mocks base method.
func (m *MockServer) Shutdown() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Shutdown")
	ret0, _ := ret[0].(error)
	return ret0
}

// Shutdown indicates an expected call of Shutdown.
func (mr *MockServerMockRecorder) Shutdown() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Shutdown", reflect.TypeOf((*MockServer)(nil).Shutdown))
}
