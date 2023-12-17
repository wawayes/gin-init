package basic

import (
	"bytes"
	"encoding/json"
	"gin-init/model/user"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

const (
	UseridKey = "user"
)

// GetUserObject 从context中取得用户对象,以支持业务逻辑层的需要
func GetUserObject(c *gin.Context) *user.User {
	userObj, exist := c.Get(UseridKey)
	if exist {
		return userObj.(*user.User)
	}
	return nil
}

func Request(requestType string, url string, param *map[string]interface{}, router *gin.Engine) []byte {
	// 将参数转化为json比特流
	var req *http.Request
	if param != nil {
		jsonByte, _ := json.Marshal(*param)
		req = httptest.NewRequest(requestType, url, bytes.NewReader(jsonByte))
	} else {
		req = httptest.NewRequest(requestType, url, bytes.NewReader(nil))
	}

	// 初始化响应
	w := httptest.NewRecorder()
	// 调用相应的handler接口
	router.ServeHTTP(w, req)
	// 提取响应
	result := w.Result()
	defer result.Body.Close()
	// 读取响应body
	body, _ := ioutil.ReadAll(result.Body)
	return body
}

// ParseJSON 解析请求JSON
func ParseJSON(c *gin.Context, obj interface{}) Error {
	if err := c.ShouldBindBodyWith(obj, binding.JSON); err != nil {
		return NewErrWithCode(InvalidParams, err)
	}
	return nil
}

// ParseQuery 解析Query参数
func ParseQuery(c *gin.Context, obj interface{}) Error {
	if err := c.ShouldBindQuery(obj); err != nil {
		return NewErrWithCode(InvalidParams, err)
	}
	return nil
}

// ParseForm 解析Form请求
func ParseForm(c *gin.Context, obj interface{}) Error {
	if err := c.ShouldBindWith(obj, binding.Form); err != nil {
		return NewErrWithCode(InvalidParams, err)
	}
	return nil
}

// ParseUri 解析Uri请求
func ParseUri(c *gin.Context, obj interface{}) Error {
	if err := c.ShouldBindUri(obj); err != nil {
		return NewErrWithCode(InvalidParams, err)
	}
	return nil
}
