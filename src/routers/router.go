package routers

import (
	"github.com/gin-gonic/gin"
)

type Routers struct {
}

// InitRouter 初始化路由
func (routers *Routers) InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	sysRouter := new(SysRouter)
	api := r.Group("/api")
	{
		sysRouter.InitSysRouter(api)
	}

	return r
}
