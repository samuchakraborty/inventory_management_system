package controllers

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"samu.com/inventory_management_system/database"
	"samu.com/inventory_management_system/models"
)

func UserController(c *gin.Context) {

	count, err := models.Users().Count(context.Background(), database.DB)
	if err != nil {
		log.Fatalln(err.Error(), "sam8")
	}

	c.JSON(http.StatusAccepted, gin.H{"hello": count})

}
