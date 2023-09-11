/**
 * @Description 用户相关的实体
 **/
package model

import (
	"time"

	"gorm.io/plugin/soft_delete"
)

// 用户表
type User struct {
	ID           uint                  `json:"id" gorm:"column:id;primarykey"`
	CreatedAt    time.Time             `json:"createdAt" gorm:"column:createdAt"`
	UpdatedAt    time.Time             `json:"updatedAt" gorm:"column:updatedAt"`
	DeletedAt    soft_delete.DeletedAt `json:"-" gorm:"column:isDeleted;softDelete:flag"`
	UserAccount  string                `json:"userAccount" gorm:"column:userAccount;index;comment:账户名"`
	Nickname     string                `json:"nickname" gorm:"column:nickname;default:gogogo;comment:用户昵称"`
	UserPassword string                `json:"userPassword" gorm:"column:userPassword;comment:用户登录密码"`
	UserRole     string                `json:"userRole" gorm:"column:userRole;default:普通用户;comment:用户角色"`
	AvatarUrl    string                `json:"avatarUrl" gorm:"column:avatarUrl;default:https://pkg.go.dev/static/shared/icon/favicon.ico:用户头像URL"`
	PhoneNumber  string                `json:"phoneNumber" gorm:"column:phoneNumber;comment:用户电话号码"`
	Email        string                `json:"email" gorm:"column:email;comment:用户邮箱地址"`
	Status       int                   `json:"status" gorm:"column:status;default:1;comment:用户状态 0-正常 1-冻结"`
}
