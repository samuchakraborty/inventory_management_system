package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var (
	DB *sql.DB
)

func Database() {
	// dsn := "root:new_password@tcp(127.0.0.1:3306)/inventory?charset=utf8mb4&parseTime=True&loc=Local"

	db, _ := sql.Open("mysql", "root:new_password@tcp(127.0.0.1:3306)/inventory?charset=utf8mb4&parseTime=True&loc=Local")

	err2 := db.Ping()
	if err2 != nil {
		fmt.Println("db.Ping failed:", err2)
	}
	DB = db

}
