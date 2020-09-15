// Code generated by MockGen. DO NOT EDIT.
// Source: internal/model/iKeyStorage.go

// Package mock_model is a generated GoMock package.
package mock_model

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
)

// MockIKeyStorage is a mock of IKeyStorage interface.
type MockIKeyStorage struct {
	ctrl     *gomock.Controller
	recorder *MockIKeyStorageMockRecorder
}

// MockIKeyStorageMockRecorder is the mock recorder for MockIKeyStorage.
type MockIKeyStorageMockRecorder struct {
	mock *MockIKeyStorage
}

// NewMockIKeyStorage creates a new mock instance.
func NewMockIKeyStorage(ctrl *gomock.Controller) *MockIKeyStorage {
	mock := &MockIKeyStorage{ctrl: ctrl}
	mock.recorder = &MockIKeyStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIKeyStorage) EXPECT() *MockIKeyStorageMockRecorder {
	return m.recorder
}

// GetStringValueByKey mocks base method.
func (m *MockIKeyStorage) GetStringValueByKey(key string) (*string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStringValueByKey", key)
	ret0, _ := ret[0].(*string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStringValueByKey indicates an expected call of GetStringValueByKey.
func (mr *MockIKeyStorageMockRecorder) GetStringValueByKey(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStringValueByKey", reflect.TypeOf((*MockIKeyStorage)(nil).GetStringValueByKey), key)
}

// GetIntValueByKey mocks base method.
func (m *MockIKeyStorage) GetIntValueByKey(key string) (*int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIntValueByKey", key)
	ret0, _ := ret[0].(*int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIntValueByKey indicates an expected call of GetIntValueByKey.
func (mr *MockIKeyStorageMockRecorder) GetIntValueByKey(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIntValueByKey", reflect.TypeOf((*MockIKeyStorage)(nil).GetIntValueByKey), key)
}

// SetTempIntKey mocks base method.
func (m *MockIKeyStorage) SetTempIntKey(key string, value int, timeout time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetTempIntKey", key, value, timeout)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetTempIntKey indicates an expected call of SetTempIntKey.
func (mr *MockIKeyStorageMockRecorder) SetTempIntKey(key, value, timeout interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTempIntKey", reflect.TypeOf((*MockIKeyStorage)(nil).SetTempIntKey), key, value, timeout)
}

// SetTempIntKeyOnTimeStamp mocks base method.
func (m *MockIKeyStorage) SetTempIntKeyOnTimeStamp(key string, value int, timestamp time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetTempIntKeyOnTimeStamp", key, value, timestamp)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetTempIntKeyOnTimeStamp indicates an expected call of SetTempIntKeyOnTimeStamp.
func (mr *MockIKeyStorageMockRecorder) SetTempIntKeyOnTimeStamp(key, value, timestamp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTempIntKeyOnTimeStamp", reflect.TypeOf((*MockIKeyStorage)(nil).SetTempIntKeyOnTimeStamp), key, value, timestamp)
}

// GetKeyLifeRest mocks base method.
func (m *MockIKeyStorage) GetKeyLifeRest(key string) (*time.Duration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetKeyLifeRest", key)
	ret0, _ := ret[0].(*time.Duration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetKeyLifeRest indicates an expected call of GetKeyLifeRest.
func (mr *MockIKeyStorageMockRecorder) GetKeyLifeRest(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetKeyLifeRest", reflect.TypeOf((*MockIKeyStorage)(nil).GetKeyLifeRest), key)
}

// DecrKey mocks base method.
func (m *MockIKeyStorage) DecrKey(key string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DecrKey", key)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DecrKey indicates an expected call of DecrKey.
func (mr *MockIKeyStorageMockRecorder) DecrKey(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DecrKey", reflect.TypeOf((*MockIKeyStorage)(nil).DecrKey), key)
}

// DelKey mocks base method.
func (m *MockIKeyStorage) DelKey(key string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DelKey", key)
	ret0, _ := ret[0].(error)
	return ret0
}

// DelKey indicates an expected call of DelKey.
func (mr *MockIKeyStorageMockRecorder) DelKey(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DelKey", reflect.TypeOf((*MockIKeyStorage)(nil).DelKey), key)
}
