package main

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Logging Gin to a file
	f, _ := os.Create("day_52/gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	// Use this to shows up logs in console and file
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "SUCCESS")
	})

	router.Run("127.0.0.1:5000")
}
