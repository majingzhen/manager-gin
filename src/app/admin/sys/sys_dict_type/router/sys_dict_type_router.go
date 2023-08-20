// Package router SysDictTypeRouter 自动生成模板
// @description <TODO description class purpose>
// @author
// @File: sys_dict_type
// @version 1.0.0
// @create 2023-08-18 13:41:26
package router

import (
	"github.com/gin-gonic/gin"
	"manager-gin/src/app/admin/sys/sys_dict_type/api"
	"manager-gin/src/middleware"
)

type SysDictTypeRouter struct{}

var sysDictTypeApi = api.SysDictTypeApiApp

// InitSysDictTypeRouter 初始化 SysDictType 路由信息
func (r *SysDictTypeRouter) InitSysDictTypeRouter(Router *gin.RouterGroup) {
	sysDictTypeRouter := Router.Group("sysDictType")
	sysDictTypeRouter.Use(middleware.JWTAuthFilter())
	sysDictTypeRouterWithoutRecord := Router.Group("sysDictType")
	{
		sysDictTypeRouter.POST("create", sysDictTypeApi.Create)        // 新建SysDictType
		sysDictTypeRouter.DELETE("delete/:ids", sysDictTypeApi.Delete) // 删除SysDictType
		sysDictTypeRouter.POST("update", sysDictTypeApi.Update)        // 更新SysDictType
	}
	{
		sysDictTypeRouterWithoutRecord.GET("get/:id", sysDictTypeApi.Get)                    // 根据ID获取SysDictType
		sysDictTypeRouterWithoutRecord.GET("list", sysDictTypeApi.List)                      // 获取SysDictType列表
		sysDictTypeRouterWithoutRecord.GET("optionSelect", sysDictTypeApi.SelectDictTypeAll) // 获取SysDictType列表
	}
}
