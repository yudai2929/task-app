package usecase

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yudai2929/task-app/pkg/entity"
	"github.com/yudai2929/task-app/pkg/lib/errors"
	"github.com/yudai2929/task-app/pkg/lib/errors/codes"
	"github.com/yudai2929/task-app/pkg/repository/mock"
)

func TestTaskUsecase_CreateTask(t *testing.T) {
	ctx := context.Background()
	in := &CreateTaskInput{
		UserID:      "user1",
		Title:       "Test Task",
		Description: "This is a test task",
		DueDate:     nil,
	}
	task := &entity.Task{
		ID:          "uuid",
		UserID:      "user1",
		Title:       "Test Task",
		Description: "This is a test task",
		Status:      entity.TaskStatusTodo,
		DueDate:     nil,
	}
	out := &CreateTaskOutput{
		Task: task,
	}
	tests := []struct {
		name    string
		in      *CreateTaskInput
		out     *CreateTaskOutput
		errcode codes.Code
		ctx     context.Context
		wantErr bool
		setup   func(mockTaskRepo *mock.MockTaskRepository)
	}{
		{
			name:    "success",
			in:      in,
			out:     out,
			ctx:     ctx,
			wantErr: false,
			setup: func(mockTaskRepo *mock.MockTaskRepository) {
				mockTaskRepo.EXPECT().CreateTask(ctx, task).Return(nil)
			},
		},
		{
			name:    "err: invalid",
			in:      &CreateTaskInput{},
			errcode: codes.CodeInvalidArgument,
			ctx:     ctx,
			wantErr: true,
		},
		{
			name:    "err: failed create task",
			in:      in,
			out:     nil,
			errcode: codes.CodeInternal,
			ctx:     context.Background(),
			wantErr: true,
			setup: func(mockTaskRepo *mock.MockTaskRepository) {
				mockTaskRepo.EXPECT().CreateTask(ctx, task).Return(errors.New(codes.CodeInternal))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mocks := newMocks(t)
			u := newTaskUsecaseMock(mocks)
			if tt.setup != nil {
				tt.setup(mocks.tr)
			}

			out, err := u.CreateTask(tt.ctx, tt.in)
			if tt.wantErr {
				assert.Equal(t, tt.errcode, errors.Code(err))
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.out, out)
		})
	}
}

func TestTaskUsecase_GetTask(t *testing.T) {
	ctx := context.Background()
	in := &GetTaskInput{
		UserID: "user1",
		TaskID: "task1",
	}
	task := &entity.Task{
		ID:          "task1",
		UserID:      "user1",
		Title:       "Test Task",
		Description: "This is a test task",
		Status:      entity.TaskStatusTodo,
		DueDate:     nil,
	}
	out := &GetTaskOutput{
		Task: task,
	}
	tests := []struct {
		name    string
		in      *GetTaskInput
		out     *GetTaskOutput
		errcode codes.Code
		ctx     context.Context
		wantErr bool
		setup   func(mockTaskRepo *mock.MockTaskRepository, mockAssigneeRepo *mock.MockTaskAssigneeRepository)
	}{
		{
			name:    "success: own task",
			in:      in,
			out:     out,
			ctx:     ctx,
			wantErr: false,
			setup: func(mockTaskRepo *mock.MockTaskRepository, mockAssigneeRepo *mock.MockTaskAssigneeRepository) {
				mockTaskRepo.EXPECT().GetTask(ctx, "task1").Return(task, nil)
			},
		},
		{
			name:    "success: assigned task",
			in:      &GetTaskInput{UserID: "user2", TaskID: "task1"},
			out:     out,
			ctx:     ctx,
			wantErr: false,
			setup: func(mockTaskRepo *mock.MockTaskRepository, mockAssigneeRepo *mock.MockTaskAssigneeRepository) {
				mockTaskRepo.EXPECT().GetTask(ctx, "task1").Return(task, nil)
				mockAssigneeRepo.EXPECT().GetTaskAssignee(ctx, "task1", "user2").Return(&entity.TaskAssignee{
					ID:     "assignee2",
					TaskID: "task1",
					UserID: "user2",
				}, nil)
			},
		},
		{
			name:    "err: invalid",
			in:      &GetTaskInput{},
			errcode: codes.CodeInvalidArgument,
			ctx:     ctx,
			wantErr: true,
		},
		{
			name:    "err: task not found",
			in:      in,
			out:     nil,
			errcode: codes.CodeNotFound,
			ctx:     context.Background(),
			wantErr: true,
			setup: func(mockTaskRepo *mock.MockTaskRepository, mockAssigneeRepo *mock.MockTaskAssigneeRepository) {
				mockTaskRepo.EXPECT().GetTask(ctx, "task1").Return(nil, errors.New(codes.CodeNotFound))
			},
		},
		{
			name:    "err: permission denied",
			in:      &GetTaskInput{UserID: "user3", TaskID: "task1"},
			out:     nil,
			errcode: codes.CodePermissionDenied,
			ctx:     ctx,
			wantErr: true,
			setup: func(mockTaskRepo *mock.MockTaskRepository, mockAssigneeRepo *mock.MockTaskAssigneeRepository) {
				mockTaskRepo.EXPECT().GetTask(ctx, "task1").Return(task, nil)
				mockAssigneeRepo.EXPECT().GetTaskAssignee(ctx, "task1", "user3").Return(nil, errors.New(codes.CodeNotFound))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mocks := newMocks(t)
			u := newTaskUsecaseMock(mocks)
			if tt.setup != nil {
				tt.setup(mocks.tr, mocks.ar)
			}

			out, err := u.GetTask(tt.ctx, tt.in)
			if tt.wantErr {
				assert.Equal(t, tt.errcode, errors.Code(err))
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.out, out)
		})
	}
}

func TestTaskUsecase_ListTasks(t *testing.T) {
	ctx := context.Background()
	in := &ListTasksInput{
		UserID: "user1",
	}
	task1 := &entity.Task{
		ID:          "task1",
		UserID:      "user1",
		Title:       "Test Task 1",
		Description: "This is a test task 1",
		Status:      entity.TaskStatusTodo,
		DueDate:     nil,
	}
	task2 := &entity.Task{
		ID:          "task2",
		UserID:      "user2",
		Title:       "Test Task 2",
		Description: "This is a test task 2",
		Status:      entity.TaskStatusTodo,
		DueDate:     nil,
	}
	out := &ListTasksOutput{
		Tasks: entity.Tasks{task1, task2},
	}
	tests := []struct {
		name    string
		in      *ListTasksInput
		out     *ListTasksOutput
		errcode codes.Code
		ctx     context.Context
		wantErr bool
		setup   func(mockTaskRepo *mock.MockTaskRepository)
	}{
		{
			name:    "success",
			in:      in,
			out:     out,
			ctx:     ctx,
			wantErr: false,
			setup: func(mockTaskRepo *mock.MockTaskRepository) {
				mockTaskRepo.EXPECT().ListMyTasks(ctx, "user1").Return(entity.Tasks{task1, task2}, nil)
			},
		},
		{
			name:    "err: invalid",
			in:      &ListTasksInput{},
			errcode: codes.CodeInvalidArgument,
			ctx:     ctx,
			wantErr: true,
		},
		{
			name:    "err: failed to list my tasks",
			in:      in,
			out:     nil,
			errcode: codes.CodeInternal,
			ctx:     context.Background(),
			wantErr: true,
			setup: func(mockTaskRepo *mock.MockTaskRepository) {
				mockTaskRepo.EXPECT().ListMyTasks(ctx, "user1").Return(nil, errors.New(codes.CodeInternal))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mocks := newMocks(t)
			u := newTaskUsecaseMock(mocks)
			if tt.setup != nil {
				tt.setup(mocks.tr)
			}

			out, err := u.ListTasks(tt.ctx, tt.in)
			if tt.wantErr {
				assert.Equal(t, tt.errcode, errors.Code(err))
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.out, out)
		})
	}
}
