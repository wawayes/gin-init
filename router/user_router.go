/**
 * @Description 用户相关的路由
 **/
package router

import (
	v1 "gin-init/api/v1"
	"github.com/gin-gonic/gin"
)

/**
 * @description: 用户相关的路由
 * @param engine
 */
func InitUserRouter(engine *gin.Engine) {
	// 不需要登录的路由
	userRouter := engine.Group("v1/user")
	{
		// 登录
		userRouter.POST("login", v1.Login)
		// 注册
		userRouter.POST("register", v1.Register)
		// 获取当前登录用户
		userRouter.POST("/current", v1.GetCurrentUser)
	}
}
