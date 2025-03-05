package usecase

import "context"

//go:generate mkdir -p mock
//go:generate mockgen -package=mock -source=./usecase.go -destination=./mock/mock.go
type AuthUsecase interface {
	Login(ctx context.Context, in *LoginInput) (*LoginOutput, error)
	SignUp(ctx context.Context, in *SignUpInput) (*SignUpOutput, error)
}

type TaskUsecase interface {
	CreateTask(ctx context.Context, in *CreateTaskInput) (*CreateTaskOutput, error)
	GetTask(ctx context.Context, in *GetTaskInput) (*GetTaskOutput, error)
	ListTasks(ctx context.Context, in *ListTasksInput) (*ListTasksOutput, error)
	UpdateTask(ctx context.Context, in *UpdateTaskInput) (*UpdateTaskOutput, error)
	DeleteTask(ctx context.Context, in *DeleteTaskInput) error
}
