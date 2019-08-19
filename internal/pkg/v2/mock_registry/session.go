// Code generated by MockGen. DO NOT EDIT.
// Source: internal/pkg/v2/registry/session.go

// Package mock_registry is a generated GoMock package.
package mock_registry

import (
	datamodel "github.com/RSE-Cambridge/data-acc/internal/pkg/v2/datamodel"
	store "github.com/RSE-Cambridge/data-acc/internal/pkg/v2/store"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockSessionRegistry is a mock of SessionRegistry interface
type MockSessionRegistry struct {
	ctrl     *gomock.Controller
	recorder *MockSessionRegistryMockRecorder
}

// MockSessionRegistryMockRecorder is the mock recorder for MockSessionRegistry
type MockSessionRegistryMockRecorder struct {
	mock *MockSessionRegistry
}

// NewMockSessionRegistry creates a new mock instance
func NewMockSessionRegistry(ctrl *gomock.Controller) *MockSessionRegistry {
	mock := &MockSessionRegistry{ctrl: ctrl}
	mock.recorder = &MockSessionRegistryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSessionRegistry) EXPECT() *MockSessionRegistryMockRecorder {
	return m.recorder
}

// GetSessionMutex mocks base method
func (m *MockSessionRegistry) GetSessionMutex(sessionName datamodel.SessionName) (store.Mutex, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSessionMutex", sessionName)
	ret0, _ := ret[0].(store.Mutex)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSessionMutex indicates an expected call of GetSessionMutex
func (mr *MockSessionRegistryMockRecorder) GetSessionMutex(sessionName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSessionMutex", reflect.TypeOf((*MockSessionRegistry)(nil).GetSessionMutex), sessionName)
}

// CreateSession mocks base method
func (m *MockSessionRegistry) CreateSession(session datamodel.Session) (datamodel.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSession", session)
	ret0, _ := ret[0].(datamodel.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSession indicates an expected call of CreateSession
func (mr *MockSessionRegistryMockRecorder) CreateSession(session interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSession", reflect.TypeOf((*MockSessionRegistry)(nil).CreateSession), session)
}

// GetSession mocks base method
func (m *MockSessionRegistry) GetSession(sessionName datamodel.SessionName) (datamodel.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSession", sessionName)
	ret0, _ := ret[0].(datamodel.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSession indicates an expected call of GetSession
func (mr *MockSessionRegistryMockRecorder) GetSession(sessionName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSession", reflect.TypeOf((*MockSessionRegistry)(nil).GetSession), sessionName)
}

// GetAllSessions mocks base method
func (m *MockSessionRegistry) GetAllSessions() ([]datamodel.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllSessions")
	ret0, _ := ret[0].([]datamodel.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllSessions indicates an expected call of GetAllSessions
func (mr *MockSessionRegistryMockRecorder) GetAllSessions() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllSessions", reflect.TypeOf((*MockSessionRegistry)(nil).GetAllSessions))
}

// UpdateSession mocks base method
func (m *MockSessionRegistry) UpdateSession(session datamodel.Session) (datamodel.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSession", session)
	ret0, _ := ret[0].(datamodel.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateSession indicates an expected call of UpdateSession
func (mr *MockSessionRegistryMockRecorder) UpdateSession(session interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSession", reflect.TypeOf((*MockSessionRegistry)(nil).UpdateSession), session)
}

// DeleteSession mocks base method
func (m *MockSessionRegistry) DeleteSession(session datamodel.Session) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSession", session)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSession indicates an expected call of DeleteSession
func (mr *MockSessionRegistryMockRecorder) DeleteSession(session interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSession", reflect.TypeOf((*MockSessionRegistry)(nil).DeleteSession), session)
}
