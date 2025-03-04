package repository

import (
	"context"
	"database/sql"

	xo "github.com/yudai2929/task-app/database/gen"
	"github.com/yudai2929/task-app/pkg/entity"
	"github.com/yudai2929/task-app/pkg/lib/errors"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) CreateUser(ctx context.Context, user *entity.User) error {
	u := &xo.User{
		ID:           user.ID,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
		Name:         user.Name,
	}
	if err := u.Insert(ctx, r.db); err != nil {
		return errors.Convert(err)
	}
	return nil
}

func (r *userRepository) GetUser(ctx context.Context, id string) (*entity.User, error) {
	u, err := xo.UserByID(ctx, r.db, id)
	if err != nil {
		return nil, errors.Convert(err)
	}

	return r.convertUser(u), nil
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	u, err := xo.UserByEmail(ctx, r.db, email)
	if err != nil {
		return nil, errors.Convert(err)
	}

	return r.convertUser(u), nil
}

func (r *userRepository) convertUser(u *xo.User) *entity.User {
	return &entity.User{
		ID:           u.ID,
		Name:         u.Name,
		Email:        u.Email,
		PasswordHash: u.PasswordHash,
	}
}
