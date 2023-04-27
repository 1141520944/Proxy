package models

type ServerRequest struct {
	ServerName       string `json:"server_name" binding:"required"`
	ShowPort         string `json:"show_port" binding:"required"`
	ConnectPort      string `json:"connect_port" binding:"required"`
	LocalProjectPort string `json:"local_project_port" binding:"required"`
	Password         string `json:"password" binding:"required"`
	ServerState      bool   `json:"server_state" binding:"required"`
	Token            string
}
