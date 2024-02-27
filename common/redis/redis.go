package redis

import (
	"context"
	"sync"

	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"

	"gin-init/config"
)

type InstanceConnection struct {
}

var (
	redisClient       *redis.Client
	redisOnce         sync.Once
	redisInstance     *InstanceConnection
	redisConnectMutex sync.Mutex
)

func GetRedisInstance() *InstanceConnection {
	if redisInstance == nil {
		redisOnce.Do(func() {
			redisInstance = &InstanceConnection{}
		})
	}
	return redisInstance
}

func (r *InstanceConnection) Init() error {
	var (
		err error
	)
	redisConnectMutex.Lock()
	defer redisConnectMutex.Unlock()

	redisClient = redis.NewClient(&redis.Options{
		Addr:     config.RedisSetting.Host,
		Password: "",
		DB:       config.RedisSetting.DB,
		PoolSize: config.RedisSetting.PoolSize,
	})

	// 测试连接
	if _, err = redisClient.Ping(context.Background()).Result(); err != nil {
		log.Errorf("Init Redis ERROR: %s", err.Error())
		return err
	}

	return nil
}

func (r *InstanceConnection) GetClient() *redis.Client {
	return redisClient
}

// 关闭Redis连接
func (r *InstanceConnection) Close() error {
	return redisClient.Close()
}
