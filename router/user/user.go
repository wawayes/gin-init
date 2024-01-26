package user

import (
	"gin-init/basic"
	"gin-init/service/user"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func Register(c *gin.Context) {
	g := basic.GetGin(c)
	var (
		req = &user.RegisterRequest{}
		u   = &user.UserCommon{}
	)
	err := basic.ParseJSON(c, req)
	if err != nil {
		log.Errorf("JSON Parse error: %s", err.Error())
		g.ResponseWithError(http.StatusOK, err)
		return
	}
	registerErr := u.Register(c, req)
	if registerErr != nil {
		log.Errorf("Register Error: %s", registerErr.Error())
		g.ResponseWithError(http.StatusOK, registerErr)
		return
	}
	g.ResponseNoPageSuccess(nil)
}

func Login(c *gin.Context) {
	g := basic.GetGin(c)
	var (
		req = &user.LoginRequest{}
		u   = &user.UserCommon{}
	)
	parseErr := basic.ParseJSON(c, req)
	if parseErr != nil {
		log.Errorf("JSON Parse Error: %s", parseErr.Error())
		g.ResponseWithError(http.StatusOK, parseErr)
		return
	}
	loginUser, loginErr := u.Login(c, req)
	if loginErr != nil {
		log.Errorf("router user Login error:%s", loginErr)
		g.ResponseWithError(http.StatusOK, loginErr)
		return
	}
	g.ResponseNoPageSuccess(loginUser)
}
