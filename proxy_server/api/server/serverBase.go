package server

import (
	"strconv"

	"github.com/1141520944/proxy/server/dao/mysql"
	"github.com/1141520944/proxy/server/pkg/models"
	"github.com/1141520944/proxy/server/pkg/util/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// ServerByShowPort 检测show_port判断服务是否存在
func (sh *ServerHandler) ServerCheckByShowPort(c *gin.Context) {
	result := new(models.ResultData)
	show_port := c.Param("show_port")
	err, code := sh.mysql.Server_checkShowPort(show_port)
	if err != nil {
		zap.L().Error("sh.mysql.Server_checkShowPort() with fail-err", zap.Error(err))
		result.ResponseError(c, models.CodeServerBusy)
		return
	}
	if code != nil {
		if code == mysql.ErrorShowPortExist {
			result.ResponseSuccess(c, models.CodeShowPortExist, 1)
			return
		} else {
			result.ResponseError(c, models.CodeShowPortNoExist)
			return
		}
	}
}

// Server_SelectByShowPort 查询服务 by - show_port
func (sh *ServerHandler) ServerSelectByShowPort(c *gin.Context) {
	result := new(models.ResultData)
	show_port := c.Param("show_port")
	re, err := sh.mysql.Server_SelectByShowPort(show_port)
	if err != nil {
		zap.L().Error("sh.mysql.Server_SelectByShowPort() with fail", zap.Error(err))
		result.ResponseError(c, models.CodeServerBusy)
		return
	}
	result.ResponseSuccess(c, re, 1)
}

func (sh *ServerHandler) ServerSelectByUserID(c *gin.Context) {
	uidStr := c.Param("uid")
	result := new(models.ResultData)
	uid, err := strconv.ParseInt(uidStr, 10, 64)
	if err != nil {
		result.ResponseError(c, models.CodeInvalidParam)
		return
	}
	page, page_size, err := request.GetPageInfo(c)
	if err != nil {
		zap.L().Error("request.GetPageInfo() failed", zap.Error(err))
		result.ResponseError(c, models.CodeInvalidParam)
		return
	}
	re, count, err2 := sh.mysql.Server_SelectAllByUserID(int(page), int(page_size), uid)
	if err2 != nil {
		zap.L().Error("sh.mysql.Server_SelectAllByUserID() failed", zap.Error(err))
		result.ResponseError(c, models.CodeServerBusy)
		return
	}
	servers := make([]*models.ServerResponse, 0, len(re))
	for _, v := range re {
		server := &models.ServerResponse{
			ServerName:       v.ServerName,
			ShowPort:         v.ShowPort,
			ConnectPort:      v.ConnectPort,
			LocalProjectPort: v.LocalProjectPort,
			Password:         v.ServerPassword,
			UserID:           strconv.FormatInt(v.UserID, 10),
			ServerID:         strconv.FormatInt(v.ServerID, 10),
			CreateTime:       v.CreatedAt,
			State:            v.ServerState,
		}
		servers = append(servers, server)
	}
	result.ResponseSuccessLayui(c, servers, int(count))
}
func (sh *ServerHandler) ServerSelectAll(c *gin.Context) {
	result := new(models.ResultData)
	page, page_size, err := request.GetPageInfo(c)
	if err != nil {
		zap.L().Error("request.GetPageInfo() failed", zap.Error(err))
		result.ResponseError(c, models.CodeInvalidParam)
		return
	}
	re, count, err2 := sh.mysql.Server_SelectAll(int(page), int(page_size))
	if err2 != nil {
		zap.L().Error("sh.mysql.Server_SelectAll() failed", zap.Error(err))
		result.ResponseError(c, models.CodeServerBusy)
		return
	}
	servers := make([]*models.ServerResponse, 0, len(re))
	for _, v := range re {
		server := &models.ServerResponse{
			ServerName:       v.ServerName,
			ShowPort:         v.ShowPort,
			ConnectPort:      v.ConnectPort,
			LocalProjectPort: v.LocalProjectPort,
			Password:         v.ServerPassword,
			UserID:           strconv.FormatInt(v.UserID, 10),
			ServerID:         strconv.FormatInt(v.ServerID, 10),
			CreateTime:       v.CreatedAt,
			State:            v.ServerState,
		}
		servers = append(servers, server)
	}
	result.ResponseSuccessLayui(c, servers, int(count))
}
