package config

import (
	"fmt"
	"jaredBlog/global"
	"jaredBlog/models"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() {
	db_user := AppConfig.Database.User
	db_password := AppConfig.Database.Password
	dsn := fmt.Sprintf("%v:%v@tcp(127.0.0.1:3306)/currency_exchange?charset=utf8mb4&parseTime=True&loc=Local", db_user, db_password)
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	sqlDB.SetMaxOpenConns(AppConfig.Database.MaxOpenConns)
	sqlDB.SetMaxIdleConns(AppConfig.Database.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(time.Duration(AppConfig.Database.ConnMaxLifetime) * time.Second) // 10 seconds

	// 自动迁移表结构
	if err := DB.AutoMigrate(&models.User{}, &models.Article{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	fmt.Println("Database connected and migrated successfully!")

	global.DB = DB

}
