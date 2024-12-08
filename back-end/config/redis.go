package config

import (
	"context"
	"log"
	"os"

	"jaredBlog/global"

	"github.com/go-redis/redis/v8"
)

func InitRedis() {
	RDB := redis.NewClient(&redis.Options{
		Addr:     AppConfig.Redis.Host + ":" + AppConfig.Redis.Port,
		Password: AppConfig.Redis.Password,
		DB:       AppConfig.Redis.DB,
		PoolSize: AppConfig.Redis.PoolSize,
	})

	_, err := RDB.Ping(context.Background()).Result()
	if err != nil {
		if os.Getenv("ENV") == "development" {
			panic("Failed to connect to Redis: " + err.Error())
		} else {
			log.Fatalf("Failed to connect to Redis: %v", err)
		}
	}

	println("Redis connected successfully!")
	global.RDB = RDB
}
