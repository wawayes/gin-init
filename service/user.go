package service

import (
	"errors"
	"gin-init/global"
	"gin-init/model"
	"gin-init/model/dto"
	"gin-init/model/request"

	"gorm.io/gorm"
)

var (
	ParamIsNull  = "参数为空"
	DataNotFound = "数据不存在"
	PasswordErr  = "密码错误"
	PasswordErr2 = "两次输入密码不相同"
)

// Login 登录
func Login(req *request.Login) (*model.User, error) {
	var user model.User
	//校验账户和密码
	if errors.Is(global.MysqlClient.Where("userAccount = ?", req.UserAccount).First(&user).Error, gorm.ErrRecordNotFound) {
		global.Logger.Sugar().Error(DataNotFound)
		return nil, errors.New(DataNotFound)
	}
	// TODO 密码 解密 登录
	if req.UserPassword != user.UserPassword {
		global.Logger.Sugar().Error(PasswordErr)
		return nil, errors.New(PasswordErr)
	}
	return &user, nil
}

// Register 注册用户
func Register(req request.Register) (uint, error) {
	// 参数校验
	if req.UserAccount == "" || req.UserPassword == "" || req.CheckPassword == "" {
		global.Logger.Sugar().Error(ParamIsNull)
		return 0, errors.New(ParamIsNull)
	}
	if req.CheckPassword != req.UserPassword {
		global.Logger.Sugar().Error(PasswordErr2)
		return 0, errors.New(PasswordErr2)
	}
	user := model.User{
		UserAccount:  req.UserAccount,
		UserPassword: req.UserPassword,
	}
	//与DB交互
	_ = global.MysqlClient.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&user).Error; err != nil {
			global.Logger.Sugar().Errorf("新增用户失败: %s", err)
			return err
		}
		return nil
	})
	return user.ID, nil
}

// FindUserByID 根据ID查找用户
func FindUserByID(uid uint) (userDTO *dto.UserDTO, err error) {
	var user model.User
	if errors.Is(global.MysqlClient.First(&user, uid).Error, gorm.ErrRecordNotFound) {
		global.Logger.Sugar().Error(DataNotFound)
		return nil, errors.New(DataNotFound)
	}
	userDTO = &dto.UserDTO{
		ID:          user.ID,
		UserAccount: user.UserAccount,
		Nickname:    user.Nickname,
		UserRole:    user.UserRole,
		AvatarUrl:   user.AvatarUrl,
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
		Status:      user.Status,
	}
	return userDTO, nil
}
