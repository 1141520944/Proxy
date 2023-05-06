package user

import (
	"github.com/1141520944/proxy/server/dao/mysql"
	"github.com/1141520944/proxy/server/pkg/mildware"
	"github.com/1141520944/proxy/server/pkg/repo"
	"github.com/1141520944/proxy/server/pkg/router"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}
type UserHandler struct {
	mysql repo.User_mysql
}

func New_handler() *UserHandler {
	return &UserHandler{
		mysql: mysql.New(),
	}
}
func init() {
	router.AddRouters(new(UserRouter))
}
func (UserRouter) Router(r *gin.Engine) {
	user := r.Group("/user")
	userHandler := New_handler()
	user.POST("/login", userHandler.LoginHandler)
	user.POST("/register", userHandler.RegisterHandler)
	user.GET("/select/one/:uid", userHandler.SelectByidHandler)
	user.GET("/select/all", userHandler.SelectAllHandler)
	user.GET("/delete/:uid", userHandler.DeleteHandler)

	user.Use(mildware.JWTAuthMiddleware())
	user.POST("/update/one", userHandler.UpdateInformationHandler)
}
