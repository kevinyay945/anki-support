// Code generated by MockGen. DO NOT EDIT.
// Source: anki-support/helper (interfaces: Configer)

// Package helper is a generated GoMock package.
package helper

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockConfiger is a mock of Configer interface.
type MockConfiger struct {
	ctrl     *gomock.Controller
	recorder *MockConfigerMockRecorder
}

// MockConfigerMockRecorder is the mock recorder for MockConfiger.
type MockConfigerMockRecorder struct {
	mock *MockConfiger
}

// NewMockConfiger creates a new mock instance.
func NewMockConfiger(ctrl *gomock.Controller) *MockConfiger {
	mock := &MockConfiger{ctrl: ctrl}
	mock.recorder = &MockConfigerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConfiger) EXPECT() *MockConfigerMockRecorder {
	return m.recorder
}

// AssetPath mocks base method.
func (m *MockConfiger) AssetPath() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssetPath")
	ret0, _ := ret[0].(string)
	return ret0
}

// AssetPath indicates an expected call of AssetPath.
func (mr *MockConfigerMockRecorder) AssetPath() *ConfigerAssetPathCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssetPath", reflect.TypeOf((*MockConfiger)(nil).AssetPath))
	return &ConfigerAssetPathCall{Call: call}
}

// ConfigerAssetPathCall wrap *gomock.Call
type ConfigerAssetPathCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *ConfigerAssetPathCall) Return(arg0 string) *ConfigerAssetPathCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *ConfigerAssetPathCall) Do(f func() string) *ConfigerAssetPathCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *ConfigerAssetPathCall) DoAndReturn(f func() string) *ConfigerAssetPathCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GoogleApiToken mocks base method.
func (m *MockConfiger) GoogleApiToken() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GoogleApiToken")
	ret0, _ := ret[0].(string)
	return ret0
}

// GoogleApiToken indicates an expected call of GoogleApiToken.
func (mr *MockConfigerMockRecorder) GoogleApiToken() *ConfigerGoogleApiTokenCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GoogleApiToken", reflect.TypeOf((*MockConfiger)(nil).GoogleApiToken))
	return &ConfigerGoogleApiTokenCall{Call: call}
}

// ConfigerGoogleApiTokenCall wrap *gomock.Call
type ConfigerGoogleApiTokenCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *ConfigerGoogleApiTokenCall) Return(arg0 string) *ConfigerGoogleApiTokenCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *ConfigerGoogleApiTokenCall) Do(f func() string) *ConfigerGoogleApiTokenCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *ConfigerGoogleApiTokenCall) DoAndReturn(f func() string) *ConfigerGoogleApiTokenCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// OpenAIToken mocks base method.
func (m *MockConfiger) OpenAIToken() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenAIToken")
	ret0, _ := ret[0].(string)
	return ret0
}

// OpenAIToken indicates an expected call of OpenAIToken.
func (mr *MockConfigerMockRecorder) OpenAIToken() *ConfigerOpenAITokenCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenAIToken", reflect.TypeOf((*MockConfiger)(nil).OpenAIToken))
	return &ConfigerOpenAITokenCall{Call: call}
}

// ConfigerOpenAITokenCall wrap *gomock.Call
type ConfigerOpenAITokenCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *ConfigerOpenAITokenCall) Return(arg0 string) *ConfigerOpenAITokenCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *ConfigerOpenAITokenCall) Do(f func() string) *ConfigerOpenAITokenCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *ConfigerOpenAITokenCall) DoAndReturn(f func() string) *ConfigerOpenAITokenCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
