// Code generated by MockGen. DO NOT EDIT.
// Source: vms/rpcchainvm/plugin_directory.go

// Package rpcchainvm is a generated GoMock package.
package rpcchainvm

import (
	reflect "reflect"

	ids "github.com/lasthyphen/dijetsnodego/ids"
	vms "github.com/lasthyphen/dijetsnodego/vms"
	gomock "github.com/golang/mock/gomock"
)

// MockPluginDirectory is a mock of PluginDirectory interface.
type MockPluginDirectory struct {
	ctrl     *gomock.Controller
	recorder *MockPluginDirectoryMockRecorder
}

// MockPluginDirectoryMockRecorder is the mock recorder for MockPluginDirectory.
type MockPluginDirectoryMockRecorder struct {
	mock *MockPluginDirectory
}

// NewMockPluginDirectory creates a new mock instance.
func NewMockPluginDirectory(ctrl *gomock.Controller) *MockPluginDirectory {
	mock := &MockPluginDirectory{ctrl: ctrl}
	mock.recorder = &MockPluginDirectoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPluginDirectory) EXPECT() *MockPluginDirectoryMockRecorder {
	return m.recorder
}

// GetVMs mocks base method.
func (m *MockPluginDirectory) GetVMs(manager vms.Manager) (map[ids.ID]vms.Factory, map[ids.ID]vms.Factory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVMs", manager)
	ret0, _ := ret[0].(map[ids.ID]vms.Factory)
	ret1, _ := ret[1].(map[ids.ID]vms.Factory)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetVMs indicates an expected call of GetVMs.
func (mr *MockPluginDirectoryMockRecorder) GetVMs(manager interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVMs", reflect.TypeOf((*MockPluginDirectory)(nil).GetVMs), manager)
}