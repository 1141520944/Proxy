package test

import (
	"net/http"

	"github.com/1141520944/proxy/client/router"
	"github.com/gin-gonic/gin"
)

type TestRouter struct{}

func init() {
	router.AddRouters(new(TestRouter))
}
func (*TestRouter) Router(r *gin.Engine) {
	test := r.Group("test")
	test.GET("test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"连接": "success",
		})
	})

}
