package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null; type:varchar(20)"`
	Password string `gorm:"not null; type:varchar(100)"`
	Email    string `gorm:"default:''; not null; type:varchar(30)"`
	Phone    string `gorm:"type:char(15); default:''; not null"`
	Role     uint8  `gorm:"default:0; not null; type:tinyint; comment:'0:user, 1:editor, 9:admin'"`
}
