package user

import "github.com/google/uuid"

type User struct {
	id       uuid.UUID
	username string
	password string
}

func (u *User) Id() uuid.UUID {
	return u.id
}

func (u *User) Username() string {
	return u.username
}

func (u *User) Authenticate(secret string) bool {
	return secret == u.password
}
