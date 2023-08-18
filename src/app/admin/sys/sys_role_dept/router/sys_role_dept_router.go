// Package router SysRoleDeptRouter 自动生成模板
// @description <TODO description class purpose>
// @author
// @File: sys_role_dept
// @version 1.0.0
// @create 2023-08-18 13:41:26
package router

import (
	"github.com/gin-gonic/gin"
	"manager-gin/src/app/admin/sys/sys_role_dept/api"
	"manager-gin/src/middleware"
)

type SysRoleDeptRouter struct{}

var sysRoleDeptApi = api.SysRoleDeptApiApp

// InitSysRoleDeptRouter 初始化 SysRoleDept 路由信息
func (r *SysRoleDeptRouter) InitSysRoleDeptRouter(Router *gin.RouterGroup) {
	sysRoleDeptRouter := Router.Group("sysRoleDept")
	sysRoleDeptRouter.Use(middleware.JWTAuthFilter())
	sysRoleDeptRouterWithoutRecord := Router.Group("sysRoleDept")
	{
		sysRoleDeptRouter.POST("create", sysRoleDeptApi.Create)             // 新建SysRoleDept
		sysRoleDeptRouter.DELETE("delete", sysRoleDeptApi.Delete)           // 删除SysRoleDept
		sysRoleDeptRouter.DELETE("deleteByIds", sysRoleDeptApi.DeleteByIds) // 批量删除SysRoleDept
		sysRoleDeptRouter.POST("update", sysRoleDeptApi.Update)             // 更新SysRoleDept
	}
	{
		sysRoleDeptRouterWithoutRecord.GET("get", sysRoleDeptApi.Get)   // 根据ID获取SysRoleDept
		sysRoleDeptRouterWithoutRecord.GET("find", sysRoleDeptApi.Find) // 获取SysRoleDept列表
	}
}
