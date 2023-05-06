package mysql

import "github.com/1141520944/proxy/server/pkg/models"

func (m *Mysql) User_SelectAll(page, size int) ([]*models.Table_user, int64, error) {
	var count int64
	var re []*models.Table_user
	err := m.DB.Model(&models.Table_user{}).Limit(size).Offset((page - 1) * size).Count(&count).Find(&re).Error
	return re, count, err
}

func (m *Mysql) User_SelectByIDOne(id int64) ([]*models.Table_user, error) {
	var re []*models.Table_user
	err := m.DB.Model(&models.Table_user{}).Where("user_id=?", id).Find(&re).Error
	return re, err
}

func (m *Mysql) User_DeleteByIDOne(id int64) error {
	return m.DB.Model(&models.Table_user{}).Unscoped().Where("user_id=?", id).Delete(&models.Table_user{}).Error
}

func (m *Mysql) User_InsertOne(val *models.Table_user) error {
	return m.DB.Model(&models.Table_user{}).Create(val).Error
}

func (m *Mysql) User_UpdateOne(val *models.Table_user) error {
	return m.DB.Model(&models.Table_user{}).Where("user_id=?", val.UserID).Updates(val).Error
}

func (m *Mysql) User_SelectByUsernameOne(val string) (*models.Table_user, error) {
	re := new(models.Table_user)
	err := m.DB.Model(&models.Table_user{}).Where("username=?", val).First(re).Error
	return re, err
}

func (m *Mysql) User_CheckByUsername(val string) (error, error) {
	var count int64
	var code error
	err := m.DB.Model(&models.Table_user{}).Where("username=?", val).Count(&count).Error
	if count > 0 {
		code = ErrorUsernameExist
		return code, err
	} else {
		code = ErrorUsernameNoExist
		return code, err
	}
}
func (m *Mysql) User_CheckByUserID(userid int64) (error, error) {
	var count int64
	var code error
	err := m.DB.Model(&models.Table_user{}).Where("user_id=?", userid).Count(&count).Error
	if count > 0 {
		code = ErrorUsernameExist
		return code, err
	} else {
		code = ErrorUsernameNoExist
		return code, err
	}
}
