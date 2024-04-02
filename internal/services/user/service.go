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

type usersService struct {
	userRepo UserRepository
}

func NewUsersService(userRepo UserRepository) *usersService {
	return &usersService{
		userRepo: userRepo,
	}
}

func (us *usersService) CreateUser(cmd CreateUserCommand) error {
	userID, err := uuid.NewV7()
	userPasswordHash, _ := bcrypt.GenerateFromPassword([]byte(cmd.Password), 12)

	if err != nil {
		// TODO: Add logger
		log.Println("failed to create id: ", err)
	}
	user := NewUser(userID, cmd.Username, string(userPasswordHash))
	return us.userRepo.Add(user)
}
