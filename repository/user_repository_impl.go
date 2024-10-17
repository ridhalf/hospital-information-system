package repository

import (
	"gorm.io/gorm"
	"hospital-information-system/model/domain"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{db}
}

func (repository UserRepositoryImpl) FindByUsername(username string) (domain.User, error) {
	user := domain.User{}
	err := repository.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (repository UserRepositoryImpl) FindById(id int) (domain.User, error) {
	user := domain.User{}
	err := repository.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (repository UserRepositoryImpl) Save(user domain.User) (domain.User, error) {
	err := repository.db.Save(&user).Error
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (repository UserRepositoryImpl) Update(user domain.User) (domain.User, error) {
	err := repository.db.Save(&user).Error
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}
