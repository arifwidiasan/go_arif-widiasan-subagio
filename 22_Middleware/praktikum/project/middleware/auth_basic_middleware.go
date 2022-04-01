package middleware

import (
	"project/config"
	"project/model"

	"github.com/labstack/echo/v4"
)

func BasicAuthDB(email, password string, c echo.Context) (bool, error) {
	var users model.User
	err := config.DB.Where("email = ? AND password = ?", email, password).First(&users).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
