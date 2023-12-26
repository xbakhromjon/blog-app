package handler

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/markbates/goth/gothic"
	"golang-project-template/internal/delivery/rest"
	"golang-project-template/internal/domain/auth"
	"log"
	"net/http"
)

type googleOauthHandler struct {
	authUseCase auth.AuthUseCase
}

func NewGoogleOauthHandler(authUseCase auth.AuthUseCase) *googleOauthHandler {

	return &googleOauthHandler{authUseCase: authUseCase}
}

func (g *googleOauthHandler) OauthLogin(write http.ResponseWriter, request *http.Request) {
	provider := chi.URLParam(request, "provider")
	request = request.WithContext(context.WithValue(context.Background(), "provider", provider))
	if gothUser, err := gothic.CompleteUserAuth(write, request); err == nil {
		log.Printf("user: %+v", gothUser)
	} else {
		gothic.BeginAuthHandler(write, request)
	}
}

func (g *googleOauthHandler) CallbackHandler(write http.ResponseWriter, request *http.Request) {
	provider := chi.URLParam(request, "provider")
	request = request.WithContext(context.WithValue(context.Background(), "provider", provider))
	user, err := gothic.CompleteUserAuth(write, request)
	if err != nil {
		fmt.Fprintln(write, err)
		return
	}
	log.Printf("%+v", user)
	signupRequest := auth.SignUpRequest{Email: user.Email, Firstname: user.FirstName, Lastname: user.LastName}
	token, err := g.authUseCase.Signup(signupRequest)
	if err != nil {
		rest.HandleError(write, request, err)
		return
	}
	render.JSON(write, request, token)
}
