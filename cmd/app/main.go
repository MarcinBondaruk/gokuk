package main

import (
	"net/http"

	"github.com/MarcinBondaruk/gokuk/config/dependencies"
	"github.com/MarcinBondaruk/gokuk/config/router"
)

func main() {
	di, err := dependencies.Initialize()
	if err != nil {
		panic(err)
	}
	defer di.TearDown()

	router := router.NewRouter(di)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	err = server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
