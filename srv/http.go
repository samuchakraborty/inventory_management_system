package srv

import "github.com/gin-gonic/gin"

// start the server
func Start(address string) error {
	return setUpRouter().Run(address)

}

func setUpRouter() *gin.Engine {

	r := gin.Default()

	return r
}
