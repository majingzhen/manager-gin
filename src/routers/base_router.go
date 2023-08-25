package routers

import (
	"github.com/gin-gonic/gin"
	"manager-gin/src/app/admin/sys/router"
)

type BaseRouter struct {
	systemRouter router.SystemRouter
}

func (r *BaseRouter) InitBaseRouter(router *gin.RouterGroup) {
	r.systemRouter.InitSystemRouter(router)
}
