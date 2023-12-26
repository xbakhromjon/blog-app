package usecase

import (
	"github.com/golang-jwt/jwt/v5"
	"golang-project-template/internal/common"
	"golang-project-template/internal/domain/auth"
	userdomain "golang-project-template/internal/domain/user"
	"log"
	"time"
)

type authUseCase struct {
	userRepository userdomain.UserRepositroy
	userFactory    userdomain.UserFactory
}

func NewAuthUseCase(userRepository userdomain.UserRepositroy, userFactory userdomain.UserFactory) auth.AuthUseCase {

	return &authUseCase{userRepository: userRepository, userFactory: userFactory}
}

// Authorize signup if user is not exist else generate token
func (a *authUseCase) Authorize(request auth.AuthorizeRequest) (string, error) {
	// checking user exist by email
	log.Printf("request to authorize: %+v", request)
	exists, err := a.userRepository.ExistsByEmail(request.Email)
	if err != nil {
		return "", err
	}
	if !exists {
		// create user if not exists
		newUser := a.userFactory.NewUser(request.Email, request.Firstname, request.Lastname)

		// save new user
		_, err := a.userRepository.Save(newUser)
		if err != nil {
			return "", err
		}
	}
	// find user to generate token
	user, err := a.userRepository.FindByEmail(request.Email)
	if err != nil {
		return "", err
	}
	// create user to generate token
	claims := common.UserClaims{
		Id:               user.Id,
		Email:            user.Email,
		Firstname:        user.Firstname,
		Lastname:         user.Lastname,
		RegisteredClaims: jwt.RegisteredClaims{ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(time.Hour)}, IssuedAt: &jwt.NumericDate{Time: time.Now()}},
	}
	token, err := common.NewAccessToken(&claims)
	if err != nil {
		return "", err
	}

	return token, nil
}
