package main

import (
	"net/http"

	"github.com/nuraziz04/golang-esta/app"
)

func main() {
	router := app.NewRouter()

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}

	server.ListenAndServe()
}
