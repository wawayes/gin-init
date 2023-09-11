package middleware

import (
	"gin-init/model/response"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SessionsHelper() gin.HandlerFunc {
	return func(c *gin.Context) {
		//store.Options(sessions.Options{
		//	Domain:   global.Config.Sessions.Domain,
		//	Path:     global.Config.Sessions.Path,
		//	MaxAge:   global.Config.Sessions.MaxAge,
		//	Secure:   global.Config.Sessions.Secure,
		//	SameSite: http.SameSite(global.Config.Sessions.SameSite),
		//})
		session := sessions.Default(c)
		userAccount := session.Get("userAccount")
		role := session.Get("role")
		id := session.Get("id")
		status := session.Get("status")
		if userAccount == nil {
			response.Error(c, "未登录")
			c.Abort()
			return
		}
		if status.(int) == 1 {
			response.Error(c, "用户已被封禁")
			c.Abort()
			return
		}
		c.Set("userAccount", userAccount)
		c.Set("role", role)
		c.Set("id", id)
		c.Next()
	}
}
