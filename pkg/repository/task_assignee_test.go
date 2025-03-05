package repository

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yudai2929/task-app/pkg/entity"
)

func TestTaskAssigneeRepository_CRUD(t *testing.T) {
	t.Parallel()
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
	userIDs := []string{"user1", "user2"}

	err = taskAssigneeRepo.UpdateTaskAssignees(ctx, "task1", userIDs)
	require.NoError(t, err)

	// Get task assignee by task ID and user ID
	assignee, err := taskAssigneeRepo.GetTaskAssignee(ctx, "task1", "user1")
	require.NoError(t, err)
	require.NotNil(t, assignee)

	// Update task assignees
	newUserIDs := []string{"user1"}

	err = taskAssigneeRepo.UpdateTaskAssignees(ctx, "task1", newUserIDs)
	require.NoError(t, err)

	// Get task assignee by task ID and user ID after update
	assignee, err = taskAssigneeRepo.GetTaskAssignee(ctx, "task1", "user1")
	require.NoError(t, err)
	require.NotNil(t, assignee)

	// Delete task assignees by task ID
	err = taskAssigneeRepo.UpdateTaskAssignees(ctx, "task1", []string{})
	require.NoError(t, err)

	// Get task assignee by task ID and user ID after delete
	assignee, err = taskAssigneeRepo.GetTaskAssignee(ctx, "task1", "user1")
	require.Error(t, err)
	require.Nil(t, assignee)
}
