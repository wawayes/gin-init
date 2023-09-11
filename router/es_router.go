/**
 * @Description 注册路由
 **/
package router

import (
	v1 "gin-init/api/v1"

	"github.com/gin-gonic/gin"
)

// 注册es相关路由
func InitESRouter(engine *gin.Engine) {
	esGroup := engine.Group("es")
	{
		esGroup.GET("create", v1.CreateIndex)
		esGroup.GET("searchById", v1.SearchById)
	}
}
