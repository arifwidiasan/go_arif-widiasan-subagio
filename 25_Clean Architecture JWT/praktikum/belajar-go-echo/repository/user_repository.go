package repository

import (
	"belajar-go-echo/config"
	"belajar-go-echo/model"
	"fmt"
)

func CreateUsers(user model.User) error {
	res := config.DB.Create(&user)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert")
	}

	return nil
}

func GetAllUsers() []model.User {
	users := []model.User{}
	config.DB.Find(&users)

	return users
}
