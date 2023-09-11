/**
 * @Description 测试路由
 **/
package router

import (
	v1 "gin-init/api/v1"

	"github.com/gin-gonic/gin"
)

// 测试路由
func InitTestRouter(engine *gin.Engine) {
	systemRouter := engine.Group("test")
	{
		// redis测试使用
		systemRouter.GET("redis", v1.RdTest)
	}
}
