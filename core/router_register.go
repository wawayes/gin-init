package core

import (
	"gin-init/router"

	"github.com/gin-gonic/gin"
)

// 注册路由入口
func RegisterRouters(engine *gin.Engine) {
	// 注册系统路由
	router.InitSystemRouter(engine)
	// 注册用户路由
	router.InitUserRouter(engine)
	// 测试路由
	router.InitTestRouter(engine)
	// 注册es路由
	router.InitESRouter(engine)
}
