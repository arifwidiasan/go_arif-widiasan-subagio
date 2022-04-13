package database

import (
	"belajar-go-echo/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//var DB *gorm.DB

func ConnectDB() *gorm.DB {
	//var err error
	DB, err := gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	DB.AutoMigrate(&model.User{})
	return DB
}

// func MigrateDB() {
// 	//DB.Migrator().DropTable(&model.User{})
// 	DB.AutoMigrate(&model.User{})
// }
