package router

import (
	"gin-init/router/user"
	"github.com/gin-gonic/gin"
)

func UserRoutes(engine *gin.Engine) {
	v1 := engine.Group("v1")
	{
		v1.POST("/user/register", user.Register)
		v1.POST("/user/login", user.Login)
	}
}
