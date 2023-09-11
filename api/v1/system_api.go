package v1

import (
	"gin-init/global"
	"gin-init/model/response"

	"github.com/gin-gonic/gin"
)

// 配置信息
func GetConfig(ctx *gin.Context) {
	response.OkWithData(ctx, global.Config)
}
