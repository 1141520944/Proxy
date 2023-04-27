package repo

import "github.com/1141520944/proxy/common/models"

type Server_mysql interface {
	Server_InsertOne(val *models.Table_Server) error

	Server_DeleteByIDOne(id int64) error
	Server_DeleteByIDShowPort(show_port string) error

	Server_UpdateOne(val *models.Table_Server) error

	Server_SelectAll(page, size int) ([]*models.Table_Server, int64, error)

	Server_ConnectCan(val *models.Table_Server) (err error, code error)

	Server_SelectByIDOne(id int64) (*models.Table_Server, error)
	Server_SelectByShowPort(show_port string) (*models.Table_Server, error)
	Server_SelectByConnectPort(connect string) (*models.Table_Server, error)

	Server_checkShowPort(show_port string) (err error, code error)
}
