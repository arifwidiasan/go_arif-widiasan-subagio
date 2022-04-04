package controller

import (
	"net/http"
	"project/model"
	"project/repository"
	"strconv"

	"github.com/labstack/echo/v4"
)

func createBookController(c echo.Context) error {
	books := model.Book{}
	if err := c.Bind(&books); err != nil {
		return err
	}

	err := repository.CreateBooks(books)
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

func getBooksController(c echo.Context) error {
	books := repository.GetAllBooks()

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "Success",
		"data":     books,
	})
}

func getBookController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	res, err := repository.GetOneByIdBook(id)
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

func deleteBookController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	err := repository.DeleteByIdBook(id)
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

func updateBookController(c echo.Context) error {
	var books model.Book
	id, _ := strconv.Atoi(c.Param("id"))
	if err := c.Bind(&books); err != nil {
		return err
	}

	//DB.Model(&users).Where("id = ?", c.Param("id")).Update("name", "hello")
	err := repository.UpdateOneByIdBook(id, books)

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
