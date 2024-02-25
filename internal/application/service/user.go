package service

import "github.com/MarcinBondaruk/gokuk/internal/domain/user"

type UserService struct {
	userRepository *user.UserRepository
}
