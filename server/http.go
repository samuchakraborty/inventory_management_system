package server

import (
	"inventory/controllers"
	"inventory/database"
	"inventory/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// gin-swagger middleware
// swagger embed files
// start the server
func Start(ip, port string) error {
	return setUpRouter().Run(ip + ":" + port)
}

// set up server config
func setUpRouter() *gin.Engine {
	// database.Database()

	router := gin.New()

	router.Use(gin.Logger())

	router.Use(database.DatabaseMiddleware())
	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Inventory Management System"
	docs.SwaggerInfo.Description = "This is a sample server for inventory."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:9000"
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Schemes = []string{"http"}
	v1 := router.Group("v1")
	{
		userGroup := v1.Group("")
		{
			// user := new(controllers.UserController)
			userGroup.GET("/getAllUser", controllers.GetAllUser)
			userGroup.POST("/createUser", controllers.CreateCustomer)

		}
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}
