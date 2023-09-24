package main

import (
	"go-gorm/db"
	initializer "go-gorm/initializers"
	model "go-gorm/models"
)

func init() {
	initializer.LoadEnv()
	db.Connection()
}

func main() {
	db.DB.AutoMigrate(&model.User{}, &model.Movie{})
}