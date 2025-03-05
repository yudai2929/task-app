package handler

import (
	"context"

	api "github.com/yudai2929/task-app/doc/gen"
)

func (h *Handler) HealthCheck(ctx context.Context) (*api.HealthCheckOK, error) {
	return &api.HealthCheckOK{Status: "ok"}, nil
}
