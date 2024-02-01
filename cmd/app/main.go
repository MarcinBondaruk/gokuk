package main

import (
	"net/http"

	"github.com/MarcinBondaruk/gokuk/configs"
	"github.com/MarcinBondaruk/gokuk/internal/application/controller"
	"github.com/MarcinBondaruk/gokuk/internal/application/router"
)

func main() {
	postgres := configs.GetPostgresConnection()
	defer configs.ClosePostgresConnection(postgres)

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
