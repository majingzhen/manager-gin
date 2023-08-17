// Package router BiRoleRouter 自动生成模板
// @description <TODO description class purpose>
// @author Majz
// @File: bi_role
// @version 1.0.0
// @create 2023-08-16 14:53:36
package router

import (
	"github.com/gin-gonic/gin"
	"manager-gin/src/app/admin/bi/bi_role/api"
	"manager-gin/src/middleware"
)

type BiRoleRouter struct{}

var biRoleApi = api.BiRoleApiApp

// InitBiRoleRouter 初始化 BiRole 路由信息
func (r *BiRoleRouter) InitBiRoleRouter(Router *gin.RouterGroup) {
	biRoleRouter := Router.Group("biRole").Use(middleware.JWTAuthFilter())
	biRoleRouterWithoutRecord := Router.Group("biRole")
	{
		biRoleRouter.POST("create", biRoleApi.Create)             // 新建BiRole
		biRoleRouter.DELETE("delete", biRoleApi.Delete)           // 删除BiRole
		biRoleRouter.DELETE("deleteByIds", biRoleApi.DeleteByIds) // 批量删除BiRole
		biRoleRouter.POST("update", biRoleApi.Update)             // 更新BiRole
	}
	{
		biRoleRouterWithoutRecord.GET("get", biRoleApi.Get)   // 根据ID获取BiRole
		biRoleRouterWithoutRecord.GET("find", biRoleApi.Find) // 获取BiRole列表
	}
}
