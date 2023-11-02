package model

import (
	"time"
)

type User struct {
	ID        int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Username  string    `gorm:"column:username" json:"username"`
	Role      string    `gorm:"column:role" json:"role"`
	CreatedAt time.Time `gorm:"column:created_at;not null;default:current_timestamp()" json:"created_at"`
	DeletedAt time.Time `gorm:"column:deleted_at;not null" json:"deleted_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
}
