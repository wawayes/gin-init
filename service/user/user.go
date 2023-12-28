package user

import (
	"errors"
	"gin-init/basic"
	"gin-init/common/database"
	"gin-init/model/user"
	"gin-init/util"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
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

var (
	LoginFail       = "登录失败"
	CreateUserFail  = "注册失败"
	QueryUserErr    = "查询失败"
	AccountHasExist = "用户名已存在"
	AccountNotExist = "用户名或密码错误"
	PasswordErr     = "密码输入异常"
)

// Register 用户注册
func Register(req *RegisterRequest) basic.Error {
	db := database.GetInstanceConnection().GetDB()
	_, err := user.QueryUserByAccount(db, req.UserAccount)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		log.Errorf("QueryUserByAccount error: %s", err.Error())
		return basic.NewErr(basic.InnerError, AccountHasExist, err)
	}
	if err != nil {
		log.Errorf("QueryUserByAccount error: %s", err.Error())
		return basic.NewErr(basic.InnerError, QueryUserErr, err)
	}
	if req.UserPassword != req.CheckPassword {
		log.Errorf("CheckPassword error: %s", PasswordErr)
		return basic.NewErr(basic.InnerError, PasswordErr, errors.New(PasswordErr))
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
func Login(req *LoginRequest) (*user.User, basic.Error) {
	db := database.GetInstanceConnection().GetDB()
	queryUser, queryErr := user.QueryUserByAccount(db, req.UserAccount)
	// 错误处理
	if errors.Is(queryErr, gorm.ErrRecordNotFound) {
		log.Errorf("QueryUserByAccount error: %s", queryErr.Error())
		return nil, basic.NewErr(basic.InnerError, AccountNotExist, queryErr)
	}
	if queryErr != nil {
		log.Errorf("QueryUserByAccount error: %s", queryErr.Error())
		return nil, basic.NewErr(basic.InnerError, LoginFail, queryErr)
	}
	// 密码校验
	if queryUser.UserPassword != req.UserPassword {
		log.Errorf("QueryUserByAccount error: %s", PasswordErr)
		return nil, basic.NewErr(basic.InnerError, AccountNotExist, errors.New(PasswordErr))
	}
	// 信息脱敏并返回
	loginUser := &user.User{
		UserID:      queryUser.UserID,
		UserAccount: queryUser.UserAccount,
		Nickname:    queryUser.Nickname,
		UserRole:    queryUser.UserRole,
		AvatarUrl:   queryUser.AvatarUrl,
		PhoneNumber: queryUser.PhoneNumber,
		Email:       queryUser.Email,
		Status:      queryUser.Status,
	}
	return loginUser, nil
}
