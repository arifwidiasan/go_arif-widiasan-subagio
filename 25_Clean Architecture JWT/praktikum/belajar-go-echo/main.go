package main

import (
	conf "belajar-go-echo/config"
	rest "belajar-go-echo/handler/restecho"

	"github.com/labstack/echo/v4"
)

func main() {
	config := conf.InitConfiguration()
	e := echo.New()

	rest.RegisterUserGroupAPI(e, config)

	e.Logger.Fatal(e.Start(":8080"))
}
