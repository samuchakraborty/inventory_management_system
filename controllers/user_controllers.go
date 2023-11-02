package controllers

import (
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
// @Router       /v1/getAllUser  [get]
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

// createUser godoc
// @Summary      Create User
// @Description  Create User
// @Tags         User
// @Accept       json
// @Param		 username	query		string	true	"username "
// @Param		 role	query		string	true	"role "
// @Produce      json
// @Success      200  {object}  []model.User
// @Failure      400  {object} object
// @Failure      404  {object} object
// @Failure      500  {object} object
// @Router       /v1/createUser  [post]
func CreateCustomer(c *gin.Context) {

	username := c.Query("username")
	role := c.Query("role")
	user := &model.User{
		Username: username,
		Role:     role,
	}

	errors := make(map[string]string)

	if username == "" {
		errors["username"] = "Username is required"
	}
	if role == "" {
		errors["role"] = "Role is required"
	}

	if len(errors) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"errors": errors})
		return
	}
	res := database.DB.Create(&user)
	if res.Error == nil {
		c.JSON(http.StatusOK, gin.H{"user": user})
	} else {
		if strings.Contains(res.Error.Error(), "Duplicate entry") {
			c.JSON(http.StatusConflict, gin.H{"error": "Username already exists"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error",	"msg":res.Error.Error()})
		}
		return
	}
}

