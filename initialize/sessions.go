package initialize

import (
	"gin-init/global"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"net/http"
)

func InitSessions() {
	store, _ := redis.NewStore(10, "tcp", global.Config.Redis.Addr, global.Config.Redis.Password, []byte("secret"))
	sessions.Sessions("mysession", store)
	store.Options(sessions.Options{
		Domain:   global.Config.Sessions.Domain,
		Path:     global.Config.Sessions.Path,
		MaxAge:   global.Config.Sessions.MaxAge,
		Secure:   global.Config.Sessions.Secure,
		SameSite: http.SameSite(global.Config.Sessions.SameSite),
	})
}
