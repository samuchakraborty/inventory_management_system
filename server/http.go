package server

import (
	"github.com/gin-gonic/gin"
	"samu.com/inventory_management_system/controllers"
)

// start the server
func Start(ip, port string) error {
	return setUpRouter().Run(ip + ":" + port)
}

// set up server config
func setUpRouter() *gin.Engine {
	router := gin.New()
	
	// router.Use(gin.Logger())

	v1 := router.Group("v1")
	{
		userGroup := v1.Group("")
		{
			// user := new(controllers.UserController)
			userGroup.GET("/samu", controllers.UserController)

			userGroup.GET("/da", controllers.UserController)

		}
	}

	return router
}
