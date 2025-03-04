package handler

import "github.com/yudai2929/task-app/pkg/usecase"

type Handler struct {
	au usecase.AuthUsecase
}

func NewHandler(au usecase.AuthUsecase) *Handler {
	return &Handler{
		au: au,
	}
}
