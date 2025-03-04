package usecase

import "context"

type AuthUsecase interface {
	Login(ctx context.Context, in *LoginInput) (*LoginOutput, error)
	SignUp(ctx context.Context, in *SingUpInput) (*SingUpOutput, error)
}

type TaskUsecase interface{}
