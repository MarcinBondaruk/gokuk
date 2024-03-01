package user

import "github.com/google/uuid"

type User struct {
	id       uuid.UUID
	username string
	password string
}

func NewUser(id uuid.UUID, username, password string) *User {
	return &User{
		id:       id,
		username: username,
		password: password,
	}
}

func (u *User) Id() uuid.UUID {
	return u.id
}

func (u *User) Username() string {
	return u.username
}

func (u *User) Password() string {
	return u.password
}

func (u *User) Authenticate(secret string) bool {
	return secret == u.password
}
