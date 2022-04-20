package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "SUCCESS",
		})
	})

	r.Run("127.0.0.1:3000")
}
