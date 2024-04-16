package user

import (
	"log"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type CreateUserCommand struct {
	Username string
	Password string
}

type UserService struct {
	userRepo UserRepository
}

func NewUserService(userRepo UserRepository) UserService {
	return UserService{
		userRepo: userRepo,
	}
}

func (us *UserService) CreateUser(cmd CreateUserCommand) error {
	userID, err := uuid.NewV7()
	userPasswordHash, _ := bcrypt.GenerateFromPassword([]byte(cmd.Password), 12)

	if err != nil {
		// TODO: Add logger
		log.Println("failed to create id: ", err)
	}
	user := NewUser(userID, cmd.Username, string(userPasswordHash))
	return us.userRepo.Add(user)
}
