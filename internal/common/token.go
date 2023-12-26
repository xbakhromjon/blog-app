package common

import (
	"github.com/golang-jwt/jwt/v5"
	"log"
	"os"
	"strings"
)

type UserClaims struct {
	Id        int64  `json:"id"`
	Email     string `json:"email"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	jwt.RegisteredClaims
}

func NewAccessToken(claims *UserClaims) (string, error) {
	log.Printf("%+v", claims)
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	log.Println("..............")
	return accessToken.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
}

func ParseAccessToken(accessToken string) (*UserClaims, error) {
	parsedAccessToken, err := jwt.ParseWithClaims(accessToken, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("TOKEN_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return parsedAccessToken.Claims.(*UserClaims), nil
}

func ValidateBearerToken(token string) bool {
	return token != "" && len(token) > 8 && strings.HasPrefix(token, "Bearer")
}
