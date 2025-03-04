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

type TaskRepository interface{}
