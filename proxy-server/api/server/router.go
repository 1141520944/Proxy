package server

import (
	"github.com/1141520944/proxy/router"
	"github.com/gin-gonic/gin"
)

type ServerRouter struct {
}

func init() {
	router.AddRouters(new(ServerRouter))
}
func (*ServerRouter) Router(r *gin.Engine) {
	server := r.Group("/server")
	serverHandler := New_handler()
	server.POST("/init", serverHandler.ServerInit)
	server.POST("/check", serverHandler.ServerCheckIsExist)
	server.GET("/check/show_port/:show_port", serverHandler.ServerCheckByShowPort)
	server.GET("/select/show_port/:show_port", serverHandler.Server_SelectByShowPort)
}
