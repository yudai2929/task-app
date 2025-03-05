package handler

import (
	"context"

	api "github.com/yudai2929/task-app/doc/gen"
	"github.com/yudai2929/task-app/pkg/lib/errors"
)

func (h *Handler) NewError(ctx context.Context, err error) *api.ErrorStatusCode {
	code := errors.Code(err)

	sc := code.HTTPStatus()
	msg := code.String()
	if code.HTTPStatus() < 500 {
		msg = err.Error()
	}
	return &api.ErrorStatusCode{
		StatusCode: sc,
		Response: api.Error{
			Message: msg,
		},
	}
}
