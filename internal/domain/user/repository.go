package user

type UserRepository interface {
	Add(user User) error
	Retrieve(id string) (User error)
}
