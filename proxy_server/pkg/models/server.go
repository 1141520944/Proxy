package models

import "time"

type ServerRequest struct {
	ServerName       string `json:"server_name" binding:"required"`
	ShowPort         string `json:"show_port" binding:"required"`
	ConnectPort      string `json:"connect_port" binding:"required"`
	LocalProjectPort string `json:"local_project_port" binding:"required"`
	Password         string `json:"password" binding:"required"`
	UserID           string `json:"user_id"`
}
type ServerResponse struct {
	ServerName       string    `json:"server_name"`
	ShowPort         string    `json:"show_port"`
	ConnectPort      string    `json:"connect_port"`
	LocalProjectPort string    `json:"local_project_port"`
	Password         string    `json:"password"`
	UserID           string    `json:"user_id"`
	ServerID         string    `json:"server_id"`
	CreateTime       time.Time `json:"create_time"`
	State            bool      `json:"state"`
}
