// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ozoncp/ocp-presentation-api/internal/common/clockwrapper (interfaces: ClockWrapper)

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
)

// MockClockWrapper is a mock of ClockWrapper interface.
type MockClockWrapper struct {
	ctrl     *gomock.Controller
	recorder *MockClockWrapperMockRecorder
}

// MockClockWrapperMockRecorder is the mock recorder for MockClockWrapper.
type MockClockWrapperMockRecorder struct {
	mock *MockClockWrapper
}

// NewMockClockWrapper creates a new mock instance.
func NewMockClockWrapper(ctrl *gomock.Controller) *MockClockWrapper {
	mock := &MockClockWrapper{ctrl: ctrl}
	mock.recorder = &MockClockWrapperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClockWrapper) EXPECT() *MockClockWrapperMockRecorder {
	return m.recorder
}

// Now mocks base method.
func (m *MockClockWrapper) Now() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Now")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// Now indicates an expected call of Now.
func (mr *MockClockWrapperMockRecorder) Now() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Now", reflect.TypeOf((*MockClockWrapper)(nil).Now))
}
