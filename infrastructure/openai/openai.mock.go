// Code generated by MockGen. DO NOT EDIT.
// Source: anki-support/infrastructure/openai (interfaces: OpenAIer)

// Package openai is a generated GoMock package.
package openai

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockOpenAIer is a mock of OpenAIer interface.
type MockOpenAIer struct {
	ctrl     *gomock.Controller
	recorder *MockOpenAIerMockRecorder
}

// MockOpenAIerMockRecorder is the mock recorder for MockOpenAIer.
type MockOpenAIerMockRecorder struct {
	mock *MockOpenAIer
}

// NewMockOpenAIer creates a new mock instance.
func NewMockOpenAIer(ctrl *gomock.Controller) *MockOpenAIer {
	mock := &MockOpenAIer{ctrl: ctrl}
	mock.recorder = &MockOpenAIerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOpenAIer) EXPECT() *MockOpenAIerMockRecorder {
	return m.recorder
}

// MakeJapaneseSentence mocks base method.
func (m *MockOpenAIer) MakeJapaneseSentence(arg0 []string, arg1, arg2 string) (string, string, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeJapaneseSentence", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(string)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// MakeJapaneseSentence indicates an expected call of MakeJapaneseSentence.
func (mr *MockOpenAIerMockRecorder) MakeJapaneseSentence(arg0, arg1, arg2 interface{}) *OpenAIerMakeJapaneseSentenceCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeJapaneseSentence", reflect.TypeOf((*MockOpenAIer)(nil).MakeJapaneseSentence), arg0, arg1, arg2)
	return &OpenAIerMakeJapaneseSentenceCall{Call: call}
}

// OpenAIerMakeJapaneseSentenceCall wrap *gomock.Call
type OpenAIerMakeJapaneseSentenceCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *OpenAIerMakeJapaneseSentenceCall) Return(arg0, arg1, arg2 string, arg3 error) *OpenAIerMakeJapaneseSentenceCall {
	c.Call = c.Call.Return(arg0, arg1, arg2, arg3)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *OpenAIerMakeJapaneseSentenceCall) Do(f func([]string, string, string) (string, string, string, error)) *OpenAIerMakeJapaneseSentenceCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *OpenAIerMakeJapaneseSentenceCall) DoAndReturn(f func([]string, string, string) (string, string, string, error)) *OpenAIerMakeJapaneseSentenceCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
