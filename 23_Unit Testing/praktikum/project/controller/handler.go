package controller

import (
	"project/constants"

	"github.com/labstack/echo/v4"
	mid "github.com/labstack/echo/v4/middleware"
)

func RegisterUserGroupAPI(e *echo.Group) {
	e.GET("", getUsersController, mid.JWT([]byte(constants.SECRET_JWT)))
	e.GET("/:id", getUserController, mid.JWT([]byte(constants.SECRET_JWT)))
	e.PUT("/:id", updateUserController, mid.JWT([]byte(constants.SECRET_JWT)))
	e.DELETE("/:id", deleteUserController, mid.JWT([]byte(constants.SECRET_JWT)))
	e.POST("", createUserController)
}

func RegisterBookGroupAPI(e *echo.Group) {
	e.GET("", getBooksController)
	e.GET("/:id", getBookController)
	e.PUT("/:id", updateBookController, mid.JWT([]byte(constants.SECRET_JWT)))
	e.DELETE("/:id", deleteBookController, mid.JWT([]byte(constants.SECRET_JWT)))
	e.POST("", createBookController, mid.JWT([]byte(constants.SECRET_JWT)))
}
