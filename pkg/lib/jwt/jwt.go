package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/yudai2929/task-app/pkg/lib/errors"
	"github.com/yudai2929/task-app/pkg/lib/errors/codes"
)

func Generate(userID string, secret string, expiry time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(expiry).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func Validate(tokenString string, secret string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.Newf(codes.CodeUnauthenticated, "unexpected signing method")
		}
		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		return nil, errors.Newf(codes.CodeUnauthenticated, "Invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.Newf(codes.CodeUnauthenticated, "Invalid token claims")
	}

	return claims, nil
}
