package user

import (
	"log"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type CreateUserCommand struct {
	Username string
	Password string
}

type UserService struct {
	userRepository userRepository
}

func NewUserService(db *pgxpool.Pool) *UserService {
	return &UserService{
		userRepository: newUserRepository(db),
	}
}

func (us *UserService) CreateUser(cmd CreateUserCommand) error {
	userID, err := uuid.NewV7()
	userPasswordHash, _ := bcrypt.GenerateFromPassword([]byte(cmd.Password), 12) // COSTLY

	if err != nil {
		// TODO: Add logger
		log.Println("failed to create id: ", err)
	}
	user := NewUser(userID, cmd.Username, string(userPasswordHash))
	return us.userRepository.Add(user)
}
