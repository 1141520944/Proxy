package models

import "gorm.io/gorm"

type Table_Server struct {
	gorm.Model
	ServerID         int64  `json:"server_id" gorm:"columm:server_id,primary"`
	ServerName       string `json:"server_name" gorm:"columm:server_name"`
	ServerPassword   string `json:"server_password" gorm:"columm:server_password"`
	ShowPort         string `json:"show_port" gorm:"columm:show_port"`
	ConnectPort      string `json:"connect_port" gorm:"columm:connect_port"`
	LocalProjectPort string `json:"local_project_port" gorm:"columm:local_project_port"`
	ServerState      bool   `json:"server_state" gorm:"columm:server_state"`
}

func (Table_Server) TableName() string {
	return "server"
}

// type Table_Linux struct {
// 	gorm.Model
// 	ServerIp string `json:"server"`
// }
