package controllers

import (
	"fmt"
	"inventory/database"
	"inventory/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserController(c *gin.Context) {

	user := &models.User{
		Username: "samu",
		Role:     "database_admin",
	}

	database.DB.Create(user)

	database.DB.Find(models.User{})

	c.JSON(http.StatusAccepted, gin.H{"hello": user})

}
func GetAllUser(c *gin.Context) {

	user := &models.User{
		Username: "samu",
		Role:     "database_admin",
	}

	res := database.DB.Create(user)
	if res.RowsAffected != 0 {
		fmt.Println(res)

	}
	fmt.Println(res)

	// database.DB.FindInBatches()

	c.JSON(http.StatusAccepted, gin.H{"hello": user})

}
