package usecase

import (
	"context"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/yudai2929/task-app/pkg/entity"
	"github.com/yudai2929/task-app/pkg/lib/errors"
	"github.com/yudai2929/task-app/pkg/lib/errors/codes"
	"github.com/yudai2929/task-app/pkg/repository"
)

type taskUsecase struct {
	tr       repository.TaskRepository
	ar       repository.TaskAssigneeRepository
	validate *validator.Validate
	uuid     func() string
}

func NewTaskUsecase(tr repository.TaskRepository, ar repository.TaskAssigneeRepository) *taskUsecase {
	return &taskUsecase{
		tr:       tr,
		ar:       ar,
		validate: validator.New(),
		uuid:     uuid.NewString,
	}
}

type CreateTaskInput struct {
	UserID      string `validate:"required"`
	Title       string `validate:"required"`
	Description string `validate:"required"`
	DueDate     *time.Time
}

type CreateTaskOutput struct {
	Task *entity.Task
}

func (u *taskUsecase) CreateTask(ctx context.Context, in *CreateTaskInput) (*CreateTaskOutput, error) {
	if err := u.validate.Struct(in); err != nil {
		return nil, errors.Convert(err)
	}

	task := &entity.Task{
		ID:          u.uuid(),
		UserID:      in.UserID,
		Title:       in.Title,
		Description: in.Description,
		Status:      entity.TaskStatusTodo,
		DueDate:     in.DueDate,
	}

	if err := u.tr.CreateTask(ctx, task); err != nil {
		return nil, errors.Convert(err)
	}

	return &CreateTaskOutput{Task: task}, nil
}

type GetTaskInput struct {
	UserID string `validate:"required"`
	TaskID string `validate:"required"`
}

type GetTaskOutput struct {
	Task *entity.Task
}

func (u *taskUsecase) GetTask(ctx context.Context, in *GetTaskInput) (*GetTaskOutput, error) {
	if err := u.validate.Struct(in); err != nil {
		return nil, errors.Convert(err)
	}

	task, err := u.tr.GetTask(ctx, in.TaskID)
	if err != nil {
		return nil, errors.Convert(err)
	}

	if task.UserID == in.UserID {
		return &GetTaskOutput{Task: task}, nil
	}

	_, err = u.ar.GetTaskAssignee(ctx, in.TaskID, in.UserID)
	if err != nil {
		if errors.EqualCode(err, codes.CodeNotFound) {
			return nil, errors.Newf(codes.CodePermissionDenied, "task permission denied")
		}
		return nil, errors.Convert(err)
	}

	return &GetTaskOutput{Task: task}, nil
}

type ListTasksInput struct {
	UserID string `validate:"required"`
}

type ListTasksOutput struct {
	Tasks entity.Tasks
}

func (u *taskUsecase) ListTasks(ctx context.Context, in *ListTasksInput) (*ListTasksOutput, error) {
	if err := u.validate.Struct(in); err != nil {
		return nil, errors.Convert(err)
	}

	tasks, err := u.tr.ListMyTasks(ctx, in.UserID)
	if err != nil {
		return nil, errors.Convert(err)
	}

	return &ListTasksOutput{Tasks: tasks}, nil
}
