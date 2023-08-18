// Package router SysOperLogRouter 自动生成模板
// @description <TODO description class purpose>
// @author
// @File: sys_oper_log
// @version 1.0.0
// @create 2023-08-18 13:41:26
package router

import (
	"github.com/gin-gonic/gin"
	"manager-gin/src/app/admin/sys/sys_oper_log/api"
	"manager-gin/src/middleware"
)

type SysOperLogRouter struct{}

var sysOperLogApi = api.SysOperLogApiApp

// InitSysOperLogRouter 初始化 SysOperLog 路由信息
func (r *SysOperLogRouter) InitSysOperLogRouter(Router *gin.RouterGroup) {
	sysOperLogRouter := Router.Group("sysOperLog")
	sysOperLogRouter.Use(middleware.JWTAuthFilter())
	sysOperLogRouterWithoutRecord := Router.Group("sysOperLog")
	{
		sysOperLogRouter.POST("create", sysOperLogApi.Create)             // 新建SysOperLog
		sysOperLogRouter.DELETE("delete", sysOperLogApi.Delete)           // 删除SysOperLog
		sysOperLogRouter.DELETE("deleteByIds", sysOperLogApi.DeleteByIds) // 批量删除SysOperLog
		sysOperLogRouter.POST("update", sysOperLogApi.Update)             // 更新SysOperLog
	}
	{
		sysOperLogRouterWithoutRecord.GET("get", sysOperLogApi.Get)   // 根据ID获取SysOperLog
		sysOperLogRouterWithoutRecord.GET("find", sysOperLogApi.Find) // 获取SysOperLog列表
	}
}
