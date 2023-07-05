package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type EstaController interface {
	Generate(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
