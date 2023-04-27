package mysql

import "errors"

var (
	ErrorServerPortExist   = errors.New("服务端show的连接已存在")
	ErrorServerPortNoExist = errors.New("服务端show的连接不存在")
	ErrorConnectPortExist  = errors.New("服务端connect的连接已存在")
	ErrorLocationPortExist = errors.New("服务端location的连接已存在")
	ErrorServerPassword    = errors.New("密码错误")
	ErrorInvalidID         = errors.New("无效的ID")
)
