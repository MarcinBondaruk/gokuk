package user

import "github.com/google/uuid"

type user struct {
	id           uuid.UUID
	username     string
	passwordHash string
}

func NewUser(id uuid.UUID, username, passwordHash string) *user {
	return &user{
		id:           id,
		username:     username,
		passwordHash: passwordHash,
	}
}
