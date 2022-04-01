package main

import (
	"project/config"
	"project/constants"
	c "project/controller"
	m "project/middleware"

	"github.com/labstack/echo/v4"
	mid "github.com/labstack/echo/v4/middleware"
)

func init() {
	config.InitDB()
	config.InitMigrate()
}

func main() {
	e := echo.New()
	//apiUser := e.Group("/users")
	apiUserAuth := e.Group("auth/users")
	apiUserJWT := e.Group("/jwt/users")
	apiUserNonJWT := e.Group("/jwt/users")
	//apiBook := e.Group("/books")
	apiBookJWT := e.Group("/jwt/books")
	apiBookNonJWT := e.Group("/jwt/books")
	e.POST("/login", c.LoginUserController)

	apiUserJWT.Use(mid.JWT([]byte(constants.SECRET_JWT)))
	apiBookJWT.Use(mid.JWT([]byte(constants.SECRET_JWT)))
	apiUserAuth.Use(mid.BasicAuth(m.BasicAuthDB))

	//c.RegisterUserGroupAPI(apiUser)
	c.RegisterUserGroupAPI(apiUserAuth)
	c.RegisterUserGroupAPIJWT(apiUserJWT)
	c.RegisterUserGroupAPINonJWT(apiUserNonJWT)

	//c.RegisterBookGroupAPI(apiBook)
	c.RegisterBookGroupAPIJWT(apiBookJWT)
	c.RegisterBookGroupAPINonJWT(apiBookNonJWT)

	m.LogMiddleware(e)

	e.Logger.Fatal(e.Start(":8000"))
}
