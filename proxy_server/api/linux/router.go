package linux

import "github.com/gin-gonic/gin"

type LinuxRouter struct {
}
type LinuxHandler struct {
}

func New_LinuxHandler() *LinuxHandler {
	return &LinuxHandler{}
}
func (LinuxRouter) Router(r *gin.Engine){
	linux:= r.Group("linux")
	linux.GET("information")
}