package main

import (
	"net/http"

	"github.com/MarcinBondaruk/gokuk/infrastructure/controller"
	"github.com/MarcinBondaruk/gokuk/infrastructure/router"
)

func main() {
	controller := controller.NewController()
	router := router.NewRouter(controller)

	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
