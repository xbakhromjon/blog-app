package router

import (
	"database/sql"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func SetupRouter(r chi.Router, db *sql.DB) {
	r.Get("/ping", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("pong\n"))
	})
	GoogleOauthRouter(r, db)
}
