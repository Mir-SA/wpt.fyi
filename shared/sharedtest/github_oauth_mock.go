// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/web-platform-tests/wpt.fyi/shared (interfaces: GitHubOAuth,GitHubAccessControl)

// Package sharedtest is a generated GoMock package.
package sharedtest

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	github "github.com/google/go-github/v31/github"
	shared "github.com/web-platform-tests/wpt.fyi/shared"
	oauth2 "golang.org/x/oauth2"
	reflect "reflect"
)

// MockGitHubOAuth is a mock of GitHubOAuth interface
type MockGitHubOAuth struct {
	ctrl     *gomock.Controller
	recorder *MockGitHubOAuthMockRecorder
}

// MockGitHubOAuthMockRecorder is the mock recorder for MockGitHubOAuth
type MockGitHubOAuthMockRecorder struct {
	mock *MockGitHubOAuth
}

// NewMockGitHubOAuth creates a new mock instance
func NewMockGitHubOAuth(ctrl *gomock.Controller) *MockGitHubOAuth {
	mock := &MockGitHubOAuth{ctrl: ctrl}
	mock.recorder = &MockGitHubOAuthMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGitHubOAuth) EXPECT() *MockGitHubOAuthMockRecorder {
	return m.recorder
}

// Context mocks base method
func (m *MockGitHubOAuth) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockGitHubOAuthMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockGitHubOAuth)(nil).Context))
}

// Datastore mocks base method
func (m *MockGitHubOAuth) Datastore() shared.Datastore {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Datastore")
	ret0, _ := ret[0].(shared.Datastore)
	return ret0
}

// Datastore indicates an expected call of Datastore
func (mr *MockGitHubOAuthMockRecorder) Datastore() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Datastore", reflect.TypeOf((*MockGitHubOAuth)(nil).Datastore))
}

// GetAccessToken mocks base method
func (m *MockGitHubOAuth) GetAccessToken() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccessToken")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetAccessToken indicates an expected call of GetAccessToken
func (mr *MockGitHubOAuthMockRecorder) GetAccessToken() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccessToken", reflect.TypeOf((*MockGitHubOAuth)(nil).GetAccessToken))
}

// GetAuthCodeURL mocks base method
func (m *MockGitHubOAuth) GetAuthCodeURL(arg0 string, arg1 ...oauth2.AuthCodeOption) string {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAuthCodeURL", varargs...)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetAuthCodeURL indicates an expected call of GetAuthCodeURL
func (mr *MockGitHubOAuthMockRecorder) GetAuthCodeURL(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAuthCodeURL", reflect.TypeOf((*MockGitHubOAuth)(nil).GetAuthCodeURL), varargs...)
}

// GetUser mocks base method
func (m *MockGitHubOAuth) GetUser(arg0 *github.Client) (*github.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0)
	ret0, _ := ret[0].(*github.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser
func (mr *MockGitHubOAuthMockRecorder) GetUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockGitHubOAuth)(nil).GetUser), arg0)
}

// NewClient mocks base method
func (m *MockGitHubOAuth) NewClient(arg0 string) (*github.Client, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewClient", arg0)
	ret0, _ := ret[0].(*github.Client)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewClient indicates an expected call of NewClient
func (mr *MockGitHubOAuthMockRecorder) NewClient(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewClient", reflect.TypeOf((*MockGitHubOAuth)(nil).NewClient), arg0)
}

// SetRedirectURL mocks base method
func (m *MockGitHubOAuth) SetRedirectURL(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetRedirectURL", arg0)
}

// SetRedirectURL indicates an expected call of SetRedirectURL
func (mr *MockGitHubOAuthMockRecorder) SetRedirectURL(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetRedirectURL", reflect.TypeOf((*MockGitHubOAuth)(nil).SetRedirectURL), arg0)
}

// MockGitHubAccessControl is a mock of GitHubAccessControl interface
type MockGitHubAccessControl struct {
	ctrl     *gomock.Controller
	recorder *MockGitHubAccessControlMockRecorder
}

// MockGitHubAccessControlMockRecorder is the mock recorder for MockGitHubAccessControl
type MockGitHubAccessControlMockRecorder struct {
	mock *MockGitHubAccessControl
}

// NewMockGitHubAccessControl creates a new mock instance
func NewMockGitHubAccessControl(ctrl *gomock.Controller) *MockGitHubAccessControl {
	mock := &MockGitHubAccessControl{ctrl: ctrl}
	mock.recorder = &MockGitHubAccessControlMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGitHubAccessControl) EXPECT() *MockGitHubAccessControlMockRecorder {
	return m.recorder
}

// IsValidWPTMember mocks base method
func (m *MockGitHubAccessControl) IsValidWPTMember() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsValidWPTMember")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsValidWPTMember indicates an expected call of IsValidWPTMember
func (mr *MockGitHubAccessControlMockRecorder) IsValidWPTMember() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsValidWPTMember", reflect.TypeOf((*MockGitHubAccessControl)(nil).IsValidWPTMember))
}
