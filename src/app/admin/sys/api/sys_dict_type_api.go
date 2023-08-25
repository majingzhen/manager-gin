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
	"manager-gin/src/app/admin/sys/service/sys_dict_type"
	"manager-gin/src/app/admin/sys/service/sys_dict_type/view"
	response "manager-gin/src/common/response"
	"manager-gin/src/framework"
	"manager-gin/src/global"
	"manager-gin/src/utils"
	"strings"
)

type SysDictTypeApi struct {
	sysDictTypeService sys_dict_type.SysDictTypeService
}

// Create 创建SysDictType
// @Summary 创建SysDictType
// @Router /sysDictType/create [post]
func (api *SysDictTypeApi) Create(c *gin.Context) {
	var sysDictTypeView view.SysDictTypeView
	_ = c.ShouldBindJSON(&sysDictTypeView)
	sysDictTypeView.Id = utils.GenUID()
	sysDictTypeView.CreateTime = utils.GetCurTimeStr()
	sysDictTypeView.UpdateTime = utils.GetCurTimeStr()
	sysDictTypeView.CreateBy = framework.GetLoginUser(c).UserName
	if err := api.sysDictTypeService.Create(&sysDictTypeView); err != nil {
		global.Logger.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// Delete 删除SysDictType
// @Summary 删除SysDictType
// @Router /sysDictType/delete [delete]
func (api *SysDictTypeApi) Delete(c *gin.Context) {
	idStr := c.Param("ids")
	ids := strings.Split(idStr, ",")
	if err := api.sysDictTypeService.DeleteByIds(ids); err != nil {
		global.Logger.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// Update 更新SysDictType
// @Summary 更新SysDictType
// @Router /sysDictType/update [put]
func (api *SysDictTypeApi) Update(c *gin.Context) {
	var sysDictTypeView view.SysDictTypeView
	_ = c.ShouldBindJSON(&sysDictTypeView)
	id := sysDictTypeView.Id
	if id == "" {
		response.FailWithMessage("更新失败", c)
		return
	}
	sysDictTypeView.UpdateTime = utils.GetCurTimeStr()
	sysDictTypeView.UpdateBy = framework.GetLoginUser(c).UserName
	if err := api.sysDictTypeService.Update(id, &sysDictTypeView); err != nil {
		global.Logger.Error("更新持久化失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// Get 用id查询SysDictType
// @Summary 用id查询SysDictType
// @Router /sysDictType/get [get]
func (api *SysDictTypeApi) Get(c *gin.Context) {
	id := c.Param("id")
	if err, sysDictTypeView := api.sysDictTypeService.Get(id); err != nil {
		global.Logger.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(sysDictTypeView, c)
	}
}

// Page 分页获取SysDictType列表
// @Summary 分页获取SysDictType列表
// @Router /sysDictType/list [get]
func (api *SysDictTypeApi) Page(c *gin.Context) {
	var pageInfo view.SysDictTypePageView
	// 绑定查询参数到 pageInfo
	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		response.FailWithMessage("获取分页数据解析失败!", c)
		return
	}
	if err, res := api.sysDictTypeService.Page(&pageInfo); err != nil {
		global.Logger.Error("获取分页信息失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(res, "获取成功", c)
	}
}

// SelectDictTypeAll 获取DictType全部数据
// @Summary 获取DictType全部数据
// @Router /sysDictType/list [get]
func (api *SysDictTypeApi) SelectDictTypeAll(c *gin.Context) {
	if err, res := api.sysDictTypeService.SelectDictTypeAll(); err != nil {
		global.Logger.Error("获取分页信息失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(res, "获取成功", c)
	}
}
