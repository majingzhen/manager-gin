// Package router SysJobLogRouter 自动生成模板
// @description <TODO description class purpose>
// @author
// @File: sys_job_log
// @version 1.0.0
// @create 2023-08-08 10:06:19
package router

import (
	"github.com/gin-gonic/gin"
	"manager-gin/src/app/admin/sys/sys_job_log/api"
	"manager-gin/src/middleware"
)

type SysJobLogRouter struct{}

var sysJobLogApi = api.SysJobLogApiApp

// InitSysJobLogRouter 初始化 SysJobLog 路由信息
func (r *SysJobLogRouter) InitSysJobLogRouter(Router *gin.RouterGroup) {
	sysJobLogRouter := Router.Group("sysJobLog").Use(middleware.JWTAuthFilter())
	sysJobLogRouterWithoutRecord := Router.Group("sysJobLog")
	{
		sysJobLogRouter.POST("create", sysJobLogApi.Create)             // 新建SysJobLog
		sysJobLogRouter.DELETE("delete", sysJobLogApi.Delete)           // 删除SysJobLog
		sysJobLogRouter.DELETE("deleteByIds", sysJobLogApi.DeleteByIds) // 批量删除SysJobLog
		sysJobLogRouter.POST("update", sysJobLogApi.Update)             // 更新SysJobLog
	}
	{
		sysJobLogRouterWithoutRecord.GET("get", sysJobLogApi.Get)   // 根据ID获取SysJobLog
		sysJobLogRouterWithoutRecord.GET("find", sysJobLogApi.Find) // 获取SysJobLog列表
	}
}
