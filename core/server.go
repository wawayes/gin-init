package core

import (
	"fmt"
	"gin-init/global"
	"gin-init/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 获取自定义HTTP SERVER
func getCustomHttpServer(engine *gin.Engine) *http.Server {
	// 创建自定义配置服务
	httpServer := &http.Server{
		//ip和端口号
		Addr: global.Config.App.Addr,
		//调用的处理器，如为nil会调用http.DefaultServeMux
		Handler: engine,
		//计算从成功建立连接到request body(或header)完全被读取的时间
		ReadTimeout: time.Second * 10,
		//计算从request body(或header)读取结束到 response write结束的时间
		WriteTimeout: time.Second * 10,
		//请求头的最大长度，如为0则用DefaultMaxHeaderBytes
		MaxHeaderBytes: 1 << 20,
	}
	return httpServer
}

// RunServer 启动服务
func RunServer() {
	// 日志强制高亮
	gin.ForceConsoleColor()
	// 原生日志
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}
	engine := gin.Default()
	// 注册公共中间件
	engine.Use(gin.Recovery())
	engine.Use(middleware.Cors())
	// sessions
	store, _ := redis.NewStore(10, "tcp", global.Config.Redis.Addr, global.Config.Redis.Password, []byte("secret"))
	engine.Use(sessions.Sessions("mysession", store))
	// 获取自定义http配置
	httpServer := getCustomHttpServer(engine)
	// 注册路由
	RegisterRouters(engine)
	// 打印服务信息
	printServerInfo()
	// 启动服务
	_ = httpServer.ListenAndServe()
}

// 打印服务信息
func printServerInfo() {
	appConfig := global.Config.App
	fmt.Printf("\n【 当前环境: %s 当前版本: %s 接口地址: http://%s 】\n", appConfig.Env, appConfig.Version, appConfig.Addr)
}
