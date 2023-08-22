// Package api  SysDictDataApi 自动生成模板
// @description <TODO description class purpose>
// @author
// @File: sys_dict_data
// @version 1.0.0
// @create 2023-08-20 19:08:42
package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"manager-gin/src/app/admin/sys/sys_dict_data/service"
	"manager-gin/src/app/admin/sys/sys_dict_data/service/view"
	response "manager-gin/src/common/response"
	"manager-gin/src/framework"
	"manager-gin/src/global"
	"manager-gin/src/utils"
	"strings"
)

type SysDictDataApi struct {
}

var sysDictDataService = service.SysDictDataServiceApp

// Create 创建SysDictData
// @Summary 创建SysDictData
// @Router /sysDictData/create [post]
func (api *SysDictDataApi) Create(c *gin.Context) {
	var sysDictDataView view.SysDictDataView
	_ = c.ShouldBindJSON(&sysDictDataView)
	sysDictDataView.Id = utils.GenUID()
	sysDictDataView.CreateTime = utils.GetCurTimeStr()
	sysDictDataView.UpdateTime = utils.GetCurTimeStr()
	sysDictDataView.CreateBy = framework.GetLoginUser(c).UserName
	if err := sysDictDataService.Create(&sysDictDataView); err != nil {
		global.Logger.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// Delete 删除SysDictData
// @Summary 删除SysDictData
// @Router /sysDictData/delete [delete]
func (api *SysDictDataApi) Delete(c *gin.Context) {
	idStr := c.Param("ids")
	ids := strings.Split(idStr, ",")
	if err := sysDictDataService.DeleteByIds(ids); err != nil {
		global.Logger.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// Update 更新SysDictData
// @Summary 更新SysDictData
// @Router /sysDictData/update [put]
func (api *SysDictDataApi) Update(c *gin.Context) {
	var sysDictDataView view.SysDictDataView
	_ = c.ShouldBindJSON(&sysDictDataView)
	id := sysDictDataView.Id
	if id == "" {
		response.FailWithMessage("更新失败", c)
		return
	}
	sysDictDataView.UpdateTime = utils.GetCurTimeStr()
	sysDictDataView.UpdateBy = framework.GetLoginUser(c).UserName
	if err := sysDictDataService.Update(id, &sysDictDataView); err != nil {
		global.Logger.Error("更新持久化失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// Get 用id查询SysDictData
// @Summary 用id查询SysDictData
// @Router /sysDictData/get [get]
func (api *SysDictDataApi) Get(c *gin.Context) {
	id := c.Param("id")
	if err, sysDictDataView := sysDictDataService.Get(id); err != nil {
		global.Logger.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(sysDictDataView, c)
	}
}

// GetByType 用字典类型查询SysDictData
// @Summary 用字典类型查询SysDictData
// @Router /sysDictData/type [get]
func (api *SysDictDataApi) GetByType(c *gin.Context) {
	dictType := c.Param("type")
	if err, sysDictDataView := sysDictDataService.GetByType(dictType); err != nil {
		global.Logger.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(sysDictDataView, c)
	}
}

// Page 分页获取SysDictData列表
// @Summary 分页获取SysDictData列表
// @Router /sysDictData/list [get]
func (api *SysDictDataApi) Page(c *gin.Context) {
	var pageInfo view.SysDictDataPageView
	// 绑定查询参数到 pageInfo
	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		response.FailWithMessage("获取分页数据解析失败!", c)
		return
	}

	if err, res := sysDictDataService.Page(&pageInfo); err != nil {
		global.Logger.Error("获取分页信息失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(res, "获取成功", c)
	}
}
