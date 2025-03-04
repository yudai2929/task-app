package handler

import "github.com/yudai2929/task-app/pkg/usecase"

type Handler struct {
	au usecase.AuthUsecase
}

func NewHandler() *Handler {
	return &Handler{}
}
