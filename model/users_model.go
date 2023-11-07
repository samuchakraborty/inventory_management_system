package model

import (
	"time"
)

type User struct {
	ID        int32     `gorm:"->;column:id;primaryKey;autoIncrement:true" json:"id"  swaggerignore:"true"`
	Username  string    `gorm:"column:username" json:"username" `
	Role      string    `gorm:"column:role" json:"role" `
	CreatedAt time.Time `gorm:"->:false;<-:create" json:"created_at" swaggerignore:"true"`
	DeletedAt time.Time `gorm:"->" json:"deleted_at" swaggerignore:"true"`
	UpdatedAt time.Time `gorm:"->" json:"updated_at"  swaggerignore:"true"`
}
