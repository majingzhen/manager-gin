// Package router SysPostRouter 自动生成模板
// @description <TODO description class purpose>
// @author
// @File: sys_post
// @version 1.0.0
// @create 2023-08-21 17:37:56
package router

import (
	"github.com/gin-gonic/gin"
	"manager-gin/src/app/admin/sys/api"
	"manager-gin/src/middleware"
)

type SysPostRouter struct {
	sysPostApi api.SysPostApi
}

// InitSysPostRouter 初始化 SysPost 路由信息
func (r *SysPostRouter) InitSysPostRouter(Router *gin.RouterGroup) {
	sysPostRouter := Router.Group("sysPost").Use(middleware.JWTAuthFilter())
	sysPostRouterWithoutRecord := Router.Group("sysPost").Use(middleware.JWTAuthFilter())
	{
		sysPostRouter.POST("create", r.sysPostApi.Create)        // 新建SysPost
		sysPostRouter.DELETE("delete/:ids", r.sysPostApi.Delete) // 删除SysPost
		sysPostRouter.POST("update", r.sysPostApi.Update)        // 更新SysPost
	}
	{
		sysPostRouterWithoutRecord.GET("get/:id", r.sysPostApi.Get) // 根据ID获取SysPost
		sysPostRouterWithoutRecord.GET("page", r.sysPostApi.Page)   // 分页获取SysPost列表
		sysPostRouterWithoutRecord.GET("list", r.sysPostApi.List)   // 分页获取SysPost列表
	}
}
