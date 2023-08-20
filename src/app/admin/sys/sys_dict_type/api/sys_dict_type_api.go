// Package api  SysDictTypeApi 自动生成模板
// @description <TODO description class purpose>
// @author
// @File: sys_dict_type
// @version 1.0.0
// @create 2023-08-18 13:41:26
package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"manager-gin/src/app/admin/sys/sys_dict_type/service"
	"manager-gin/src/app/admin/sys/sys_dict_type/service/view"
	response "manager-gin/src/common/response"
	"manager-gin/src/global"
	"manager-gin/src/utils"
	"strings"
)

type SysDictTypeApi struct {
}

var sysDictTypeService = service.SysDictTypeServiceApp

// Create 创建SysDictType
// @Summary 创建SysDictType
// @Router /sysDictType/create [post]
func (sysDictTypeApi *SysDictTypeApi) Create(c *gin.Context) {
	var sysDictTypeView view.SysDictTypeView
	_ = c.ShouldBindJSON(&sysDictTypeView)
	sysDictTypeView.Id = utils.GenUID()
	sysDictTypeView.CreateTime = utils.GetCurTimeStr()
	sysDictTypeView.UpdateTime = utils.GetCurTimeStr()
	if err := sysDictTypeService.Create(&sysDictTypeView); err != nil {
		global.Logger.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// Delete 删除SysDictType
// @Summary 删除SysDictType
// @Router /sysDictType/delete [delete]
func (sysDictTypeApi *SysDictTypeApi) Delete(c *gin.Context) {
	idStr := c.Param("ids")
	ids := strings.Split(idStr, ",")
	if err := sysDictTypeService.DeleteByIds(ids); err != nil {
		global.Logger.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// Update 更新SysDictType
// @Summary 更新SysDictType
// @Router /sysDictType/update [put]
func (sysDictTypeApi *SysDictTypeApi) Update(c *gin.Context) {
	var sysDictTypeView view.SysDictTypeView
	_ = c.ShouldBindJSON(&sysDictTypeView)
	id := sysDictTypeView.Id
	if id == "" {
		response.FailWithMessage("更新失败", c)
	}
	sysDictTypeView.UpdateTime = utils.GetCurTimeStr()
	if err := sysDictTypeService.Update(id, &sysDictTypeView); err != nil {
		global.Logger.Error("更新持久化失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// Get 用id查询SysDictType
// @Summary 用id查询SysDictType
// @Router /sysDictType/get [get]
func (sysDictTypeApi *SysDictTypeApi) Get(c *gin.Context) {
	id := c.Param("id")
	if err, sysDictTypeView := sysDictTypeService.Get(id); err != nil {
		global.Logger.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(sysDictTypeView, c)
	}
}

// List 分页获取SysDictType列表
// @Summary 分页获取SysDictType列表
// @Router /sysDictType/list [get]
func (sysDictTypeApi *SysDictTypeApi) List(c *gin.Context) {
	var pageInfo view.SysDictTypePageView
	// 绑定查询参数到 pageInfo
	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		response.FailWithMessage("获取分页数据解析失败!", c)
	}
	if err, res := sysDictTypeService.List(&pageInfo); err != nil {
		global.Logger.Error("获取分页信息失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(res, "获取成功", c)
	}
}

// SelectDictTypeAll 获取DictType全部数据
// @Summary 获取DictType全部数据
// @Router /sysDictType/list [get]
func (sysDictTypeApi *SysDictTypeApi) SelectDictTypeAll(c *gin.Context) {
	if err, res := sysDictTypeService.SelectDictTypeAll(); err != nil {
		global.Logger.Error("获取分页信息失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(res, "获取成功", c)
	}
}
