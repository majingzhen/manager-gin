// Package api  SysNoticeApi 自动生成模板
// @description <TODO description class purpose>
// @author
// @File: sys_notice
// @version 1.0.0
// @create 2023-08-08 10:06:19
package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"manager-gin/src/app/admin/sys/sys_notice/service"
	"manager-gin/src/app/admin/sys/sys_notice/service/view"
	"manager-gin/src/common"
	response "manager-gin/src/common/response"
	"manager-gin/src/global"
	"manager-gin/src/utils"
	"strconv"
)

type SysNoticeApi struct {
}

var sysNoticeService = service.SysNoticeServiceApp

// Create 创建SysNotice
// @Summary 创建SysNotice
// @Router /sysNotice/create [post]
func (sysNoticeApi *SysNoticeApi) Create(c *gin.Context) {
	var sysNoticeView view.SysNoticeView
	_ = c.ShouldBindJSON(&sysNoticeView)
	sysNoticeView.CreateTime = utils.GetCurTimeStr()
	sysNoticeView.UpdateTime = utils.GetCurTimeStr()
	if err := sysNoticeService.Create(&sysNoticeView); err != nil {
		global.Logger.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// Delete 删除SysNotice
// @Summary 删除SysNotice
// @Router /sysNotice/delete [delete]
func (sysNoticeApi *SysNoticeApi) Delete(c *gin.Context) {
	var id int
	_ = c.ShouldBindJSON(&id)
	if err := sysNoticeService.Delete(id); err != nil {
		global.Logger.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteByIds 批量删除SysNotice
// @Summary 批量删除SysNotice
// @Router /sysNotice/deleteByIds [delete]
func (sysNoticeApi *SysNoticeApi) DeleteByIds(c *gin.Context) {
	var ids []int
	_ = c.ShouldBindJSON(&ids)
	if err := sysNoticeService.DeleteByIds(ids); err != nil {
		global.Logger.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// Update 更新SysNotice
// @Summary 更新SysNotice
// @Router /sysNotice/update [put]
func (sysNoticeApi *SysNoticeApi) Update(c *gin.Context) {
	id := c.Query("id")
	atoi, err := strconv.Atoi(id)
	if err != nil {
		global.Logger.Error("更新解析上报数据失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	}
	sysNoticeViewJson := c.Query("sysNoticeView")
	var sysNoticeView view.SysNoticeView
	err = json.Unmarshal([]byte(sysNoticeViewJson), &sysNoticeView)
	sysNoticeView.UpdateTime = utils.GetCurTimeStr()
	if err := sysNoticeService.Update(atoi, &sysNoticeView); err != nil {
		global.Logger.Error("更新解析上报数据失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	}
	if err = sysNoticeService.Update(atoi, &sysNoticeView); err != nil {
		global.Logger.Error("更新持久化失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// Get 用id查询SysNotice
// @Summary 用id查询SysNotice
// @Router /sysNotice/get [get]
func (sysNoticeApi *SysNoticeApi) Get(c *gin.Context) {
	id := c.Query("id")
	atoi, err := strconv.Atoi(id)
	if err != nil {
		global.Logger.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	}
	if err, sysNoticeView := sysNoticeService.Get(atoi); err != nil {
		global.Logger.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"sysNoticeView": sysNoticeView}, c)
	}
}

// Find 分页获取SysNotice列表
// @Summary 分页获取SysNotice列表
// @Router /sysNotice/find [get]
func (sysNoticeApi *SysNoticeApi) Find(c *gin.Context) {
	var pageInfo common.PageInfoV2
	_ = c.ShouldBindQuery(&pageInfo)
	if err := sysNoticeService.Find(&pageInfo); err != nil {
		global.Logger.Error("获取分页信息失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(pageInfo, "获取成功", c)
	}
}
