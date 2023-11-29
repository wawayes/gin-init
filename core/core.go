package core

import (
	"flag"
	"gin-init/common/database"
	"gin-init/config"
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

	router := gin.Default()

	serverPort := config.ServerSetting.ServerPort
	router.Run("0.0.0.0:" + serverPort)
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
