package config

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
	"strings"
)

var cfg *ini.File
var ServerSetting = &Server{}
var DatabaseSetting = &Database{}
var RedisSetting = &Redis{}

// LoadConfig load config by env
func LoadConfig(env string) error {
	var err error

	// dev mapping test config test.ini
	if strings.EqualFold(env, "dev") {
		cfg, err = ini.Load("conf/dev.ini")
	} else if strings.EqualFold(env, "test") {
		// prod mapping prod config prod.ini
		cfg, err = ini.Load("conf/test.ini")
	} else if strings.EqualFold(env, "prod") {
		// prod mapping prod config prod.ini
		cfg, err = ini.Load("conf/prod.ini")
	}

	if err != nil {
		log.Fatalf("Setup, fail to parse config file: %v", err.Error())
		return err
	}

	createSettingMap()

	return nil
}

func createSettingMap() {
	mapTo("server", ServerSetting)
	mapTo("database", DatabaseSetting)
	mapTo("redis", RedisSetting)

}

// LoadConfigByPath load config by path
func LoadConfigByPath(confPath string) error {
	var err error
	cfg, err = ini.Load(confPath)
	if err != nil {
		log.Fatalf("Setup, fail to parse config file: %v", err.Error())
		return err
	}
	createSettingMap()
	return nil
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo Setting err: %v", err.Error())
	}
}
