// Code generated by MockGen. DO NOT EDIT.
// Source: advertise.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	model "github.com/JavaHutt/crud-api/internal/model"
	gomock "github.com/golang/mock/gomock"
)

// MockAdvertiseRepository is a mock of AdvertiseRepository interface.
type MockAdvertiseRepository struct {
	ctrl     *gomock.Controller
	recorder *MockAdvertiseRepositoryMockRecorder
}

// MockAdvertiseRepositoryMockRecorder is the mock recorder for MockAdvertiseRepository.
type MockAdvertiseRepositoryMockRecorder struct {
	mock *MockAdvertiseRepository
}

// NewMockAdvertiseRepository creates a new mock instance.
func NewMockAdvertiseRepository(ctrl *gomock.Controller) *MockAdvertiseRepository {
	mock := &MockAdvertiseRepository{ctrl: ctrl}
	mock.recorder = &MockAdvertiseRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAdvertiseRepository) EXPECT() *MockAdvertiseRepositoryMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockAdvertiseRepository) Delete(ctx context.Context, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockAdvertiseRepositoryMockRecorder) Delete(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockAdvertiseRepository)(nil).Delete), ctx, id)
}

// Get mocks base method.
func (m *MockAdvertiseRepository) Get(ctx context.Context, id int) (*model.Advertise, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, id)
	ret0, _ := ret[0].(*model.Advertise)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockAdvertiseRepositoryMockRecorder) Get(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockAdvertiseRepository)(nil).Get), ctx, id)
}

// GetAll mocks base method.
func (m *MockAdvertiseRepository) GetAll(ctx context.Context) ([]model.Advertise, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", ctx)
	ret0, _ := ret[0].([]model.Advertise)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockAdvertiseRepositoryMockRecorder) GetAll(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockAdvertiseRepository)(nil).GetAll), ctx)
}

// Insert mocks base method.
func (m *MockAdvertiseRepository) Insert(ctx context.Context, advertise model.Advertise) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", ctx, advertise)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockAdvertiseRepositoryMockRecorder) Insert(ctx, advertise interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockAdvertiseRepository)(nil).Insert), ctx, advertise)
}

// InsertBulk mocks base method.
func (m *MockAdvertiseRepository) InsertBulk(ctx context.Context, ads []model.Advertise) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertBulk", ctx, ads)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertBulk indicates an expected call of InsertBulk.
func (mr *MockAdvertiseRepositoryMockRecorder) InsertBulk(ctx, ads interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertBulk", reflect.TypeOf((*MockAdvertiseRepository)(nil).InsertBulk), ctx, ads)
}

// Update mocks base method.
func (m *MockAdvertiseRepository) Update(ctx context.Context, advertise model.Advertise) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, advertise)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockAdvertiseRepositoryMockRecorder) Update(ctx, advertise interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockAdvertiseRepository)(nil).Update), ctx, advertise)
}
