package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null; type:varchar(20)"`
	Password string `gorm:"not null; type:varchar(100)" json:"-"`
	Email    string `gorm:"default:''; not null; type:varchar(30)"`
	Phone    string `gorm:"type:char(15); default:''; not null"`
	Role     string `gorm:"default:'user'; not null; type:varchar(10) comment:'user, editor, admin'"`
}
