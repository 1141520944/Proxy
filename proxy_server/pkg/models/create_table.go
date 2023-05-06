package models

import (
	"time"

	"gorm.io/gorm"
)

type Table_Server struct {
	gorm.Model
	ServerID         int64  `json:"server_id" gorm:"columm:server_id,primary"`
	ServerName       string `json:"server_name" gorm:"columm:server_name"`
	ServerPassword   string `json:"server_password" gorm:"columm:server_password"`
	ShowPort         string `json:"show_port" gorm:"columm:show_port"`
	ConnectPort      string `json:"connect_port" gorm:"columm:connect_port"`
	LocalProjectPort string `json:"local_project_port" gorm:"columm:local_project_port"`
	UserID           int64  `json:"user_id" gorm:"columm:user_id"`
	ServerState      bool   `json:"server_state" gorm:"columm:server_state"`
}

func (Table_Server) TableName() string {
	return "server"
}

type Table_Linux struct {
	gorm.Model
	OptiationStream string `json:"Optiation_stream"  gorm:"columm:Optiation_stream"`
	ServerIp        string `json:"server_ip"  gorm:"columm:server_ip"`
	ProxyRunPort    string `json:"proxy_run_port" gorm:"columm:proxy_run_port"`
	ProxyMysqlPort  string `json:"proxy_mysql_port" gorm:"columm:proxy_mysql_port"`
	ProxyRedisPort  string `json:"proxy_redis_port" gorm:"columm:proxy_redis_port"`
	ServerNum       int    `json:"server_num" gorm:"columm:server_num"`
}

func (Table_Linux) TableName() string {
	return "linux"
}

// Table_user 用户
type Table_user struct {
	gorm.Model
	Name          string    `json:"name" gorm:"columm:name"`
	Username      string    `json:"username" gorm:"columm:username"`
	Password      string    `json:"password" gorm:"columm:password"`
	Phone         string    `json:"phone" gorm:"columm:phone"`
	UserID        int64     `json:"user_id" gorm:"columm:user_id"`
	PassLoginTime time.Time `json:"pass_login_time" gorm:"columm:pass_login_time"`
	ServerNum     int       `json:"server_num" gorm:"columm:server_num"`
}

// TableName 指定表名
func (Table_user) TableName() string {
	return "user"
}
