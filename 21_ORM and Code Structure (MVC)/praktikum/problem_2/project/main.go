package main

import (
	"project/config"
	"project/route"
)

func init() {
	config.InitDB()
	config.InitMigrate()
}

func main() {
	e := route.New()

	e.Logger.Fatal(e.Start(":8000"))
}
