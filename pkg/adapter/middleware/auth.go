package middleware

import (
	"context"
	"strings"

	"github.com/ogen-go/ogen/middleware"
	"github.com/yudai2929/task-app/pkg/lib/errors"
	"github.com/yudai2929/task-app/pkg/lib/errors/codes"
	"github.com/yudai2929/task-app/pkg/lib/jwt"
)

type contextKey string

const userIDContextKey contextKey = "userID"

func RequireAuth(secretKey string, excludePaths ...string) middleware.Middleware {
	exclude := make(map[string]struct{})
	for _, path := range excludePaths {
		exclude[path] = struct{}{}
	}

	return func(req middleware.Request, next middleware.Next) (middleware.Response, error) {
		if _, ok := exclude[req.Raw.URL.Path]; ok {
			return next(req)
		}

		authHeader := req.Raw.Header.Get("Authorization")
		if authHeader == "" {
			return middleware.Response{}, errors.Newf(codes.CodeUnauthenticated, "Authorization header is required")
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := jwt.ValidateToken(tokenString, secretKey)
		if err != nil {
			return middleware.Response{}, errors.Newf(codes.CodeUnauthenticated, err.Error())
		}

		userID, ok := claims["user_id"].(string)
		if !ok {
			return middleware.Response{}, errors.Newf(codes.CodeUnauthenticated, "Invalid user ID in token claims")
		}

		ctx := context.WithValue(req.Context, userIDContextKey, userID)
		req.Raw = req.Raw.WithContext(ctx)
		return next(req)
	}
}

func UserIDFromContext(ctx context.Context) string {
	if userID, ok := ctx.Value(userIDContextKey).(string); ok {
		return userID
	}
	return ""
}
