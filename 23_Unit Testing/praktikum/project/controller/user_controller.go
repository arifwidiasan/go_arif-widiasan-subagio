package controller

import (
	"net/http"
	"strconv"

	"project/config"
	"project/middleware"
	"project/model"
	"project/repository"

	"github.com/labstack/echo/v4"
)

func createUserController(c echo.Context) error {
	users := model.User{}
	if err := c.Bind(&users); err != nil {
		return err
	}

	err := repository.CreateUsers(users)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"messages": "Success create new user",
		"user":     users,
	})
}

func getUsersController(c echo.Context) error {
	users := repository.GetAllUsers()

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "Success",
		"data":     users,
	})
}

func getUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	res, err := repository.GetOneByIdUser(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "id not found",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "Success",
		"data":     res,
	})
}

func deleteUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	err := repository.DeleteByIdUser(id)
	//err := config.DB.Delete(&users, c.Param("id")).Error
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": "id not found",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "Data is deleted",
	})
}

func updateUserController(c echo.Context) error {
	var users model.User
	id, _ := strconv.Atoi(c.Param("id"))
	if err := c.Bind(&users); err != nil {
		return err
	}

	//DB.Model(&users).Where("id = ?", c.Param("id")).Update("name", "hello")
	err := repository.UpdateOneByIdUser(id, users)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": "update error",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "Data is updated",
		"id user":  c.Param("id"),
	})
}

func LoginUserController(c echo.Context) error {
	users := model.User{}
	if err := c.Bind(&users); err != nil {
		return err
	}

	err := config.DB.Where("email = ? AND password = ?", users.Email, users.Password).First(&users).Error
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": "Fail Login",
			"error":    err.Error(),
		})
	}

	token, err := middleware.CreateToken(int(users.ID), users.Name)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": "Fail Login",
			"error":    err.Error(),
		})
	}

	userResponse := model.UserResponse{int(users.ID), users.Name, users.Email, token}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"messages": "Success create new user",
		"user":     userResponse,
	})
}
