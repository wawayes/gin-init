package config

import "time"

// server config
type Server struct {
	ServerName string
	ServerPort string
}

// database config
type Database struct {
	DbType                          string
	Host                            string
	Name                            string
	User                            string
	Password                        string
	MaxIdleConns                    int
	MaxOpenConns                    int
	ConnMaxLifetime                 int
	ConnMaxIdleTime                 int
	TablePrefix                     string
	TestDbPath                      string
	BatchInsertLimit                int
	StatementTimeout                int
	LockTimeout                     int
	IdleInTransactionSessionTimeout int
	ConnectTimeout                  int
	CommDebug                       bool
	SingularTable                   bool
}

// redis config
type Redis struct {
	Host        string
	Passwd      string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
	PoolSize    int
}

// page config
type PaginateConfig struct {
	PageNo      int
	PageSizeMax int
}
