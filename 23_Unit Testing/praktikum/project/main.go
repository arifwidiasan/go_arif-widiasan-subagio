package main

import (
	"project/config"
	//"project/constants"
	c "project/controller"
	m "project/middleware"

	"github.com/labstack/echo/v4"
	//mid "github.com/labstack/echo/v4/middleware"
)

func init() {
	config.InitDB()
	config.InitMigrate()
}

func main() {
	e := echo.New()
	apiUser := e.Group("/users")
	apiBook := e.Group("/books")
	e.POST("/login", c.LoginUserController)

	c.RegisterUserGroupAPI(apiUser)

	c.RegisterBookGroupAPI(apiBook)

	m.LogMiddleware(e)

	e.Logger.Fatal(e.Start(":8000"))
}
