// Package router SysConfigRouter 自动生成模板
// @description <TODO description class purpose>
// @author
// @File: sys_config
// @version 1.0.0
// @create 2023-08-08 10:06:19
package router

import (
	"github.com/gin-gonic/gin"
	"manager-gin/src/app/admin/sys/sys_config/api"
	"manager-gin/src/middleware"
)

type SysConfigRouter struct{}

var sysConfigApi = api.SysConfigApiApp

// InitSysConfigRouter 初始化 SysConfig 路由信息
func (r *SysConfigRouter) InitSysConfigRouter(Router *gin.RouterGroup) {
	sysConfigRouter := Router.Group("sysConfig").Use(middleware.JWTAuthFilter())
	sysConfigRouterWithoutRecord := Router.Group("sysConfig")
	{
		sysConfigRouter.POST("create", sysConfigApi.Create)             // 新建SysConfig
		sysConfigRouter.DELETE("delete", sysConfigApi.Delete)           // 删除SysConfig
		sysConfigRouter.DELETE("deleteByIds", sysConfigApi.DeleteByIds) // 批量删除SysConfig
		sysConfigRouter.POST("update", sysConfigApi.Update)             // 更新SysConfig
	}
	{
		sysConfigRouterWithoutRecord.GET("get", sysConfigApi.Get)   // 根据ID获取SysConfig
		sysConfigRouterWithoutRecord.GET("find", sysConfigApi.Find) // 获取SysConfig列表
	}
}
