package server

import (
	"github.com/1141520944/proxy/common/models"
	"github.com/1141520944/proxy/dao/mysql"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// ServerByShowPort 检测show_port --Server
func (sh *ServerHandler) ServerCheckByShowPort(c *gin.Context) {
	result := new(models.ResultData)
	show_port := c.Param("show_port")
	err, code := sh.mysql.Server_checkShowPort(show_port)
	if err != nil {
		zap.L().Error("sh.mysql.Server_ConnectCan() with fail", zap.Error(err))
		// result.ResponseError(c, models.CodeServerBusy)
	}
	if code != nil {
		if code == mysql.ErrorServerPortExist {
			result.ResponseSuccess(c, models.CodeConnectPortExist)
			return
		} else {
			result.ResponseErrorWithMsg(c, models.CodeServerBusy, "server不存在")
			return
		}
	}
}
func (sh *ServerHandler) Server_SelectByShowPort(c *gin.Context) {
	result := new(models.ResultData)
	show_port := c.Param("show_port")
	re, err := sh.mysql.Server_SelectByShowPort(show_port)
	if err != nil {
		zap.L().Error("sh.mysql.Server_ConnectCan() with fail", zap.Error(err))
		result.ResponseError(c, models.CodeServerBusy)
		return
	}
	result.ResponseSuccess(c, re)
}
