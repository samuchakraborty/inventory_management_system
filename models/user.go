package models

import (
	"time"
)

// gorm.Model definition
type User struct {
	// *gorm.Model
	Id        uint      `gorm:"primaryKey, autoIncrement" json:"id" query:"id" form:"id"`
	Username  string    `gorm:"unique" json:"user_name" query:"user_name" form:"user_name"`
	Role      string    `json:"role" query:"role" form:"role"`
	CreatedAt time.Time `json:"created_at" query:"created_at" form:"created_at"`
	
}
