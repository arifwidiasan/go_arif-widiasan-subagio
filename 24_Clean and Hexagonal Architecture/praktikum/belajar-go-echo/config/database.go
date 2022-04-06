package config

import (
	"belajar-go-echo/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	MigrateDB()
}

func MigrateDB() {
	DB.Migrator().DropTable(&model.User{})
	DB.AutoMigrate(&model.User{})
}
