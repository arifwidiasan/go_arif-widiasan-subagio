package repository

import (
	"fmt"
	conn "project/config"
	"project/model"
)

func CreateUsers(user model.User) error {
	res := conn.DB.Create(&user)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert")
	}

	return nil
}

func GetAllUsers() []model.User {
	users := []model.User{}
	conn.DB.Find(&users)

	return users
}

func GetOneByIdUser(id int) (user model.User, err error) {
	res := conn.DB.Where("id = ?", id).Find(&user)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return
}

func DeleteByIdUser(id int) error {
	var users model.User
	res := conn.DB.Unscoped().Delete(&users, id)

	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete")
	}

	return nil
}

func UpdateOneByIdUser(id int, user model.User) error {
	res := conn.DB.Where("id = ?", id).UpdateColumns(&user)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error update")
	}

	return nil
}
