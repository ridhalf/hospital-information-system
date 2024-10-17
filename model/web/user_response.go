package web

import (
	"hospital-information-system/model/domain"
)

type UserRegisterResponse struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Token    string `json:"token"`
}

func ToRegisterResponse(user domain.User, token string) UserRegisterResponse {
	return UserRegisterResponse{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
		Token:    token,
	}
}

type UserFindByIdResponse struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

func ToFindByIdResponse(user domain.User) UserFindByIdResponse {
	return UserFindByIdResponse{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
	}
}
