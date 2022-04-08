package controller

import (
	"belajar-go-echo/model"
	"belajar-go-echo/repository"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
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

func login(c echo.Context) error {
	request := model.LoginRequest{}
	if err := c.Bind(&request); err != nil {
		return c.JSON(400, echo.Map{
			"error": err.Error(),
		})
	}
	if request.Username != "admin" && request.Password != "admin" {
		return c.JSON(http.StatusUnauthorized, "username dan password salah")
	}

	claims := jwt.MapClaims{}
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwt, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "login success",
		"data":    jwt,
	})
}
