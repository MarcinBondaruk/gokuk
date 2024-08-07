package router

import (
	"net/http"

	"github.com/MarcinBondaruk/gokuk/config/middleware"
	"github.com/MarcinBondaruk/gokuk/internal/app_interface/http/user"
)

func RegisterUserRoutes(router *http.ServeMux, handler *user.UserHandler) *http.ServeMux {
	router.HandleFunc("POST /api/v1/users", middleware.Pipe(handler.CreateUser, middleware.Logger()))

	return router
}
