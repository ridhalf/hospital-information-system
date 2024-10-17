package repository

import "hospital-information-system/model/domain"

type UserRepository interface {
	FindByUsername(username string) (domain.User, error)
	FindById(id int) (domain.User, error)
	Save(user domain.User) (domain.User, error)
	Update(user domain.User) (domain.User, error)
}
