// Package api  SysNoticeApi 自动生成模板
// @description <TODO description class purpose>
// @author
// @File: sys_notice
// @version 1.0.0
// @create 2023-08-18 13:41:26
package api

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"manager-gin/src/app/admin/sys/sys_notice/service"
	"manager-gin/src/app/admin/sys/sys_notice/service/view"
	"manager-gin/src/common"
	response "manager-gin/src/common/response"
	"manager-gin/src/global"
	"manager-gin/src/utils"
	"strings"
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
	sysNoticeView.Id = strings.ReplaceAll(uuid.NewV4().String(), "-", "")
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
	var id common.Id
	_ = c.ShouldBindJSON(&id)
	if err := sysNoticeService.Delete(id.ID); err != nil {
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
	var ids common.Ids
	_ = c.ShouldBindJSON(&ids)
	if err := sysNoticeService.DeleteByIds(ids.Ids); err != nil {
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
	var sysNoticeView view.SysNoticeView
	_ = c.ShouldBindJSON(&sysNoticeView)
	id := sysNoticeView.Id
	if id == "" {
		response.FailWithMessage("更新失败", c)
	}
	sysNoticeView.UpdateTime = utils.GetCurTimeStr()
	if err := sysNoticeService.Update(id, &sysNoticeView); err != nil {
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
	if err, sysNoticeView := sysNoticeService.Get(id); err != nil {
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
	// 绑定查询参数到 pageInfo
	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		response.FailWithMessage("获取分页数据解析失败!", c)
	}
	// 调用 Calculate 方法自动计算 Limit 和 Offset
	pageInfo.Calculate()
	if err := sysNoticeService.Find(&pageInfo); err != nil {
		global.Logger.Error("获取分页信息失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(pageInfo, "获取成功", c)
	}
}
