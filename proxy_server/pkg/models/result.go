package models

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CodeError int

const (
	CodeSuccess CodeError = 1000 + iota
	CodeInvalidParam
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeServerBusy

	CodeNeedLogin
	CodeInvalidToken

	CodeConnectPortExist
	CodeShowPortExist
	CodeLocationPortExist

	CodeShowPortNoExist

	CodeUsernameExist
	CodeUsernameNoExist
)

var codeMsgMap = map[CodeError]string{
	CodeSuccess:         "success",
	CodeInvalidParam:    "请求参数错误",
	CodeUserExist:       "用户名已存在",
	CodeUserNotExist:    "用户名不存在",
	CodeInvalidPassword: "密码错误",
	CodeServerBusy:      "服务器繁忙",

	CodeNeedLogin:    "需要登录",
	CodeInvalidToken: "无效的token",

	CodeConnectPortExist:  "服务端connect的连接已存在",
	CodeShowPortExist:     "服务端show的连接已存在",
	CodeLocationPortExist: "服务端location的连接已存在",

	CodeShowPortNoExist: "服务端show的连接不存在",

	CodeUsernameExist:   "该用户名已存在",
	CodeUsernameNoExist: "该用户名不存在",
}

func (r CodeError) GetCode() string {
	v, ok := codeMsgMap[r]
	if !ok {
		return v
	}
	return v
}

type ResultData struct {
	Code  CodeError   `json:"code"`
	Count int         `json:"count"`
	Msg   interface{} `json:"msg"`
	Data  interface{} `json:"data"`
}

func (ResultData) ResponseSuccessLayui(c *gin.Context, val interface{}, count int) {
	resp := &ResultData{
		Code:  0,
		Msg:   CodeSuccess.GetCode(),
		Count: count,
		Data:  val,
	}
	c.JSON(http.StatusOK, resp)
}
func (ResultData) ResponseSuccess(c *gin.Context, val interface{}, count int) {
	resp := &ResultData{
		Code:  CodeSuccess,
		Msg:   CodeSuccess.GetCode(),
		Count: count,
		Data:  val,
	}
	c.JSON(http.StatusOK, resp)
}

func (ResultData) ResponseError(c *gin.Context, code CodeError) {
	resp := &ResultData{
		Code:  code,
		Msg:   code.GetCode(),
		Count: 0,
		Data:  nil,
	}
	c.JSON(http.StatusOK, resp)
}
func (ResultData) ResponseErrorWithMsg(c *gin.Context, code CodeError, msg interface{}) {
	resp := &ResultData{
		Code:  code,
		Msg:   msg,
		Count: 0,
		Data:  nil,
	}
	c.JSON(http.StatusOK, resp)
}
