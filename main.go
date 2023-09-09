package main

import (
	"go-gorm/controllers"
	"go-gorm/db"
	initializer "go-gorm/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializer.LoadEnv()
	db.Connection()
}

func main() {
	r := gin.Default()

	r.POST("/users", controllers.CreateUser)
	r.GET("/users", controllers.GetUsers)

	r.Run()
}