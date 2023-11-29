package service

import (
	"errors"
	"gin-init/basic"
	"gin-init/model/user"
	"gin-init/util"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	ParamIsNull  = "参数为空"
	DataNotFound = "数据不存在"
	PasswordErr  = "密码错误"
	PasswordErr2 = "两次输入密码不相同"
)

// RegisterRequest 用户注册
type RegisterRequest struct {
	UserAccount   string `json:"userAccount" example:"账户名"`
	UserPassword  string `json:"userPassword" example:"密码"`
	CheckPassword string `json:"checkPassword" example:"二次输入密码"`
}

// LoginRequest 用户登录
type LoginRequest struct {
	UserAccount  string `json:"userAccount"`  // 账户名
	UserPassword string `json:"userPassword"` // 密码
}

const (
	CreateUserFail  = "注册失败"
	AccountHasExist = "用户名已存在"
)

// Register 用户注册
func Register(db *gorm.DB, req *RegisterRequest) error {
	_, err := user.QueryUserByAccount(db, req.UserAccount)
	if !errors.Is(err, gorm.ErrRecordNotFound) && err != nil {
		log.Errorf("QueryUserByAccount error: %s", err.Error())
		return basic.NewErr(basic.InnerError, AccountHasExist, err)
	}
	createUser := user.User{
		UserID:       util.NewShortIDString("user"),
		UserAccount:  req.UserAccount,
		UserPassword: req.UserPassword,
	}
	createErr := user.CreateUser(db, &createUser)
	if createErr != nil {
		log.Errorf("CreateUser error: %s", err.Error())
		return basic.NewErr(basic.InnerError, CreateUserFail, createErr)
	}
	return nil
}

// Login 用户登录
func Login(req *LoginRequest) error {
	return nil
}
