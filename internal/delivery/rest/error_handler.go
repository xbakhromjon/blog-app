package rest

import (
	"github.com/go-chi/render"
	"net/http"
)

func HandleError(write http.ResponseWriter, request *http.Request, err error) {
	render.JSON(write, request, err.Error())
}
