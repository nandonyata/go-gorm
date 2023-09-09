package controllers

import (
	"go-gorm/db"
	model "go-gorm/models"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	// get data from req.body
	var dataFromBody struct{
		Name string
		Email string
	}

	c.Bind(&dataFromBody)

	newUser := model.User{
		Name: dataFromBody.Name, 
		Email: dataFromBody.Email,
	}

	// create
	result := db.DB.Create(&newUser)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// result
	c.JSON(201, gin.H{
		"message": "Success create new user",
		"user": newUser,
	})
}