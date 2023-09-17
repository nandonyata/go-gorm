package main

import (
	"go-gorm/controllers"
	"go-gorm/db"
	initializer "go-gorm/initializers"
	"go-gorm/middlewares"

	"github.com/gin-gonic/gin"
)

func init() {
	initializer.LoadEnv()
	db.Connection()
}

func main() {
	r := gin.Default()

	// r.Use(middlewares.Authentication())
	r.POST("/users/login", controllers.Login)
	r.POST("/users", controllers.CreateUser)
	r.GET("/users", controllers.GetUsers)
	r.GET("/users/:id", controllers.GetUserById)
	r.PATCH("/users/:id", controllers.UpdateUser)
	r.DELETE("/users/:id", middlewares.Authentication(), controllers.DeleteUser)

	r.Run()
}