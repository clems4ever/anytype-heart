// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/anyproto/anytype-heart/util/builtintemplate (interfaces: BuiltinTemplate)

// Package mockBuiltinTemplate is a generated GoMock package.
package mockBuiltinTemplate

import (
	context "context"
	reflect "reflect"

	app "github.com/anyproto/any-sync/app"
	gomock "go.uber.org/mock/gomock"
)

// MockBuiltinTemplate is a mock of BuiltinTemplate interface.
type MockBuiltinTemplate struct {
	ctrl     *gomock.Controller
	recorder *MockBuiltinTemplateMockRecorder
}

// MockBuiltinTemplateMockRecorder is the mock recorder for MockBuiltinTemplate.
type MockBuiltinTemplateMockRecorder struct {
	mock *MockBuiltinTemplate
}

// NewMockBuiltinTemplate creates a new mock instance.
func NewMockBuiltinTemplate(ctrl *gomock.Controller) *MockBuiltinTemplate {
	mock := &MockBuiltinTemplate{ctrl: ctrl}
	mock.recorder = &MockBuiltinTemplateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBuiltinTemplate) EXPECT() *MockBuiltinTemplateMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockBuiltinTemplate) Close(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockBuiltinTemplateMockRecorder) Close(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockBuiltinTemplate)(nil).Close), arg0)
}

// Hash mocks base method.
func (m *MockBuiltinTemplate) Hash() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Hash")
	ret0, _ := ret[0].(string)
	return ret0
}

// Hash indicates an expected call of Hash.
func (mr *MockBuiltinTemplateMockRecorder) Hash() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Hash", reflect.TypeOf((*MockBuiltinTemplate)(nil).Hash))
}

// Init mocks base method.
func (m *MockBuiltinTemplate) Init(arg0 *app.App) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init.
func (mr *MockBuiltinTemplateMockRecorder) Init(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockBuiltinTemplate)(nil).Init), arg0)
}

// Name mocks base method.
func (m *MockBuiltinTemplate) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockBuiltinTemplateMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockBuiltinTemplate)(nil).Name))
}

// Run mocks base method.
func (m *MockBuiltinTemplate) Run(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run.
func (mr *MockBuiltinTemplateMockRecorder) Run(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockBuiltinTemplate)(nil).Run), arg0)
}
