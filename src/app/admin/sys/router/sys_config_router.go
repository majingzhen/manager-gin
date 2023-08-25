// Package router SysConfigRouter 自动生成模板
// @description <TODO description class purpose>
// @author
// @File: sys_config
// @version 1.0.0
// @create 2023-08-21 14:20:37
package router

import (
	"github.com/gin-gonic/gin"
	"manager-gin/src/app/admin/sys/api"
	"manager-gin/src/middleware"
)

type SysConfigRouter struct {
	sysConfigApi api.SysConfigApi
}

// InitSysConfigRouter 初始化 SysConfig 路由信息
func (r *SysConfigRouter) InitSysConfigRouter(Router *gin.RouterGroup) {
	sysConfigRouter := Router.Group("sysConfig").Use(middleware.JWTAuthFilter())
	sysConfigRouterWithoutRecord := Router.Group("sysConfig").Use(middleware.JWTAuthFilter())
	{
		sysConfigRouter.POST("create", r.sysConfigApi.Create)        // 新建SysConfig
		sysConfigRouter.DELETE("delete/:ids", r.sysConfigApi.Delete) // 删除SysConfig
		sysConfigRouter.POST("update", r.sysConfigApi.Update)        // 更新SysConfig
	}
	{
		sysConfigRouterWithoutRecord.GET("get/:id", r.sysConfigApi.Get)                      // 根据ID获取SysConfig
		sysConfigRouterWithoutRecord.GET("configKey/:key", r.sysConfigApi.SelectConfigByKey) // 根据ID获取SysConfig
		sysConfigRouterWithoutRecord.GET("page", r.sysConfigApi.Page)                        // 分页获取SysConfig列表
		sysConfigRouterWithoutRecord.GET("list", r.sysConfigApi.List)                        // 分页获取SysConfig列表
	}
}
