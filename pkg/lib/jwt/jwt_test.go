package jwt

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/yudai2929/task-app/pkg/lib/errors"
	"github.com/yudai2929/task-app/pkg/lib/errors/codes"
)

func TestGenerate(t *testing.T) {
	t.Parallel()
	secret := "secret"
	userID := "user123"
	expiry := time.Hour

	tests := []struct {
		name    string
		userID  string
		secret  string
		expiry  time.Duration
		wantErr bool
	}{
		{
			name:    "success",
			userID:  userID,
			secret:  secret,
			expiry:  expiry,
			wantErr: false,
		},
		{
			name:    "expired token",
			userID:  userID,
			secret:  secret,
			expiry:  -time.Hour,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			token, err := Generate(tt.userID, tt.secret, tt.expiry)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.NotEmpty(t, token)
		})
	}
}

func TestValidateToken(t *testing.T) {
	t.Parallel()
	secret := "secret"
	userID := "user123"
	expiry := time.Hour

	validToken, _ := Generate(userID, secret, expiry)
	expiredToken, _ := Generate(userID, secret, -time.Hour)

	tests := []struct {
		name    string
		token   string
		secret  string
		wantErr bool
		errCode codes.Code
		userID  string
	}{
		{
			name:    "success",
			token:   validToken,
			secret:  secret,
			wantErr: false,
			userID:  userID,
		},
		{
			name:    "invalid signing method",
			token:   validToken,
			secret:  "wrong_secret",
			wantErr: true,
			errCode: codes.CodeUnauthenticated,
		},
		{
			name:    "expired token",
			token:   expiredToken,
			secret:  secret,
			wantErr: true,
			errCode: codes.CodeUnauthenticated,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			claims, err := Validate(tt.token, tt.secret)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.errCode, errors.Code(err))
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.userID, claims["user_id"])
		})
	}
}
