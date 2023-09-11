/**
 * @Author Mr.LiuQH
 * @Description 系统接口
 * @Date 2021/7/6 4:43 下午
 **/
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
