package user

import (
	"context"
	"errors"
	"gin-init/basic"
	"gin-init/common/database"
	"gin-init/model/user"
	"gin-init/test"
	"github.com/agiledragon/gomonkey/v2"
	"gorm.io/gorm"
	"reflect"
	"testing"
)

func TestUserCommon_Register(t *testing.T) {
	cu := &UserCommon{
		UserModel: user.User{
			UserAccount: "test_account",
		},
	}

	ctx := context.Background()
	test.InitSetting()
	database.GetInstanceConnection().Init()
	req1 := &RegisterRequest{
		UserAccount:   "test",
		UserPassword:  "test_password",
		CheckPassword: "test_password",
	}

	// 使用gomonkey打桩QueryUserByAccount函数
	patch := gomonkey.ApplyMethod(reflect.TypeOf(&user.User{}), "QueryUserByAccount", func(_ *user.User, _ *gorm.DB, _ string) (*user.User, error) {
		return nil, nil // 模拟用户不存在
	})
	defer patch.Reset()

	err := cu.Register(ctx, req1)
	if err != nil {
		t.Errorf("Expected nil error when user account does not exist, got %v", err)
	}

	// 更改打桩行为以模拟账号已存在的情况
	patch.Reset() // 重置之前的打桩
	patch = gomonkey.ApplyMethod(reflect.TypeOf(&cu.UserModel), "QueryUserByAccount", func(_ *user.User, _ *gorm.DB, _ string) (*user.User, error) {
		return &user.User{}, nil // 模拟用户已存在
	})

	err = cu.Register(ctx, req1)
	if err == nil || err.Error() != AccountHasExist {
		t.Errorf("Expected 'account has exist' error, got %v", err)
	}

	// 更改打桩行为以创建用户失败的情况
	patch.Reset()
	patch = gomonkey.ApplyMethod(reflect.TypeOf(&user.User{}), "QueryUserByAccount", func(_ *user.User, _ *gorm.DB, _ string) (*user.User, error) {
		return nil, nil // 模拟用户不存在
	})
	patch = gomonkey.ApplyMethod(reflect.TypeOf(&cu.UserModel), "CreateUser", func(_ *user.User, _ *gorm.DB) error {
		return basic.NewErr(basic.InnerError, CreateUserFail, errors.New(CreateUserFail)) // 模拟创建用户失败
	})
	err = cu.Register(ctx, req1)
	if err == nil || err.Error() != CreateUserFail {
		t.Errorf("Expected CreateUserFail error, got %v", err)
	}

	// 更改打桩行为以两次密码不正确的情况
	patch.Reset()
	req2 := &RegisterRequest{
		UserAccount:   "test_account12322",
		UserPassword:  "test_password",
		CheckPassword: "error_password",
	}
	err = cu.Register(ctx, req2)
	if err == nil || err.Error() != PasswordErr {
		t.Errorf("Expected PasswordErr error, got %v", err)
	}

}
