package handler

import (
	"context"

	api "github.com/yudai2929/task-app/doc/gen"
	"github.com/yudai2929/task-app/pkg/lib/errors"
)

func (h *Handler) NewError(ctx context.Context, err error) *api.ErrorStatusCode {
	code := errors.Code(err)
	return &api.ErrorStatusCode{
		StatusCode: code.HTTPStatus(),
		Response: api.Error{
			Message: code.String(),
		},
	}
}
