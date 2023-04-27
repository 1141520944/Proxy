package server

import (
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	"github.com/1141520944/proxy/common/models"
	validateG "github.com/1141520944/proxy/common/validate"
	"github.com/1141520944/proxy/dao/mysql"
	"github.com/1141520944/proxy/pkg/snowflake"
	"github.com/1141520944/proxy/repo"
	"github.com/1141520944/proxy/util"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"go.uber.org/zap"
)

type ServerHandler struct {
	mysql repo.Server_mysql
}

func New_handler() *ServerHandler {
	return &ServerHandler{
		mysql: mysql.New(),
	}
}

// 启动一个server
func (sh *ServerHandler) ServerInit(c *gin.Context) {
	serverDie := make(chan struct{})
	ClienDie := make(chan struct{})
	goroutineDie := make(chan struct{})
	serverCreate := make(chan struct{})
	//post 请求
	sr := new(models.ServerRequest)
	result := new(models.ResultData)
	if err := c.ShouldBindJSON(sr); err != nil {
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
	go func() {
		//服务创建
		select {
		case <-serverCreate:
			//创建server
			server := &models.Table_Server{
				ServerID:         snowflake.GenID(),
				ServerName:       sr.ServerName,
				ServerPassword:   sr.Password,
				ShowPort:         sr.ShowPort,
				ConnectPort:      sr.ConnectPort,
				LocalProjectPort: sr.LocalProjectPort,
				ServerState:      true,
			}
			if err := sh.mysql.Server_InsertOne(server); err != nil {
				zap.L().Error("sh.mysql.Server_InsertOne() with fail", zap.Error(err))
			}
		case <-time.After(time.Second * 10):
			result.ResponseErrorWithMsg(c, models.CodeServerBusy, "连接失败")
		}
	}()
	//http连接
	chSession := make(chan net.Conn, 100)
	// go listenC(sr.ServerPort, chSession, onClientConnect, sr, ConnDie)
	// go listenG(sr.ShowPort, chSession, onUserConnect, sr, ClienDie, goroutineDie)
	if nil != listenC(sr.ConnectPort, chSession, onClientConnect, sr, ClienDie, serverDie, goroutineDie, serverCreate) {
		return
	}
	if nil != listenG(sr.ShowPort, chSession, onUserConnect, sr, ClienDie, goroutineDie) {
		return
	}
	select {
	case <-serverDie:
		//删除数据库
		if err := sh.mysql.Server_DeleteByIDShowPort(sr.ShowPort); err != nil {
			zap.L().Error("sh.mysql.Server_DeleteByIDShowPort() with fail", zap.Error(err))
		}
		// result.ResponseSuccess(c, "连接结束")
	case <-time.After(time.Hour * 3):
	}
}

// 执行Handler的方法
type OnConnectFunc func(net.Conn, chan net.Conn, *models.ServerRequest, chan struct{})

// 监测是否有客户端连接
func listenC(port string, chSession chan net.Conn, onConnect OnConnectFunc, re *models.ServerRequest, ClienDie chan struct{}, serverDie chan struct{}, goroutineDie chan struct{}, serverCreate chan struct{}) error {
	l, err := net.Listen("tcp", net.JoinHostPort("0.0.0.0", port))
	if err != nil {
		zap.L().Error("CAN'T LISTEN: ", zap.Error(err))
		return err
	}
	zap.L().Info(fmt.Sprintf("listen port:%s", port))
	serverCreate <- struct{}{}
	go func() {
		defer l.Close()
		go func() {
			select {
			case <-ClienDie:
				l.Close()
				zap.L().Error(fmt.Sprintf("l.Close(): %s", port), zap.Error(err))
				goroutineDie <- struct{}{}
				time.Sleep(time.Second)
				serverDie <- struct{}{}
			case <-time.After(time.Hour * 3):
			}
		}()
		for {
			conn, err := l.Accept()
			if err != nil {
				zap.L().Info("Can't Accept: -listener 断开", zap.Error(err))
				break
			}
			go onConnect(conn, chSession, re, ClienDie)
		}
	}()
	return nil
}
func listenG(port string, chSession chan net.Conn, onConnect OnConnectFunc, re *models.ServerRequest, ClienDie chan struct{}, goroutineDie chan struct{}) error {
	l, err := net.Listen("tcp", net.JoinHostPort("0.0.0.0", port))
	if err != nil {
		zap.L().Error("CAN'T LISTEN: ", zap.Error(err))
		return err
	}
	zap.L().Info(fmt.Sprintf("listen port:%s", port))
	go func() {
		go func() {
			select {
			case <-goroutineDie:
				l.Close()
				zap.L().Error(fmt.Sprintf("l.Close(): %s", port), zap.Error(err))
			case <-time.After(time.Hour * 3):
			}
		}()
		for {
			conn, err := l.Accept()
			if err != nil {
				zap.L().Error("Can't Accept: %s", zap.Error(err))
				break
			}
			go onConnect(conn, chSession, re, ClienDie)
		}
	}()
	return nil
}

// 客户端listern创建监听链接后执行的操作
func onClientConnect(conn net.Conn, chSession chan net.Conn, re *models.ServerRequest, ClienDie chan struct{}) {
	strConn := util.Conn2Str(conn)
	zap.L().Info(fmt.Sprintf("Proxy Client Connect:%s", strConn))
	//设置read的截至时间
	conn.SetReadDeadline(time.Now().Add(5 * time.Second))
	//读取字符
	msg, err := util.ReadString(conn)
	conn.SetReadDeadline(time.Time{})
	// println("Read:", string(buf[0:n]))
	if err != nil {
		zap.L().Error("util.ReadString() with fail", zap.Error(err))
		conn.Close()
		return
	}
	msgs := strings.Split(msg, "\n")
	pwd := msgs[0]

	//--查password
	//
	if re.Password != pwd {
		util.CloseConn(conn)
		return
	}
	token := msgs[1]
	log.Println("token=", token)
	if token == util.C2P_CONNECT {
		//内网服务器启动时连接代理，建立长连接
		clientConnect(conn, re, ClienDie)
		return
	} else if token == util.C2P_SESSION {
		//为客户端的单次连接请求建立一个临时的"内网服务器<->代理"的连接
		initUserSession(conn, chSession)
		return
	}

}

// 代理客户端连接
var _clientProxy net.Conn = nil

// 处理golocproxy client的连接
func clientConnect(conn net.Conn, re *models.ServerRequest, ConnDie chan struct{}) {
	defer func() {
		ConnDie <- struct{}{}
		_clientProxy = nil
		util.CloseConn(conn)
	}()
	if _clientProxy != nil {
		log.Println("_clientProxy is nil")
		util.WriteString(conn, "SERVICE EXIST")
		util.CloseConn(conn)
		return
	}
	zap.L().Info("REG SERVICE")
	_clientProxy = conn
	for {
		_, err := util.ReadString(_clientProxy)
		if err != nil {
			zap.L().Info("UNREG SERVICE---C")
			break
		}
	}
}

func initUserSession(conn net.Conn, chSession chan net.Conn) {
	chSession <- conn
}

// 处理最终用户的连接
func onUserConnect(conn net.Conn, chSession chan net.Conn, re *models.ServerRequest, ConnDie chan struct{}) {
	if _clientProxy == nil {
		conn.Write([]byte("NO SERVICE"))
		util.CloseConn(conn)
		return
	}
	_, err := util.WriteString(_clientProxy, util.P2C_NEW_SESSION)
	if err != nil {
		conn.Write([]byte("SERVICE FAIL"))
		util.CloseConn(conn)
		return
	}
	connSession := recvSession(chSession) // := <-chSession
	if connSession == nil {
		util.CloseConn(conn)
		return
	}
	zap.L().Info("Transfer...")
	go util.CopyFromTo(conn, connSession, nil)
	go util.CopyFromTo(connSession, conn, nil)
}

// 加入超时
func recvSession(ch chan net.Conn) net.Conn {
	var conn net.Conn = nil
	select {
	case conn = <-ch:
	case <-time.After(time.Second * 5):
		conn = nil
	}
	return conn
}
