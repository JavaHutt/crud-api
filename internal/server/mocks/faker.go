// Code generated by MockGen. DO NOT EDIT.
// Source: faker.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	model "github.com/JavaHutt/crud-api/internal/model"
	gomock "github.com/golang/mock/gomock"
)

// MockfakerService is a mock of fakerService interface.
type MockfakerService struct {
	ctrl     *gomock.Controller
	recorder *MockfakerServiceMockRecorder
}

// MockfakerServiceMockRecorder is the mock recorder for MockfakerService.
type MockfakerServiceMockRecorder struct {
	mock *MockfakerService
}

// NewMockfakerService creates a new mock instance.
func NewMockfakerService(ctrl *gomock.Controller) *MockfakerService {
	mock := &MockfakerService{ctrl: ctrl}
	mock.recorder = &MockfakerServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockfakerService) EXPECT() *MockfakerServiceMockRecorder {
	return m.recorder
}

// Fake mocks base method.
func (m *MockfakerService) Fake(num int) []model.SlowestQuery {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fake", num)
	ret0, _ := ret[0].([]model.SlowestQuery)
	return ret0
}

// Fake indicates an expected call of Fake.
func (mr *MockfakerServiceMockRecorder) Fake(num interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fake", reflect.TypeOf((*MockfakerService)(nil).Fake), num)
}

// MockquerService is a mock of querService interface.
type MockquerService struct {
	ctrl     *gomock.Controller
	recorder *MockquerServiceMockRecorder
}

// MockquerServiceMockRecorder is the mock recorder for MockquerService.
type MockquerServiceMockRecorder struct {
	mock *MockquerService
}

// NewMockquerService creates a new mock instance.
func NewMockquerService(ctrl *gomock.Controller) *MockquerService {
	mock := &MockquerService{ctrl: ctrl}
	mock.recorder = &MockquerServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockquerService) EXPECT() *MockquerServiceMockRecorder {
	return m.recorder
}

// InsertBulk mocks base method.
func (m *MockquerService) InsertBulk(ctx context.Context, queries []model.SlowestQuery) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertBulk", ctx, queries)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertBulk indicates an expected call of InsertBulk.
func (mr *MockquerServiceMockRecorder) InsertBulk(ctx, queries interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertBulk", reflect.TypeOf((*MockquerService)(nil).InsertBulk), ctx, queries)
}
