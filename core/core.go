package core

import (
	"flag"
	"gin-init/common/database"
	"gin-init/common/redis"
	"gin-init/config"
	"gin-init/router"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"os"
)

var (
	envMode = flag.String("env", "dev", "env: dev | test | prod")
)

func InitService() {
	// 加载配置
	err := loadConfig(*envMode)
	if err != nil {
		log.Errorf("load config error: %s", err.Error())
		os.Exit(1)
	}

	// 初始化数据库
	err = initDatabase()
	if err != nil {
		log.Errorf("init database error: %s", err.Error())
		os.Exit(1)
	}
	log.Infof("init database success: %s", config.ServerSetting.ServerPort)

	// 初始化Redis
	err = initRedis()
	if err != nil {
		log.Errorf("init redis error: %s", err.Error())
		os.Exit(1)
	}
	log.Infof("init redis success: %s, DB num: %d", config.RedisSetting.Host, config.RedisSetting.DB)

	r := gin.Default()

	router.UserRoutes(r)

	serverPort := config.ServerSetting.ServerPort
	r.Run("0.0.0.0:" + serverPort)
}

// 初始化配置文件
func loadConfig(env string) error {
	err := config.LoadConfig(*envMode)
	if err != nil {
		return err
	}

	return nil
}

// 初始化DB
func initDatabase() error {
	err := database.GetInstanceConnection().Init()
	if err != nil {
		return err
	}
	return nil
}

// 初始化Redis
func initRedis() error {
	err := redis.GetRedisInstance().Init()
	if err != nil {
		return err
	}
	return nil
}
