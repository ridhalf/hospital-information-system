package service

import (
	"hospital-information-system/model/domain"
	"hospital-information-system/model/web"
)

type UserService interface {
	Register(request web.UserRegisterRequest) (domain.User, error)
	FindById(request web.UserFindByIdRequest) (web.UserFindByIdResponse, error)
}
