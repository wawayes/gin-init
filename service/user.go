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

/**
 * @description: 账户密码登录
 * @param user
 */
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

// 注册用户
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
	// 与DB交互
	//_ = global.MysqlClient.Transaction(func(tx *gorm.DB) error {
	//	if err := tx.Create(&user).Error; err != nil {
	//		global.Logger.Sugar().Errorf("新增用户失败: %s", err)
	//		return err
	//	}
	//	return nil
	//})
	if err := global.MysqlClient.Create(&user).Error; err != nil {
		global.Logger.Sugar().Errorf("新增用户失败: %s", err)
		return 0, err
	}
	return user.ID, nil
}

// 根据ID查找用户
func FindUserByID(uid uint) (user *dto.UserDTO, err error) {
	if errors.Is(global.MysqlClient.Select("id = ?", uid).First(user).Error, gorm.ErrRecordNotFound) {
		global.Logger.Sugar().Error(DataNotFound)
		return nil, errors.New(DataNotFound)
	}
	return user, nil
}
