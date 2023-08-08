// Package api  SysJobApi 自动生成模板
// @description <TODO description class purpose>
// @author
// @File: sys_job
// @version 1.0.0
// @create 2023-08-08 10:06:19
package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"manager-gin/src/app/admin/sys/sys_job/service"
	"manager-gin/src/app/admin/sys/sys_job/service/view"
	"manager-gin/src/common"
	response "manager-gin/src/common/response"
	"manager-gin/src/global"
	"manager-gin/src/utils"
	"strconv"
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
	var id int
	_ = c.ShouldBindJSON(&id)
	if err := sysJobService.Delete(id); err != nil {
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
	var ids []int
	_ = c.ShouldBindJSON(&ids)
	if err := sysJobService.DeleteByIds(ids); err != nil {
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
	id := c.Query("id")
	atoi, err := strconv.Atoi(id)
	if err != nil {
		global.Logger.Error("更新解析上报数据失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	}
	sysJobViewJson := c.Query("sysJobView")
	var sysJobView view.SysJobView
	err = json.Unmarshal([]byte(sysJobViewJson), &sysJobView)
	sysJobView.UpdateTime = utils.GetCurTimeStr()
	if err := sysJobService.Update(atoi, &sysJobView); err != nil {
		global.Logger.Error("更新解析上报数据失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	}
	if err = sysJobService.Update(atoi, &sysJobView); err != nil {
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
	atoi, err := strconv.Atoi(id)
	if err != nil {
		global.Logger.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	}
	if err, sysJobView := sysJobService.Get(atoi); err != nil {
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
	_ = c.ShouldBindQuery(&pageInfo)
	if err := sysJobService.Find(&pageInfo); err != nil {
		global.Logger.Error("获取分页信息失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(pageInfo, "获取成功", c)
	}
}
