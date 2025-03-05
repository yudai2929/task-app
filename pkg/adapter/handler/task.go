package handler

import (
	"context"

	api "github.com/yudai2929/task-app/doc/gen"
)

func (h *Handler) GetTask(ctx context.Context, params api.GetTaskParams) (*api.Task, error) {
	return nil, nil
}

func (h *Handler) ListTasks(ctx context.Context) ([]api.Task, error) {
	return nil, nil
}

func (h *Handler) CreateTask(ctx context.Context, req *api.CreateTaskReq) (*api.Task, error) {
	return nil, nil
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
