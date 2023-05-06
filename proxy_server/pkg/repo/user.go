package repo

import "github.com/1141520944/proxy/server/pkg/models"

type User_mysql interface {
	User_SelectAll(page, size int) ([]*models.Table_user, int64, error)

	User_SelectByIDOne(id int64) ([]*models.Table_user, error)
	User_SelectByUsernameOne(val string) (*models.Table_user, error)

	User_DeleteByIDOne(id int64) error

	User_InsertOne(val *models.Table_user) error

	User_UpdateOne(val *models.Table_user) error

	User_CheckByUsername(val string) (error, error)
	User_CheckByUserID(userid int64) (error, error)
}
