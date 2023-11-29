package test

import "gin-init/config"

func InitSetting() {
	config.DatabaseSetting.DbType = "mysql"
	config.DatabaseSetting.User = "root"
	config.DatabaseSetting.Password = "123456"
	config.DatabaseSetting.Host = "127.0.0.1:3306"
	config.DatabaseSetting.Name = "db"
	config.DatabaseSetting.TablePrefix = "db_"
	config.DatabaseSetting.SingularTable = true
}
