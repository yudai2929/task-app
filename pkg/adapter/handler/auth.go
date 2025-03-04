package handler

import (
	"context"

	api "github.com/yudai2929/task-app/doc/gen"
	"github.com/yudai2929/task-app/pkg/lib/errors"
	"github.com/yudai2929/task-app/pkg/usecase"
)

func (h *Handler) Login(ctx context.Context, req *api.LoginReq) (*api.LoginOK, error) {
	in := usecase.LoginInput{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}
	out, err := h.au.Login(ctx, &in)
	if err != nil {
		return nil, errors.Convert(err)
	}

	return &api.LoginOK{
		Jwt: out.JWT,
	}, nil
}

func (h *Handler) SignUp(ctx context.Context, req *api.SignUpReq) (*api.SignUpCreated, error) {
	in := usecase.SingUpInput{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
		Name:     req.GetName(),
	}
	out, err := h.au.SignUp(ctx, &in)
	if err != nil {
		return nil, errors.Convert(err)
	}

	return &api.SignUpCreated{
		Jwt: out.JWT,
		User: api.User{
			ID:    out.User.ID,
			Name:  out.User.Name,
			Email: out.User.Email,
		},
	}, nil
}
