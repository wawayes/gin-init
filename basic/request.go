package basic

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

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
