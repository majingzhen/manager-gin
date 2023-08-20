package routers

import (
	"github.com/gin-gonic/gin"
	"manager-gin/src/middleware"
)

type Routers struct {
}

// InitRouter 初始化路由
func (routers *Routers) InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// 跨域处理
	// 使用Cors中间件处理跨域请求
	r.Use(middleware.Cors())

	sysRouter := new(SysRouter)
	baseRouter := new(BaseRouter)
	r.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "hello")
	})
	api := r.Group("/api")
	{
		sysRouter.InitSysRouter(api)
		baseRouter.InitBaseRouter(api)
	}

	return r
}
