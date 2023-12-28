package user

import (
	"encoding/json"
	"errors"
	"gin-init/basic"
	userModel "gin-init/model/user"
	userService "gin-init/service/user"
	"gin-init/test"
	"github.com/agiledragon/gomonkey/v2"
	"github.com/gin-gonic/gin"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRouter(t *testing.T) {
	test.InitSetting()

	router := gin.Default()
	router.POST("/v1/user/register", Register)
	router.POST("/v1/user/login", Login)
	response := basic.ResponseSuccess{}

	convey.Convey("TestRegister", t, func() {
		patches := gomonkey.ApplyFunc(basic.ParseJSON, func(c *gin.Context, i interface{}) basic.Error {
			return nil
		})
		defer patches.Reset()
		patches = patches.ApplyFunc(userService.Register, func(req *userService.RegisterRequest) error {
			return nil
		})

		url := "/v1/user/register"

		body := basic.Request("POST", url, &map[string]interface{}{"userAccount": "test_user", "userPassword": "111111", "checkPassword": "111111"}, router)
		err := json.Unmarshal(body, &response)
		if err != nil {
			t.Errorf("json unmarshal error: %s", err.Error())
		}
		assert.Equal(t, true, response.Success)
	})

	// fail
	convey.Convey("TestRegisterErr1", t, func() {
		patches := gomonkey.ApplyFunc(basic.ParseJSON, func(c *gin.Context, i interface{}) basic.Error {
			return basic.NewErrWithCode(basic.InnerError, errors.New("json marshal error"))
		})
		defer patches.Reset()
		patches = patches.ApplyFunc(userService.Register, func(req *userService.RegisterRequest) error {
			return nil
		})

		url := "/v1/user/register"

		body := basic.Request("POST", url, &map[string]interface{}{"userAccount": "test_user", "userPassword": "111111", "checkPassword": "111111"}, router)
		err := json.Unmarshal(body, &response)
		if err != nil {
			t.Errorf("json unmarshal error: %s", err.Error())
		}
		assert.Equal(t, false, response.Success)
	})

	convey.Convey("TestRegisterErr2", t, func() {
		patches := gomonkey.ApplyFunc(basic.ParseJSON, func(c *gin.Context, i interface{}) basic.Error {
			return nil
		})
		defer patches.Reset()
		patches = patches.ApplyFunc(userService.Register, func(req *userService.RegisterRequest) basic.Error {
			return basic.NewErrWithCode(basic.InnerError, errors.New("register error"))
		})

		url := "/v1/user/register"

		body := basic.Request("POST", url, &map[string]interface{}{"userAccount": "test_user", "userPassword": "111111", "checkPassword": "222222"}, router)
		err := json.Unmarshal(body, &response)
		if err != nil {
			t.Errorf("json unmarshal error: %s", err.Error())
		}
		assert.Equal(t, false, response.Success)
	})

	convey.Convey("TestLoginSuccess", t, func() {
		patches := gomonkey.ApplyFunc(basic.ParseJSON, func(c *gin.Context, obj interface{}) basic.Error {
			return nil
		})
		defer patches.Reset()
		patches = patches.ApplyFunc(userService.Login, func(req *userService.LoginRequest) (*userModel.User, basic.Error) {
			return &userModel.User{
				UserID:      "user-123123",
				UserAccount: "admin",
			}, nil
		})
		url := "/v1/user/login"
		body := basic.Request("POST", url, &map[string]interface{}{"userAccount": "admin", "userPassword": "111111"}, router)
		err := json.Unmarshal(body, &response)
		if err != nil {
			t.Errorf("json unmarshal error: %s", err.Error())
		}
		assert.Equal(t, true, response.Success)
	})

	convey.Convey("TestLoginErr1", t, func() {
		patches := gomonkey.ApplyFunc(basic.ParseJSON, func(c *gin.Context, obj interface{}) basic.Error {
			return basic.NewErrWithCode(basic.InnerError, errors.New("json marshal error"))
		})
		defer patches.Reset()
		patches = patches.ApplyFunc(userService.Login, func(req *userService.LoginRequest) (*userModel.User, basic.Error) {
			return &userModel.User{
				UserID:      "user-123123",
				UserAccount: "admin",
			}, nil
		})
		url := "/v1/user/login"
		body := basic.Request("POST", url, &map[string]interface{}{"userAccount": "test_user", "userPassword": "111111"}, router)
		err := json.Unmarshal(body, &response)
		if err != nil {
			t.Errorf("json unmarshal error: %s", err.Error())
		}
		assert.Equal(t, false, response.Success)
	})

	convey.Convey("TestLoginErr2", t, func() {
		patches := gomonkey.ApplyFunc(basic.ParseJSON, func(c *gin.Context, obj interface{}) basic.Error {
			return nil
		})
		defer patches.Reset()
		patches = patches.ApplyFunc(userService.Login, func(req *userService.LoginRequest) (*userModel.User, basic.Error) {
			return nil, basic.NewErrWithCode(basic.InnerError, errors.New("login error"))
		})
		url := "/v1/user/login"
		body := basic.Request("POST", url, &map[string]interface{}{"userAccount": "test_user", "userPassword": "111111"}, router)
		err := json.Unmarshal(body, &response)
		if err != nil {
			t.Errorf("json unmarshal error: %s", err.Error())
		}
		assert.Equal(t, false, response.Success)
	})
}
