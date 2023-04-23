package test

import (
	"net/http"

	"github.com/1141520944/proxy/router"
	"github.com/gin-gonic/gin"
)

type TestRouter struct {
}

func init() {
	router.AddRouters(new(TestRouter))
}
func (*TestRouter) Router(r *gin.Engine) {
	t := r.Group("/test")
	t.GET("/ping", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"连接信号": "success"}) })
}
