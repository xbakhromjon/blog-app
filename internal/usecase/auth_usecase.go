package usecase

import (
	"errors"
	"golang-project-template/internal/domain/auth"
	"golang-project-template/internal/domain/user"
	"strconv"
)

type authUseCase struct {
	userRepository user.UserRepositroy
	userFactory    user.UserFactory
}

func NewAuthUseCase(userRepository user.UserRepositroy, userFactory user.UserFactory) auth.AuthUseCase {

	return &authUseCase{userRepository: userRepository, userFactory: userFactory}
}

func (a *authUseCase) Signup(request auth.SignUpRequest) (string, error) {
	// checking user exist by email
	exists, err := a.userRepository.ExistsByEmail(request.Email)
	if err != nil {
		return "", err
	}
	if exists {
		return "", errors.New("user already signed up")
	}
	// create new user by request
	newUser := a.userFactory.NewUser(request.Email, request.Firstname, request.Lastname)

	// save new user
	id, err := a.userRepository.Save(newUser)
	if err != nil {
		return "", err
	}

	return strconv.FormatInt(id, 10), nil
}
