package web

type UserRegisterRequest struct {
	Email                string `json:"email"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
	Role                 string `json:"role"`
}
type UserFindByIdRequest struct {
	Id int `uri:"id" json:"id"`
}
type UserLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
