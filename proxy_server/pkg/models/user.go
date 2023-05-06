package models

import "time"

type UserLoginRequest struct {
	RemoteAddr string `json:"remote_addr" binding:"required"`
	UserName   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
}
type UserRegisterRequest struct {
	UserName   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}
type UserInformationRequest struct {
	Name     string `json:"name" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type UserInformationResponse struct {
	Name          string    `json:"name"`
	Username      string    `json:"username"`
	Password      string    `json:"password"`
	Phone         string    `json:"phone"`
	UserID        string    `json:"user_id"`
	PassLoginTime time.Time `json:"pass_login_time"`
	ServerNum     int       `json:"server_num"`
}
