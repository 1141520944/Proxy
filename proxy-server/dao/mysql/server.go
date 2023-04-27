package mysql

import (
	"github.com/1141520944/proxy/common/models"
)

func (m *Mysql) Server_InsertOne(val *models.Table_Server) error {
	return m.DB.Model(&models.Table_Server{}).Create(val).Error
}
func (m *Mysql) Server_UpdateOne(val *models.Table_Server) error {
	return m.DB.Model(&models.Table_Server{}).Where("server_id=?", val.ServerID).Updates(val).Error
}
func (m *Mysql) Server_SelectByIDOne(id int64) (*models.Table_Server, error) {
	re := new(models.Table_Server)
	err := m.DB.Model(&models.Table_Server{}).Where("server_id=?", id).First(re).Error
	return re, err
}
func (m *Mysql) Server_SelectAll(page, size int) ([]*models.Table_Server, int64, error) {
	var re []*models.Table_Server
	var count int64
	err := m.DB.Model(&models.Table_Server{}).Count(&count).Limit(size).Offset((page - 1) * size).Find(&re).Error
	return re, count, err
}
func (m *Mysql) Server_DeleteByIDOne(id int64) error {
	return m.DB.Model(&models.Table_Server{}).Unscoped().Where("server_id=?", id).Delete(&models.Table_Server{}).Error
}
func (m *Mysql) Server_ConnectCan(val *models.Table_Server) (err error, code error) {
	var Count int64
	err = m.DB.Model(&models.Table_Server{}).Where("connect_port=?", val.ConnectPort).Count(&Count).Error
	if err != nil {
		return
	} else if Count > 0 {
		code = ErrorConnectPortExist
		return
	}
	err = m.DB.Model(&models.Table_Server{}).Where("show_port=?", val.ShowPort).Count(&Count).Error
	if err != nil {
		return
	} else if Count > 0 {
		code = ErrorServerPortExist
		return
	}
	err = m.DB.Model(&models.Table_Server{}).Where("local_project_port=?", val.LocalProjectPort).Count(&Count).Error
	if err != nil {
		return
	} else if Count > 0 {
		code = ErrorLocationPortExist
		return
	}
	return
}
func (m *Mysql) Server_SelectByConnectPort(connect string) (*models.Table_Server, error) {
	re := new(models.Table_Server)
	err := m.DB.Model(&models.Table_Server{}).Where("connect_port=?", connect).First(re).Error
	return re, err
}
func (m *Mysql) Server_SelectByShowPort(show_port string) (*models.Table_Server, error) {
	re := new(models.Table_Server)
	err := m.DB.Model(&models.Table_Server{}).Where("show_port=?", show_port).First(re).Error
	return re, err
}
func (m *Mysql) Server_checkShowPort(show_port string) (err error, code error) {
	var Count int64
	err = m.DB.Model(&models.Table_Server{}).Where("show_port=?", show_port).Count(&Count).Error
	if err != nil {
		return
	} else if Count > 0 {
		code = ErrorServerPortExist
		return
	} else {
		code = ErrorServerPortNoExist
		return
	}
}
func (m *Mysql) Server_DeleteByIDShowPort(show_port string) error {
	return m.DB.Model(&models.Table_Server{}).Unscoped().Where("show_port=?", show_port).Delete(&models.Table_Server{}).Error
}
