// Code generated by MockGen. DO NOT EDIT.
// Source: internal/entity/interface.go
//
// Generated by this command:
//
//	mockgen -source=internal/entity/interface.go -package=mocks -destination=test/mocks/interface_mock.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	entity "todoList/internal/entity"

	uuid "github.com/google/uuid"
	gomock "go.uber.org/mock/gomock"
)

// MockTodoRepositoryInterface is a mock of TodoRepositoryInterface interface.
type MockTodoRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockTodoRepositoryInterfaceMockRecorder
}

// MockTodoRepositoryInterfaceMockRecorder is the mock recorder for MockTodoRepositoryInterface.
type MockTodoRepositoryInterfaceMockRecorder struct {
	mock *MockTodoRepositoryInterface
}

// NewMockTodoRepositoryInterface creates a new mock instance.
func NewMockTodoRepositoryInterface(ctrl *gomock.Controller) *MockTodoRepositoryInterface {
	mock := &MockTodoRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockTodoRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTodoRepositoryInterface) EXPECT() *MockTodoRepositoryInterfaceMockRecorder {
	return m.recorder
}

// Complete mocks base method.
func (m *MockTodoRepositoryInterface) Complete(todo *entity.Todo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Complete", todo)
	ret0, _ := ret[0].(error)
	return ret0
}

// Complete indicates an expected call of Complete.
func (mr *MockTodoRepositoryInterfaceMockRecorder) Complete(todo any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Complete", reflect.TypeOf((*MockTodoRepositoryInterface)(nil).Complete), todo)
}

// DeleteByID mocks base method.
func (m *MockTodoRepositoryInterface) DeleteByID(id uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByID", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByID indicates an expected call of DeleteByID.
func (mr *MockTodoRepositoryInterfaceMockRecorder) DeleteByID(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockTodoRepositoryInterface)(nil).DeleteByID), id)
}

// ExistsByID mocks base method.
func (m *MockTodoRepositoryInterface) ExistsByID(id uuid.UUID) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExistsByID", id)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExistsByID indicates an expected call of ExistsByID.
func (mr *MockTodoRepositoryInterfaceMockRecorder) ExistsByID(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExistsByID", reflect.TypeOf((*MockTodoRepositoryInterface)(nil).ExistsByID), id)
}

// GetAll mocks base method.
func (m *MockTodoRepositoryInterface) GetAll() ([]*entity.Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]*entity.Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockTodoRepositoryInterfaceMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockTodoRepositoryInterface)(nil).GetAll))
}

// GetByID mocks base method.
func (m *MockTodoRepositoryInterface) GetByID(id uuid.UUID) (*entity.Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", id)
	ret0, _ := ret[0].(*entity.Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockTodoRepositoryInterfaceMockRecorder) GetByID(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockTodoRepositoryInterface)(nil).GetByID), id)
}

// Save mocks base method.
func (m *MockTodoRepositoryInterface) Save(todo *entity.Todo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", todo)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockTodoRepositoryInterfaceMockRecorder) Save(todo any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockTodoRepositoryInterface)(nil).Save), todo)
}

// Undo mocks base method.
func (m *MockTodoRepositoryInterface) Undo(todo *entity.Todo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Undo", todo)
	ret0, _ := ret[0].(error)
	return ret0
}

// Undo indicates an expected call of Undo.
func (mr *MockTodoRepositoryInterfaceMockRecorder) Undo(todo any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Undo", reflect.TypeOf((*MockTodoRepositoryInterface)(nil).Undo), todo)
}

// Update mocks base method.
func (m *MockTodoRepositoryInterface) Update(todo *entity.Todo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", todo)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockTodoRepositoryInterfaceMockRecorder) Update(todo any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockTodoRepositoryInterface)(nil).Update), todo)
}
