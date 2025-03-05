package handler

import (
	"context"

	api "github.com/yudai2929/task-app/doc/gen"
	"github.com/yudai2929/task-app/pkg/usecase"
)

func (h *Handler) GetTask(ctx context.Context, params api.GetTaskParams) (*api.Task, error) {
	return nil, nil
}

func (h *Handler) ListTasks(ctx context.Context) ([]api.Task, error) {
	return nil, nil
}

func (h *Handler) CreateTask(ctx context.Context, req *api.CreateTaskReq) (*api.Task, error) {
	userID := UserIDFromContext(ctx)

	dueDate := req.GetDueDate().Value
	in := &usecase.CreateTaskInput{
		UserID:      userID,
		Title:       req.GetTitle(),
		Description: req.GetDescription(),
		DueDate:     &dueDate,
	}
	out, err := h.tu.CreateTask(ctx, in)
	if err != nil {
		return nil, err
	}

	dueDateResp := api.NewOptDateTime(*out.Task.DueDate)
	if out.Task.DueDate == nil {
		dueDateResp.Reset()
	}
	return &api.Task{
		ID:          out.Task.ID,
		UserID:      out.Task.UserID,
		Title:       out.Task.Title,
		Description: out.Task.Description,
		Status:      out.Task.Status.Int(),
		DueDate:     dueDateResp,
	}, nil
}

func (h *Handler) UpdateTask(ctx context.Context, req *api.UpdateTaskReq, params api.UpdateTaskParams) (*api.Task, error) {
	return nil, nil
}

func (h *Handler) DeleteTask(ctx context.Context, params api.DeleteTaskParams) error {
	return nil
}

func (h *Handler) AssignTask(ctx context.Context, req *api.AssignTaskReq, params api.AssignTaskParams) error {
	return nil
}
