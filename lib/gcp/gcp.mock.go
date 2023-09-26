// Code generated by MockGen. DO NOT EDIT.
// Source: anki-support/lib/gcp (interfaces: GCPer)

// Package gcp is a generated GoMock package.
package gcp

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockGCPer is a mock of GCPer interface.
type MockGCPer struct {
	ctrl     *gomock.Controller
	recorder *MockGCPerMockRecorder
}

// MockGCPerMockRecorder is the mock recorder for MockGCPer.
type MockGCPerMockRecorder struct {
	mock *MockGCPer
}

// NewMockGCPer creates a new mock instance.
func NewMockGCPer(ctrl *gomock.Controller) *MockGCPer {
	mock := &MockGCPer{ctrl: ctrl}
	mock.recorder = &MockGCPerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGCPer) EXPECT() *MockGCPerMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockGCPer) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockGCPerMockRecorder) Close() *GCPerCloseCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockGCPer)(nil).Close))
	return &GCPerCloseCall{Call: call}
}

// GCPerCloseCall wrap *gomock.Call
type GCPerCloseCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *GCPerCloseCall) Return() *GCPerCloseCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *GCPerCloseCall) Do(f func()) *GCPerCloseCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *GCPerCloseCall) DoAndReturn(f func()) *GCPerCloseCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GenerateAudioByText mocks base method.
func (m *MockGCPer) GenerateAudioByText(arg0, arg1, arg2 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateAudioByText", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateAudioByText indicates an expected call of GenerateAudioByText.
func (mr *MockGCPerMockRecorder) GenerateAudioByText(arg0, arg1, arg2 interface{}) *GCPerGenerateAudioByTextCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateAudioByText", reflect.TypeOf((*MockGCPer)(nil).GenerateAudioByText), arg0, arg1, arg2)
	return &GCPerGenerateAudioByTextCall{Call: call}
}

// GCPerGenerateAudioByTextCall wrap *gomock.Call
type GCPerGenerateAudioByTextCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *GCPerGenerateAudioByTextCall) Return(arg0 string, arg1 error) *GCPerGenerateAudioByTextCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *GCPerGenerateAudioByTextCall) Do(f func(string, string, string) (string, error)) *GCPerGenerateAudioByTextCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *GCPerGenerateAudioByTextCall) DoAndReturn(f func(string, string, string) (string, error)) *GCPerGenerateAudioByTextCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// setClientByToken mocks base method.
func (m *MockGCPer) setClientByToken(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "setClientByToken", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// setClientByToken indicates an expected call of setClientByToken.
func (mr *MockGCPerMockRecorder) setClientByToken(arg0 interface{}) *GCPersetClientByTokenCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "setClientByToken", reflect.TypeOf((*MockGCPer)(nil).setClientByToken), arg0)
	return &GCPersetClientByTokenCall{Call: call}
}

// GCPersetClientByTokenCall wrap *gomock.Call
type GCPersetClientByTokenCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *GCPersetClientByTokenCall) Return(arg0 error) *GCPersetClientByTokenCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *GCPersetClientByTokenCall) Do(f func(string) error) *GCPersetClientByTokenCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *GCPersetClientByTokenCall) DoAndReturn(f func(string) error) *GCPersetClientByTokenCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}