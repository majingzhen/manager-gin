// Package router SysDictTypeRouter 自动生成模板
// @description <TODO description class purpose>
// @author
// @File: sys_dict_type
// @version 1.0.0
// @create 2023-08-18 13:41:26
package router

import (
	"github.com/gin-gonic/gin"
	"manager-gin/src/app/admin/sys/api"
	"manager-gin/src/middleware"
)

type SysDictTypeRouter struct {
	sysDictTypeApi api.SysDictTypeApi
}

// InitSysDictTypeRouter 初始化 SysDictType 路由信息
func (r *SysDictTypeRouter) InitSysDictTypeRouter(Router *gin.RouterGroup) {
	sysDictTypeRouter := Router.Group("sysDictType").Use(middleware.JWTAuthFilter())
	sysDictTypeRouterWithoutRecord := Router.Group("sysDictType").Use(middleware.JWTAuthFilter())
	{
		sysDictTypeRouter.POST("create", r.sysDictTypeApi.Create)        // 新建SysDictType
		sysDictTypeRouter.DELETE("delete/:ids", r.sysDictTypeApi.Delete) // 删除SysDictType
		sysDictTypeRouter.POST("update", r.sysDictTypeApi.Update)        // 更新SysDictType
	}
	{
		sysDictTypeRouterWithoutRecord.GET("get/:id", r.sysDictTypeApi.Get)                    // 根据ID获取SysDictType
		sysDictTypeRouterWithoutRecord.GET("page", r.sysDictTypeApi.Page)                      // 获取SysDictType列表
		sysDictTypeRouterWithoutRecord.GET("optionSelect", r.sysDictTypeApi.SelectDictTypeAll) // 获取SysDictType列表
	}
}
