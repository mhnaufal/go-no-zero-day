package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type MainController struct{}

func (h *MainController) Default(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "SUCCESS"})
}
