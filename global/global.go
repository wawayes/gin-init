package global

import (
	"gin-init/config"

	"github.com/go-redis/redis/v8"
	"github.com/olivere/elastic/v7"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// 常量
const (
	ConfigFile = "./config.yaml" // 配置文件
)

// 变量
var (
	Config        config.ServerConfig // 全局配置
	Logger        *zap.Logger         // 日志
	MysqlClient   *gorm.DB            //Mysql客户端
	RedisClient   *redis.Client       //Redis客户端
	ElasticClient *elastic.Client     // ES客户端
)
