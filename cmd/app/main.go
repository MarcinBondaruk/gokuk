package main

import (
	"log/slog"
	"net/http"

	"github.com/MarcinBondaruk/gokuk/config/dependencies"
	"github.com/MarcinBondaruk/gokuk/config/router"
)

func main() {
	slog.Info("initializig app")
	di, err := dependencies.Initialize()
	// TODO: maybe shouldn't panic
	if err != nil {
		panic(err)
	}
	defer di.TearDown()
	slog.Info("dependencies initialized")

	router := router.NewRouter(di)
	slog.Info("routes registered")

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	slog.Info("listening on", "addr", server.Addr)
	err = server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
