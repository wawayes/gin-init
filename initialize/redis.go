// Package initialize /**
package initialize

import (
	"context"
	"gin-init/global"

	"github.com/go-redis/redis/v8"
)

// InitRedis 初始化redis客户端
func InitRedis() {
	if !global.Config.Redis.Enable {
		return
	}
	// 创建
	redisClient := redis.NewClient(&redis.Options{
		Addr:     global.Config.Redis.Addr,
		Password: global.Config.Redis.Password,
		DB:       global.Config.Redis.DefaultDB,
	})
	// 使用超时上下文，验证redis
	timeoutCtx, cancelFunc := context.WithTimeout(context.Background(), global.Config.Redis.DialTimeout)
	defer cancelFunc()
	_, err := redisClient.Ping(timeoutCtx).Result()
	if err != nil {
		panic("redis初始化失败! " + err.Error())
	}
	global.Logger.Info("Redis初始化成功")
	global.RedisClient = redisClient
}
