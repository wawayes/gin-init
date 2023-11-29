package config

func InitSetting() {
	DatabaseSetting.User = "root"
	DatabaseSetting.Password = "123456"
	DatabaseSetting.Host = "127.0.0.1:3306"
	DatabaseSetting.Name = "database"
	DatabaseSetting.TablePrefix = "db_"
	DatabaseSetting.BatchInsertLimit = 2
	DatabaseSetting.CommDebug = true

	RedisSetting.Host = "127.0.0.1:6379"
	RedisSetting.Passwd = ""
	RedisSetting.IdleTimeout = 200
	RedisSetting.MaxIdle = 30
	RedisSetting.MaxActive = 30
}
