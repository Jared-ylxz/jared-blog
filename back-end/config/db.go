package config

import (
	"database/sql"
	"fmt"
	"jaredBlog/global"
	"jaredBlog/models"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() {
	dbName := "jaredBlog"
	dbUser := AppConfig.Database.User
	dbPassword := AppConfig.Database.Password

	// 连接到MySQL服务器，不指定数据库
	dsn := fmt.Sprintf("%v:%v@tcp(127.0.0.1:3306)/?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword)
	sqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database server: %v", err)
	}
	defer sqlDB.Close()

	// 检查并创建数据库
	_, err = sqlDB.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci", dbName))
	if err != nil {
		log.Fatalf("Failed to create database: %v", err)
	}

	// 连接到指定的数据库
	dsn = fmt.Sprintf("%v:%v@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbName)
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// 设置连接池
	sqlDB, err = DB.DB()
	if err != nil {
		log.Fatalf("Failed to get database object: %v", err)
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
