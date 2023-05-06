package mysql

import "github.com/1141520944/proxy/server/pkg/models"

func (m *Mysql) Linux_GetInformation() (*models.Table_Linux, error) {
	re := new(models.Table_Linux)
	err := m.DB.Model(&models.Table_Linux{}).First(re).Error
	return re, err
}
