package routers

import (
	"github.com/gin-gonic/gin"
	config "manager-gin/src/app/admin/sys/sys_config/router"
	job "manager-gin/src/app/admin/sys/sys_job/router"
	jobLog "manager-gin/src/app/admin/sys/sys_job_log/router"
	logininfor "manager-gin/src/app/admin/sys/sys_logininfor/router"
	menu "manager-gin/src/app/admin/sys/sys_menu/router"
	notice "manager-gin/src/app/admin/sys/sys_notice/router"
	operLog "manager-gin/src/app/admin/sys/sys_oper_log/router"
	org "manager-gin/src/app/admin/sys/sys_organization/router"
	post "manager-gin/src/app/admin/sys/sys_post/router"
	role "manager-gin/src/app/admin/sys/sys_role/router"
	user "manager-gin/src/app/admin/sys/sys_user/router"
)

type SysRouter struct{}

var userRouter user.SysUserRouter
var configRouter config.SysConfigRouter
var orgRouter org.SysOrganizationRouter
var jobRouter job.SysJobRouter
var jobLogRouter jobLog.SysJobLogRouter
var logininForRouter logininfor.SysLogininforRouter
var menuRouter menu.SysMenuRouter
var noticeRouter notice.SysNoticeRouter
var operLogRouter operLog.SysOperLogRouter
var postRouter post.SysPostRouter
var roleRouter role.SysRoleRouter

// InitSysRouter 初始化 InitSys 路由信息
func (r *SysRouter) InitSysRouter(Router *gin.RouterGroup) {
	sys := Router.Group("sys")
	{
		userRouter.InitSysUserRouter(sys)
		configRouter.InitSysConfigRouter(sys)
		orgRouter.InitSysOrganizationRouter(sys)
		jobRouter.InitSysJobRouter(sys)
		jobLogRouter.InitSysJobLogRouter(sys)
		logininForRouter.InitSysLogininforRouter(sys)
		menuRouter.InitSysMenuRouter(sys)
		noticeRouter.InitSysNoticeRouter(sys)
		operLogRouter.InitSysOperLogRouter(sys)
		postRouter.InitSysPostRouter(sys)
		roleRouter.InitSysRoleRouter(sys)

	}
}
