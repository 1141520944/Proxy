package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/1141520944/proxy/client/pkg/models"
	"github.com/1141520944/proxy/client/pkg/util"
	"github.com/gin-gonic/gin"
)

var ClientConnectDie = make(chan struct{})

// 建立连接
func (ch *ClientHandler) ClientConnectServer(c *gin.Context) {
	sspr := new(models.ServerShowPortResponse)
	result := new(models.ResultData)
	show_port := c.Param("show_port")
	clientConnectSuccess := make(chan struct{})
	remoteIp := ch.remoteIp
	remotePort := ch.remotePort
	//查找对应结果-server
	src := fmt.Sprintf("http://%s:%s/server/select/show_port/%s", remoteIp, remotePort, show_port)
	response, err := http.Get(src)
	if err != nil || response.StatusCode != http.StatusOK {
		c.Status(http.StatusServiceUnavailable)
		return
	}
	r, _ := ioutil.ReadAll(response.Body)
	//返回结果
	json.Unmarshal(r, sspr)

	client := &models.ClientConnect{
		ConnectAddr:      fmt.Sprintf("%s:%s", remoteIp, sspr.Data.Connect_port),
		LocalProjectAddr: fmt.Sprintf("%s:%s", "127.0.0.1", sspr.Data.Local_project_port),
		Password:         sspr.Data.Server_password,
	}

	go func() {
		connectServer(client, clientConnectSuccess)
		time.Sleep(time.Second * 3)
	}()
	select {
	case <-clientConnectSuccess:
		result.ResponseSuccess(c, "连接成功")
	case <-time.After(time.Second * 5):
		result.ResponseSuccess(c, "连接超时")
	}
}

// 远程连接服务器
func connectServer(cl *models.ClientConnect, success chan struct{}) {
	proxy, err := net.DialTimeout("tcp", cl.ConnectAddr, 5*time.Second)
	if err != nil {
		log.Println("CAN'T CONNECT:", cl.ConnectAddr, " err:", err)
		return
	}
	success <- struct{}{}
	// defer proxy.Close()
	go func() {
		select {
		case <-ClientConnectDie:
			proxy.Close()
		case <-time.After(time.Hour * 3):
			proxy.Close()
		}
	}()
	util.WriteString(proxy, cl.Password+"\n"+util.C2P_CONNECT)

	for {
		proxy.SetReadDeadline(time.Now().Add(2 * time.Second))
		msg, err := util.ReadString(proxy)
		//	proxy.SetReadDeadline(time.Time{})
		if err == nil {
			if msg == util.P2C_NEW_SESSION {
				go session(cl.ConnectAddr, cl.LocalProjectAddr, cl.Password)
			} else {
				log.Println(msg)
			}
		} else {
			if nerr, ok := err.(net.Error); ok && nerr.Timeout() {
				//log.Println("Timeout")
				proxy.SetWriteDeadline(time.Now().Add(2 * time.Second))
				_, werr := util.WriteString(proxy, util.C2P_KEEP_ALIVE) //send KeepAlive msg
				if werr != nil {
					log.Println("CAN'T WRITE, err:", werr)
					return
				}
				continue
			} else {
				log.Println("SERVER CLOSE, err:", err)
				return
			}
		}
		//time.Sleep(2*time.Second)
	}

}

// 客户端单次连接处理
func session(remote, local, pwd string) {
	log.Println("Create Session")
	rp, err := net.Dial("tcp", remote)
	if err != nil {
		log.Println("Can't' connect:", remote, " err:", err)
		return
	}
	//defer util.CloseConn(rp)
	util.WriteString(rp, pwd+"\n"+util.C2P_SESSION)
	lp, err := net.Dial("tcp", local)
	if err != nil {
		log.Println("Can't' connect:", local, " err:", err)
		rp.Close()
		return
	}
	go util.CopyFromTo(rp, lp, nil)
	go util.CopyFromTo(lp, rp, nil)
}
