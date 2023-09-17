package middlewares

import (
	"go-gorm/db"
	"go-gorm/helpers"
	model "go-gorm/models"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get token from header
		accessToken := c.GetHeader("accessToken")

		if len(accessToken) == 0 {
			c.AbortWithStatusJSON(401, gin.H{
				"code": 401,
				"message": "Invalid Token",
			})
		}

		// check & validate token
		checkToken, err := helpers.VerifyToken(accessToken)

		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{
				"code": 401,
				"message": "Invalid Token",
			})
		}

		// find user
		var user model.User
		db.DB.First(&user, "id = ?", checkToken["id"])

		if user.ID == 0 {
			c.AbortWithStatusJSON(401, gin.H{
				"code": 401,
				"message": "Invalid Token",
			})
		}

		// set header
		c.Set("user", user)
		c.Next()

		/*
		// use this to call the header that we've set
		user, bool:= c.Get("user")
		*/
	}
}