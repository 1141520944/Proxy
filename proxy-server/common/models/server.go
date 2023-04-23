package models

type ServerRequest struct {
	UserPort   string `json:"user_port" binding:"required"`
	ServerPort string `json:"server_port" binding:"required"`
	Password   string `json:"password" binding:"required"`
	Exist      bool   `json:"exist" binding:"required"`
	Token      string `json:"token"`
	Die        chan string
}
