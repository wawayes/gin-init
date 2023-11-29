package basic

import (
	"errors"
)

const (
	SUCCESS       = 200
	InvalidParams = 400
	FailAuth      = 401
	InnerError    = 500
)

var errMsg = map[int]string{
	InvalidParams: "输入参数有误",
}

// Error 错误接口
type Error interface {
	error
	Code() int
	GetMsg() string
}

// error struct
type errorString struct {
	code int
	msg  string
	error
}

func (e *errorString) GetMsg() string {
	return e.msg
}
func (e *errorString) Code() int {
	return e.code
}

func NewErr(code int, msg string, err error) Error {
	if err == nil {
		err = errors.New(msg)
	}
	return &errorString{code, msg, err}
}

func NewErrWithCode(code int, err error) Error {
	if msg, ok := errMsg[code]; ok {
		return &errorString{code, msg, err}
	}
	return &errorString{code, err.Error(), err}
}
