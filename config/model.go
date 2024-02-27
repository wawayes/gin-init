package config

import "time"

// server config
type Server struct {
	ServerName string
	ServerPort string
}

// database config
type Database struct {
	DbType                          string // 数据库类型（如：MySQL, PostgreSQL等）
	Host                            string // 数据库服务器的主机名或IP地址
	Name                            string // 数据库名
	User                            string // 用于数据库连接的用户名
	Password                        string // 用于数据库连接的密码
	MaxIdleConns                    int    // 数据库连接池中的最大空闲连接数
	MaxOpenConns                    int    // 数据库连接池中的最大打开连接数
	ConnMaxLifetime                 int    // 数据库连接的最大生存时间（秒）
	ConnMaxIdleTime                 int    // 数据库连接的最大空闲时间（秒）
	TablePrefix                     string // 表名前缀，用于数据库中的表
	TestDbPath                      string // 测试数据库的路径（用于单元测试等）
	BatchInsertLimit                int    // 批量插入操作的限制数
	StatementTimeout                int    // SQL语句的超时时间（秒）
	LockTimeout                     int    // 数据库锁定操作的超时时间（秒）
	IdleInTransactionSessionTimeout int    // 事务空闲会话的最大超时时间（秒）
	ConnectTimeout                  int    // 数据库连接的超时时间（秒）
	CommDebug                       bool   // 是否开启通讯调试，通常用于记录数据库操作日志
	SingularTable                   bool   // 是否使用单数形式的表名（默认是复数形式）
}

// redis config
type Redis struct {
	Host        string        // Redis服务器的主机名或IP地址
	Passwd      string        // 用于Redis连接的密码
	DB          int           // Redis数据库的编号
	MaxIdle     int           // Redis连接池中的最大空闲连接数
	MaxActive   int           // Redis连接池中的最大活动（打开）连接数
	IdleTimeout time.Duration // Redis连接的最大空闲时间，超时的连接将被关闭
	PoolSize    int           // Redis连接池的大小
}
