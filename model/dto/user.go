package dto

type UserDTO struct {
	ID          uint   `json:"id" gorm:"primarykey"`
	UserAccount string `json:"userAccount" gorm:"index;comment:账户名"`
	Nickname    string `json:"nickname" gorm:"default:gogogo;comment:用户昵称"`
	UserRole    string `json:"userRole" gorm:"default:普通用户;comment:用户角色"`
	AvatarUrl   string `json:"avatarUrl" gorm:"default:https://pkg.go.dev/static/shared/icon/favicon.ico:用户头像URL"`
	PhoneNumber string `json:"phoneNumber" gorm:"comment:用户电话号码"`
	Email       string `json:"email" gorm:"comment:用户邮箱地址"`
	Status      int    `json:"status" gorm:"default:1;comment:用户状态 0-正常 1-冻结"`
}
