package routers

import (
	"github.com/gin-gonic/gin"
	config "manager-gin/src/app/admin/sys/sys_config/router"
	dept "manager-gin/src/app/admin/sys/sys_dept/router"
	dictData "manager-gin/src/app/admin/sys/sys_dict_data/router"
	dictType "manager-gin/src/app/admin/sys/sys_dict_type/router"
	post "manager-gin/src/app/admin/sys/sys_post/router"

	//job "manager-gin/src/app/admin/sys/sys_job/router"
	//jobLog "manager-gin/src/app/admin/sys/sys_job_log/router"
	//logininfor "manager-gin/src/app/admin/sys/sys_logininfor/router"
	menu "manager-gin/src/app/admin/sys/sys_menu/router"

	role "manager-gin/src/app/admin/sys/sys_role/router"
	user "manager-gin/src/app/admin/sys/sys_user/router"
)

type SysRouter struct{}

var userRouter user.SysUserRouter

var menuRouter menu.SysMenuRouter

// var noticeRouter notice.SysNoticeRouter
// var operLogRouter operLog.SysOperLogRouter
var postRouter post.SysPostRouter
var configRouter config.SysConfigRouter
var deptRouter dept.SysDeptRouter

// var jobRouter job.SysJobRouter
// var jobLogRouter jobLog.SysJobLogRouter
// var logininForRouter logininfor.SysLogininforRouter
var roleRouter role.SysRoleRouter
var dictTypeRouter dictType.SysDictTypeRouter
var dictDataRouter dictData.SysDictDataRouter

// InitSysRouter 初始化 InitSys 路由信息
func (r *SysRouter) InitSysRouter(Router *gin.RouterGroup) {
	sys := Router.Group("sys")
	{
		userRouter.InitSysUserRouter(sys)
		menuRouter.InitSysMenuRouter(sys)
		roleRouter.InitSysRoleRouter(sys)
		dictTypeRouter.InitSysDictTypeRouter(sys)
		dictDataRouter.InitSysDictDataRouter(sys)
		//noticeRouter.InitSysNoticeRouter(sys)
		//operLogRouter.InitSysOperLogRouter(sys)
		postRouter.InitSysPostRouter(sys)
		configRouter.InitSysConfigRouter(sys)
		deptRouter.InitSysDeptRouter(sys)
		//jobRouter.InitSysJobRouter(sys)
		//jobLogRouter.InitSysJobLogRouter(sys)
		//logininForRouter.InitSysLogininforRouter(sys)
	}
}
