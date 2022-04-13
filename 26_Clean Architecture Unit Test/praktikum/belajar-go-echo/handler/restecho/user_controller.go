package restecho

import (
	"belajar-go-echo/domain"
	"belajar-go-echo/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

type EchoController struct {
	svc domain.AdapterService
}

func (ce *EchoController) CreateUserController(c echo.Context) error {
	user := model.User{}
	c.Bind(&user)

	err := ce.svc.CreateUserService(user)
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	return c.JSON(201, map[string]interface{}{
		"messages": "success",
		"users":    user,
	})
}

func (ce *EchoController) GetUsersController(c echo.Context) error {
	users := ce.svc.GetAllUsersService()

	return c.JSON(200, map[string]interface{}{
		"messages": "success",
		"users":    users,
	})
}

func (ce *EchoController) LoginUserController(c echo.Context) error {
	userLogin := model.LoginRequest{}

	c.Bind(&userLogin)

	token, statusCode := ce.svc.LoginUser(userLogin.Email, userLogin.Password)
	switch statusCode {
	case http.StatusUnauthorized:
		return c.JSONPretty(http.StatusUnauthorized, map[string]interface{}{
			"messages": "email atau password salah",
		}, "  ")

	case http.StatusInternalServerError:
		return c.JSONPretty(http.StatusInternalServerError, map[string]interface{}{
			"messages": "internal",
		}, "  ")
	}

	return c.JSONPretty(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"token":    token,
	}, "  ")
}
