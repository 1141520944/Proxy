package server

import (
	"github.com/1141520944/proxy/server/dao/mysql"
	"github.com/1141520944/proxy/server/pkg/models"
	validateG "github.com/1141520944/proxy/server/pkg/util/validate"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// ServerCheckIsExist 检查服务是否存在--创建端口
func (sh *ServerHandler) ServerCheckIsExist(c *gin.Context) {
	sr := new(models.ServerRequest)
	result := new(models.ResultData)
	if err := c.ShouldBindJSON(sr); err != nil {
		zap.L().Error("ServerCheckIsExist with invalid fail", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			result.ResponseError(c, models.CodeServerBusy)
			return
		} else {
			result.ResponseErrorWithMsg(c, models.CodeInvalidParam, validateG.RemoveTopStruct(errs.Translate(validateG.Trans)))
			return
		}
	}
	server := &models.Table_Server{
		ServerName:       sr.ServerName,
		ServerPassword:   sr.Password,
		ShowPort:         sr.ShowPort,
		ConnectPort:      sr.ConnectPort,
		LocalProjectPort: sr.LocalProjectPort,
		ServerState:      true,
	}
	err, code := sh.mysql.Server_ConnectCan(server)
	if err != nil {
		zap.L().Error("sh.mysql.Server_ConnectCan() with fail-err", zap.Error(err))
		result.ResponseError(c, models.CodeServerBusy)
		return
	}
	if code != nil {
		zap.L().Error("sh.mysql.Server_ConnectCan() with fail-code", zap.Error(code))
		if code == mysql.ErrorConnectPortExist {
			result.ResponseError(c, models.CodeConnectPortExist)
			return
		} else if code == mysql.ErrorShowPortExist {
			result.ResponseError(c, models.CodeShowPortExist)
			return
		} else if code == mysql.ErrorLocationPortExist {
			result.ResponseError(c, models.CodeLocationPortExist)
			return
		}
	}
	result.ResponseSuccess(c, models.CodeSuccess, 1)
}
