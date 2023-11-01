package controllers

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"samu.com/inventory_management_system/models"
)

var (
	// DB *gorm.DB
	DB *sql.DB
)

func UserController(c *gin.Context) {

	// user, err := models.Users(qm.Load(models.UserRels)).One(context.Background(), server.DB)
	// if err != nil {
	// 	// Handle the error
	// }
	count, err := models.Users().Count(context.Background(), DB)
	if err != nil {
		// Handle the error
	}
	// user, err := models.FindUser(context.Background(), server.DB)

	c.JSON(http.StatusAccepted, gin.H{"hello": count})

}
