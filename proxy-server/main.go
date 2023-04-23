package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/1141520944/proxy/api"
	"github.com/1141520944/proxy/common/run"
	validateG "github.com/1141520944/proxy/common/validate"
	"github.com/1141520944/proxy/logger"
	"github.com/1141520944/proxy/router"
	"github.com/1141520944/proxy/setting"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	if err := setting.Init(); err != nil {
		log.Fatal("setting.Init()", err)
	}
	if err := logger.Init(setting.Conf.LogConfig, setting.Conf.Mode); err != nil {
		log.Fatal("logger.Init()", err)
	}
	defer zap.L().Sync()
	//注册翻译器
	if trans_err := validateG.InitTrans("zh"); trans_err != nil {
		zap.L().Error("init() validator failed;err: %v\n", zap.Error(trans_err))
		return
	}
	//gin
	e := gin.Default()
	r := router.Init(e)
	server := &http.Server{Addr: fmt.Sprintf(":%d", setting.Conf.Port), Handler: r}
	run.Run(server)
}
