// Package router SysDeptRouter 自动生成模板
// @description <TODO description class purpose>
// @author
// @File: sys_dept
// @version 1.0.0
// @create 2023-08-21 10:27:01
package router

import (
	"github.com/gin-gonic/gin"
	"manager-gin/src/app/admin/sys/sys_dept/api"
	"manager-gin/src/middleware"
)

type SysDeptRouter struct{}

var sysDeptApi = api.SysDeptApiApp

// InitSysDeptRouter 初始化 SysDept 路由信息
func (r *SysDeptRouter) InitSysDeptRouter(Router *gin.RouterGroup) {
	sysDeptRouter := Router.Group("sysDept").Use(middleware.JWTAuthFilter())
	sysDeptRouterWithoutRecord := Router.Group("sysDept").Use(middleware.JWTAuthFilter())
	{
		sysDeptRouter.POST("create", sysDeptApi.Create)       // 新建SysDept
		sysDeptRouter.DELETE("delete/:id", sysDeptApi.Delete) // 删除SysDept
		sysDeptRouter.POST("update", sysDeptApi.Update)       // 更新SysDept
	}
	{
		sysDeptRouterWithoutRecord.GET("get/:id", sysDeptApi.Get)                  // 根据ID获取SysDept
		sysDeptRouterWithoutRecord.GET("page", sysDeptApi.Page)                    // 分页获取SysDept列表
		sysDeptRouterWithoutRecord.GET("list", sysDeptApi.List)                    // 分页获取SysDept列表
		sysDeptRouterWithoutRecord.GET("list/exclude/:id", sysDeptApi.ListExclude) // 分页获取SysDept列表
		sysDeptRouterWithoutRecord.GET("tree", sysDeptApi.SelectDeptTree)          // 获取部门树
	}
}
