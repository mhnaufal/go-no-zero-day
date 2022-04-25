package main

import (
	"net/http"
	"no-zero-day/api/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		home := new(controllers.MainController)

		v1.GET("/", home.Default)
	}

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "FAILED"})
	})

	router.Run("127.0.0.1:5000")
}
