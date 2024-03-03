package repository

import (
	"context"

	"github.com/MarcinBondaruk/gokuk/internal/domain/user"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepositoryImpl struct {
	Connection *pgxpool.Pool
}

func (u UserRepositoryImpl) Add(newUser *user.User) error {
	_, err := u.Connection.Exec(context.Background(), "INSERT INTO users(id, username, password) VALUES($1, $2, $3)", newUser.Id(), newUser.Username(), newUser.Password())

	if err != nil {
		return err
	}

	return nil
}

func (u UserRepositoryImpl) Retrieve(id string) (*user.User, error) {
	var userId string
	var username string
	var password string

	err := u.Connection.QueryRow(context.Background(), "SELECT id, username, password FROM users WHERE id = $1", id).Scan(&userId, &username, &password)

	if err != nil {
		return nil, err
	}

	userIdUUID, err := uuid.Parse(userId)

	if err != nil {
		return nil, err
	}

	user := user.NewUser(userIdUUID, username, password)
	return user, nil
}
