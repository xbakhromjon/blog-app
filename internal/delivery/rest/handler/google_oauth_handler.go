package handler

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/markbates/goth/gothic"
	"golang-project-template/internal/domain/auth"
	"log"
	"net/http"
)

type GoogleOauthHandler struct {
	authUseCase auth.AuthUseCase
}

func NewGoogleOauthHandler(authUseCase auth.AuthUseCase) *GoogleOauthHandler {

	return &GoogleOauthHandler{authUseCase: authUseCase}
}

func (g *GoogleOauthHandler) OauthLogin(write http.ResponseWriter, request *http.Request) {
	provider := chi.URLParam(request, "provider")
	request = request.WithContext(context.WithValue(context.Background(), "provider", provider))
	if gothUser, err := gothic.CompleteUserAuth(write, request); err == nil {
		log.Printf("user: %+v", gothUser)
	} else {
		gothic.BeginAuthHandler(write, request)
	}
}

func (g *GoogleOauthHandler) CallbackHandler(write http.ResponseWriter, request *http.Request) {
	log.Println("authusecase: ", g.authUseCase)
	provider := chi.URLParam(request, "provider")
	request = request.WithContext(context.WithValue(context.Background(), "provider", provider))

	user, err := gothic.CompleteUserAuth(write, request)
	if err != nil {
		fmt.Fprintln(write, err)
		return
	}
	log.Printf("%+v", user)
	log.Println(":OK")
	authorizeRequest := auth.AuthorizeRequest{Email: user.Email, Firstname: user.FirstName, Lastname: user.LastName}
	token, err := g.authUseCase.Authorize(authorizeRequest)

	if err != nil {
		HandleError(write, request, err)
		return
	}
	render.JSON(write, request, token)
}
