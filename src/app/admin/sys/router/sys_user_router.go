// Package router SysUserRouter 自动生成模板
// @description <TODO description class purpose>
// @author
// @File: sys_user
// @version 1.0.0
// @create 2023-08-18 14:02:24
package router

import (
	"github.com/gin-gonic/gin"
	"manager-gin/src/app/admin/sys/api"
	"manager-gin/src/middleware"
)

type SysUserRouter struct {
	sysUserApi api.SysUserApi
}

// InitSysUserRouter 初始化 SysUser 路由信息
func (r *SysUserRouter) InitSysUserRouter(Router *gin.RouterGroup) {
	sysUserRouter := Router.Group("sysUser").Use(middleware.JWTAuthFilter())
	sysUserRouterWithoutRecord := Router.Group("sysUser").Use(middleware.JWTAuthFilter())
	{
		sysUserRouter.POST("create", r.sysUserApi.Create)            // 新建SysUser
		sysUserRouter.DELETE("delete/:ids", r.sysUserApi.Delete)     // 删除SysUser
		sysUserRouter.PUT("update", r.sysUserApi.Update)             // 更新SysUser
		sysUserRouter.PUT("changeStatus", r.sysUserApi.ChangeStatus) // 修改状态
		sysUserRouter.PUT("resetPwd", r.sysUserApi.ResetPwd)         // 重置密码
		sysUserRouter.PUT("authRole", r.sysUserApi.AuthRole)         // 重置密码
	}
	{
		sysUserRouterWithoutRecord.GET("get", r.sysUserApi.Get)
		sysUserRouterWithoutRecord.GET("get/:id", r.sysUserApi.Get)                           // 根据ID获取SysUser
		sysUserRouterWithoutRecord.GET("authRole/:id", r.sysUserApi.GetAuthRole)              // 根据ID获取授权集合
		sysUserRouterWithoutRecord.GET("page", r.sysUserApi.Page)                             // 分页获取SysUser列表
		sysUserRouterWithoutRecord.GET("list", r.sysUserApi.List)                             // 分页获取SysUser列表
		sysUserRouterWithoutRecord.GET("allocatedList", r.sysUserApi.SelectAllocatedList)     // 分页获取角色已授权用户列表
		sysUserRouterWithoutRecord.GET("unallocatedList", r.sysUserApi.SelectUnallocatedList) // 分页获取角色未授权用户列表
	}
}
