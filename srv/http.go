package srv

import (
	"fmt"
	"os/exec"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
)

var wg sync.WaitGroup

// start the server
func Start(ip, port string) error {
	killPort(ip, port) // Check and kill the port if needed
	fmt.Println("samu")
	return setUpRouter().Run(ip + ":" + port) // Use "ip:port" format for the address
}

// set up server config
func setUpRouter() *gin.Engine {
	r := gin.Default()
	return r
}

//

// Kill all processes using the specified port
func killPort(ip, port string) {
	pids, err := getPIDsUsingPort(ip, port)
	if err != nil {
		fmt.Println("Error getting PIDs:", err)
		return
	}

	// for _, pid := range pids {
	// 	fmt.Println("Killing process with PID:", pid)
	// 	cmd := exec.Command("kill", "-9", pid)
	// 	err := cmd.Run()
	// 	if err != nil {
	// 		fmt.Printf("Error killing process with PID %s: %v\n", pid, err)
	// 	}
	// }
	fmt.Println(&pids)
	for i := 0; i < len(pids); i++ {
		wg.Add(i + 1)

		pid := pids[i]
		fmt.Println("Killing process with PID:", pid)
		cmd := exec.Command("kill", "-9", pid)
		err := cmd.Run()
		if err != nil {
			fmt.Printf("Error killing process with PID %s: %v\n", pid, err)
		} else {
			fmt.Printf("Killed process with PID %s\n", pid)

		}
	}
	wg.Done()

}

// Get a list of PIDs of processes using the specified port
func getPIDsUsingPort(ip, port string) ([]string, error) {
	cmd := exec.Command("lsof", "-i", "tcp:"+port)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(output), "\n")
	var pids []string
	for _, line := range lines[1:] {
		fields := strings.Fields(line)
		if len(fields) >= 2 {
			pids = append(pids, fields[1])
		}
	}

	return pids, nil
}
