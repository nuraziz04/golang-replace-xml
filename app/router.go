package app

import (
	"github.com/julienschmidt/httprouter"
	"github.com/nuraziz04/golang-esta/controller"
	"github.com/nuraziz04/golang-esta/exception"
)

func NewRouter() *httprouter.Router {
	router := httprouter.New()

	estaController := controller.NewEstaController()

	router.POST("/api/generate", estaController.Generate)
	router.PanicHandler = exception.ErrorHandler

	return router
}
