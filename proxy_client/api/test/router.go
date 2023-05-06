package test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/1141520944/proxy/client/pkg/models"
	"github.com/1141520944/proxy/client/pkg/router"
	"github.com/gin-gonic/gin"
)

type TestRouter struct{}

func init() {
	router.AddRouters(new(TestRouter))
}
func (*TestRouter) Router(r *gin.Engine) {
	test := r.Group("test")
	test.GET("ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"连接": "success",
		})
	})
	test.GET("http", func(c *gin.Context) {
		sspr := new(models.ServerShowPortResponse)
		result := new(models.ResultData)
		src := "http://127.0.0.1:8088/server/select/show_port/8888"
		response, err := http.Get(src)
		if err != nil || response.StatusCode != http.StatusOK {
			c.Status(http.StatusServiceUnavailable)
			return
		}
		r, _ := ioutil.ReadAll(response.Body)
		//返回结果
		json.Unmarshal(r, sspr)
		fmt.Printf("sspr: %+v\n", sspr)
		result.ResponseSuccess(c, sspr)
	})

}
