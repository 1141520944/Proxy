package mysql

import (
	"fmt"

	"github.com/1141520944/proxy/server/pkg/models"
	"github.com/1141520944/proxy/server/pkg/setting"
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
	d.AutoMigrate(&models.Table_Server{}, &models.Table_Linux{}, &models.Table_user{})
	db = d
	//纠正
	if ok, re := checkServerIsExist(); ok {
		clearServerData(re)
	}
	return nil
}
func New() *Mysql {
	return &Mysql{DB: db}
}

func checkServerIsExist() (bool, []*models.Table_Server) {
	var count int64
	var re []*models.Table_Server
	db.Model(&models.Table_Server{}).Count(&count)
	if count > 0 {
		db.Model(&models.Table_Server{}).Find(&re)
		return true, re
	} else {
		return false, nil
	}
}

func clearServerData(val []*models.Table_Server) error {
	for _, v := range val {
		if err := db.Unscoped().Model(&models.Table_Server{}).Where("id=?", v.ID).Delete(&models.Table_Server{}).Error; err != nil {
			return err
		}
	}
	return nil
}
