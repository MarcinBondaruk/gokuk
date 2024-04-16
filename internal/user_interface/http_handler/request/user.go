package request

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r *CreateUserRequest) IsValid() bool {
	if r.Username == "" {
		return false
	}

	if r.Password == "" {
		return false
	}

	return true
}
