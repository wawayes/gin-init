package main

import (
	"gin-init/core"
	"gin-init/initialize"
)

func main() {
	// 初始化配置
	initialize.InitConfig()
	// 启动服务
	core.RunServer()

	// 释放
	defer initialize.CloseResource()
}
