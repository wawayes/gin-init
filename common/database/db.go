package database

import (
	"fmt"
	"gin-init/config"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"sync"
	"time"
)

type InstanceConnection struct {
}

var (
	db                 *gorm.DB
	once               sync.Once
	instanceConnection *InstanceConnection
)

func GetInstanceConnection() *InstanceConnection {
	if instanceConnection == nil {
		once.Do(func() {
			instanceConnection = &InstanceConnection{}
		})
	}
	return instanceConnection
}

func (c *InstanceConnection) Init() error {
	var (
		err error
	)
	switch config.DatabaseSetting.DbType {
	case "mysql":
		db, err = gorm.Open(mysql.Open(
			fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
				config.DatabaseSetting.User,
				config.DatabaseSetting.Password,
				config.DatabaseSetting.Host,
				config.DatabaseSetting.Name,
			)), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix:   config.DatabaseSetting.TablePrefix,   // 设置表名前缀
				SingularTable: config.DatabaseSetting.SingularTable, // 禁用表名复数
			},
		})

	default:
		// TODO 其他数据库
	}
	if err != nil {
		log.Errorf("Init DB ERROR err: %s", err.Error())
	}

	sqlSet, err := db.DB()
	if err != nil {
		log.Errorf("Init DB ERROR err: %s", err.Error())
		return err
	}

	// 连接池的空闲数大小
	sqlSet.SetMaxIdleConns(config.DatabaseSetting.MaxIdleConns)
	// 最大打开的连接数
	sqlSet.SetMaxOpenConns(config.DatabaseSetting.MaxOpenConns)
	// 连接最大生存时间
	if config.DatabaseSetting.ConnMaxLifetime != 0 {
		sqlSet.SetConnMaxLifetime(time.Duration(config.DatabaseSetting.ConnMaxLifetime) * time.Minute)
	}
	// 连接超时时间
	if config.DatabaseSetting.ConnMaxIdleTime != 0 {
		sqlSet.SetConnMaxIdleTime(time.Duration(config.DatabaseSetting.ConnMaxIdleTime) * time.Second)
	}

	return nil
}

/*
 * 对外获取数据库连接对象db
 */
func (c *InstanceConnection) GetDB() *gorm.DB {
	return db
}
