// Package router BiUserRouter 自动生成模板
// @description <TODO description class purpose>
// @author Administrator
// @File: bi_user
// @version 1.0.0
// @create 2023-08-16 14:53:36
package router

import (
	"github.com/gin-gonic/gin"
	"manager-gin/src/app/admin/bi/bi_user/api"
)

type BiUserRouter struct{}

var biUserApi = api.BiUserApiApp

// InitBiUserRouter 初始化 BiUser 路由信息
func (r *BiUserRouter) InitBiUserRouter(Router *gin.RouterGroup) {
	// biUserRouter := Router.Group("biUser").Use(middleware.JWTAuthFilter())
	biUserRouter := Router.Group("biUser")
	biUserRouterWithoutRecord := Router.Group("biUser")
	{
		biUserRouter.POST("create", biUserApi.Create)             // 新建BiUser
		biUserRouter.DELETE("delete", biUserApi.Delete)           // 删除BiUser
		biUserRouter.DELETE("deleteByIds", biUserApi.DeleteByIds) // 批量删除BiUser
		biUserRouter.POST("update", biUserApi.Update)             // 更新BiUser
	}
	{
		biUserRouterWithoutRecord.GET("get", biUserApi.Get)   // 根据ID获取BiUser
		biUserRouterWithoutRecord.GET("find", biUserApi.Find) // 获取BiUser列表
	}
}
