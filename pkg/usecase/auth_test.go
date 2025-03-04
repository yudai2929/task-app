package usecase

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yudai2929/task-app/pkg/entity"
	"github.com/yudai2929/task-app/pkg/lib/errors"
	"github.com/yudai2929/task-app/pkg/lib/errors/codes"
	"github.com/yudai2929/task-app/pkg/lib/password"
	"github.com/yudai2929/task-app/pkg/repository/mock"
)

func TestAuthUsecase_SignUp(t *testing.T) {
	ctx := context.Background()
	in := &SingUpInput{
		Email:    "hoge@hoge.com",
		Name:     "hoge",
		Password: "password",
	}
	user := &entity.User{
		ID:           "uuid",
		Email:        "hoge@hoge.com",
		Name:         "hoge",
		PasswordHash: "hashed",
	}
	tests := []struct {
		name    string
		in      *SingUpInput
		out     *SingUpOutput
		errcode codes.Code
		ctx     context.Context
		wantErr bool
		setup   func(mockUserRepo *mock.MockUserRepository)
	}{
		{
			name:    "err: invalid",
			in:      &SingUpInput{},
			errcode: codes.CodeInvalidArgument,
			ctx:     ctx,
			wantErr: true,
		},
		{
			name:    "err: already exists",
			in:      in,
			errcode: codes.CodeAlreadyExists,
			ctx:     ctx,
			wantErr: true,
			setup: func(mockUserRepo *mock.MockUserRepository) {
				mockUserRepo.EXPECT().GetUserByEmail(ctx, "hoge@hoge.com").Return(nil, errors.New(codes.CodeAlreadyExists))
			},
		},
		{
			name:    "err: failed create user",
			in:      in,
			errcode: codes.CodeInternal,
			ctx:     ctx,
			wantErr: true,
			setup: func(mockUserRepo *mock.MockUserRepository) {
				mockUserRepo.EXPECT().GetUserByEmail(ctx, "hoge@hoge.com").Return(nil, errors.New(codes.CodeNotFound))
				mockUserRepo.EXPECT().CreateUser(ctx, user).Return(errors.New(codes.CodeInternal))
			},
		},
		{
			name: "success",
			in:   in,
			out: &SingUpOutput{
				User: user,
				JWT:  "jwt",
			},
			ctx: ctx,
			setup: func(mockUserRepo *mock.MockUserRepository) {
				mockUserRepo.EXPECT().GetUserByEmail(ctx, "hoge@hoge.com").Return(nil, errors.New(codes.CodeNotFound))
				mockUserRepo.EXPECT().CreateUser(ctx, user).Return(nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mocks := newMocks(t)
			u := newAuthUsecaseMock(mocks)
			if tt.setup != nil {
				tt.setup(mocks.ur)
			}
			out, err := u.SingUp(tt.ctx, tt.in)
			if tt.wantErr {
				assert.Equal(t, tt.errcode, errors.Code(err))
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.out, out)
		})
	}

}

func TestAuthUsecase_Login(t *testing.T) {
	ctx := context.Background()
	in := &LoginInput{
		Email:    "hoge@hoge.com",
		Password: "password",
	}
	hash, _ := password.Hash("password")
	user := &entity.User{
		ID:           "uuid",
		Email:        "hoge@hoge.com",
		Name:         "hoge",
		PasswordHash: hash,
	}
	tests := []struct {
		name    string
		in      *LoginInput
		out     *LoginOutput
		errcode codes.Code
		ctx     context.Context
		wantErr bool
		setup   func(mockUserRepo *mock.MockUserRepository)
	}{
		{
			name:    "err: invalid",
			in:      &LoginInput{},
			errcode: codes.CodeInvalidArgument,
			ctx:     ctx,
			wantErr: true,
		},
		{
			name:    "err: not found",
			in:      in,
			errcode: codes.CodeNotFound,
			ctx:     ctx,
			wantErr: true,
			setup: func(mockUserRepo *mock.MockUserRepository) {
				mockUserRepo.EXPECT().GetUserByEmail(ctx, "hoge@hoge.com").Return(nil, errors.New(codes.CodeNotFound))
			},
		},
		{
			name:    "err: password not match",
			in:      in,
			errcode: codes.CodeUnauthenticated,
			ctx:     ctx,
			wantErr: true,
			setup: func(mockUserRepo *mock.MockUserRepository) {
				mockUserRepo.EXPECT().GetUserByEmail(ctx, "hoge@hoge.com").Return(&entity.User{
					ID:           "uuid",
					Email:        "hoge@hoge.com",
					Name:         "hoge",
					PasswordHash: "hashed",
				}, nil)
			},
		},
		{
			name: "success",
			in:   in,
			ctx:  ctx,
			setup: func(mockUserRepo *mock.MockUserRepository) {
				mockUserRepo.EXPECT().GetUserByEmail(ctx, "hoge@hoge.com").Return(user, nil)
			},
			out: &LoginOutput{
				JWT: "jwt",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mocks := newMocks(t)
			u := newAuthUsecaseMock(mocks)
			if tt.setup != nil {
				tt.setup(mocks.ur)
			}
			out, err := u.Login(tt.ctx, tt.in)
			if tt.wantErr {
				assert.Equal(t, tt.errcode, errors.Code(err))
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.out, out)
		})
	}
}
