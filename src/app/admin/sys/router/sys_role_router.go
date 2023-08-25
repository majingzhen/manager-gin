// Package router SysRoleRouter 自动生成模板
// @description <TODO description class purpose>
// @author
// @File: sys_role
// @version 1.0.0
// @create 2023-08-21 17:37:56
package router

import (
	"github.com/gin-gonic/gin"
	"manager-gin/src/app/admin/sys/api"
	"manager-gin/src/middleware"
)

type SysRoleRouter struct {
	sysRoleApi api.SysRoleApi
}

// InitSysRoleRouter 初始化 SysRole 路由信息
func (r *SysRoleRouter) InitSysRoleRouter(Router *gin.RouterGroup) {
	sysRoleRouter := Router.Group("sysRole").Use(middleware.JWTAuthFilter())
	sysRoleRouterWithoutRecord := Router.Group("sysRole").Use(middleware.JWTAuthFilter())
	{
		sysRoleRouter.POST("create", r.sysRoleApi.Create)                      // 新建SysRole
		sysRoleRouter.DELETE("delete/:ids", r.sysRoleApi.Delete)               // 删除SysRole
		sysRoleRouter.POST("update", r.sysRoleApi.Update)                      // 更新SysRole
		sysRoleRouter.PUT("changeStatus", r.sysRoleApi.ChangeStatus)           // 更新SysRole状态
		sysRoleRouter.PUT("dataScope", r.sysRoleApi.DataScope)                 // 数据授权
		sysRoleRouter.PUT("cancelAuth", r.sysRoleApi.CancelAuthUser)           // 取消授权
		sysRoleRouter.PUT("batchCancelAuth", r.sysRoleApi.BatchCancelAuthUser) // 批量取消授权
		sysRoleRouter.PUT("batchSelectAuth", r.sysRoleApi.BatchSelectAuthUser) // 批量选择授权

	}
	{
		sysRoleRouterWithoutRecord.GET("get/:id", r.sysRoleApi.Get) // 根据ID获取SysRole
		sysRoleRouterWithoutRecord.GET("page", r.sysRoleApi.Page)   // 分页获取SysRole列表
		sysRoleRouterWithoutRecord.GET("list", r.sysRoleApi.List)   // 分页获取SysRole列表
	}
}
