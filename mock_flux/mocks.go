// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/dave/flux (interfaces: DispatcherInterface,NotifierInterface,AppInterface,WatcherInterface)

// Package mock_flux is a generated GoMock package.
package mock_flux

import (
	flux "github.com/dave/flux"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockDispatcherInterface is a mock of DispatcherInterface interface
type MockDispatcherInterface struct {
	ctrl     *gomock.Controller
	recorder *MockDispatcherInterfaceMockRecorder
}

// MockDispatcherInterfaceMockRecorder is the mock recorder for MockDispatcherInterface
type MockDispatcherInterfaceMockRecorder struct {
	mock *MockDispatcherInterface
}

// NewMockDispatcherInterface creates a new mock instance
func NewMockDispatcherInterface(ctrl *gomock.Controller) *MockDispatcherInterface {
	mock := &MockDispatcherInterface{ctrl: ctrl}
	mock.recorder = &MockDispatcherInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDispatcherInterface) EXPECT() *MockDispatcherInterfaceMockRecorder {
	return m.recorder
}

// Dispatch mocks base method
func (m *MockDispatcherInterface) Dispatch(arg0 flux.ActionInterface) chan struct{} {
	ret := m.ctrl.Call(m, "Dispatch", arg0)
	ret0, _ := ret[0].(chan struct{})
	return ret0
}

// Dispatch indicates an expected call of Dispatch
func (mr *MockDispatcherInterfaceMockRecorder) Dispatch(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Dispatch", reflect.TypeOf((*MockDispatcherInterface)(nil).Dispatch), arg0)
}

// MockNotifierInterface is a mock of NotifierInterface interface
type MockNotifierInterface struct {
	ctrl     *gomock.Controller
	recorder *MockNotifierInterfaceMockRecorder
}

// MockNotifierInterfaceMockRecorder is the mock recorder for MockNotifierInterface
type MockNotifierInterfaceMockRecorder struct {
	mock *MockNotifierInterface
}

// NewMockNotifierInterface creates a new mock instance
func NewMockNotifierInterface(ctrl *gomock.Controller) *MockNotifierInterface {
	mock := &MockNotifierInterface{ctrl: ctrl}
	mock.recorder = &MockNotifierInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNotifierInterface) EXPECT() *MockNotifierInterfaceMockRecorder {
	return m.recorder
}

// Notify mocks base method
func (m *MockNotifierInterface) Notify() chan struct{} {
	ret := m.ctrl.Call(m, "Notify")
	ret0, _ := ret[0].(chan struct{})
	return ret0
}

// Notify indicates an expected call of Notify
func (mr *MockNotifierInterfaceMockRecorder) Notify() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Notify", reflect.TypeOf((*MockNotifierInterface)(nil).Notify))
}

// MockAppInterface is a mock of AppInterface interface
type MockAppInterface struct {
	ctrl     *gomock.Controller
	recorder *MockAppInterfaceMockRecorder
}

// MockAppInterfaceMockRecorder is the mock recorder for MockAppInterface
type MockAppInterfaceMockRecorder struct {
	mock *MockAppInterface
}

// NewMockAppInterface creates a new mock instance
func NewMockAppInterface(ctrl *gomock.Controller) *MockAppInterface {
	mock := &MockAppInterface{ctrl: ctrl}
	mock.recorder = &MockAppInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAppInterface) EXPECT() *MockAppInterfaceMockRecorder {
	return m.recorder
}

// Delete mocks base method
func (m *MockAppInterface) Delete(arg0 interface{}) {
	m.ctrl.Call(m, "Delete", arg0)
}

// Delete indicates an expected call of Delete
func (mr *MockAppInterfaceMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockAppInterface)(nil).Delete), arg0)
}

// Dispatch mocks base method
func (m *MockAppInterface) Dispatch(arg0 flux.ActionInterface) chan struct{} {
	ret := m.ctrl.Call(m, "Dispatch", arg0)
	ret0, _ := ret[0].(chan struct{})
	return ret0
}

// Dispatch indicates an expected call of Dispatch
func (mr *MockAppInterfaceMockRecorder) Dispatch(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Dispatch", reflect.TypeOf((*MockAppInterface)(nil).Dispatch), arg0)
}

// Watch mocks base method
func (m *MockAppInterface) Watch(arg0 interface{}, arg1 func(chan struct{})) {
	m.ctrl.Call(m, "Watch", arg0, arg1)
}

// Watch indicates an expected call of Watch
func (mr *MockAppInterfaceMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockAppInterface)(nil).Watch), arg0, arg1)
}

// MockWatcherInterface is a mock of WatcherInterface interface
type MockWatcherInterface struct {
	ctrl     *gomock.Controller
	recorder *MockWatcherInterfaceMockRecorder
}

// MockWatcherInterfaceMockRecorder is the mock recorder for MockWatcherInterface
type MockWatcherInterfaceMockRecorder struct {
	mock *MockWatcherInterface
}

// NewMockWatcherInterface creates a new mock instance
func NewMockWatcherInterface(ctrl *gomock.Controller) *MockWatcherInterface {
	mock := &MockWatcherInterface{ctrl: ctrl}
	mock.recorder = &MockWatcherInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockWatcherInterface) EXPECT() *MockWatcherInterfaceMockRecorder {
	return m.recorder
}

// Delete mocks base method
func (m *MockWatcherInterface) Delete(arg0 interface{}) {
	m.ctrl.Call(m, "Delete", arg0)
}

// Delete indicates an expected call of Delete
func (mr *MockWatcherInterfaceMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockWatcherInterface)(nil).Delete), arg0)
}

// Watch mocks base method
func (m *MockWatcherInterface) Watch(arg0 interface{}, arg1 func(chan struct{})) {
	m.ctrl.Call(m, "Watch", arg0, arg1)
}

// Watch indicates an expected call of Watch
func (mr *MockWatcherInterfaceMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockWatcherInterface)(nil).Watch), arg0, arg1)
}
