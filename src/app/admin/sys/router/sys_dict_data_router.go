// Package router SysDictDataRouter 自动生成模板
// @description <TODO description class purpose>
// @author
// @File: sys_dict_data
// @version 1.0.0
// @create 2023-08-20 19:08:42
package router

import (
	"github.com/gin-gonic/gin"
	"manager-gin/src/app/admin/sys/api"
	"manager-gin/src/middleware"
)

type SysDictDataRouter struct {
	sysDictDataApi api.SysDictDataApi
}

// InitSysDictDataRouter 初始化 SysDictData 路由信息
func (r *SysDictDataRouter) InitSysDictDataRouter(Router *gin.RouterGroup) {
	sysDictDataRouter := Router.Group("sysDictData").Use(middleware.JWTAuthFilter())
	sysDictDataRouterWithoutRecord := Router.Group("sysDictData").Use(middleware.JWTAuthFilter())
	{
		sysDictDataRouter.POST("create", r.sysDictDataApi.Create)        // 新建SysDictData
		sysDictDataRouter.DELETE("delete/:ids", r.sysDictDataApi.Delete) // 删除SysDictData
		sysDictDataRouter.POST("update", r.sysDictDataApi.Update)        // 更新SysDictData
	}
	{
		sysDictDataRouterWithoutRecord.GET("get/:id", r.sysDictDataApi.Get)          // 根据ID获取SysDictData
		sysDictDataRouterWithoutRecord.GET("type/:type", r.sysDictDataApi.GetByType) // 根据type获取SysDictData
		sysDictDataRouterWithoutRecord.GET("page", r.sysDictDataApi.Page)            // 获取SysDictData列表
	}
}
