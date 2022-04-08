package controller

import (
	m "belajar-go-echo/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterUserGroupAPI(e *echo.Group) {
	e.GET("", getAllUsers, m.JWTMiddleware())
	e.POST("", createUser, m.JWTMiddleware())
}

func RegisterLoginGroupAPI(e *echo.Group) {
	e.POST("", login)
}
