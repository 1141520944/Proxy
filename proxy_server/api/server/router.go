package server

import (
	"github.com/1141520944/proxy/server/api/user"
	"github.com/1141520944/proxy/server/dao/mysql"
	"github.com/1141520944/proxy/server/pkg/repo"
	"github.com/1141520944/proxy/server/pkg/router"
	"github.com/1141520944/proxy/server/pkg/setting"
	"github.com/gin-gonic/gin"
)

type ServerRouter struct {
}
type ServerHandler struct {
	mysql        repo.Server_mysql
	user         user.UserHandler
	maxExistTime int64
}

func New_handler() *ServerHandler {
	mysql := mysql.New()
	return &ServerHandler{
		mysql:        mysql,
		user:         *user.New_handler(),
		maxExistTime: setting.Conf.MaxExistTime,
	}
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
	server.GET("/select/show_port/:show_port", serverHandler.ServerSelectByShowPort)
	server.GET("/close/:show_port", serverHandler.ServerClose)
	server.GET("/select/all/:uid", serverHandler.ServerSelectByUserID)
	server.GET("/select/all", serverHandler.ServerSelectAll)
}
