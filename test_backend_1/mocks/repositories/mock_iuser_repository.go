// Code generated by MockGen. DO NOT EDIT.
// Source: test_backend_1/domain/repository (interfaces: IUser)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	model "test_backend_1/domain/model"

	gomock "github.com/golang/mock/gomock"
)

// MockIUser is a mock of IUser interface.
type MockIUser struct {
	ctrl     *gomock.Controller
	recorder *MockIUserMockRecorder
}

// MockIUserMockRecorder is the mock recorder for MockIUser.
type MockIUserMockRecorder struct {
	mock *MockIUser
}

// NewMockIUser creates a new mock instance.
func NewMockIUser(ctrl *gomock.Controller) *MockIUser {
	mock := &MockIUser{ctrl: ctrl}
	mock.recorder = &MockIUserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUser) EXPECT() *MockIUserMockRecorder {
	return m.recorder
}

// GetById mocks base method.
func (m *MockIUser) GetById(arg0 context.Context, arg1 int) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", arg0, arg1)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockIUserMockRecorder) GetById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockIUser)(nil).GetById), arg0, arg1)
}
