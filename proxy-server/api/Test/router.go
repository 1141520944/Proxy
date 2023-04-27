package test

import (
	"context"
	"net/http"
	"time"

	"github.com/1141520944/proxy/common/models"
	validateG "github.com/1141520944/proxy/common/validate"
	"github.com/1141520944/proxy/dao/mysql"
	"github.com/1141520944/proxy/dao/redis"
	"github.com/1141520944/proxy/pkg/snowflake"
	"github.com/1141520944/proxy/router"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"go.uber.org/zap"
)

type TestRouter struct {
}

func init() {
	router.AddRouters(new(TestRouter))
}
func (*TestRouter) Router(r *gin.Engine) {
	t := r.Group("/test")
	t.GET("/ping", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"连接信号": "success"}) })
	t.POST("/serverinit", func(c *gin.Context) {
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
		result.ResponseSuccess(c, sr)
	})
	t.GET("/redis", func(c *gin.Context) {
		r := redis.New()
		result := new(models.ResultData)
		if err := r.Put_string(context.Background(), "测试", "值123", time.Minute*1); err != nil {
			zap.L().Error("rdis.Put_string with invalid fail", zap.Error(err))
			result.ResponseErrorWithMsg(c, models.CodeServerBusy, "put")
			return
		}
		re, err_string := r.Get_string(context.Background(), "测试")
		if err_string != nil {
			zap.L().Error("rdis.Get_string with invalid fail", zap.Error(err_string))
			result.ResponseErrorWithMsg(c, models.CodeServerBusy, "get")
			return
		}
		result.ResponseSuccess(c, re)
	})
	t.GET("/mysql", func(c *gin.Context) {
		m := mysql.New()
		result := new(models.ResultData)

		//insert
		id := snowflake.GenID()
		server := &models.Table_Server{
			ServerID:         id,
			ServerName:       "demo",
			ServerPassword:   "123456",
			ShowPort:         "8888",
			ConnectPort:      "9999",
			LocalProjectPort: "12000",
			ServerState:      true,
		}
		if err := m.Server_InsertOne(server); err != nil {
			zap.L().Error("mysql.Server_InsertOne with invalid fail", zap.Error(err))
			result.ResponseErrorWithMsg(c, models.CodeServerBusy, "insert")
			return
		}

		//selectone
		// i, _ := strconv.ParseInt("7700532210700288", 10, 64)
		// re, err := m.Server_SelectByIDOne(i)
		// if err != nil {
		// 	zap.L().Error("mysql.Server_SelectByIDOne with invalid fail", zap.Error(err))
		// 	result.ResponseErrorWithMsg(c, models.CodeServerBusy, "selectone")
		// 	return
		// }

		//selectall
		// re, i, err := m.Server_SelectAll(1, 10)
		// if err != nil {
		// 	zap.L().Error("mysql.Server_SelectAll with invalid fail", zap.Error(err))
		// 	result.ResponseErrorWithMsg(c, models.CodeServerBusy, "selectall")
		// 	return
		// }
		// result.ResponseSuccess(c, gin.H{
		// 	"data":  re,
		// 	"count": i,
		// })

		//update-delete
		result.ResponseSuccess(c, "success")
		return
	})
}
