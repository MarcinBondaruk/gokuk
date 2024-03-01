package service

import (
	"github.com/MarcinBondaruk/gokuk/internal/domain/user"
	"github.com/google/uuid"
)

type UserService struct {
	UserRepository user.UserRepository
}

func NewUserService(userRepository user.UserRepository) *UserService {
	return &UserService{UserRepository: userRepository}
}

func (s UserService) CreateUser(username, password string) (string, error) {
	userId, err := uuid.NewV7()

	if err != nil {
		return "", err
	}

	newUser := user.NewUser(userId, username, password)

	s.UserRepository.Add(newUser)

	return userId.String(), nil
}

func (s UserService) GetUser(id string) (*user.User, error) {
	user, err := s.UserRepository.Retrieve(id)

	if err != nil {
		return nil, err
	}

	return user, nil
}
