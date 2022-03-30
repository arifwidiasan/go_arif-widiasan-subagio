package controller

import (
	"net/http"

	"project/config"
	"project/model"

	"github.com/labstack/echo/v4"
)

func CreateUserController(c echo.Context) error {
	users := model.User{}
	if err := c.Bind(&users); err != nil {
		return err
	}

	err := config.DB.Save(&users).Error
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

func GetUsersController(c echo.Context) error {
	var users []model.User
	err := config.DB.Find(&users).Error
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "Success",
		"data":     users,
	})
}

func GetUserController(c echo.Context) error {
	var users model.User
	err := config.DB.First(&users, c.Param("id")).Error
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "Success",
		"data":     users,
	})
}

func DeleteUserController(c echo.Context) error {
	var users model.User
	err := config.DB.Unscoped().Delete(&users, c.Param("id")).Error
	//err := config.DB.Delete(&users, c.Param("id")).Error
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "Data is deleted",
	})
}

func UpdateUserController(c echo.Context) error {
	var users model.User
	if err := c.Bind(&users); err != nil {
		return err
	}

	//DB.Model(&users).Where("id = ?", c.Param("id")).Update("name", "hello")
	err := config.DB.Model(&users).Where("id = ?", c.Param("id")).
		Updates(map[string]interface{}{
			"name":     users.Name,
			"email":    users.Email,
			"password": users.Password,
		}).Error

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "Data is updated",
		"id user":  c.Param("id"),
	})
}
