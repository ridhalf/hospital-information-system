package service

import (
	"hospital-information-system/model/domain"
	"hospital-information-system/model/web"
)

type UserService interface {
	Register(request web.UserRegisterRequest) (domain.User, error)
	FindById(request web.UserFindByIdRequest) (domain.User, error)
	Login(request web.UserLoginRequest) (domain.User, error)
}
