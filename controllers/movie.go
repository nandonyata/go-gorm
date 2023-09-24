package controllers

import (
	"go-gorm/db"
	model "go-gorm/models"

	"github.com/gin-gonic/gin"
)

func AddMovie(c *gin.Context) {
	var dataFromBody struct {
		Title string
		UserID float64
	}

	c.Bind(&dataFromBody)

	newMovie := model.Movie {
		Title: dataFromBody.Title,
		UserID: uint(dataFromBody.UserID),
	}

	db.DB.Create(&newMovie)

	c.JSON(201, gin.H{
		"message": "Success create new movie",
		"movie": newMovie,
	})
}

func GetMovies(c *gin.Context) {
	userId := c.Param("userId")

	var user model.User

	db.DB.Preload("Movies").First(&user, userId)

	if user.ID == 0 {
		c.AbortWithStatusJSON(404, gin.H{
			"code": 404,
			"message": "Data Not Found",
		})
		return
	}

	c.JSON(200, gin.H{
		"user": user,
	})
}