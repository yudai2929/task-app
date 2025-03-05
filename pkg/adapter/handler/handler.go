package handler

import "github.com/yudai2929/task-app/pkg/usecase"

type Handler struct {
	au usecase.AuthUsecase
	tu usecase.TaskUsecase
}

func NewHandler(au usecase.AuthUsecase, tu usecase.TaskUsecase) *Handler {
	return &Handler{
		au: au,
		tu: tu,
	}
}
