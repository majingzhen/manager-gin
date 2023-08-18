// Package router SysJobRouter 自动生成模板
// @description <TODO description class purpose>
// @author
// @File: sys_job
// @version 1.0.0
// @create 2023-08-18 13:41:26
package router

import (
	"github.com/gin-gonic/gin"
	"manager-gin/src/app/admin/sys/sys_job/api"
	"manager-gin/src/middleware"
)

type SysJobRouter struct{}

var sysJobApi = api.SysJobApiApp

// InitSysJobRouter 初始化 SysJob 路由信息
func (r *SysJobRouter) InitSysJobRouter(Router *gin.RouterGroup) {
	sysJobRouter := Router.Group("sysJob")
	sysJobRouter.Use(middleware.JWTAuthFilter())
	sysJobRouterWithoutRecord := Router.Group("sysJob")
	{
		sysJobRouter.POST("create", sysJobApi.Create)             // 新建SysJob
		sysJobRouter.DELETE("delete", sysJobApi.Delete)           // 删除SysJob
		sysJobRouter.DELETE("deleteByIds", sysJobApi.DeleteByIds) // 批量删除SysJob
		sysJobRouter.POST("update", sysJobApi.Update)             // 更新SysJob
	}
	{
		sysJobRouterWithoutRecord.GET("get", sysJobApi.Get)   // 根据ID获取SysJob
		sysJobRouterWithoutRecord.GET("find", sysJobApi.Find) // 获取SysJob列表
	}
}
