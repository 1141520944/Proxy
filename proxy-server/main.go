package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/1141520944/proxy/api"
	"github.com/1141520944/proxy/common/run"
	validateG "github.com/1141520944/proxy/common/validate"
	"github.com/1141520944/proxy/dao/mysql"
	"github.com/1141520944/proxy/dao/redis"
	"github.com/1141520944/proxy/pkg/snowflake"

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
	//初始redis
	if err := redis.Init(); err != nil {
		zap.L().Error("redis.Init() failed;err: %v\n", zap.Error(err))
		return
	}
	//初始redismysql
	if err := mysql.Init(); err != nil {
		zap.L().Error("mysql.Init() failed;err: %v\n", zap.Error(err))
		return
	}
	//初始snowflake
	if err := snowflake.Init(setting.Conf.StartTime, setting.Conf.MachineID); err != nil {
		zap.L().Error("snowflake.Init() failed;err: %v\n", zap.Error(err))
		return
	}
	//gin
	e := gin.Default()
	r := router.Init(e)
	server := &http.Server{Addr: fmt.Sprintf(":%d", setting.Conf.Port), Handler: r}
	run.Run(server)
}
