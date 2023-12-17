package basic

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// Page 定义分页对象
type Page struct {
	PageSize   int         `json:"pageSize"`          // 分页大小，默认10， 必须
	PageNo     int         `json:"pageNo"`            // 当前页码 ，必须
	TotalCount int         `json:"totalCount"`        // 总数，必须
	Order      string      `json:"order,omitempty"`   // 排序方式，可选
	OrderBy    string      `json:"orderBy,omitempty"` // 排序字段，可选
	Result     interface{} `json:"result"`            // 分页结果，必须
}

type Gin struct {
	C *gin.Context
}

// ResponseSuccessByPage 分页response 成功返回
type ResponseSuccessByPage struct {
	Success bool        `json:"success"`
	Page    interface{} `json:"page"`
}

// ResponseSuccess 非分页response 成功返回
type ResponseSuccess struct {
	Success bool        `json:"success"`
	Result  interface{} `json:"result"`
}

// ResponseFail 失败返回
type ResponseFail struct {
	Success bool        `json:"success"`
	Code    string      `json:"code"`
	Message interface{} `json:"message"`
}

// GetGin 获取Gin
func GetGin(c *gin.Context) Gin {
	return Gin{c}
}

// ResponseSuccess 返回成功
func (g *Gin) ResponseSuccess(data interface{}) {
	if data != nil {
		g.C.JSON(http.StatusOK, gin.H{
			"code":    SUCCESS,
			"message": "successful",
			"data":    data,
		})
		return
	} else {
		g.C.JSON(http.StatusOK, gin.H{
			"code":    SUCCESS,
			"message": "successful",
		})
		return
	}
}

// ResponseWithError 失败响应
func (g *Gin) ResponseWithError(httpCode int, err Error) {
	resp := &ResponseFail{
		Success: false,
		Code:    strconv.Itoa(err.Code()),
		Message: map[string]interface{}{
			"global":      err.GetMsg(),
			"serverError": err.Error(),
		},
	}

	g.C.JSON(httpCode, resp)
	return
}

// ResponseSuccessWithOk 成功响应
func (g *Gin) ResponseSuccessWithOk(data interface{}) {
	resp := &ResponseSuccess{
		Success: true,
		Result:  data,
	}
	g.C.JSON(http.StatusOK, resp)
	return
}
