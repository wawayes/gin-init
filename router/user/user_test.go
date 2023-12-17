package user

import (
	"encoding/json"
	"errors"
	"gin-init/basic"
	"gin-init/test"
	"github.com/agiledragon/gomonkey/v2"
	"github.com/gin-gonic/gin"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRegister(t *testing.T) {
	test.InitSetting()

	router := gin.Default()
	router.POST("/v1/user/register", Register)
	response := basic.ResponseSuccess{}

	convey.Convey("TestRegister", t, func() {
		patches := gomonkey.ApplyFunc(basic.ParseJSON, func(c *gin.Context, i interface{}) basic.Error {
			return nil
		})
		defer patches.Reset()

		url := "/v1/user/register"

		body := basic.Request("POST", url, &map[string]interface{}{"userAccount": "test_user", "userPassword": "111111", "checkPassword": "111111"}, router)
		err := json.Unmarshal(body, &response)
		if err != nil {
			t.Errorf("json unmarshal error: %s", err.Error())
		}
		assert.Equal(t, true, response.Success)
	})

	// fail
	convey.Convey("TestRegister", t, func() {
		patches := gomonkey.ApplyFunc(basic.ParseJSON, func(c *gin.Context, i interface{}) basic.Error {
			return basic.NewErrWithCode(basic.InnerError, errors.New("json marshal error"))
		})
		defer patches.Reset()

		url := "/v1/user/register"

		body := basic.Request("POST", url, &map[string]interface{}{"userAccount": "test_user", "userPassword": "111111", "checkPassword": "111111"}, router)
		err := json.Unmarshal(body, &response)
		if err != nil {
			t.Errorf("json unmarshal error: %s", err.Error())
		}
		assert.Equal(t, false, response.Success)
	})

	convey.Convey("TestRegister", t, func() {
		patches := gomonkey.ApplyFunc(basic.ParseJSON, func(c *gin.Context, i interface{}) basic.Error {
			return nil
		})
		defer patches.Reset()

		url := "/v1/user/register"

		body := basic.Request("POST", url, &map[string]interface{}{"userAccount": "test_user", "userPassword": "111111", "checkPassword": "222222"}, router)
		err := json.Unmarshal(body, &response)
		if err != nil {
			t.Errorf("json unmarshal error: %s", err.Error())
		}
		assert.Equal(t, false, response.Success)
	})
}
