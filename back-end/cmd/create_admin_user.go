package main

import (
	"flag"
	"fmt"
	"jaredBlog/config"
	"jaredBlog/global"
	"jaredBlog/models"
	"jaredBlog/utils"
	"log"
)

func main() {
	// 从命令行参数获取用户名和密码，默认值可改
	username := flag.String("username", "Jared", "Username for admin")
	password := flag.String("password", "12345", "Password for admin")
	flag.Parse()

	// 先哈希密码
	hashedPassword, err := utils.HashPassword(*password)
	if err != nil {
		log.Fatalf("Failed to hash password: %v", err)
	}

	// 创建用户对象并赋值哈希后的密码
	user := models.User{
		Username: *username,
		Password: hashedPassword,
		Role:     "admin",
	}

	// 连接数据库
	config.InitConfig()

	if err := global.DB.Create(&user).Error; err != nil {
		log.Fatalf("Failed to create superadmin: %v", err)
	}

	fmt.Println("Superadmin user created successfully!")
}
