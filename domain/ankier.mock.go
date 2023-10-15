// Code generated by MockGen. DO NOT EDIT.
// Source: anki-support/domain (interfaces: Ankier)

// Package domain is a generated GoMock package.
package domain

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockAnkier is a mock of Ankier interface.
type MockAnkier struct {
	ctrl     *gomock.Controller
	recorder *MockAnkierMockRecorder
}

// MockAnkierMockRecorder is the mock recorder for MockAnkier.
type MockAnkierMockRecorder struct {
	mock *MockAnkier
}

// NewMockAnkier creates a new mock instance.
func NewMockAnkier(ctrl *gomock.Controller) *MockAnkier {
	mock := &MockAnkier{ctrl: ctrl}
	mock.recorder = &MockAnkierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAnkier) EXPECT() *MockAnkierMockRecorder {
	return m.recorder
}

// AddNoteTagFromNoteId mocks base method.
func (m *MockAnkier) AddNoteTagFromNoteId(arg0 int64, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddNoteTagFromNoteId", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddNoteTagFromNoteId indicates an expected call of AddNoteTagFromNoteId.
func (mr *MockAnkierMockRecorder) AddNoteTagFromNoteId(arg0, arg1 interface{}) *AnkierAddNoteTagFromNoteIdCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddNoteTagFromNoteId", reflect.TypeOf((*MockAnkier)(nil).AddNoteTagFromNoteId), arg0, arg1)
	return &AnkierAddNoteTagFromNoteIdCall{Call: call}
}

// AnkierAddNoteTagFromNoteIdCall wrap *gomock.Call
type AnkierAddNoteTagFromNoteIdCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *AnkierAddNoteTagFromNoteIdCall) Return(arg0 error) *AnkierAddNoteTagFromNoteIdCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *AnkierAddNoteTagFromNoteIdCall) Do(f func(int64, string) error) *AnkierAddNoteTagFromNoteIdCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *AnkierAddNoteTagFromNoteIdCall) DoAndReturn(f func(int64, string) error) *AnkierAddNoteTagFromNoteIdCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// DeleteNoteTagFromNoteId mocks base method.
func (m *MockAnkier) DeleteNoteTagFromNoteId(arg0 int64, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteNoteTagFromNoteId", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteNoteTagFromNoteId indicates an expected call of DeleteNoteTagFromNoteId.
func (mr *MockAnkierMockRecorder) DeleteNoteTagFromNoteId(arg0, arg1 interface{}) *AnkierDeleteNoteTagFromNoteIdCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNoteTagFromNoteId", reflect.TypeOf((*MockAnkier)(nil).DeleteNoteTagFromNoteId), arg0, arg1)
	return &AnkierDeleteNoteTagFromNoteIdCall{Call: call}
}

// AnkierDeleteNoteTagFromNoteIdCall wrap *gomock.Call
type AnkierDeleteNoteTagFromNoteIdCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *AnkierDeleteNoteTagFromNoteIdCall) Return(arg0 error) *AnkierDeleteNoteTagFromNoteIdCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *AnkierDeleteNoteTagFromNoteIdCall) Do(f func(int64, string) error) *AnkierDeleteNoteTagFromNoteIdCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *AnkierDeleteNoteTagFromNoteIdCall) DoAndReturn(f func(int64, string) error) *AnkierDeleteNoteTagFromNoteIdCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetNoteById mocks base method.
func (m *MockAnkier) GetNoteById(arg0 int64) (AnkiNote, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNoteById", arg0)
	ret0, _ := ret[0].(AnkiNote)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNoteById indicates an expected call of GetNoteById.
func (mr *MockAnkierMockRecorder) GetNoteById(arg0 interface{}) *AnkierGetNoteByIdCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNoteById", reflect.TypeOf((*MockAnkier)(nil).GetNoteById), arg0)
	return &AnkierGetNoteByIdCall{Call: call}
}

// AnkierGetNoteByIdCall wrap *gomock.Call
type AnkierGetNoteByIdCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *AnkierGetNoteByIdCall) Return(arg0 AnkiNote, arg1 error) *AnkierGetNoteByIdCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *AnkierGetNoteByIdCall) Do(f func(int64) (AnkiNote, error)) *AnkierGetNoteByIdCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *AnkierGetNoteByIdCall) DoAndReturn(f func(int64) (AnkiNote, error)) *AnkierGetNoteByIdCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetNoteListByDeckName mocks base method.
func (m *MockAnkier) GetNoteListByDeckName(arg0 string) ([]AnkiNote, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNoteListByDeckName", arg0)
	ret0, _ := ret[0].([]AnkiNote)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNoteListByDeckName indicates an expected call of GetNoteListByDeckName.
func (mr *MockAnkierMockRecorder) GetNoteListByDeckName(arg0 interface{}) *AnkierGetNoteListByDeckNameCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNoteListByDeckName", reflect.TypeOf((*MockAnkier)(nil).GetNoteListByDeckName), arg0)
	return &AnkierGetNoteListByDeckNameCall{Call: call}
}

// AnkierGetNoteListByDeckNameCall wrap *gomock.Call
type AnkierGetNoteListByDeckNameCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *AnkierGetNoteListByDeckNameCall) Return(arg0 []AnkiNote, arg1 error) *AnkierGetNoteListByDeckNameCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *AnkierGetNoteListByDeckNameCall) Do(f func(string) ([]AnkiNote, error)) *AnkierGetNoteListByDeckNameCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *AnkierGetNoteListByDeckNameCall) DoAndReturn(f func(string) ([]AnkiNote, error)) *AnkierGetNoteListByDeckNameCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetTodoNoteFromDeck mocks base method.
func (m *MockAnkier) GetTodoNoteFromDeck(arg0 string) ([]AnkiNote, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTodoNoteFromDeck", arg0)
	ret0, _ := ret[0].([]AnkiNote)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTodoNoteFromDeck indicates an expected call of GetTodoNoteFromDeck.
func (mr *MockAnkierMockRecorder) GetTodoNoteFromDeck(arg0 interface{}) *AnkierGetTodoNoteFromDeckCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTodoNoteFromDeck", reflect.TypeOf((*MockAnkier)(nil).GetTodoNoteFromDeck), arg0)
	return &AnkierGetTodoNoteFromDeckCall{Call: call}
}

// AnkierGetTodoNoteFromDeckCall wrap *gomock.Call
type AnkierGetTodoNoteFromDeckCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *AnkierGetTodoNoteFromDeckCall) Return(arg0 []AnkiNote, arg1 error) *AnkierGetTodoNoteFromDeckCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *AnkierGetTodoNoteFromDeckCall) Do(f func(string) ([]AnkiNote, error)) *AnkierGetTodoNoteFromDeckCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *AnkierGetTodoNoteFromDeckCall) DoAndReturn(f func(string) ([]AnkiNote, error)) *AnkierGetTodoNoteFromDeckCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// UpdateNoteById mocks base method.
func (m *MockAnkier) UpdateNoteById(arg0 int64, arg1 AnkiNote, arg2 []AnkiAudio) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateNoteById", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateNoteById indicates an expected call of UpdateNoteById.
func (mr *MockAnkierMockRecorder) UpdateNoteById(arg0, arg1, arg2 interface{}) *AnkierUpdateNoteByIdCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateNoteById", reflect.TypeOf((*MockAnkier)(nil).UpdateNoteById), arg0, arg1, arg2)
	return &AnkierUpdateNoteByIdCall{Call: call}
}

// AnkierUpdateNoteByIdCall wrap *gomock.Call
type AnkierUpdateNoteByIdCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *AnkierUpdateNoteByIdCall) Return(arg0 error) *AnkierUpdateNoteByIdCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *AnkierUpdateNoteByIdCall) Do(f func(int64, AnkiNote, []AnkiAudio) error) *AnkierUpdateNoteByIdCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *AnkierUpdateNoteByIdCall) DoAndReturn(f func(int64, AnkiNote, []AnkiAudio) error) *AnkierUpdateNoteByIdCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
