// Package api  SysJobApi 自动生成模板
// @description <TODO description class purpose>
// @author
// @File: sys_job
// @version 1.0.0
// @create 2023-08-18 13:41:26
package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"manager-gin/src/app/admin/sys/sys_job/service"
	"manager-gin/src/app/admin/sys/sys_job/service/view"
	"manager-gin/src/common"
	response "manager-gin/src/common/response"
	"manager-gin/src/global"
	"manager-gin/src/utils"
)

type SysJobApi struct {
}

var sysJobService = service.SysJobServiceApp

// Create 创建SysJob
// @Summary 创建SysJob
// @Router /sysJob/create [post]
func (sysJobApi *SysJobApi) Create(c *gin.Context) {
	var sysJobView view.SysJobView
	_ = c.ShouldBindJSON(&sysJobView)
	sysJobView.Id = utils.GenUID()
	sysJobView.CreateTime = utils.GetCurTimeStr()
	sysJobView.UpdateTime = utils.GetCurTimeStr()
	if err := sysJobService.Create(&sysJobView); err != nil {
		global.Logger.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// Delete 删除SysJob
// @Summary 删除SysJob
// @Router /sysJob/delete [delete]
func (sysJobApi *SysJobApi) Delete(c *gin.Context) {
	var id common.Id
	_ = c.ShouldBindJSON(&id)
	if err := sysJobService.Delete(id.ID); err != nil {
		global.Logger.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteByIds 批量删除SysJob
// @Summary 批量删除SysJob
// @Router /sysJob/deleteByIds [delete]
func (sysJobApi *SysJobApi) DeleteByIds(c *gin.Context) {
	var ids common.Ids
	_ = c.ShouldBindJSON(&ids)
	if err := sysJobService.DeleteByIds(ids.Ids); err != nil {
		global.Logger.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// Update 更新SysJob
// @Summary 更新SysJob
// @Router /sysJob/update [put]
func (sysJobApi *SysJobApi) Update(c *gin.Context) {
	var sysJobView view.SysJobView
	_ = c.ShouldBindJSON(&sysJobView)
	id := sysJobView.Id
	if id == "" {
		response.FailWithMessage("更新失败", c)
	}
	sysJobView.UpdateTime = utils.GetCurTimeStr()
	if err := sysJobService.Update(id, &sysJobView); err != nil {
		global.Logger.Error("更新持久化失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// Get 用id查询SysJob
// @Summary 用id查询SysJob
// @Router /sysJob/get [get]
func (sysJobApi *SysJobApi) Get(c *gin.Context) {
	id := c.Query("id")
	if err, sysJobView := sysJobService.Get(id); err != nil {
		global.Logger.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"sysJobView": sysJobView}, c)
	}
}

// Find 分页获取SysJob列表
// @Summary 分页获取SysJob列表
// @Router /sysJob/find [get]
func (sysJobApi *SysJobApi) Find(c *gin.Context) {
	var pageInfo common.PageInfoV2
	// 绑定查询参数到 pageInfo
	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		response.FailWithMessage("获取分页数据解析失败!", c)
	}
	// 调用 Calculate 方法自动计算 Limit 和 Offset
	pageInfo.Calculate()
	if err := sysJobService.Find(&pageInfo); err != nil {
		global.Logger.Error("获取分页信息失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(pageInfo, "获取成功", c)
	}
}
