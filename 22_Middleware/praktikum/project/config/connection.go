package config

import (
	"fmt"

	"project/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	conf := initConfiguration()
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.DB_Username,
		conf.DB_Password,
		conf.DB_Host,
		conf.DB_Port,
		conf.DB_Name,
	)
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
}

func InitMigrate() {
	DB.AutoMigrate(&model.User{})
	DB.AutoMigrate(&model.Book{})
}
