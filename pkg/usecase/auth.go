package usecase

import (
	"context"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/yudai2929/task-app/pkg/entity"
	"github.com/yudai2929/task-app/pkg/lib/errors"
	"github.com/yudai2929/task-app/pkg/lib/errors/codes"
	"github.com/yudai2929/task-app/pkg/lib/jwt"
	"github.com/yudai2929/task-app/pkg/lib/password"
	"github.com/yudai2929/task-app/pkg/repository"
)

type authUsecase struct {
	ur           repository.UserRepository
	uuid         func() string
	hashPassword func(p string) (string, error)
	validate     *validator.Validate
	jwt          func(id string) (string, error)
}

func NewAuthUsecase(
	ur repository.UserRepository,
	jwtSecret string,
	tokenExpiry time.Duration,
) *authUsecase {
	return &authUsecase{
		ur:       ur,
		uuid:     uuid.NewString,
		validate: validator.New(),
		hashPassword: func(p string) (string, error) {
			return password.Hash(p)
		},
		jwt: func(id string) (string, error) {
			return jwt.Generate(id, jwtSecret, tokenExpiry)
		},
	}
}

type SingUpInput struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8"`
	Name     string `validate:"required"`
}
type SingUpOutput struct {
	User *entity.User
	JWT  string
}

func (u *authUsecase) SignUp(ctx context.Context, in *SingUpInput) (*SingUpOutput, error) {
	if err := u.validate.Struct(in); err != nil {
		return nil, errors.Convert(err)
	}

	user, err := u.ur.GetUserByEmail(ctx, in.Email)
	if err != nil {
		if !errors.EqualCode(err, codes.CodeNotFound) {
			return nil, errors.Convert(err)
		}
	}
	if user != nil {
		return nil, errors.Newf(codes.CodeAlreadyExists, "user already exists")
	}

	hashPassword, err := u.hashPassword(in.Password)
	if err != nil {
		return nil, errors.Newf(codes.CodeInternal, "failed to hash password: %w", err)
	}

	user = &entity.User{
		ID:           u.uuid(),
		Email:        in.Email,
		PasswordHash: hashPassword,
		Name:         in.Name,
	}

	if err := u.ur.CreateUser(ctx, user); err != nil {
		return nil, errors.Convert(err)
	}

	token, err := u.jwt(user.ID)
	if err != nil {
		return nil, errors.Newf(codes.CodeInternal, "failed to generate jwt: %w", err)
	}

	return &SingUpOutput{User: user, JWT: token}, nil
}

type LoginInput struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8"`
}
type LoginOutput struct {
	JWT string
}

func (u *authUsecase) Login(ctx context.Context, in *LoginInput) (*LoginOutput, error) {
	if err := u.validate.Struct(in); err != nil {
		return nil, errors.Convert(err)
	}

	user, err := u.ur.GetUserByEmail(ctx, in.Email)
	if err != nil {
		return nil, errors.Convert(err)
	}

	if !password.Equal(user.PasswordHash, in.Password) {
		return nil, errors.Newf(codes.CodeUnauthenticated, "failed to compare password: %w", err)
	}

	token, err := u.jwt(user.ID)
	if err != nil {
		return nil, errors.Newf(codes.CodeInternal, "failed to generate jwt: %w", err)
	}

	return &LoginOutput{
		JWT: token,
	}, nil
}
