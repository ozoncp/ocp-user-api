// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ozoncp/ocp-user-api/internal/repo (interfaces: Repo)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/ozoncp/ocp-user-api/internal/models"
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

// CreateUser mocks base method.
func (m *MockRepo) CreateUser(arg0 context.Context, arg1 *models.User) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockRepoMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockRepo)(nil).CreateUser), arg0, arg1)
}

// CreateUsers mocks base method.
func (m *MockRepo) CreateUsers(arg0 context.Context, arg1 []models.User) ([]uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUsers", arg0, arg1)
	ret0, _ := ret[0].([]uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUsers indicates an expected call of CreateUsers.
func (mr *MockRepoMockRecorder) CreateUsers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUsers", reflect.TypeOf((*MockRepo)(nil).CreateUsers), arg0, arg1)
}

// GetUser mocks base method.
func (m *MockRepo) GetUser(arg0 context.Context, arg1 uint64) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockRepoMockRecorder) GetUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockRepo)(nil).GetUser), arg0, arg1)
}

// GetUsers mocks base method.
func (m *MockRepo) GetUsers(arg0 context.Context, arg1 []uint64) ([]models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsers", arg0, arg1)
	ret0, _ := ret[0].([]models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsers indicates an expected call of GetUsers.
func (mr *MockRepoMockRecorder) GetUsers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsers", reflect.TypeOf((*MockRepo)(nil).GetUsers), arg0, arg1)
}

// RemoveUser mocks base method.
func (m *MockRepo) RemoveUser(arg0 context.Context, arg1 uint64) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveUser", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveUser indicates an expected call of RemoveUser.
func (mr *MockRepoMockRecorder) RemoveUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveUser", reflect.TypeOf((*MockRepo)(nil).RemoveUser), arg0, arg1)
}

// SearchUsers mocks base method.
func (m *MockRepo) SearchUsers(arg0 context.Context, arg1 models.UserSearchParams) (*models.UserSearchResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchUsers", arg0, arg1)
	ret0, _ := ret[0].(*models.UserSearchResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchUsers indicates an expected call of SearchUsers.
func (mr *MockRepoMockRecorder) SearchUsers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchUsers", reflect.TypeOf((*MockRepo)(nil).SearchUsers), arg0, arg1)
}

// UpdateUser mocks base method.
func (m *MockRepo) UpdateUser(arg0 context.Context, arg1 *models.User) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockRepoMockRecorder) UpdateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockRepo)(nil).UpdateUser), arg0, arg1)
}
