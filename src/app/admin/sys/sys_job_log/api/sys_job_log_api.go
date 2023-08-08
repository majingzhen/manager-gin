// Package api  SysJobLogApi 自动生成模板
// @description <TODO description class purpose>
// @author
// @File: sys_job_log
// @version 1.0.0
// @create 2023-08-08 10:06:19
package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"manager-gin/src/app/admin/sys/sys_job_log/service"
	"manager-gin/src/app/admin/sys/sys_job_log/service/view"
	"manager-gin/src/common"
	response "manager-gin/src/common/response"
	"manager-gin/src/global"
	"manager-gin/src/utils"
	"strconv"
)

type SysJobLogApi struct {
}

var sysJobLogService = service.SysJobLogServiceApp

// Create 创建SysJobLog
// @Summary 创建SysJobLog
// @Router /sysJobLog/create [post]
func (sysJobLogApi *SysJobLogApi) Create(c *gin.Context) {
	var sysJobLogView view.SysJobLogView
	_ = c.ShouldBindJSON(&sysJobLogView)
	sysJobLogView.CreateTime = utils.GetCurTimeStr()
	if err := sysJobLogService.Create(&sysJobLogView); err != nil {
		global.Logger.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// Delete 删除SysJobLog
// @Summary 删除SysJobLog
// @Router /sysJobLog/delete [delete]
func (sysJobLogApi *SysJobLogApi) Delete(c *gin.Context) {
	var id int
	_ = c.ShouldBindJSON(&id)
	if err := sysJobLogService.Delete(id); err != nil {
		global.Logger.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteByIds 批量删除SysJobLog
// @Summary 批量删除SysJobLog
// @Router /sysJobLog/deleteByIds [delete]
func (sysJobLogApi *SysJobLogApi) DeleteByIds(c *gin.Context) {
	var ids []int
	_ = c.ShouldBindJSON(&ids)
	if err := sysJobLogService.DeleteByIds(ids); err != nil {
		global.Logger.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// Update 更新SysJobLog
// @Summary 更新SysJobLog
// @Router /sysJobLog/update [put]
func (sysJobLogApi *SysJobLogApi) Update(c *gin.Context) {
	id := c.Query("id")
	atoi, err := strconv.Atoi(id)
	if err != nil {
		global.Logger.Error("更新解析上报数据失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	}
	sysJobLogViewJson := c.Query("sysJobLogView")
	var sysJobLogView view.SysJobLogView
	err = json.Unmarshal([]byte(sysJobLogViewJson), &sysJobLogView)
	if err := sysJobLogService.Update(atoi, &sysJobLogView); err != nil {
		global.Logger.Error("更新解析上报数据失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	}
	if err = sysJobLogService.Update(atoi, &sysJobLogView); err != nil {
		global.Logger.Error("更新持久化失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// Get 用id查询SysJobLog
// @Summary 用id查询SysJobLog
// @Router /sysJobLog/get [get]
func (sysJobLogApi *SysJobLogApi) Get(c *gin.Context) {
	id := c.Query("id")
	atoi, err := strconv.Atoi(id)
	if err != nil {
		global.Logger.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	}
	if err, sysJobLogView := sysJobLogService.Get(atoi); err != nil {
		global.Logger.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"sysJobLogView": sysJobLogView}, c)
	}
}

// Find 分页获取SysJobLog列表
// @Summary 分页获取SysJobLog列表
// @Router /sysJobLog/find [get]
func (sysJobLogApi *SysJobLogApi) Find(c *gin.Context) {
	var pageInfo common.PageInfoV2
	_ = c.ShouldBindQuery(&pageInfo)
	if err := sysJobLogService.Find(&pageInfo); err != nil {
		global.Logger.Error("获取分页信息失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(pageInfo, "获取成功", c)
	}
}
