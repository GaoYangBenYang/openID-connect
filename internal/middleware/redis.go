package middleware

import (
	"errors"
	"log"
	"time"

	"github.com/go-redis/redis"
)

// redis分组定义
const (
	OIDC         string = "oidc"
	CLIENT       string = "client"
	COOKIE       string = "cookie"
	ACCESS_TOKEN string = "access_token"
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

/*
Setnx 只有当key不存在时进行存储,
kay:存储键，
value:存储值，
expiration：过期时间 0 不设置过期时间 单位ms
*/
func SetString(key, value string, expiration time.Duration) error {
	// 判断key是否存在
	if RedisClient.Exists(key).Val() == 1 {
		return errors.New("key已存在")
	}
	//存储
	if RedisClient.Set(key, value, expiration*1000*1000).Val() == "OK" {
		return nil
	}
	return errors.New("缓存失败")
}
