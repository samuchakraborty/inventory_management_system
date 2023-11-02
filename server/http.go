package server

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:9000
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

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

	router.Use(database.DatabaseMiddleware())
	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:9000/v1"
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
