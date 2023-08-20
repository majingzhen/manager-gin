package routers

import (
	"github.com/gin-gonic/gin"
	"manager-gin/src/app/admin/sys/system/router"
)

type BaseRouter struct{}

var systemRouter router.SystemRouter

func (r *BaseRouter) InitBaseRouter(router *gin.RouterGroup) {
	systemRouter.InitSystemRouter(router)
}
