package main

import (
	"net/http"

	"github.com/MarcinBondaruk/gokuk/configs"
)

func main() {
	// check envs
	// boonstrap repositories, services
	// register routes
	postgres := configs.GetPostgresConnection()
	defer configs.ClosePostgresConnection(postgres)

	router := router.NewRouter()

	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
