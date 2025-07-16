package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique; not null; type:varchar(20)"`
	Password string `gorm:"not null; type:varchar(100)"; json:"-"`
	Email    string `gorm:"default:''; not null; type:varchar(30)"`
	Phone    string `gorm:"type:varchar(15); default:''; not null"`
	Role     string `gorm:"type:varchar(10); not null default:'user'; value:'user, editor, admin'"`
}
