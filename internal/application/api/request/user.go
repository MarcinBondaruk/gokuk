package request

type UserRequest struct {
	username string `json:"username"`
	password string `json:"password"`
}
