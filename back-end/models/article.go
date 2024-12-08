package models

import "time"

type Article struct {
	// gorm.Model
	ID          uint      `gorm:"primarykey" json:"id"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   time.Time `gorm:"index" json:"deleted_at"`
	Title       string    `gorm:"type:varchar(255);not null" json:"title" binding:"required"`
	Description string    `gorm:"type:text" json:"description"`
	Content     string    `gorm:"type:text;not null" json:"content" binding:"required"`
	AuthorID    uint      `gorm:"not null;index:idx_author_id" json:"author_id"` // 会成为数据库的author_id字段,成为实际起作用的外键
	Author      User      `gorm:"foreignKey:AuthorID" json:"author"`             // 用于关联查询，不会出现在数据库字段中
}
