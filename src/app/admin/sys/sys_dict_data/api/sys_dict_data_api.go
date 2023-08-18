// Package api  SysDictDataApi 自动生成模板
// @description <TODO description class purpose>
// @author
// @File: sys_dict_data
// @version 1.0.0
// @create 2023-08-18 13:41:26
package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"manager-gin/src/app/admin/sys/sys_dict_data/service"
	"manager-gin/src/app/admin/sys/sys_dict_data/service/view"
	"manager-gin/src/common"
	response "manager-gin/src/common/response"
	"manager-gin/src/global"
	"manager-gin/src/utils"
	"strconv"
)

type SysDictDataApi struct {
}

var sysDictDataService = service.SysDictDataServiceApp

// Create 创建SysDictData
// @Summary 创建SysDictData
// @Router /sysDictData/create [post]
func (sysDictDataApi *SysDictDataApi) Create(c *gin.Context) {
	var sysDictDataView view.SysDictDataView
	_ = c.ShouldBindJSON(&sysDictDataView)
	sysDictDataView.Id = strings.ReplaceAll(uuid.NewV4().String(), "-", "")
	sysDictDataView.CreateTime = utils.GetCurTimeStr()
	sysDictDataView.UpdateTime = utils.GetCurTimeStr()
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
func (sysDictDataApi *SysDictDataApi) Delete(c *gin.Context) {
	var id common.Id
	_ = c.ShouldBindJSON(&id)
	if err := sysDictDataService.Delete(id.ID); err != nil {
		global.Logger.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteByIds 批量删除SysDictData
// @Summary 批量删除SysDictData
// @Router /sysDictData/deleteByIds [delete]
func (sysDictDataApi *SysDictDataApi) DeleteByIds(c *gin.Context) {
	var ids common.Ids
	_ = c.ShouldBindJSON(&ids)
	if err := sysDictDataService.DeleteByIds(ids.Ids); err != nil {
		global.Logger.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// Update 更新SysDictData
// @Summary 更新SysDictData
// @Router /sysDictData/update [put]
func (sysDictDataApi *SysDictDataApi) Update(c *gin.Context) {
	var sysDictDataView view.SysDictDataView
	_ = c.ShouldBindJSON(&sysDictDataView)
	id := sysDictDataView.Id
	if id == "" {
		response.FailWithMessage("更新失败", c)
	}
	sysDictDataView.UpdateTime = utils.GetCurTimeStr()
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
func (sysDictDataApi *SysDictDataApi) Get(c *gin.Context) {
	id := c.Query("id")
	if err, sysDictDataView := sysDictDataService.Get(id); err != nil {
		global.Logger.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"sysDictDataView": sysDictDataView}, c)
	}
}

// Find 分页获取SysDictData列表
// @Summary 分页获取SysDictData列表
// @Router /sysDictData/find [get]
func (sysDictDataApi *SysDictDataApi) Find(c *gin.Context) {
	var pageInfo common.PageInfoV2
	// 绑定查询参数到 pageInfo
	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		response.FailWithMessage("获取分页数据解析失败!", c)
	}
	// 调用 Calculate 方法自动计算 Limit 和 Offset
	pageInfo.Calculate()
	if err := sysDictDataService.Find(&pageInfo); err != nil {
		global.Logger.Error("获取分页信息失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(pageInfo, "获取成功", c)
	}
}
