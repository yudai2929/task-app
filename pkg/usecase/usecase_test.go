package usecase

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/yudai2929/task-app/pkg/repository/mock"
	"go.uber.org/mock/gomock"
)

func newAuthUsecaseMock(mocks *mocks) *authUsecase {
	return &authUsecase{
		uuid:     func() string { return "uuid" },
		validate: validator.New(),
		ur:       mocks.ur,
		hashPassword: func(p string) (string, error) {
			return "hashed", nil
		},
		jwt: func(id string) (string, error) {
			return "jwt", nil
		},
	}
}

func newTaskUsecaseMock(mocks *mocks) *taskUsecase {
	return &taskUsecase{
		uuid:     func() string { return "uuid" },
		tr:       mocks.tr,
		ar:       mocks.ar,
		validate: validator.New(),
	}
}

type mocks struct {
	ur *mock.MockUserRepository
	tr *mock.MockTaskRepository
	ar *mock.MockTaskAssigneeRepository
}

func newMocks(t *testing.T) *mocks {
	ctrl := gomock.NewController(t)
	return &mocks{
		ur: mock.NewMockUserRepository(ctrl),
		tr: mock.NewMockTaskRepository(ctrl),
		ar: mock.NewMockTaskAssigneeRepository(ctrl),
	}
}
