package server

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"samu.com/inventory_management_system/controllers"
)

func Database() {

	//dsn := "root:new_password@tcp(127.0.0.1:3306)/inventory?charset=utf8mb4&parseTime=True&loc=Local"
	// database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }

	db, err := sql.Open("mysql", "root:new_password@tcp(127.0.0.1:3306)/inventory?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err.Error())
	}
	controllers.DB = db
}
