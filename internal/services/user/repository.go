package user

import (
	"context"
	_ "embed"

	"github.com/jackc/pgx/v5/pgxpool"
)

type userRepository struct {
	Connection *pgxpool.Pool
}

func newUserRepository(connection *pgxpool.Pool) userRepository {
	return userRepository{
		Connection: connection,
	}
}

//go:embed sql/CreateUser.sql
var createUserQuery string

func (u *userRepository) Add(ctx context.Context, newUser *user) error {
	_, err := u.Connection.Exec(ctx, createUserQuery, newUser.id, newUser.username, newUser.passwordHash)

	if err != nil {
		return err
	}

	return nil
}

// TODO: reimplement this method
func (u *userRepository) Retrieve(ctx context.Context, id string) (*user, error) {
	var userId string
	var username string
	var passwordHash string

	err := u.Connection.QueryRow(ctx, "SELECT id, username, password FROM users WHERE id = $1", id).Scan(&userId, &username, &passwordHash)

	if err != nil {
		return nil, err
	}

	user := NewUser(username, passwordHash)
	return user, nil
}
