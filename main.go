package main

import (
	"log"

	"samu.com/inventory_management_system/configs"
	"samu.com/inventory_management_system/srv"
)

func main() {

	error := srv.Start(configs.Config.Server, configs.Config.Addr)

	if error != nil {

		log.Println("Start", error)

	}

}
