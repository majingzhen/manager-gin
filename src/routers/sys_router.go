package routers

import (
	"github.com/gin-gonic/gin"
	"manager-gin/src/app/admin/sys/router"
)

type SysRouter struct {
	userRouter     router.SysUserRouter
	menuRouter     router.SysMenuRouter
	postRouter     router.SysPostRouter
	configRouter   router.SysConfigRouter
	deptRouter     router.SysDeptRouter
	roleRouter     router.SysRoleRouter
	dictTypeRouter router.SysDictTypeRouter
	dictDataRouter router.SysDictDataRouter
}

// var jobRouter job.SysJobRouter
// var jobLogRouter jobLog.SysJobLogRouter
// var logininForRouter logininfor.SysLogininforRouter

// InitSysRouter 初始化 InitSys 路由信息
func (r *SysRouter) InitSysRouter(Router *gin.RouterGroup) {
	sys := Router.Group("sys")
	{
		r.userRouter.InitSysUserRouter(sys)
		r.menuRouter.InitSysMenuRouter(sys)
		r.roleRouter.InitSysRoleRouter(sys)
		r.dictTypeRouter.InitSysDictTypeRouter(sys)
		r.dictDataRouter.InitSysDictDataRouter(sys)
		//noticeRouter.InitSysNoticeRouter(sys)
		//operLogRouter.InitSysOperLogRouter(sys)
		r.postRouter.InitSysPostRouter(sys)
		r.configRouter.InitSysConfigRouter(sys)
		r.deptRouter.InitSysDeptRouter(sys)
		//jobRouter.InitSysJobRouter(sys)
		//jobLogRouter.InitSysJobLogRouter(sys)
		//logininForRouter.InitSysLogininforRouter(sys)
	}
}
