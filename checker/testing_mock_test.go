// Code generated by MockGen. DO NOT EDIT.
// Source: checker/testing.go

// Package checker is a generated GoMock package.
package checker

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MocktIntf is a mock of tIntf interface.
type MocktIntf struct {
	ctrl     *gomock.Controller
	recorder *MocktIntfMockRecorder
}

// MocktIntfMockRecorder is the mock recorder for MocktIntf.
type MocktIntfMockRecorder struct {
	mock *MocktIntf
}

// NewMocktIntf creates a new mock instance.
func NewMocktIntf(ctrl *gomock.Controller) *MocktIntf {
	mock := &MocktIntf{ctrl: ctrl}
	mock.recorder = &MocktIntfMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MocktIntf) EXPECT() *MocktIntfMockRecorder {
	return m.recorder
}

// Cleanup mocks base method.
func (m *MocktIntf) Cleanup(arg0 func()) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Cleanup", arg0)
}

// Cleanup indicates an expected call of Cleanup.
func (mr *MocktIntfMockRecorder) Cleanup(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cleanup", reflect.TypeOf((*MocktIntf)(nil).Cleanup), arg0)
}

// Error mocks base method.
func (m *MocktIntf) Error(args ...any) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Error", varargs...)
}

// Error indicates an expected call of Error.
func (mr *MocktIntfMockRecorder) Error(args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MocktIntf)(nil).Error), args...)
}

// Errorf mocks base method.
func (m *MocktIntf) Errorf(format string, args ...any) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Errorf", varargs...)
}

// Errorf indicates an expected call of Errorf.
func (mr *MocktIntfMockRecorder) Errorf(format interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Errorf", reflect.TypeOf((*MocktIntf)(nil).Errorf), varargs...)
}

// Fail mocks base method.
func (m *MocktIntf) Fail() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Fail")
}

// Fail indicates an expected call of Fail.
func (mr *MocktIntfMockRecorder) Fail() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fail", reflect.TypeOf((*MocktIntf)(nil).Fail))
}

// FailNow mocks base method.
func (m *MocktIntf) FailNow() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "FailNow")
}

// FailNow indicates an expected call of FailNow.
func (mr *MocktIntfMockRecorder) FailNow() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FailNow", reflect.TypeOf((*MocktIntf)(nil).FailNow))
}

// Failed mocks base method.
func (m *MocktIntf) Failed() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Failed")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Failed indicates an expected call of Failed.
func (mr *MocktIntfMockRecorder) Failed() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Failed", reflect.TypeOf((*MocktIntf)(nil).Failed))
}

// Fatal mocks base method.
func (m *MocktIntf) Fatal(args ...any) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Fatal", varargs...)
}

// Fatal indicates an expected call of Fatal.
func (mr *MocktIntfMockRecorder) Fatal(args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fatal", reflect.TypeOf((*MocktIntf)(nil).Fatal), args...)
}

// Fatalf mocks base method.
func (m *MocktIntf) Fatalf(format string, args ...any) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Fatalf", varargs...)
}

// Fatalf indicates an expected call of Fatalf.
func (mr *MocktIntfMockRecorder) Fatalf(format interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fatalf", reflect.TypeOf((*MocktIntf)(nil).Fatalf), varargs...)
}

// Helper mocks base method.
func (m *MocktIntf) Helper() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Helper")
}

// Helper indicates an expected call of Helper.
func (mr *MocktIntfMockRecorder) Helper() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Helper", reflect.TypeOf((*MocktIntf)(nil).Helper))
}

// Log mocks base method.
func (m *MocktIntf) Log(args ...any) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Log", varargs...)
}

// Log indicates an expected call of Log.
func (mr *MocktIntfMockRecorder) Log(args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Log", reflect.TypeOf((*MocktIntf)(nil).Log), args...)
}

// Logf mocks base method.
func (m *MocktIntf) Logf(format string, args ...any) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Logf", varargs...)
}

// Logf indicates an expected call of Logf.
func (mr *MocktIntfMockRecorder) Logf(format interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logf", reflect.TypeOf((*MocktIntf)(nil).Logf), varargs...)
}

// Name mocks base method.
func (m *MocktIntf) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MocktIntfMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MocktIntf)(nil).Name))
}

// Setenv mocks base method.
func (m *MocktIntf) Setenv(key, value string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Setenv", key, value)
}

// Setenv indicates an expected call of Setenv.
func (mr *MocktIntfMockRecorder) Setenv(key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Setenv", reflect.TypeOf((*MocktIntf)(nil).Setenv), key, value)
}

// Skip mocks base method.
func (m *MocktIntf) Skip(args ...any) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Skip", varargs...)
}

// Skip indicates an expected call of Skip.
func (mr *MocktIntfMockRecorder) Skip(args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Skip", reflect.TypeOf((*MocktIntf)(nil).Skip), args...)
}

// SkipNow mocks base method.
func (m *MocktIntf) SkipNow() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SkipNow")
}

// SkipNow indicates an expected call of SkipNow.
func (mr *MocktIntfMockRecorder) SkipNow() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SkipNow", reflect.TypeOf((*MocktIntf)(nil).SkipNow))
}

// Skipf mocks base method.
func (m *MocktIntf) Skipf(format string, args ...any) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Skipf", varargs...)
}

// Skipf indicates an expected call of Skipf.
func (mr *MocktIntfMockRecorder) Skipf(format interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Skipf", reflect.TypeOf((*MocktIntf)(nil).Skipf), varargs...)
}

// Skipped mocks base method.
func (m *MocktIntf) Skipped() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Skipped")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Skipped indicates an expected call of Skipped.
func (mr *MocktIntfMockRecorder) Skipped() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Skipped", reflect.TypeOf((*MocktIntf)(nil).Skipped))
}

// TempDir mocks base method.
func (m *MocktIntf) TempDir() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TempDir")
	ret0, _ := ret[0].(string)
	return ret0
}

// TempDir indicates an expected call of TempDir.
func (mr *MocktIntfMockRecorder) TempDir() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TempDir", reflect.TypeOf((*MocktIntf)(nil).TempDir))
}
