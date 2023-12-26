package usecase

import (
	"golang-project-template/internal/domain/user"
	"golang-project-template/internal/repository/mock"
	"testing"
)

func TestAuthUseCase_Signup(t *testing.T) {
	_ = authUseCase{
		userRepository: mock.NewUserMockRepository(),
		userFactory:    user.UserFactory{},
	}

	t.Run("new user", func(t *testing.T) {

	})

	t.Run("exists user", func(t *testing.T) {

	})
}
