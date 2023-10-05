package middleware

import (
	"log"

	"github.com/go-redis/redis"
)

// 声明一个全局的rdb变量
var RedisClient *redis.Client

// 初始化连接
func init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := RedisClient.Ping().Result()
	if err != nil {
		log.Println("Redis连接失败!", err)
	} else {
		log.Println("Redis连接成功!")
	}
}
