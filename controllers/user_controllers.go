package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserController(c *gin.Context) {

	c.JSON(http.StatusAccepted, gin.H{"hello": "samu"})

}
