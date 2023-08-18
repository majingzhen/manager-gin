// Package router SysOrganizationRouter 自动生成模板
// @description <TODO description class purpose>
// @author
// @File: sys_organization
// @version 1.0.0
// @create 2023-08-18 14:00:54
package router

import (
	"github.com/gin-gonic/gin"
	"manager-gin/src/app/admin/sys/sys_organization/api"
	"manager-gin/src/middleware"
)

type SysOrganizationRouter struct{}

var sysOrganizationApi = api.SysOrganizationApiApp

// InitSysOrganizationRouter 初始化 SysOrganization 路由信息
func (r *SysOrganizationRouter) InitSysOrganizationRouter(Router *gin.RouterGroup) {
	sysOrganizationRouter := Router.Group("sysOrganization")
	sysOrganizationRouter.Use(middleware.JWTAuthFilter())
	sysOrganizationRouterWithoutRecord := Router.Group("sysOrganization")
	{
		sysOrganizationRouter.POST("create", sysOrganizationApi.Create)             // 新建SysOrganization
		sysOrganizationRouter.DELETE("delete", sysOrganizationApi.Delete)           // 删除SysOrganization
		sysOrganizationRouter.DELETE("deleteByIds", sysOrganizationApi.DeleteByIds) // 批量删除SysOrganization
		sysOrganizationRouter.POST("update", sysOrganizationApi.Update)             // 更新SysOrganization
	}
	{
		sysOrganizationRouterWithoutRecord.GET("get", sysOrganizationApi.Get)   // 根据ID获取SysOrganization
		sysOrganizationRouterWithoutRecord.GET("find", sysOrganizationApi.Find) // 获取SysOrganization列表
	}
}
