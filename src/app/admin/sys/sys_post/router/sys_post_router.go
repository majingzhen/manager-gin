// Package router SysPostRouter 自动生成模板
// @description <TODO description class purpose>
// @author
// @File: sys_post
// @version 1.0.0
// @create 2023-08-21 17:37:56
package router

import (
	"github.com/gin-gonic/gin"
	"manager-gin/src/app/admin/sys/sys_post/api"
	"manager-gin/src/middleware"
)

type SysPostRouter struct{}

var sysPostApi = api.SysPostApiApp

// InitSysPostRouter 初始化 SysPost 路由信息
func (r *SysPostRouter) InitSysPostRouter(Router *gin.RouterGroup) {
	sysPostRouter := Router.Group("sysPost").Use(middleware.JWTAuthFilter())
	sysPostRouterWithoutRecord := Router.Group("sysPost").Use(middleware.JWTAuthFilter())
	{
		sysPostRouter.POST("create", sysPostApi.Create)        // 新建SysPost
		sysPostRouter.DELETE("delete/:ids", sysPostApi.Delete) // 删除SysPost
		sysPostRouter.POST("update", sysPostApi.Update)        // 更新SysPost
	}
	{
		sysPostRouterWithoutRecord.GET("get/:id", sysPostApi.Get) // 根据ID获取SysPost
		sysPostRouterWithoutRecord.GET("page", sysPostApi.Page)   // 分页获取SysPost列表
		sysPostRouterWithoutRecord.GET("list", sysPostApi.List)   // 分页获取SysPost列表
	}
}
