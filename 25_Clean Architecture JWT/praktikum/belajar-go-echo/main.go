package main

import (
	"belajar-go-echo/config"
	c "belajar-go-echo/controller"

	"github.com/labstack/echo/v4"
)

func init() {
	config.ConnectDB()
	config.MigrateDB()
}

func main() {
	app := echo.New()
	apiUser := app.Group("/users")
	apiLogin := app.Group("/login")
	c.RegisterUserGroupAPI(apiUser)
	c.RegisterLoginGroupAPI(apiLogin)

	app.Start(":8080")
}
