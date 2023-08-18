// Package router SysDictDataRouter 自动生成模板
// @description <TODO description class purpose>
// @author
// @File: sys_dict_data
// @version 1.0.0
// @create 2023-08-18 13:41:26
package router

import (
	"github.com/gin-gonic/gin"
	"manager-gin/src/app/admin/sys/sys_dict_data/api"
	"manager-gin/src/middleware"
)

type SysDictDataRouter struct{}

var sysDictDataApi = api.SysDictDataApiApp

// InitSysDictDataRouter 初始化 SysDictData 路由信息
func (r *SysDictDataRouter) InitSysDictDataRouter(Router *gin.RouterGroup) {
	sysDictDataRouter := Router.Group("sysDictData")
	sysDictDataRouter.Use(middleware.JWTAuthFilter())
	sysDictDataRouterWithoutRecord := Router.Group("sysDictData")
	{
		sysDictDataRouter.POST("create", sysDictDataApi.Create)             // 新建SysDictData
		sysDictDataRouter.DELETE("delete", sysDictDataApi.Delete)           // 删除SysDictData
		sysDictDataRouter.DELETE("deleteByIds", sysDictDataApi.DeleteByIds) // 批量删除SysDictData
		sysDictDataRouter.POST("update", sysDictDataApi.Update)             // 更新SysDictData
	}
	{
		sysDictDataRouterWithoutRecord.GET("get", sysDictDataApi.Get)   // 根据ID获取SysDictData
		sysDictDataRouterWithoutRecord.GET("find", sysDictDataApi.Find) // 获取SysDictData列表
	}
}
