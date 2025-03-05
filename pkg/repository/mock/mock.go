// Code generated by MockGen. DO NOT EDIT.
// Source: ./repository.go
//
// Generated by this command:
//
//	mockgen -package=mock -source=./repository.go -destination=./mock/mock.go
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	entity "github.com/yudai2929/task-app/pkg/entity"
	gomock "go.uber.org/mock/gomock"
)

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockUserRepository) CreateUser(ctx context.Context, user *entity.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserRepositoryMockRecorder) CreateUser(ctx, user any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserRepository)(nil).CreateUser), ctx, user)
}

// GetUser mocks base method.
func (m *MockUserRepository) GetUser(ctx context.Context, id string) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", ctx, id)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockUserRepositoryMockRecorder) GetUser(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockUserRepository)(nil).GetUser), ctx, id)
}

// GetUserByEmail mocks base method.
func (m *MockUserRepository) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByEmail", ctx, email)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByEmail indicates an expected call of GetUserByEmail.
func (mr *MockUserRepositoryMockRecorder) GetUserByEmail(ctx, email any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmail", reflect.TypeOf((*MockUserRepository)(nil).GetUserByEmail), ctx, email)
}

// MockTaskRepository is a mock of TaskRepository interface.
type MockTaskRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTaskRepositoryMockRecorder
}

// MockTaskRepositoryMockRecorder is the mock recorder for MockTaskRepository.
type MockTaskRepositoryMockRecorder struct {
	mock *MockTaskRepository
}

// NewMockTaskRepository creates a new mock instance.
func NewMockTaskRepository(ctrl *gomock.Controller) *MockTaskRepository {
	mock := &MockTaskRepository{ctrl: ctrl}
	mock.recorder = &MockTaskRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTaskRepository) EXPECT() *MockTaskRepositoryMockRecorder {
	return m.recorder
}

// CreateTask mocks base method.
func (m *MockTaskRepository) CreateTask(ctx context.Context, task *entity.Task) (*entity.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTask", ctx, task)
	ret0, _ := ret[0].(*entity.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTask indicates an expected call of CreateTask.
func (mr *MockTaskRepositoryMockRecorder) CreateTask(ctx, task any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTask", reflect.TypeOf((*MockTaskRepository)(nil).CreateTask), ctx, task)
}

// DeleteTask mocks base method.
func (m *MockTaskRepository) DeleteTask(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTask", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTask indicates an expected call of DeleteTask.
func (mr *MockTaskRepositoryMockRecorder) DeleteTask(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTask", reflect.TypeOf((*MockTaskRepository)(nil).DeleteTask), ctx, id)
}

// GetTask mocks base method.
func (m *MockTaskRepository) GetTask(ctx context.Context, id string) (*entity.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTask", ctx, id)
	ret0, _ := ret[0].(*entity.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTask indicates an expected call of GetTask.
func (mr *MockTaskRepositoryMockRecorder) GetTask(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTask", reflect.TypeOf((*MockTaskRepository)(nil).GetTask), ctx, id)
}

// ListMyTasks mocks base method.
func (m *MockTaskRepository) ListMyTasks(ctx context.Context, userID string) (entity.Tasks, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListMyTasks", ctx, userID)
	ret0, _ := ret[0].(entity.Tasks)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListMyTasks indicates an expected call of ListMyTasks.
func (mr *MockTaskRepositoryMockRecorder) ListMyTasks(ctx, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMyTasks", reflect.TypeOf((*MockTaskRepository)(nil).ListMyTasks), ctx, userID)
}

// ListTasks mocks base method.
func (m *MockTaskRepository) ListTasks(ctx context.Context) (entity.Tasks, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTasks", ctx)
	ret0, _ := ret[0].(entity.Tasks)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTasks indicates an expected call of ListTasks.
func (mr *MockTaskRepositoryMockRecorder) ListTasks(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTasks", reflect.TypeOf((*MockTaskRepository)(nil).ListTasks), ctx)
}

// UpdateTask mocks base method.
func (m *MockTaskRepository) UpdateTask(ctx context.Context, task *entity.Task) (*entity.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTask", ctx, task)
	ret0, _ := ret[0].(*entity.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTask indicates an expected call of UpdateTask.
func (mr *MockTaskRepositoryMockRecorder) UpdateTask(ctx, task any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTask", reflect.TypeOf((*MockTaskRepository)(nil).UpdateTask), ctx, task)
}

// MockTaskAssigneeRepository is a mock of TaskAssigneeRepository interface.
type MockTaskAssigneeRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTaskAssigneeRepositoryMockRecorder
}

// MockTaskAssigneeRepositoryMockRecorder is the mock recorder for MockTaskAssigneeRepository.
type MockTaskAssigneeRepositoryMockRecorder struct {
	mock *MockTaskAssigneeRepository
}

// NewMockTaskAssigneeRepository creates a new mock instance.
func NewMockTaskAssigneeRepository(ctrl *gomock.Controller) *MockTaskAssigneeRepository {
	mock := &MockTaskAssigneeRepository{ctrl: ctrl}
	mock.recorder = &MockTaskAssigneeRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTaskAssigneeRepository) EXPECT() *MockTaskAssigneeRepositoryMockRecorder {
	return m.recorder
}

// BatchCreate mocks base method.
func (m *MockTaskAssigneeRepository) BatchCreate(ctx context.Context, assignees entity.TaskAssignees) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchCreate", ctx, assignees)
	ret0, _ := ret[0].(error)
	return ret0
}

// BatchCreate indicates an expected call of BatchCreate.
func (mr *MockTaskAssigneeRepositoryMockRecorder) BatchCreate(ctx, assignees any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchCreate", reflect.TypeOf((*MockTaskAssigneeRepository)(nil).BatchCreate), ctx, assignees)
}

// BatchDeleteByTaskID mocks base method.
func (m *MockTaskAssigneeRepository) BatchDeleteByTaskID(ctx context.Context, taskID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchDeleteByTaskID", ctx, taskID)
	ret0, _ := ret[0].(error)
	return ret0
}

// BatchDeleteByTaskID indicates an expected call of BatchDeleteByTaskID.
func (mr *MockTaskAssigneeRepositoryMockRecorder) BatchDeleteByTaskID(ctx, taskID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchDeleteByTaskID", reflect.TypeOf((*MockTaskAssigneeRepository)(nil).BatchDeleteByTaskID), ctx, taskID)
}

// GetTaskAssignee mocks base method.
func (m *MockTaskAssigneeRepository) GetTaskAssignee(ctx context.Context, taskID, userID string) (*entity.TaskAssignee, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTaskAssignee", ctx, taskID, userID)
	ret0, _ := ret[0].(*entity.TaskAssignee)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTaskAssignee indicates an expected call of GetTaskAssignee.
func (mr *MockTaskAssigneeRepositoryMockRecorder) GetTaskAssignee(ctx, taskID, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTaskAssignee", reflect.TypeOf((*MockTaskAssigneeRepository)(nil).GetTaskAssignee), ctx, taskID, userID)
}

// MockTransactionRepository is a mock of TransactionRepository interface.
type MockTransactionRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionRepositoryMockRecorder
}

// MockTransactionRepositoryMockRecorder is the mock recorder for MockTransactionRepository.
type MockTransactionRepositoryMockRecorder struct {
	mock *MockTransactionRepository
}

// NewMockTransactionRepository creates a new mock instance.
func NewMockTransactionRepository(ctrl *gomock.Controller) *MockTransactionRepository {
	mock := &MockTransactionRepository{ctrl: ctrl}
	mock.recorder = &MockTransactionRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransactionRepository) EXPECT() *MockTransactionRepositoryMockRecorder {
	return m.recorder
}

// Run mocks base method.
func (m *MockTransactionRepository) Run(ctx context.Context, fn func(context.Context) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", ctx, fn)
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run.
func (mr *MockTransactionRepositoryMockRecorder) Run(ctx, fn any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockTransactionRepository)(nil).Run), ctx, fn)
}
