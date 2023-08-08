// Package router SysPostRouter 自动生成模板
// @description <TODO description class purpose>
// @author
// @File: sys_post
// @version 1.0.0
// @create 2023-08-08 10:06:19
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
	sysPostRouterWithoutRecord := Router.Group("sysPost")
	{
		sysPostRouter.POST("create", sysPostApi.Create)             // 新建SysPost
		sysPostRouter.DELETE("delete", sysPostApi.Delete)           // 删除SysPost
		sysPostRouter.DELETE("deleteByIds", sysPostApi.DeleteByIds) // 批量删除SysPost
		sysPostRouter.POST("update", sysPostApi.Update)             // 更新SysPost
	}
	{
		sysPostRouterWithoutRecord.GET("get", sysPostApi.Get)   // 根据ID获取SysPost
		sysPostRouterWithoutRecord.GET("find", sysPostApi.Find) // 获取SysPost列表
	}
}
