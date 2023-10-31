package server

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func Database() *gorm.DB {

	dsn := "root:new_password@tcp(127.0.0.1:3306)/inventory?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	// db.SetConnMaxLifetime(time.Minute * 3)
	// db.SetMaxOpenConns(10)
	// db.SetMaxIdleConns(10)

	return db
}

func GenerateAllTable() *gen.Generator {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./models",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	// gormdb, _ := gorm.Open(mysql.Open("root:@(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"))
	// g.UseDB(gormdb) // reuse your gorm db

	// g.ApplyBasic(
	// 	g.GenerateAllTable()...,
	// )
	// g.Execute()
	return g
}
