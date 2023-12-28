package user

import (
	"errors"
	"gin-init/basic"
	"gin-init/common/database"
	"gin-init/model/user"
	"gin-init/test"
	"github.com/agiledragon/gomonkey/v2"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

func TestUserService(t *testing.T) {
	test.InitSetting()
	database.GetInstanceConnection().Init()

	convey.Convey("TestRegisterSuccess", t, func() {
		patches := gomonkey.ApplyFunc(user.QueryUserByAccount, func(db *gorm.DB, userAccount string) (*user.User, error) {
			return &user.User{}, nil
		})
		defer patches.Reset()

		req := RegisterRequest{
			UserAccount:   "test",
			UserPassword:  "123123",
			CheckPassword: "123123",
		}
		err := Register(&req)
		assert.Nil(t, err)
	})

	convey.Convey("TestRegisterErr1", t, func() {
		patches := gomonkey.ApplyFunc(user.QueryUserByAccount, func(db *gorm.DB, userAccount string) (*user.User, error) {
			return &user.User{}, gorm.ErrRecordNotFound
		})
		defer patches.Reset()

		req := RegisterRequest{
			UserAccount:   "test",
			UserPassword:  "123123",
			CheckPassword: "123123",
		}
		err := Register(&req)
		assert.NotNil(t, err)
	})

	convey.Convey("TestRegisterErr2", t, func() {
		patches := gomonkey.ApplyFunc(user.QueryUserByAccount, func(db *gorm.DB, userAccount string) (*user.User, error) {
			return &user.User{}, errors.New("query error")
		})
		defer patches.Reset()

		req := RegisterRequest{
			UserAccount:   "test",
			UserPassword:  "123123",
			CheckPassword: "123123",
		}
		err := Register(&req)
		assert.NotNil(t, err)
	})

	convey.Convey("TestRegisterErr3", t, func() {
		patches := gomonkey.ApplyFunc(user.QueryUserByAccount, func(db *gorm.DB, userAccount string) (*user.User, error) {
			return &user.User{}, nil
		})
		defer patches.Reset()

		req := RegisterRequest{
			UserAccount:   "test",
			UserPassword:  "123123000",
			CheckPassword: "123123",
		}
		err := Register(&req)
		assert.NotNil(t, err)
	})

	convey.Convey("TestLoginSuccess", t, func() {
		patches := gomonkey.ApplyFunc(user.QueryUserByAccount, func(db *gorm.DB, account string) (*user.User, error) {
			return &user.User{
				UserID:       "user-123123",
				UserAccount:  "admin",
				UserPassword: "123456",
				Nickname:     "nickname",
			}, nil
		})
		defer patches.Reset()
		req := &LoginRequest{
			UserAccount:  "admin",
			UserPassword: "123456",
		}
		loginUser, loginErr := Login(req)
		assert.Nil(t, loginErr)
		assert.NotNil(t, loginUser)
	})

	convey.Convey("TestLoginSuccess", t, func() {
		patches := gomonkey.ApplyFunc(user.QueryUserByAccount, func(db *gorm.DB, account string) (*user.User, error) {
			return &user.User{
				UserID:       "user-123123",
				UserAccount:  "admin",
				UserPassword: "123456",
				Nickname:     "nickname",
			}, nil
		})
		defer patches.Reset()
		req := &LoginRequest{
			UserAccount:  "admin",
			UserPassword: "123456",
		}
		loginUser, loginErr := Login(req)
		assert.Nil(t, loginErr)
		assert.NotNil(t, loginUser)
	})

	convey.Convey("TestLoginErr1", t, func() {
		patches := gomonkey.ApplyFunc(user.QueryUserByAccount, func(db *gorm.DB, account string) (*user.User, error) {
			return &user.User{
				UserID:       "user-123123",
				UserAccount:  "admin",
				UserPassword: "123456",
				Nickname:     "nickname",
			}, nil
		})
		defer patches.Reset()
		req := &LoginRequest{
			UserAccount:  "admin",
			UserPassword: "654321",
		}
		loginUser, loginErr := Login(req)
		assert.NotNil(t, loginErr)
		assert.Nil(t, loginUser)
	})

	convey.Convey("TestLoginErr2", t, func() {
		patches := gomonkey.ApplyFunc(user.QueryUserByAccount, func(db *gorm.DB, account string) (*user.User, error) {
			return nil, gorm.ErrRecordNotFound
		})
		defer patches.Reset()
		req := &LoginRequest{
			UserAccount:  "admin",
			UserPassword: "123456",
		}
		loginUser, loginErr := Login(req)
		assert.NotNil(t, loginErr)
		assert.Nil(t, loginUser)
	})

	convey.Convey("TestLoginErr3", t, func() {
		patches := gomonkey.ApplyFunc(user.QueryUserByAccount, func(db *gorm.DB, account string) (*user.User, error) {
			return nil, basic.NewErr(basic.InnerError, AccountNotExist, errors.New("query Err"))
		})
		defer patches.Reset()
		req := &LoginRequest{
			UserAccount:  "admin",
			UserPassword: "123456",
		}
		loginUser, loginErr := Login(req)
		assert.NotNil(t, loginErr)
		assert.Nil(t, loginUser)
	})
}
