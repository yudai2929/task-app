package repository

import (
	"context"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
	"github.com/yudai2929/task-app/pkg/entity"
	"github.com/yudai2929/task-app/pkg/lib/errors"
	"github.com/yudai2929/task-app/pkg/lib/errors/codes"
)

func TestUserRepository_CRUD(t *testing.T) {
	ctx := context.Background()
	repo := NewUserRepository(db)

	// Get (not found)
	user, err := repo.GetUser(ctx, "id")
	require.Error(t, err)
	require.Nil(t, user)
	require.Equal(t, codes.CodeNotFound, errors.Code(err))

	// Create
	user = &entity.User{
		ID:           "id",
		Email:        "hoge@hoge.com",
		PasswordHash: "password",
		Name:         "hoge",
	}

	err = repo.CreateUser(ctx, user)
	require.NoError(t, err)

	// Get
	user, err = repo.GetUser(ctx, "id")
	require.NoError(t, err)
	require.NotNil(t, user)
	require.Equal(t, "id", user.ID)

	// Get by email
	user, err = repo.GetUserByEmail(ctx, "hoge@hoge.com")
	require.NoError(t, err)
	require.NotNil(t, user)
	require.Equal(t, "id", user.ID)
}
