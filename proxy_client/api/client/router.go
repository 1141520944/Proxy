package client

import (
	"github.com/1141520944/proxy/client/pkg/router"
	"github.com/1141520944/proxy/client/pkg/setting"
	"github.com/gin-gonic/gin"
)

type ClientRouter struct {
}
type ClientHandler struct {
	remoteIp   string
	remotePort string
}

func New_clientHandler() *ClientHandler {
	return &ClientHandler{
		remoteIp:   setting.Conf.RemoteIp,
		remotePort: setting.Conf.RemotePort,
	}
}
func init() {
	router.AddRouters(new(ClientRouter))
}
func (*ClientRouter) Router(r *gin.Engine) {
	client := r.Group("client")
	clientHandler := New_clientHandler()
	client.POST("/server/init", clientHandler.ServerCheckAndInit)
	client.GET("/server/connect/:show_port", clientHandler.ClientConnectServer)
	client.GET("/server/disconnect/:show_port", clientHandler.ClientDisConnect)
}
