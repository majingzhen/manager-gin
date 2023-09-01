package router

import (
	"github.com/gin-gonic/gin"
)

type GenRouter struct {
	tableRouter       TableRouter
	tableColumnRouter TableColumnRouter
}

// InitGenRouter 初始化 Init 路由信息
func (r *GenRouter) InitGenRouter(Router *gin.RouterGroup) {
	gen := Router.Group("gen")
	{
		r.tableRouter.InitTableRouter(gen)
		r.tableColumnRouter.InitTableColumnRouter(gen)
	}
}
