// Package router SysUserRouter 自动生成模板
// @description <TODO description class purpose>
// @author
// @File: sys_user
// @version 1.0.0
// @create 2023-08-18 14:02:24
package router

import (
	"github.com/gin-gonic/gin"
	"manager-gin/src/app/admin/sys/sys_user/api"
	"manager-gin/src/middleware"
)

type SysUserRouter struct{}

var sysUserApi = api.SysUserApiApp

// InitSysUserRouter 初始化 SysUser 路由信息
func (r *SysUserRouter) InitSysUserRouter(Router *gin.RouterGroup) {
	sysUserRouter := Router.Group("sysUser").Use(middleware.JWTAuthFilter())
	sysUserRouterWithoutRecord := Router.Group("sysUser").Use(middleware.JWTAuthFilter())
	{
		sysUserRouter.POST("create", sysUserApi.Create)            // 新建SysUser
		sysUserRouter.DELETE("delete/:ids", sysUserApi.Delete)     // 删除SysUser
		sysUserRouter.PUT("update", sysUserApi.Update)             // 更新SysUser
		sysUserRouter.PUT("changeStatus", sysUserApi.ChangeStatus) // 修改状态
		sysUserRouter.PUT("resetPwd", sysUserApi.ResetPwd)         // 重置密码
		sysUserRouter.PUT("authRole", sysUserApi.AuthRole)         // 重置密码
	}
	{
		sysUserRouterWithoutRecord.GET("get", sysUserApi.Get)
		sysUserRouterWithoutRecord.GET("get/:id", sysUserApi.Get)              // 根据ID获取SysUser
		sysUserRouterWithoutRecord.GET("authRole/:id", sysUserApi.GetAuthRole) // 根据ID获取授权集合
		sysUserRouterWithoutRecord.GET("page", sysUserApi.Page)                // 分页获取SysUser列表
		sysUserRouterWithoutRecord.GET("list", sysUserApi.List)                // 分页获取SysUser列表
	}
}
