package database

import (
	"context"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	Ctx = context.Background()
)

func Database() {
	dsn := "root:new_password@tcp(127.0.0.1:3306)/inventory?charset=utf8mb4&parseTime=True&loc=Local"

	// db, _ := sql.Open("mysql", "root:new_password@tcp(127.0.0.1:3306)/inventory?charset=utf8mb4&parseTime=True&loc=Local")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("db.Ping failed:", err)
	}

	DB = db
	// return db

}
