package controllers

import (
	"fmt"
	"inventory/database"
	"inventory/model"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// GetAllUser godoc
// @Summary      List User
// @Description  get User
// @Tags         User
// @Accept       json
// @Produce      json
// @Success      200  {object}  []model.User
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /getAllUser  [get]
func GetAllUser(c *gin.Context) {

	user := []model.User{}

	// result := database.DB.Table("users").Find(&user)
	// fmt.Println(result.Statement.RowsAffected)
	if err := database.DB.Find(&user).Error; err != nil {
		log.Printf("Error retrieving users: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": user})

}

func CreateCustomer(c *gin.Context) {

	username := c.Query("username")
	role := c.Query("role")

	fmt.Println(username)

	user := &model.User{
		Username: username,
		Role:     role,
	}

	res := database.DB.Create(&user)
	if res.Error == nil {
		c.JSON(http.StatusOK, gin.H{"user": user})
	} else {
		if strings.Contains(res.Error.Error(), "Duplicate entry") {
			c.JSON(http.StatusConflict, gin.H{"error": "Username already exists"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}

}
