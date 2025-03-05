package repository

import (
	"context"

	"github.com/yudai2929/task-app/pkg/entity"
)

//go:generate mkdir -p mock
//go:generate mockgen -package=mock -source=./repository.go -destination=./mock/mock.go
type UserRepository interface {
	CreateUser(ctx context.Context, user *entity.User) error
	GetUser(ctx context.Context, id string) (*entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
}

type TaskRepository interface {
	GetTask(ctx context.Context, id string) (*entity.Task, error)
	CreateTask(ctx context.Context, task *entity.Task) error
	ListTasks(ctx context.Context) (entity.Tasks, error)
	UpdateTask(ctx context.Context, task *entity.Task) error
	DeleteTask(ctx context.Context, id string) error
}

type TaskAssigneeRepository interface {
	BatchCreate(ctx context.Context, assignees entity.TaskAssignees) error
	BatchDeleteByTaskID(ctx context.Context, taskID string) error
	GetTaskAssignee(ctx context.Context, taskID, userID string) (*entity.TaskAssignee, error)
}
