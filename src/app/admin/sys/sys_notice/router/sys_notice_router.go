// Package router SysNoticeRouter 自动生成模板
// @description <TODO description class purpose>
// @author
// @File: sys_notice
// @version 1.0.0
// @create 2023-08-08 10:06:19
package router

import (
	"github.com/gin-gonic/gin"
	"manager-gin/src/app/admin/sys/sys_notice/api"
	"manager-gin/src/middleware"
)

type SysNoticeRouter struct{}

var sysNoticeApi = api.SysNoticeApiApp

// InitSysNoticeRouter 初始化 SysNotice 路由信息
func (r *SysNoticeRouter) InitSysNoticeRouter(Router *gin.RouterGroup) {
	sysNoticeRouter := Router.Group("sysNotice").Use(middleware.JWTAuthFilter())
	sysNoticeRouterWithoutRecord := Router.Group("sysNotice")
	{
		sysNoticeRouter.POST("create", sysNoticeApi.Create)             // 新建SysNotice
		sysNoticeRouter.DELETE("delete", sysNoticeApi.Delete)           // 删除SysNotice
		sysNoticeRouter.DELETE("deleteByIds", sysNoticeApi.DeleteByIds) // 批量删除SysNotice
		sysNoticeRouter.POST("update", sysNoticeApi.Update)             // 更新SysNotice
	}
	{
		sysNoticeRouterWithoutRecord.GET("get", sysNoticeApi.Get)   // 根据ID获取SysNotice
		sysNoticeRouterWithoutRecord.GET("find", sysNoticeApi.Find) // 获取SysNotice列表
	}
}
