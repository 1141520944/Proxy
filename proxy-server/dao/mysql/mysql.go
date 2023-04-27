package mysql

import (
	"fmt"

	"github.com/1141520944/proxy/common/models"
	"github.com/1141520944/proxy/setting"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

type Mysql struct {
	DB *gorm.DB
}

func Init() error {
	cfg := setting.Conf.MySQLConfig
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DbName,
	)
	//建立连接
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		zap.L().Error("connect DB failed", zap.Error(err))
		return err
	}
	d.AutoMigrate(&models.Table_Server{})
	db = d
	return nil
}
func New() *Mysql {
	return &Mysql{DB: db}
}
