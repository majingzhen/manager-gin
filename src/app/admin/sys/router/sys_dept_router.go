// Package router SysDeptRouter 自动生成模板
// @description <TODO description class purpose>
// @author
// @File: sys_dept
// @version 1.0.0
// @create 2023-08-21 10:27:01
package router

import (
	"github.com/gin-gonic/gin"
	"manager-gin/src/app/admin/sys/api"
	"manager-gin/src/middleware"
)

type SysDeptRouter struct {
	sysDeptApi api.SysDeptApi
}

// InitSysDeptRouter 初始化 SysDept 路由信息
func (r *SysDeptRouter) InitSysDeptRouter(Router *gin.RouterGroup) {
	sysDeptRouter := Router.Group("sysDept").Use(middleware.JWTAuthFilter())
	sysDeptRouterWithoutRecord := Router.Group("sysDept").Use(middleware.JWTAuthFilter())
	{
		sysDeptRouter.POST("create", r.sysDeptApi.Create)       // 新建SysDept
		sysDeptRouter.DELETE("delete/:id", r.sysDeptApi.Delete) // 删除SysDept
		sysDeptRouter.POST("update", r.sysDeptApi.Update)       // 更新SysDept
	}
	{
		sysDeptRouterWithoutRecord.GET("get/:id", r.sysDeptApi.Get) // 根据ID获取SysDept
		sysDeptRouterWithoutRecord.GET("list", r.sysDeptApi.List)   // 获取SysDept列表
		sysDeptRouterWithoutRecord.GET("list/exclude/:id", r.sysDeptApi.ListExclude)
		sysDeptRouterWithoutRecord.GET("tree", r.sysDeptApi.SelectDeptTree)                     // 获取部门树
		sysDeptRouterWithoutRecord.GET("treeByRole/:roleId", r.sysDeptApi.SelectDeptTreeByRole) // 根据角色ID获取部门树
	}
}
