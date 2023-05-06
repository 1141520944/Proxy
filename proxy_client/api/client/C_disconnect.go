package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/1141520944/proxy/client/pkg/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (ch *ClientHandler) ClientDisConnect(c *gin.Context) {
	ClientConnectDie <- struct{}{}
	result := new(models.ResultData)
	remoteIp := ch.remoteIp
	remotePort := ch.remotePort
	show_port := c.Param("show_port")
	src := fmt.Sprintf("http://%s:%s/server/check/show_port/%s", remoteIp, remotePort, show_port)
	response, err := http.Get(src)
	if err != nil || response.StatusCode != http.StatusOK {
		zap.L().Error("http.Get() with fail", zap.Error(err))
		return
	}
	r, _ := ioutil.ReadAll(response.Body)
	//返回结果
	json.Unmarshal(r, result)
	fmt.Printf("result: %+v\n", result)
	if result.Code != models.CodeServerBusy {
		//不存在--连接断开
		result.ResponseSuccess(c, "连接断开")
		return
	} else {
		result.ResponseError(c, models.CodeServerBusy)
		return
	}
}
