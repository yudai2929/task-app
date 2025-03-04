package repository

import (
	"context"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
	"github.com/yudai2929/task-app/pkg/entity"
	"github.com/yudai2929/task-app/pkg/lib/errors"
	"github.com/yudai2929/task-app/pkg/lib/errors/codes"
)

func TestTaskRepository_CRUD(t *testing.T) {
	ctx := context.Background()
	taskRepo := NewTaskRepository(db)
	userRepo := NewUserRepository(db)

	// Create user
	user := &entity.User{
		ID:           "user1",
		Email:        "user1@example.com",
		PasswordHash: "password",
		Name:         "User One",
	}
	err := userRepo.CreateUser(ctx, user)
	require.NoError(t, err)

	// Create task
	task := &entity.Task{
		ID:          "task1",
		UserID:      "user1",
		Title:       "Test Task",
		Description: "This is a test task",
		Status:      entity.TaskStatusTodo,
		DueDate:     nil,
	}

	err = taskRepo.CreateTask(ctx, task)
	require.NoError(t, err)

	// Get task
	task, err = taskRepo.GetTask(ctx, "task1")
	require.NoError(t, err)
	require.NotNil(t, task)
	require.Equal(t, "task1", task.ID)

	// Update task
	task.Title = "Updated Test Task"
	err = taskRepo.UpdateTask(ctx, task)
	require.NoError(t, err)

	// Get task after update
	task, err = taskRepo.GetTask(ctx, "task1")
	require.NoError(t, err)
	require.NotNil(t, task)
	require.Equal(t, "Updated Test Task", task.Title)

	// List tasks
	tasks, err := taskRepo.ListTasks(ctx)
	require.NoError(t, err)
	require.NotEmpty(t, tasks)

	// Delete task
	err = taskRepo.DeleteTask(ctx, "task1")
	require.NoError(t, err)

	// Get task after delete
	task, err = taskRepo.GetTask(ctx, "task1")
	require.Error(t, err)
	require.Nil(t, task)
	require.Equal(t, codes.CodeNotFound, errors.Code(err))
}
