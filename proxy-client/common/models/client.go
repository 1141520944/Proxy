package models

import (
	"gorm.io/gorm"
)

type ServerInitRequest struct {
	ServerName       string `json:"server_name" binding:"required"`
	ShowPort         string `json:"show_port" binding:"required"`
	ConnectPort      string `json:"connect_port" binding:"required"`
	LocalProjectPort string `json:"local_project_port" binding:"required"`
	Password         string `json:"password" binding:"required"`
	ServerState      bool   `json:"server_state" binding:"required"`
	Token            string
}

// ./client -l 127.0.0.1:80 -r 61.1.1.1:8010 -pwd mypassword
type ClientConnect struct {
	LocalProjectAddr string `json:"local_project_addr" binding:"required"`
	ConnectAddr      string `json:"connect_addr" binding:"required"`
	Password         string `json:"password" binding:"required"`
}

type Tabler_Server struct {
	gorm.Model
	Server_id          int64  `json:"server_id,string"`
	Server_name        string `json:"server_name"`
	Server_password    string `json:"server_password"`
	Show_port          string `json:"show_port"`
	Connect_port       string `json:"connect_port"`
	Local_project_port string `json:"local_project_port"`
	Server_state       string `json:"server_state"`
}
type ServerShowPortResponse struct {
	Code CodeError     `json:"code"`
	Msg  interface{}   `json:"msg"`
	Data Tabler_Server `json:"data,omitempty"`
}
