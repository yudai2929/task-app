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
	txr      repository.TransactionRepository
	validate *validator.Validate
	uuid     func() string
}

func NewTaskUsecase(tr repository.TaskRepository, ar repository.TaskAssigneeRepository, txr repository.TransactionRepository) *taskUsecase {
	return &taskUsecase{
		tr:       tr,
		ar:       ar,
		txr:      txr,
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

	createdTask, err := u.tr.CreateTask(ctx, task)
	if err != nil {
		return nil, errors.Convert(err)
	}

	return &CreateTaskOutput{Task: createdTask}, nil
}

type GetTaskInput struct {
	UserID string `validate:"required"`
	TaskID string `validate:"required"`
}

type GetTaskOutput struct {
	Task *entity.Task
}

func (u *taskUsecase) hasPermission(ctx context.Context, task *entity.Task, userID string) (bool, error) {
	if task.UserID == userID {
		return true, nil
	}

	_, err := u.ar.GetTaskAssignee(ctx, task.ID, userID)
	if err != nil {
		if errors.EqualCode(err, codes.CodeNotFound) {
			return false, nil
		}
		return false, errors.Convert(err)
	}

	return true, nil
}

func (u *taskUsecase) GetTask(ctx context.Context, in *GetTaskInput) (*GetTaskOutput, error) {
	if err := u.validate.Struct(in); err != nil {
		return nil, errors.Convert(err)
	}
	task, err := u.tr.GetTask(ctx, in.TaskID)
	if err != nil {
		return nil, errors.Convert(err)
	}

	ok, err := u.hasPermission(ctx, task, in.UserID)
	if err != nil {
		return nil, errors.Convert(err)
	}
	if !ok {
		return nil, errors.Newf(codes.CodePermissionDenied, "task permission denied")
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

type UpdateTaskInput struct {
	UserID      string `validate:"required"`
	TaskID      string `validate:"required"`
	Title       string `validate:"required"`
	Description string `validate:"required"`
	Status      int    `validate:"oneof=1 2 3"`
	DueDate     *time.Time
}

type UpdateTaskOutput struct {
	Task *entity.Task
}

func (u *taskUsecase) UpdateTask(ctx context.Context, in *UpdateTaskInput) (*UpdateTaskOutput, error) {
	if err := u.validate.Struct(in); err != nil {
		return nil, errors.Convert(err)
	}
	task, err := u.tr.GetTask(ctx, in.TaskID)
	if err != nil {
		return nil, errors.Convert(err)
	}

	ok, err := u.hasPermission(ctx, task, in.UserID)
	if err != nil {
		return nil, errors.Convert(err)
	}
	if !ok {
		return nil, errors.Newf(codes.CodePermissionDenied, "task permission denied")
	}

	task.Update(in.Title, in.Description, entity.TaskStatus(in.Status), in.DueDate)
	updated, err := u.tr.UpdateTask(ctx, task)
	if err != nil {
		return nil, errors.Convert(err)
	}
	return &UpdateTaskOutput{Task: updated}, nil
}

type DeleteTaskInput struct {
	UserID string `validate:"required"`
	TaskID string `validate:"required"`
}

func (u *taskUsecase) DeleteTask(ctx context.Context, in *DeleteTaskInput) error {
	if err := u.validate.Struct(in); err != nil {
		return errors.Convert(err)
	}
	task, err := u.tr.GetTask(ctx, in.TaskID)
	if err != nil {
		return errors.Convert(err)
	}

	if task.UserID != in.UserID {
		return errors.Newf(codes.CodePermissionDenied, "task permission denied")
	}

	if err := u.tr.DeleteTask(ctx, in.TaskID); err != nil {
		return errors.Convert(err)
	}
	return nil
}

type AssignTaskInput struct {
	UserID      string   `validate:"required"`
	TaskID      string   `validate:"required"`
	AssigneeIDs []string `validate:"required"`
}

func (u *taskUsecase) AssignTask(ctx context.Context, in *AssignTaskInput) error {
	if err := u.validate.Struct(in); err != nil {
		return errors.Convert(err)
	}
	task, err := u.tr.GetTask(ctx, in.TaskID)
	if err != nil {
		return errors.Convert(err)
	}

	if task.UserID != in.UserID {
		return errors.Newf(codes.CodePermissionDenied, "task assign permission denied")
	}

	if err := u.ar.UpdateTaskAssignees(ctx, in.TaskID, in.AssigneeIDs); err != nil {
		return errors.Convert(err)
	}
	return nil
}
