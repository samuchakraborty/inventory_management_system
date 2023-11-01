package main

import (
	"inventory/configs"
	"inventory/server"
	"log"
)

func main() {

	error := server.Start(configs.Config.Server, configs.Config.Addr)

	if error != nil {
		log.Println("Start", error)
	}

}
