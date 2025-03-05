package handler

import (
	"context"

	api "github.com/yudai2929/task-app/doc/gen"
	"github.com/yudai2929/task-app/pkg/entity"
	"github.com/yudai2929/task-app/pkg/usecase"
)

func (h *Handler) GetTask(ctx context.Context, params api.GetTaskParams) (*api.Task, error) {
	userID := UserIDFromContext(ctx)
	in := &usecase.GetTaskInput{
		TaskID: params.ID,
		UserID: userID,
	}
	out, err := h.tu.GetTask(ctx, in)
	if err != nil {
		return nil, err
	}

	return h.convertToAPITask(out.Task), nil
}

func (h *Handler) ListTasks(ctx context.Context) ([]api.Task, error) {
	userID := UserIDFromContext(ctx)
	in := &usecase.ListTasksInput{
		UserID: userID,
	}
	out, err := h.tu.ListTasks(ctx, in)
	if err != nil {
		return nil, err
	}

	return h.convertToAPITasks(out.Tasks), nil
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

	return h.convertToAPITask(out.Task), nil
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

func (h *Handler) convertToAPITask(task *entity.Task) *api.Task {
	dueDateResp := api.NewOptDateTime(*task.DueDate)
	if task.DueDate == nil {
		dueDateResp.Reset()
	}
	return &api.Task{
		ID:          task.ID,
		UserID:      task.UserID,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status.Int(),
		DueDate:     dueDateResp,
	}
}

func (h *Handler) convertToAPITasks(tasks entity.Tasks) []api.Task {
	var apiTasks []api.Task
	for _, task := range tasks {
		apiTasks = append(apiTasks, *h.convertToAPITask(task))
	}
	return apiTasks
}
