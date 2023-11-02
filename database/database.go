package database

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	Ctx = context.Background()
)

func Database(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}

func DatabaseMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !isDatabaseConnected() {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database is not connected"})
			c.Abort()
			return
		}
		c.Next()
	}
}

// Function to check if the database is connected
func isDatabaseConnected() bool {

	dsn := "root:new_password@tcp(127.0.0.1:3306)/inventory?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := Database(dsn)
	if err != nil {
		return false
	}

	DB = db
	return true

}
