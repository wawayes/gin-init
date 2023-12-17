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
	var req user.RegisterRequest
	err := basic.ParseJSON(c, req)
	if err != nil {
		log.Errorf("JSON Parse error: %s", err.Error())
		g.ResponseWithError(http.StatusOK, err)
		return
	}
	registerErr := user.Register(&req)
	if registerErr != nil {
		log.Errorf("Register Error: %s", err.Error())
		g.ResponseWithError(http.StatusOK, err)
		return
	}
	g.ResponseSuccessWithOk(nil)
	return
}
