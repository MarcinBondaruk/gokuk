package http_handler

import (
	"encoding/json"
	"net/http"

	"github.com/MarcinBondaruk/gokuk/internal/services/user"
	"github.com/MarcinBondaruk/gokuk/internal/user_interface/http_handler/request"
)

type UserHandler struct {
	userService user.UserService
}

func NewUserHandler(
	userService user.UserService,
) UserHandler {
	return UserHandler{
		userService: userService,
	}
}

func (u *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var request request.CreateUserRequest

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !request.IsValid() {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	cmd := user.CreateUserCommand{
		Username: request.Username,
		Password: request.Password,
	}

	err = u.userService.CreateUser(cmd)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("User created"))
}
