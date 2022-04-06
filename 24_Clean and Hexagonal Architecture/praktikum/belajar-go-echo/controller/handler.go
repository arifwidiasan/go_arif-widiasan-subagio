package controller

import "github.com/labstack/echo/v4"

func RegisterUserGroupAPI(e *echo.Group) {
	e.GET("", getAllUsers)
	e.POST("", createUser)
}
