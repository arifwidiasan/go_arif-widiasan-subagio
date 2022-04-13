package domain

import "belajar-go-echo/model"

type AdapterRepository interface {
	CreateUsers(user model.User) error
	GetAll() []model.User
	GetOneByEmail(email string) (user model.User, err error)
}

type AdapterService interface {
	CreateUserService(user model.User) error
	GetAllUsersService() []model.User
	LoginUser(email, password string) (string, int)
}
