package routers

import (
	"github.com/gin-gonic/gin"
	"manager-gin/src/app/admin/sys/router"
)

type SysRouter struct {
	userRouter     router.UserRouter
	menuRouter     router.MenuRouter
	postRouter     router.PostRouter
	configRouter   router.ConfigRouter
	deptRouter     router.DeptRouter
	roleRouter     router.RoleRouter
	dictTypeRouter router.DictTypeRouter
	dictDataRouter router.DictDataRouter
}

// var jobRouter job.JobRouter
// var jobLogRouter jobLog.JobLogRouter
// var logininForRouter logininfor.LogininforRouter

// InitSysRouter 初始化 Init 路由信息
func (r *SysRouter) InitSysRouter(Router *gin.RouterGroup) {
	sys := Router.Group("sys")
	{
		r.userRouter.InitUserRouter(sys)
		r.menuRouter.InitMenuRouter(sys)
		r.roleRouter.InitRoleRouter(sys)
		r.dictTypeRouter.InitDictTypeRouter(sys)
		r.dictDataRouter.InitDictDataRouter(sys)
		//noticeRouter.InitNoticeRouter(sys)
		//operLogRouter.InitOperLogRouter(sys)
		r.postRouter.InitPostRouter(sys)
		r.configRouter.InitConfigRouter(sys)
		r.deptRouter.InitDeptRouter(sys)
		//jobRouter.InitJobRouter(sys)
		//jobLogRouter.InitJobLogRouter(sys)
		//logininForRouter.InitLogininforRouter(sys)
	}
}
