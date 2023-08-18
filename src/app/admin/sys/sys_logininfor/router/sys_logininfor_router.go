// Package router SysLogininforRouter 自动生成模板
// @description <TODO description class purpose>
// @author
// @File: sys_logininfor
// @version 1.0.0
// @create 2023-08-18 13:41:26
package router

import (
	"github.com/gin-gonic/gin"
	"manager-gin/src/app/admin/sys/sys_logininfor/api"
	"manager-gin/src/middleware"
)

type SysLogininforRouter struct{}

var sysLogininforApi = api.SysLogininforApiApp

// InitSysLogininforRouter 初始化 SysLogininfor 路由信息
func (r *SysLogininforRouter) InitSysLogininforRouter(Router *gin.RouterGroup) {
	sysLogininforRouter := Router.Group("sysLogininfor")
	sysLogininforRouter.Use(middleware.JWTAuthFilter())
	sysLogininforRouterWithoutRecord := Router.Group("sysLogininfor")
	{
		sysLogininforRouter.POST("create", sysLogininforApi.Create)             // 新建SysLogininfor
		sysLogininforRouter.DELETE("delete", sysLogininforApi.Delete)           // 删除SysLogininfor
		sysLogininforRouter.DELETE("deleteByIds", sysLogininforApi.DeleteByIds) // 批量删除SysLogininfor
		sysLogininforRouter.POST("update", sysLogininforApi.Update)             // 更新SysLogininfor
	}
	{
		sysLogininforRouterWithoutRecord.GET("get", sysLogininforApi.Get)   // 根据ID获取SysLogininfor
		sysLogininforRouterWithoutRecord.GET("find", sysLogininforApi.Find) // 获取SysLogininfor列表
	}
}
