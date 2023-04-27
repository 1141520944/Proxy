package redis

import (
	"context"
	"fmt"

	"github.com/1141520944/proxy/setting"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

var db *redis.Client

type Redis struct {
	DB *redis.Client
}

func Init() error {
	cfg := setting.Conf.RedisConfig
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			cfg.Host,
			cfg.Port,
		),
		Password: cfg.Password,
		DB:       cfg.DB,
		PoolSize: cfg.PoolSize,
	})
	//测试连接是否成功
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		zap.L().Error("connect DB failed", zap.Error(err))
		return err
	}
	db = client
	return nil
}
func New() *Redis {
	return &Redis{
		DB: db,
	}
}
