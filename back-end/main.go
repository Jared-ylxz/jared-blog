package main

import (
	"fmt"
	"jaredBlog/config"
	"jaredBlog/router"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// 根据环境加载对应的 .env 文件
	env := os.Getenv("RUNNING_ENV") // 获取环境变量
	if env == "" {
		env = "development" // 启动后端项目时，如果没有设置环境变量，则默认使用 development 环境
	}
	err := godotenv.Load(".env." + env) // 根据环境加载对应的 .env 文件
	if err != nil {
		log.Println("Error loading .env file")
	}

	config.InitConfig()
	r := router.SetupRouter()

	// 启动服务
	// port := config.AppConfig.App.Port
	// host := config.AppConfig.App.Host
	host := os.Getenv("APP_HOST")
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080" // 默认端口
	}
	r.Run(fmt.Sprintf("%s:%s", host, port))

	// 优雅地关闭服务器: Ctrl + Shift + C
}
