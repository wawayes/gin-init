/**
 * @Description 用户相关表
 **/
package v1

import (
	"gin-init/global"
	"gin-init/model/dto"
	"gin-init/model/request"
	"gin-init/model/response"
	"gin-init/service"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

/**
 * @description: 用户注册
 * @param c
 */
func Register(c *gin.Context) {
	// 绑定参数
	var registerParam request.Register
	_ = c.ShouldBindJSON(&registerParam)
	// 调用注册
	registerId, err := service.Register(registerParam)
	if err != nil {
		response.Error(c, err.Error())
		return
	}
	response.OkWithData(c, registerId)
}

/**
 * @description: 用户登录
 * @param c
 */
func Login(c *gin.Context) {
	// 绑定参数
	var req request.Login
	_ = c.ShouldBindJSON(&req)
	// 调用登录服务
	user, err := service.Login(&req)
	if err != nil {
		global.Logger.Error("登录失败:", zap.Any("user", user))
		response.Error(c, err.Error())
		return
	}
	userDTO := &dto.UserDTO{
		ID:          user.ID,
		UserAccount: user.UserAccount,
		Nickname:    user.Nickname,
		UserRole:    user.UserRole,
		AvatarUrl:   user.AvatarUrl,
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
		Status:      user.Status,
	}
	session := sessions.Default(c)

	session.Set("id", userDTO.ID)
	session.Set("userAccount", userDTO.UserAccount)
	session.Set("role", userDTO.UserRole)
	session.Set("status", userDTO.Status)
	session.Save()
	response.OkWithData(c, session.Get("id"))
}

// 查询用户信息
func GetCurrentUser(c *gin.Context) {
	var req request.GetCurrentUser
	c.ShouldBindJSON(&req)
	session := sessions.Default(c)
	userId := req.ID
	v := session.Get(userId)
	if v == nil {
		response.Error(c, "获取登录态失败")
	}
	response.OkWithData(c, v)
}
