package repo

import "github.com/1141520944/proxy/server/pkg/models"

type Linux_mysql interface {
	Linux_GetInformation() (*models.Table_Linux, error)
}
