package router

import (
	"database/sql"
	"github.com/go-chi/chi/v5"
	"golang-project-template/internal/delivery/rest/handler"
	"golang-project-template/internal/domain/user"
	"golang-project-template/internal/repository"
	"golang-project-template/internal/usecase"
)

func GoogleOauthRouter(r chi.Router, db *sql.DB) {
	authUseCase := usecase.NewAuthUseCase(repository.NewUserRepository(db), user.UserFactory{})
	googleOauthHandler := handler.NewGoogleOauthHandler(authUseCase)
	r.Route("/auth", func(r chi.Router) {
		r.Get("/{provider}", googleOauthHandler.OauthLogin)
		r.Get("/{provider}/callback", googleOauthHandler.CallbackHandler)
	})
}
