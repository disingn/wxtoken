package dao

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"wxlogin/cfg"
)

var RedisClient *redis.Client
var Ctx = context.Background()

func Init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Config.Sever.RdsHost, cfg.Config.Sever.RdsPort),
		Password: cfg.Config.Sever.RdsPass,
		DB:       cfg.Config.Sever.RdsDB,
	})

	pong, err := RedisClient.Ping(Ctx).Result()
	if err != nil {
		log.Println("Error connecting to Redis:", err)
		return
	}
	log.Println("Connected to Redis:", pong)
}
