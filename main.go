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
	r.GET("/user/:id", controllers.GetUserById)
	r.PATCH("/user/:id", controllers.UpdateUser)
	r.DELETE("/user/:id", controllers.DeleteUser)

	r.Run()
}