package controller

import (
	"belajar-go-echo/model"
	"belajar-go-echo/repository"

	"github.com/labstack/echo/v4"
)

func getAllUsers(c echo.Context) error {
	users := repository.GetAllUsers()

	return c.JSON(200, echo.Map{
		"data": users,
	})
}

func createUser(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(400, echo.Map{
			"error": err.Error(),
		})
	}
	err := repository.CreateUsers(user)
	if err != nil {
		return c.JSON(500, echo.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(201, echo.Map{
		"message": "success create user",
		"data":    user,
	})
}
