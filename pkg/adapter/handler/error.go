package handler

import (
	"context"

	api "github.com/yudai2929/task-app/doc/gen"
)

func (h *Handler) NewError(ctx context.Context, err error) *api.ErrorStatusCode {
	return nil
}
