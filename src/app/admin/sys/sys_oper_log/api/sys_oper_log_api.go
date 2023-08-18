// Package api  SysOperLogApi 自动生成模板
// @description <TODO description class purpose>
// @author
// @File: sys_oper_log
// @version 1.0.0
// @create 2023-08-18 13:41:26
package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"manager-gin/src/app/admin/sys/sys_oper_log/service"
	"manager-gin/src/app/admin/sys/sys_oper_log/service/view"
	"manager-gin/src/common"
	response "manager-gin/src/common/response"
	"manager-gin/src/global"
	"manager-gin/src/utils"
)

type SysOperLogApi struct {
}

var sysOperLogService = service.SysOperLogServiceApp

// Create 创建SysOperLog
// @Summary 创建SysOperLog
// @Router /sysOperLog/create [post]
func (sysOperLogApi *SysOperLogApi) Create(c *gin.Context) {
	var sysOperLogView view.SysOperLogView
	_ = c.ShouldBindJSON(&sysOperLogView)
	sysOperLogView.Id = utils.GenUID()
	if err := sysOperLogService.Create(&sysOperLogView); err != nil {
		global.Logger.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// Delete 删除SysOperLog
// @Summary 删除SysOperLog
// @Router /sysOperLog/delete [delete]
func (sysOperLogApi *SysOperLogApi) Delete(c *gin.Context) {
	var id common.Id
	_ = c.ShouldBindJSON(&id)
	if err := sysOperLogService.Delete(id.ID); err != nil {
		global.Logger.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteByIds 批量删除SysOperLog
// @Summary 批量删除SysOperLog
// @Router /sysOperLog/deleteByIds [delete]
func (sysOperLogApi *SysOperLogApi) DeleteByIds(c *gin.Context) {
	var ids common.Ids
	_ = c.ShouldBindJSON(&ids)
	if err := sysOperLogService.DeleteByIds(ids.Ids); err != nil {
		global.Logger.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// Update 更新SysOperLog
// @Summary 更新SysOperLog
// @Router /sysOperLog/update [put]
func (sysOperLogApi *SysOperLogApi) Update(c *gin.Context) {
	var sysOperLogView view.SysOperLogView
	_ = c.ShouldBindJSON(&sysOperLogView)
	id := sysOperLogView.Id
	if id == "" {
		response.FailWithMessage("更新失败", c)
	}
	if err := sysOperLogService.Update(id, &sysOperLogView); err != nil {
		global.Logger.Error("更新持久化失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// Get 用id查询SysOperLog
// @Summary 用id查询SysOperLog
// @Router /sysOperLog/get [get]
func (sysOperLogApi *SysOperLogApi) Get(c *gin.Context) {
	id := c.Query("id")
	if err, sysOperLogView := sysOperLogService.Get(id); err != nil {
		global.Logger.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"sysOperLogView": sysOperLogView}, c)
	}
}

// Find 分页获取SysOperLog列表
// @Summary 分页获取SysOperLog列表
// @Router /sysOperLog/find [get]
func (sysOperLogApi *SysOperLogApi) Find(c *gin.Context) {
	var pageInfo common.PageInfoV2
	// 绑定查询参数到 pageInfo
	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		response.FailWithMessage("获取分页数据解析失败!", c)
	}
	// 调用 Calculate 方法自动计算 Limit 和 Offset
	pageInfo.Calculate()
	if err := sysOperLogService.Find(&pageInfo); err != nil {
		global.Logger.Error("获取分页信息失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(pageInfo, "获取成功", c)
	}
}
