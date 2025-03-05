//go:build integration || !unit

package repository

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yudai2929/task-app/pkg/entity"
)

func TestTaskAssigneeRepository_CRUD(t *testing.T) {
	ctx := context.Background()
	taskAssigneeRepo := NewTaskAssigneeRepository(db)
	taskRepo := NewTaskRepository(db)
	userRepo := NewUserRepository(db)

	// Create user
	user1 := &entity.User{
		ID:           "user1",
		Email:        "user1@example.com",
		PasswordHash: "password",
		Name:         "User One",
	}
	user2 := &entity.User{
		ID:           "user2",
		Email:        "user2@example.com",
		PasswordHash: "password",
		Name:         "User Two",
	}
	err := userRepo.CreateUser(ctx, user1)
	require.NoError(t, err)
	err = userRepo.CreateUser(ctx, user2)
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
	_, err = taskRepo.CreateTask(ctx, task)
	require.NoError(t, err)

	// Create task assignees
	assignees := entity.TaskAssignees{
		&entity.TaskAssignee{
			ID:     "assignee1",
			TaskID: "task1",
			UserID: "user1",
		},
		&entity.TaskAssignee{
			ID:     "assignee2",
			TaskID: "task1",
			UserID: "user2",
		},
	}

	err = taskAssigneeRepo.BatchCreate(ctx, assignees)
	require.NoError(t, err)

	// Get task assignee by task ID and user ID
	assignee, err := taskAssigneeRepo.GetTaskAssignee(ctx, "task1", "user1")
	require.NoError(t, err)
	require.NotNil(t, assignee)

	// Delete task assignees by task ID
	err = taskAssigneeRepo.BatchDeleteByTaskID(ctx, "task1")
	require.NoError(t, err)

	// Get task assignee by task ID and user ID after delete
	assignee, err = taskAssigneeRepo.GetTaskAssignee(ctx, "task1", "user1")
	require.Error(t, err)
	require.Nil(t, assignee)
}
