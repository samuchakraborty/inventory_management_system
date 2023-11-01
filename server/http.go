package server

import (
	"inventory/controllers"
	"inventory/database"

	"github.com/gin-gonic/gin"
)

// start the server
func Start(ip, port string) error {
	return setUpRouter().Run(ip + ":" + port)
}

// set up server config
func setUpRouter() *gin.Engine {
	database.Database()

	router := gin.New()

	v1 := router.Group("v1")
	{
		userGroup := v1.Group("")
		{
			// user := new(controllers.UserController)
			userGroup.GET("/samu", controllers.UserController)

			userGroup.GET("/da", controllers.GetAllUser)

		}
	}

	return router
}
