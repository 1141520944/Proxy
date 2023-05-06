package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/1141520944/proxy/client/pkg/models"
	validateG "github.com/1141520944/proxy/client/pkg/util/validate"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"go.uber.org/zap"
)

// var src string = "http://127.0.0.1:8088/server/check"

// ServerInit 服务端建立连接
func (ch *ClientHandler) ServerCheckAndInit(c *gin.Context) {
	result := new(models.ResultData)
	client := new(models.ClientInitRequest)
	serverInit := make(chan struct{})
	serverInitError := make(chan struct{})
	remoteIp := ch.remoteIp
	remotePort := ch.remotePort
	fmt.Println("进入请求")
	fmt.Printf("remoteIp: %v\n", remoteIp)
	fmt.Printf("remotePort: %v\n", remotePort)
	// initSuccess := make(chan struct{})
	if err := c.ShouldBindJSON(client); err != nil {
		zap.L().Error("ServerHandler with invalid fail", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			result.ResponseError(c, models.CodeServerBusy)
			return
		} else {
			result.ResponseErrorWithMsg(c, models.CodeInvalidParam, validateG.RemoveTopStruct(errs.Translate(validateG.Trans)))
			return
		}
	}
	// id, _ := strconv.ParseInt(client.UserID, 10, 64)
	server := &models.ServerInitRequest{
		ServerName:       client.ServerName,
		ShowPort:         client.ShowPort,
		ConnectPort:      client.ConnectPort,
		LocalProjectPort: client.LocalProjectPort,
		ServerState:      true,
		Password:         client.Password,
		UserID:           client.UserID,
	}
	go func() {
		src := fmt.Sprintf("http://%s:%s/server/check", remoteIp, remotePort)
		serverByte, _ := json.Marshal(&server)
		var body = bytes.NewReader(serverByte)
		response, err := http.Post(src, "application/json; charset=utf-8", body)
		if err != nil || response.StatusCode != http.StatusOK {
			c.Status(http.StatusServiceUnavailable)
			return
		}
		r, _ := ioutil.ReadAll(response.Body)
		//返回结果
		json.Unmarshal(r, result)
		if result.Code != models.CodeSuccess {
			serverInitError <- struct{}{}
		} else {
			//创建serverInit
			serverInit <- struct{}{}
		}
	}()
	select {
	case <-serverInitError:
		result.ResponseErrorWithMsg(c, models.CodeServerBusy, result.Msg)
	case <-serverInit:
		ch.ServerInit(server, c)
		// result.ResponseSuccess(c, models.CodeSuccess)
	}
}

// ServerInit 请求建立
func (ch *ClientHandler) ServerInit(server *models.ServerInitRequest, c *gin.Context) {
	result := new(models.ResultData)
	goroutineDie := make(chan struct{})
	serverInitSign := make(chan struct{})
	remoteIp := ch.remoteIp
	remotePort := ch.remotePort
	go func() {
		src := fmt.Sprintf("http://%s:%s/server/init", remoteIp, remotePort)
		serverByte, _ := json.Marshal(&server)
		var body = bytes.NewReader(serverByte)
		c2 := http.Client{Timeout: time.Second * 10}
		c2.Post(src, "application/json; charset=utf-8", body)
	}()
	go func() {
		//查看数据库字段--依据show_port判断创建server
		src := fmt.Sprintf("http://%s:%s/server/check/show_port/%s", remoteIp, remotePort, server.ShowPort)
		// log.Println(src)
		//结束协程
		for i := 0; i < 3; i++ {
			go func() {
				response, err := http.Get(src)
				if err != nil || response.StatusCode != http.StatusOK {
					zap.L().Error("http.Get() with fail", zap.Error(err))
					return
				}
				r, _ := ioutil.ReadAll(response.Body)
				//返回结果
				json.Unmarshal(r, result)
				// fmt.Printf("%dresult: %+v\n", i, result)
				if result.Code == models.CodeSuccess {
					//成功查询
					serverInitSign <- struct{}{}
					i = 2
					return
				}
			}()
			select {
			case <-goroutineDie:
			case <-time.After(time.Second):
				continue
			}
		}
	}()
	select {
	case <-serverInitSign:
		//关闭协程
		goroutineDie <- struct{}{}
		result.ResponseSuccess(c, models.CodeSuccess)
	case <-time.After(time.Second * 5):
		//关闭协程
		goroutineDie <- struct{}{}
		result.ResponseError(c, models.CodeServerBusy)
	}
}
