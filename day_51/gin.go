package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Private user data
var users = gin.H{
	"Joker":      gin.H{"email": "joker@asylum.com", "phone": "08123"},
	"The Riddle": gin.H{"email": "riddle@asylum.com", "phone": "08321"},
}

func main() {
	r := gin.Default()

	// BasicAuth => return basic HTTP middleare
	// => take a gin.Accounts, where key is user, value is the password
	// => gin.Accounts form a map[string]string
	authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"Joker":      "mysecretpassword",
		"The Riddle": "pass123",
	}))

	// MustGet => check if the given key is exists
	authorized.GET("/secret", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)
		if secret, ok := users[user]; ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET"})
		}
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "SUCCESS",
		})
	})

	r.Run("127.0.0.1:5000")
}
