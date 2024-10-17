package service

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"hospital-information-system/model/domain"
	"hospital-information-system/model/web"
	"hospital-information-system/repository"
	"strings"
)

type UserServiceImpl struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &UserServiceImpl{
		userRepository: userRepository,
	}
}

func (service UserServiceImpl) Register(request web.UserRegisterRequest) (domain.User, error) {
	if request.Password != request.PasswordConfirmation {
		return domain.User{}, errors.New("password incorrect")
	}
	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.MinCost)
	if err != nil {
		return domain.User{}, err
	}
	user := domain.User{
		Username: strings.Split(request.Email, "@")[0],
		Password: string(password),
		Email:    request.Email,
		Role:     request.Role,
	}
	save, err := service.userRepository.Save(user)
	if err != nil {
		return domain.User{}, err
	}
	return save, nil
}
func (service UserServiceImpl) Login(request web.UserLoginRequest) (domain.User, error) {
	user, err := service.userRepository.FindByUsername(request.Username)
	if err != nil {
		return domain.User{}, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (service UserServiceImpl) FindById(request web.UserFindByIdRequest) (domain.User, error) {
	user, err := service.userRepository.FindById(request.Id)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}
