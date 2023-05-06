package router

import (
	"github.com/1141520944/proxy/client/pkg/logger"
	"github.com/1141520944/proxy/client/pkg/mildware"
	"github.com/gin-gonic/gin"
)

type Router interface {
	Router(r *gin.Engine)
}

type RegistRouter struct {
	R *gin.Engine
}

func New(r *gin.Engine) *RegistRouter {
	return &RegistRouter{
		R: r,
	}
}
func (rg *RegistRouter) Router(r Router) {
	r.Router(rg.R)
}

var Routers []Router

func AddRouters(router ...Router) {
	Routers = append(Routers, router...)
}

func Init(r *gin.Engine) *gin.Engine {
	gr := New(r)
	gr.R.Use(logger.GinLogger(), logger.GinRecovery(true), mildware.CORSMiddleware())
	for _, router := range Routers {
		gr.Router(router)
	}
	return gr.R
}
