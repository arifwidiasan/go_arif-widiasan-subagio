package service

import (
	"net/http"

	"belajar-go-echo/config"
	"belajar-go-echo/domain"
	"belajar-go-echo/helper"
	"belajar-go-echo/model"
)

type svcUser struct {
	c    config.Config
	repo domain.AdapterRepository
}

func (s *svcUser) CreateUserService(user model.User) error {
	return s.repo.CreateUsers(user)
}

func (s *svcUser) GetAllUsersService() []model.User {
	return s.repo.GetAll()
}

func (s *svcUser) LoginUser(email, password string) (string, int) {
	user, _ := s.repo.GetOneByEmail(email)

	if (user.Password != password) || (user == model.User{}) {
		return "", http.StatusUnauthorized
	}

	token, err := helper.CreateToken(int(user.ID), user.Email, s.c.JWT_KEY)
	if err != nil {
		return "", http.StatusInternalServerError
	}

	return token, http.StatusOK
}

func NewServiceUser(repo domain.AdapterRepository, c config.Config) domain.AdapterService {
	return &svcUser{
		repo: repo,
		c:    c,
	}
}
