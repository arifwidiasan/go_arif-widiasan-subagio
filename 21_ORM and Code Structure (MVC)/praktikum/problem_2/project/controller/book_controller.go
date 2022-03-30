package controller

import (
	"net/http"
	"project/config"
	"project/model"

	"github.com/labstack/echo/v4"
)

func CreateBookController(c echo.Context) error {
	books := model.Book{}
	if err := c.Bind(&books); err != nil {
		return err
	}

	err := config.DB.Save(&books).Error
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"messages": "Success create new book",
		"user":     books,
	})
}

func GetBooksController(c echo.Context) error {
	var books []model.Book
	err := config.DB.Find(&books).Error
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "Success",
		"data":     books,
	})
}

func GetBookController(c echo.Context) error {
	var books model.Book
	err := config.DB.First(&books, c.Param("id")).Error
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "Success",
		"data":     books,
	})
}

func DeleteBookController(c echo.Context) error {
	var books model.Book
	err := config.DB.Unscoped().Delete(&books, c.Param("id")).Error
	//err := config.DB.Delete(&books, c.Param("id")).Error
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "Data is deleted",
	})
}

func UpdateBookController(c echo.Context) error {
	var books model.Book
	if err := c.Bind(&books); err != nil {
		return err
	}

	//DB.Model(&users).Where("id = ?", c.Param("id")).Update("name", "hello")
	err := config.DB.Model(&books).Where("id = ?", c.Param("id")).
		Updates(map[string]interface{}{
			"title":  books.Title,
			"author": books.Author,
			"pages":  books.Pages,
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
