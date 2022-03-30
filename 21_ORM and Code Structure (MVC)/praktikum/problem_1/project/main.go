package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var DB *gorm.DB

type Config struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}

func InitDB() {
	config := Config{
		DB_Username: "root",
		DB_Password: "",
		DB_Port:     "3306",
		DB_Host:     "localhost",
		DB_Name:     "crud_go",
	}
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Name,
	)
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
}

func InitMigrate() {
	DB.AutoMigrate(&User{})
}

func init() {
	InitDB()
	InitMigrate()
}

func main() {
	e := echo.New()
	// routing with query parameter
	e.GET("/users", GetUsersController)
	e.GET("/users/:id", GetUserController)
	e.POST("/users", CreateUserController)
	e.DELETE("/users/:id", DeleteUserController)
	e.PUT("/users/:id", UpdateUserController)

	e.Logger.Fatal(e.Start(":8000"))
}

func CreateUserController(c echo.Context) error {
	users := User{}
	if err := c.Bind(&users); err != nil {
		return err
	}

	err := DB.Save(&users).Error
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
	var users []User
	err := DB.Find(&users).Error
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
	var users User
	err := DB.First(&users, c.Param("id")).Error
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
	var users User
	err := DB.Unscoped().Delete(&users, c.Param("id")).Error
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
	var users User
	if err := c.Bind(&users); err != nil {
		return err
	}

	//DB.Model(&users).Where("id = ?", c.Param("id")).Update("name", "hello")
	err := DB.Model(&users).Where("id = ?", c.Param("id")).
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
