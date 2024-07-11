package user

import (
	"context"

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

func (us *UserService) CreateUser(ctx context.Context, cmd CreateUserCommand) error {
	userPasswordHash, _ := bcrypt.GenerateFromPassword([]byte(cmd.Password), 12) // COSTLY

	user := NewUser(cmd.Username, string(userPasswordHash))
	return us.userRepository.Add(ctx, user)
}
