package middleware

import (
	"golang-project-template/internal/common"
	"log"
	"net/http"
)

func SecureEndpoints(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		// validate bearer token
		if common.ValidateBearerToken(token) {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("StatusUnauthorized"))
			return
		}

		// remove Bearer from token
		token = token[7:]

		// extract claims from token
		userClaims, err := common.ParseAccessToken(token)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Invalid token"))
			return
		}

		// check user by userClaims.Id
		log.Printf("userID: %s", userClaims.Id)

		next.ServeHTTP(w, r)
	})
}
