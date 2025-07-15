package main

import (
    "fmt"
    "log"
	"jaredBlog/global"
    "jaredBlog/models"
    "golang.org/x/crypto/bcrypt"
)

func main() {
    username := "my_username" // 可通过命令行参数输入
    password := "my_secure_password" // 可通过命令行参数输入

    user := models.User{
        Username: username,
        Password: password,
        Role:     "admin",
    }

    // 加密密码
    if err := user.HashPassword(); err != nil {
        log.Fatalf("Failed to hash password: %v", err)
    }

    // 保存到数据库
    db, err := db.Connect() // 假设你有一个数据库连接函数
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    result := global.DB.Create(&user)
    if result.Error != nil {
        log.Fatalf("Failed to create superadmin: %v", result.Error)
    }

    fmt.Println("Superadmin user created successfully!")
}