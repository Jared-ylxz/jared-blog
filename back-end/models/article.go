package models

import "time"

type Article struct {
	// gorm.Model
	ID          uint       `gorm:"primarykey" json:"id"`
	CreatedAt   time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   *time.Time `gorm:"index" json:"deleted_at"` // 修改为指针类型，允许为空
	Title       string     `gorm:"type:varchar(255);not null" json:"title" binding:"required"`
	Description string     `gorm:"type:varchar(255)" json:"description"`
	Content     string     `gorm:"type:longtext;not null" json:"content" binding:"required"`
	IsPublished bool       `gorm:"not null" default:"true" json:"is_published"`
	AuthorID    uint       `gorm:"not null" index:"idx_author_id" json:"author_id"` // 会成为数据库的author_id字段,成为实际起作用的外键
	Author      User       `gorm:"foreignKey:AuthorID" json:"author"`               // 用于关联查询，不会出现在数据库字段中
	// CategoryID  uint       `gorm:"not null" index:"idx_category_id" json:"category_id"`
	// Category    Category   `gorm:"foreignKey:CategoryID" json:"category"`
}
