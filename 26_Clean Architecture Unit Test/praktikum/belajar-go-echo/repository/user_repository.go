package repository

import (
	"belajar-go-echo/domain"
	"belajar-go-echo/model"
	"fmt"

	"gorm.io/gorm"
)

type repositoryMysqlLayer struct {
	DB *gorm.DB
}

func (r *repositoryMysqlLayer) CreateUsers(user model.User) error {
	res := r.DB.Create(&user)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert")
	}

	return nil
}

func (r *repositoryMysqlLayer) GetAll() []model.User {
	users := []model.User{}
	r.DB.Find(&users)

	return users
}

func (r *repositoryMysqlLayer) GetOneByEmail(email string) (user model.User, err error) {
	res := r.DB.Where("email = ?", email).Find(&user)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return
}

func NewMysqlRepository(db *gorm.DB) domain.AdapterRepository {
	return &repositoryMysqlLayer{
		DB: db,
	}
}
