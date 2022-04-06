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
	c.RegisterUserGroupAPI(apiUser)

	app.Start(":8080")
}
