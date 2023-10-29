package srv

import (
	"fmt"
	"net"
	"os/exec"

	"github.com/gin-gonic/gin"
)

// start the server
func Start(ip, port string) error {
	killPort(ip, port)
	return setUpRouter().Run(port)

}

// set up server config
func setUpRouter() *gin.Engine {

	r := gin.Default()

	return r
}

//

func killPort(ip, port string) {

	if isPortInUse(ip, port) {
		// Kill the process using the port
		cmd := exec.Command("kill", "-9", "<PID>")
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error killing the process:", err)
			return
		}
		fmt.Println("Killed the process using port", port)
	}

}
func isPortInUse(ip, port string) bool {
	address := ip + port
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}
