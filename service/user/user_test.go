package user

import (
	"errors"
	"gin-init/common/database"
	"gin-init/model/user"
	"gin-init/test"
	"github.com/agiledragon/gomonkey/v2"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

func TestRegister(t *testing.T) {
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
}
