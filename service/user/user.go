package user

import (
	"context"
	"encoding/json"
	"errors"
	"gin-init/common/redis"
	"time"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"gin-init/basic"
	"gin-init/common/database"
	"gin-init/model/user"
	"gin-init/util"
)

type UserCommon struct {
	UserModel user.User
}

type UserService interface {
	Register(ctx context.Context, req *RegisterRequest) basic.Error
	Login(ctx context.Context, req *LoginRequest) (*user.User, basic.Error)
}

var (
	LoginFail       = "登录失败"
	CreateUserFail  = "注册失败"
	QueryUserErr    = "查询失败"
	AccountHasExist = "用户名已存在"
	AccountNotExist = "用户名或密码错误"
	PasswordErr     = "密码输入异常"
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

// Register 用户注册
func (cu *UserCommon) Register(ctx context.Context, req *RegisterRequest) basic.Error {
	db := database.GetInstanceConnection().GetDB()
	u, err := cu.UserModel.QueryUserByAccount(db, req.UserAccount)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Errorf("QueryUserByAccount error: %s, req: %v", err.Error(), req)
		return basic.NewErr(basic.InnerError, QueryUserErr, err)
	}
	// 如果用户名已存在，则返回 AccountHasExist 用户名已存在
	if u != nil {
		log.Errorf("AccountHasExist error: %s, req: %v", AccountHasExist, req)
		return basic.NewErr(basic.InnerError, AccountHasExist, err)
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
	createErr := cu.UserModel.CreateUser(db, &createUser)
	if createErr != nil {
		log.Errorf("CreateUser error: %s", createErr.Error())
		return basic.NewErr(basic.InnerError, CreateUserFail, createErr)
	}
	return nil
}

// Login 用户登录
func (cu *UserCommon) Login(ctx context.Context, req *LoginRequest) (*user.User, basic.Error) {
	db := database.GetInstanceConnection().GetDB()
	r := redis.GetRedisInstance().GetClient()
	queryUser, queryErr := cu.UserModel.QueryUserByAccount(db, req.UserAccount)
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
	// 登录状态存入redis
	userData, err := json.Marshal(loginUser)
	if err != nil {
		log.Errorf("json marshal error login service, err: %s", err.Error())
		return nil, basic.NewErr(basic.InnerError, LoginFail, err)
	}
	err = r.Set(ctx, "login"+loginUser.UserID, userData, time.Hour).Err()
	if err != nil {
		log.Errorf("SetNX error: %s", err.Error())
		return nil, basic.NewErr(basic.InnerError, LoginFail, err)
	}
	return loginUser, nil
}
