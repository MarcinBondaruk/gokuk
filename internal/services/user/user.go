package user

type user struct {
	id           int64
	username     string
	passwordHash string
}

func NewUser(username, passwordHash string) *user {
	return &user{
		username:     username,
		passwordHash: passwordHash,
	}
}
