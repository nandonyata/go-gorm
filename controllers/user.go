package controllers

import (
	"go-gorm/db"
	"go-gorm/helpers"
	model "go-gorm/models"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	// get data from req.body
	var dataFromBody struct{
		Name string
		Email string
		Password string
	}

	c.Bind(&dataFromBody)

	newUser := model.User{
		Name: dataFromBody.Name, 
		Email: dataFromBody.Email,
		Password: helpers.HashPassword([]byte(dataFromBody.Password)),
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

func GetUsers(c *gin.Context) {
	var users []model.User
	result := db.DB.Find(&users)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"code": 200,
		"user": users,
	})
}

func GetUserById(c *gin.Context) {
	id := c.Param("id")
	var user model.User
	
	db.DB.First(&user, id)

	if user.ID == 0 {
		c.JSON(404, gin.H{
			"code": 404,
			"message": "user not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"code": 200,
		"user": user,
	})
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var dataFromBody struct{
		Name string
		Email string
		Password string
	}
	c.Bind(&dataFromBody)

	// find data by id
	var user model.User
	db.DB.First(&user, id)

	// err 404 if not found
	if user.ID == 0 {
		c.JSON(404, gin.H{
			"code": 404,
			"message": "user not found",
		})
		return
	}

	// update
	db.DB.Model(&user).Updates(model.User{
		Name: dataFromBody.Name,
		Email: dataFromBody.Email,
		Password: helpers.HashPassword([]byte(dataFromBody.Password)),
	})
	
	c.JSON(200, gin.H{
		"code": 200,
		"user": user,
	})
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var user model.User
	
	db.DB.First(&user, id)

	if user.ID == 0 {
		c.JSON(404, gin.H{
			"code": 404,
			"message": "user not found",
		})
		return
	}

	db.DB.Delete(&user, id)

	c.JSON(200, gin.H{
		"code": 200,
		"message": "Success delete user",
	})
}

func Login(c *gin.Context) {
	var dataFromBody struct{
		Email string
		Password string
	}
	c.Bind(&dataFromBody)

	// find user
	var user model.User
	db.DB.First(&user, "email = ?", dataFromBody.Email)

	// err if not found
	if user.ID == 0 {
		c.JSON(404, gin.H{
			"code": 404,
			"message": "Incorrect email/password",
		})
		return
	}

	// check if password correct
	compare := helpers.ComparePassword([]byte(user.Password), []byte(dataFromBody.Password))

	if !compare {
		c.JSON(404, gin.H{
			"code": 404,
			"message": "Incorrect email/password",
		})
		return
	}

	// sign jwt token
	newToken := helpers.SignToken(user.ID)

	// response
	response := model.LoginResponse{
		User: user,
		AccessToken: newToken,
	}

	c.JSON(200, gin.H{
		"code": 200,
		"response": response,
	})
}