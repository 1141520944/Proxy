package router

import (
	"github.com/1141520944/proxy/logger"
	"github.com/gin-gonic/gin"
)

type Router interface {
	Router(*gin.Engine)
}

type RegistRouter struct {
	r *gin.Engine
}

func New(r *gin.Engine) *RegistRouter {
	return &RegistRouter{r: r}
}
func (rr *RegistRouter) Router(ro Router) {
	ro.Router(rr.r)
}
func AddRouters(router ...Router) {
	Routers = append(Routers, router...)
}

var Routers []Router

func Init(e *gin.Engine) *gin.Engine {
	rr := New(e)
	rr.r.Use(logger.GinLogger(), logger.GinRecovery(true))
	for _, router := range Routers {
		rr.Router(router)
	}
	return rr.r
}
