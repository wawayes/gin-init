package user

import (
	"errors"
	"gorm.io/gorm"
)

// User 用户表
type User struct {
	gorm.Model
	UserID       string `gorm:"column:user_id;NOT NULL" json:"user_id"`                                                        // 用户ID
	UserAccount  string `gorm:"column:user_account" json:"user_account"`                                                       // 账户名
	UserPassword string `gorm:"column:user_password" json:"user_password"`                                                     // 密码
	Nickname     string `gorm:"column:nickname;default:gogogo" json:"nickname"`                                                // 用户昵称
	UserRole     string `gorm:"column:user_role;default:user;NOT NULL" json:"user_role"`                                       // 用户角色
	AvatarUrl    string `gorm:"column:avatar_url;default:https://pkg.go.dev/static/shared/icon/favicon.ico" json:"avatar_url"` // 用户头像地址
	PhoneNumber  string `gorm:"column:phone_number" json:"phone_number"`                                                       // 电话
	Email        string `gorm:"column:email" json:"email"`                                                                     // 邮箱地址
	Status       int    `gorm:"column:status;default:1;NOT NULL" json:"status"`                                                // 账户状态 0-正常 1-冻结
}

type UserRepository interface {
	CreateUser(db *gorm.DB, user *User) error
	DeleteUserByUserId(db *gorm.DB, userId string) error
	BatchDeleteUserByUserId(db *gorm.DB, userIds []string) error
	UpdateUser(db *gorm.DB, user *User) error
	QueryUserByAccount(db *gorm.DB, account string) (*User, error)
	QueryUserDetailByUserId(db *gorm.DB, userId string) (*User, error)
	QueryUserInfoByUserId(db *gorm.DB, userId string) (*User, error)
	QueryUserList(db *gorm.DB, page int, size int) ([]*User, error)
}

// CreateUser 创建用户
func (*User) CreateUser(db *gorm.DB, user *User) error {
	return db.Create(user).Error
}

// DeleteUserByUserId 删除用户
func (*User) DeleteUserByUserId(db *gorm.DB, userId string) error {
	return db.Where("user_id = ?", userId).Delete(&User{}).Error
}

// BatchDeleteUserByUserId 批量删除用户
func (*User) BatchDeleteUserByUserId(db *gorm.DB, userIds []string) error {
	return db.Where("user_id in ?", userIds).Delete(&User{}).Error
}

// UpdateUser 更新用户
func (*User) UpdateUser(db *gorm.DB, user *User) error {
	return db.Save(user).Error
}

// QueryUserByAccount 获取用户信息
func (*User) QueryUserByAccount(db *gorm.DB, account string) (user *User, err error) {
	if err := db.Where("user_account = ?", account).First(&user).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return user, err
}

// QueryUserDetailByUserId 获取用户详情
func (*User) QueryUserDetailByUserId(db *gorm.DB, userId string) (*User, error) {
	var user User
	err := db.Where("user_id = ?", userId).First(&user).Error
	return &user, err
}

// QueryUserInfoByUserId 获取用户信息
func (*User) QueryUserInfoByUserId(db *gorm.DB, userId string) (*User, error) {
	var user User
	err := db.Select(
		"user_id",
		"user_account",
		"nickname",
		"user_role",
		"avatar_url",
		"phone_number",
		"email",
		"status",
	).Where("user_id = ?", userId).First(&user).Error
	return &user, err
}

// QueryUserList 获取用户列表
func (*User) QueryUserList(db *gorm.DB, page int, size int) ([]*User, error) {
	var users []*User
	err := db.Limit(size).Offset((page - 1) * size).Find(&users).Error
	return users, err
}
