package handler

import (
	"context"

	api "github.com/yudai2929/task-app/doc/gen"
	"github.com/yudai2929/task-app/pkg/lib/errors"
	"github.com/yudai2929/task-app/pkg/lib/errors/codes"
	"github.com/yudai2929/task-app/pkg/lib/jwt"
	"github.com/yudai2929/task-app/pkg/repository"
)

type contextKey string

const userIDContextKey contextKey = "userID"

type SecurityHandler struct {
	ur        repository.UserRepository
	secretKey string
}

func NewSecurityHandler(secretKey string, ur repository.UserRepository) *SecurityHandler {
	return &SecurityHandler{
		ur:        ur,
		secretKey: secretKey,
	}
}

func (sh *SecurityHandler) HandleBearerAuth(ctx context.Context, operationName api.OperationName, t api.BearerAuth) (context.Context, error) {
	claims, err := jwt.Validate(t.GetToken(), sh.secretKey)
	if err != nil {
		return ctx, errors.Newf(codes.CodeUnauthenticated, err.Error())
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		return ctx, errors.Newf(codes.CodeUnauthenticated, "Invalid user ID in token claims")
	}

	ctx = context.WithValue(ctx, userIDContextKey, userID)

	return ctx, nil
}

func UserIDFromContext(ctx context.Context) string {
	if userID, ok := ctx.Value(userIDContextKey).(string); ok {
		return userID
	}
	return ""
}
