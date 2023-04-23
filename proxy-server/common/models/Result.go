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
}

func (r CodeError) GetCode() string {
	v, ok := codeMsgMap[r]
	if !ok {
		return v
	}
	return v
}

type ResultData struct {
	Code CodeError   `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func (ResultData) ResponseSuccess(c *gin.Context, val interface{}) {
	resp := &ResultData{
		Code: CodeSuccess,
		Msg:  CodeSuccess.GetCode(),
		Data: val,
	}
	c.JSON(http.StatusOK, resp)
}

func (ResultData) ResponseError(c *gin.Context, code CodeError) {
	resp := &ResultData{
		Code: code,
		Msg:  code.GetCode(),
		Data: nil,
	}
	c.JSON(http.StatusOK, resp)
}
func (ResultData) ResponseErrorWithMsg(c *gin.Context, code CodeError, msg interface{}) {
	resp := &ResultData{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
	c.JSON(http.StatusOK, resp)
}
