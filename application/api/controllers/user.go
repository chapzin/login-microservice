package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func TestUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "Ola mundo usuario",
	})
	return
}
