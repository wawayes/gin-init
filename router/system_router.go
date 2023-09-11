/**
 * @Author Mr.LiuQH
 * @Description 不需要验证登录
 * @Date 2021/7/5 3:44 下午
 **/
package router

import (
	v1 "gin-init/api/v1"

	"github.com/gin-gonic/gin"
)

// 系统路由
func InitSystemRouter(engine *gin.Engine) {
	systemRouter := engine.Group("system")
	{
		// 获取全局变量
		systemRouter.GET("config", v1.GetConfig)
	}
}
