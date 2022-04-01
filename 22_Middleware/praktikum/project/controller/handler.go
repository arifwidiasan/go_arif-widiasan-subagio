package controller

import "github.com/labstack/echo/v4"

func RegisterUserGroupAPI(e *echo.Group) {
	e.GET("", getUsersController)
	e.GET("/:id", getUserController)
	e.PUT("/:id", updateUserController)
	e.DELETE("/:id", deleteUserController)
	e.POST("", createUserController)
}

func RegisterUserGroupAPIJWT(e *echo.Group) {
	e.GET("", getUsersController)
	e.GET("/:id", getUserController)
	e.PUT("/:id", updateUserController)
	e.DELETE("/:id", deleteUserController)
}

func RegisterUserGroupAPINonJWT(e *echo.Group) {
	e.POST("", createUserController)
}

func RegisterBookGroupAPI(e *echo.Group) {
	e.GET("", getBooksController)
	e.GET("/:id", getBookController)
	e.PUT("/:id", updateBookController)
	e.DELETE("/:id", deleteBookController)
	e.POST("", createBookController)
}

func RegisterBookGroupAPIJWT(e *echo.Group) {
	e.PUT("/:id", updateBookController)
	e.DELETE("/:id", deleteBookController)
	e.POST("", createBookController)
}

func RegisterBookGroupAPINonJWT(e *echo.Group) {
	e.GET("", getBooksController)
	e.GET("/:id", getBookController)
}
