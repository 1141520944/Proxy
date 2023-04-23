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
	server.POST("/init", ServerInit)
}
