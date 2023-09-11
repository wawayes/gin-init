package initialize

import (
	"gin-init/global"
)

func InitConfig() {
	// 捕获启动时错误
	defer global.CatchError()
	// 程序退出前释放资源
	//defer CloseResource()
	// 初始化全局配置文件
	InitViperConfig()
	// 初始化zap日志
	InitLogger()
	// 初始化gorm
	InitGorm()
	// 初始化redis
	InitRedis()
	// 初始化es
	InitES()
	// 初始化sessions
	InitSessions()
}

// 程序退出前释放资源
func CloseResource() {
	// 关闭数据库连接
	if global.MysqlClient != nil {
		db, _ := global.MysqlClient.DB()
		_ = db.Close()
	}
}
