// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ozoncp/ocp-presentation-api/internal/ocp-slide-api/repo (interfaces: Repo)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/ozoncp/ocp-presentation-api/internal/ocp-slide-api/model"
)

// MockRepo is a mock of Repo interface.
type MockRepo struct {
	ctrl     *gomock.Controller
	recorder *MockRepoMockRecorder
}

// MockRepoMockRecorder is the mock recorder for MockRepo.
type MockRepoMockRecorder struct {
	mock *MockRepo
}

// NewMockRepo creates a new mock instance.
func NewMockRepo(ctrl *gomock.Controller) *MockRepo {
	mock := &MockRepo{ctrl: ctrl}
	mock.recorder = &MockRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepo) EXPECT() *MockRepoMockRecorder {
	return m.recorder
}

// CreateSlide mocks base method.
func (m *MockRepo) CreateSlide(arg0 context.Context, arg1 model.Slide) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSlide", arg0, arg1)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSlide indicates an expected call of CreateSlide.
func (mr *MockRepoMockRecorder) CreateSlide(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSlide", reflect.TypeOf((*MockRepo)(nil).CreateSlide), arg0, arg1)
}

// DescribeSlide mocks base method.
func (m *MockRepo) DescribeSlide(arg0 context.Context, arg1 uint64) (*model.Slide, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeSlide", arg0, arg1)
	ret0, _ := ret[0].(*model.Slide)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeSlide indicates an expected call of DescribeSlide.
func (mr *MockRepoMockRecorder) DescribeSlide(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeSlide", reflect.TypeOf((*MockRepo)(nil).DescribeSlide), arg0, arg1)
}

// ListSlides mocks base method.
func (m *MockRepo) ListSlides(arg0 context.Context, arg1, arg2 uint64) ([]model.Slide, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSlides", arg0, arg1, arg2)
	ret0, _ := ret[0].([]model.Slide)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSlides indicates an expected call of ListSlides.
func (mr *MockRepoMockRecorder) ListSlides(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSlides", reflect.TypeOf((*MockRepo)(nil).ListSlides), arg0, arg1, arg2)
}

// MultiCreateSlides mocks base method.
func (m *MockRepo) MultiCreateSlides(arg0 context.Context, arg1 []model.Slide) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MultiCreateSlides", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MultiCreateSlides indicates an expected call of MultiCreateSlides.
func (mr *MockRepoMockRecorder) MultiCreateSlides(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MultiCreateSlides", reflect.TypeOf((*MockRepo)(nil).MultiCreateSlides), arg0, arg1)
}

// RemoveSlide mocks base method.
func (m *MockRepo) RemoveSlide(arg0 context.Context, arg1 uint64) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveSlide", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveSlide indicates an expected call of RemoveSlide.
func (mr *MockRepoMockRecorder) RemoveSlide(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveSlide", reflect.TypeOf((*MockRepo)(nil).RemoveSlide), arg0, arg1)
}

// UpdateSlide mocks base method.
func (m *MockRepo) UpdateSlide(arg0 context.Context, arg1 model.Slide) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSlide", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateSlide indicates an expected call of UpdateSlide.
func (mr *MockRepoMockRecorder) UpdateSlide(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSlide", reflect.TypeOf((*MockRepo)(nil).UpdateSlide), arg0, arg1)
}
