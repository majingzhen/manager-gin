package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	r.Use(cors.New(config))

	sysRouter := new(SysRouter)
	r.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "hello")
	})
	api := r.Group("/api")
	{
		sysRouter.InitSysRouter(api)
	}

	return r
}
