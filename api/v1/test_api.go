/**
 * @Description TODO
 **/
package v1

import (
	"gin-init/global"
	"gin-init/model/response"
	"time"

	"github.com/gin-gonic/gin"
)

// 验证redis
func RdTest(ctx *gin.Context) {
	method, _ := ctx.GetQuery("type")
	var result string
	var err error
	switch method {
	case "get":
		result, err = global.RedisClient.Get(ctx, "test").Result()
	case "set":
		result, err = global.RedisClient.Set(ctx, "test", "hello word!", time.Hour).Result()
	}
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.OkWithData(ctx, result)
}
