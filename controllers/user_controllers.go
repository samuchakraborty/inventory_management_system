package controllers

import (
	"inventory/database"
	"inventory/model"
	"inventory/utils"
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
// @Failure      400  {object}  utils.HTTPError
// @Failure      404  {object}  utils.HTTPError
// @Failure      500  {object}  utils.HTTPError
// @Router       /v1/getAllUser  [get]
func GetAllUser(c *gin.Context) {

	user := []model.User{}
	if err := database.DB.Find(&user).Error; err != nil {
		log.Printf("Error retrieving users: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": user})
	c.Abort()
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
	user := &model.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, &utils.HTTPError{Code: http.StatusBadRequest,
			Message: "Username, Role is required"})
		return
	}

	errors := validateUser(*user)

	if len(errors) > 0 {
		c.JSON(http.StatusBadRequest, &utils.HTTPError{Code: http.StatusBadRequest, Message: errors})
		return
	}
	res := database.DB.Create(&user)

	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, &utils.HTTPError{Code: http.StatusBadRequest, Message: "Username already exists"})
		return

	} else if res.Error == nil {
		c.JSON(http.StatusOK, gin.H{"user": user})
		return
	} else {

		if strings.Contains(res.Error.Error(), "Duplicate entry") {
			c.JSON(http.StatusConflict, &utils.HTTPError{Code: http.StatusBadRequest, Message: "Username already exists"})
			return

		} else {
			c.JSON(http.StatusInternalServerError,
				&utils.HTTPError{Code: http.StatusInternalServerError, Message: res.Error.Error()})
			return

		}
		c.Abort()
	}
}

func validateUser(user model.User) []string {
	var errors []string

	if user.Username == "" {
		errors = append(errors, "Username is required")
	}

	if user.Role == "" {
		errors = append(errors, "Role is required")
	}

	return errors
}
