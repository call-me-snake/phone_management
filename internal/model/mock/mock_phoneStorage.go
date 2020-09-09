// Code generated by MockGen. DO NOT EDIT.
// Source: internal/model/iPhoneStorage.go

// Package mock_model is a generated GoMock package.
package mock_model

import (
	reflect "reflect"

	model "github.com/call-me-snake/phone_management/internal/model"
	gomock "github.com/golang/mock/gomock"
)

// MockIPhoneStorage is a mock of IPhoneStorage interface.
type MockIPhoneStorage struct {
	ctrl     *gomock.Controller
	recorder *MockIPhoneStorageMockRecorder
}

// MockIPhoneStorageMockRecorder is the mock recorder for MockIPhoneStorage.
type MockIPhoneStorageMockRecorder struct {
	mock *MockIPhoneStorage
}

// NewMockIPhoneStorage creates a new mock instance.
func NewMockIPhoneStorage(ctrl *gomock.Controller) *MockIPhoneStorage {
	mock := &MockIPhoneStorage{ctrl: ctrl}
	mock.recorder = &MockIPhoneStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIPhoneStorage) EXPECT() *MockIPhoneStorageMockRecorder {
	return m.recorder
}

// GetPhone mocks base method.
func (m *MockIPhoneStorage) GetPhone(owner string) (*model.PhoneOwner, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPhone", owner)
	ret0, _ := ret[0].(*model.PhoneOwner)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPhone indicates an expected call of GetPhone.
func (mr *MockIPhoneStorageMockRecorder) GetPhone(owner interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPhone", reflect.TypeOf((*MockIPhoneStorage)(nil).GetPhone), owner)
}

// CreateOwner mocks base method.
func (m *MockIPhoneStorage) CreateOwner(data model.PhoneOwner) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOwner", data)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateOwner indicates an expected call of CreateOwner.
func (mr *MockIPhoneStorageMockRecorder) CreateOwner(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOwner", reflect.TypeOf((*MockIPhoneStorage)(nil).CreateOwner), data)
}

// UpdatePhone mocks base method.
func (m *MockIPhoneStorage) UpdatePhone(data model.PhoneOwner) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePhone", data)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePhone indicates an expected call of UpdatePhone.
func (mr *MockIPhoneStorageMockRecorder) UpdatePhone(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePhone", reflect.TypeOf((*MockIPhoneStorage)(nil).UpdatePhone), data)
}

// DeleteOwner mocks base method.
func (m *MockIPhoneStorage) DeleteOwner(owner string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOwner", owner)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteOwner indicates an expected call of DeleteOwner.
func (mr *MockIPhoneStorageMockRecorder) DeleteOwner(owner interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOwner", reflect.TypeOf((*MockIPhoneStorage)(nil).DeleteOwner), owner)
}
