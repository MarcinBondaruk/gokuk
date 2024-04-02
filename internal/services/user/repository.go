package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	Connection *pgxpool.Pool
}

func (u UserRepository) Add(newUser *user) error {
	_, err := u.Connection.Exec(context.Background(), "INSERT INTO users(id, username, password) VALUES($1, $2, $3)", newUser.id.String(), newUser.username, newUser.passwordHash)

	if err != nil {
		return err
	}

	return nil
}

func (u UserRepository) Retrieve(id string) (*user, error) {
	var userId string
	var username string
	var passwordHash string

	err := u.Connection.QueryRow(context.Background(), "SELECT id, username, password FROM users WHERE id = $1", id).Scan(&userId, &username, &passwordHash)

	if err != nil {
		return nil, err
	}

	userIdUUID, err := uuid.Parse(userId)

	if err != nil {
		return nil, err
	}

	user := NewUser(userIdUUID, username, passwordHash)
	return user, nil
}
