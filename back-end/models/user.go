package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Password string
	Email    string
	Phone    string `gorm:"type:char(15)"`
	Role     uint8  `gorm:"default:0;not null;type:tinyint;comment:'0:user, 1:admin, 2:editor'"`
}
