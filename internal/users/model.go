package users

type Users struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
