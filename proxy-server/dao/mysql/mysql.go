package mysql

import "gorm.io/gorm"

type Mysql struct {
	DB *gorm.DB
}

func New(db *gorm.DB) *Mysql {
	return &Mysql{DB: db}
}
