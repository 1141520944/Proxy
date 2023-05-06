package server

import (
	"github.com/1141520944/proxy/server/dao/mysql"
	"github.com/1141520944/proxy/server/pkg/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 服务状态
var server_state chan struct{}

func (sh *ServerHandler) ServerClose(c *gin.Context) {
	showPort := c.Param("show_port")
	result := new(models.ResultData)
	server_state <- struct{}{}
	code, err := sh.mysql.Server_checkShowPort(showPort)
	if err != nil {
		zap.L().Error("sh.mysql.Server_checkShowPort() with fail")
		result.ResponseError(c, models.CodeServerBusy)
		return
	}
	if code == mysql.ErrorServerPortNoExist {
		result.ResponseSuccess(c, models.CodeSuccess, 1)
		return
	}
}
