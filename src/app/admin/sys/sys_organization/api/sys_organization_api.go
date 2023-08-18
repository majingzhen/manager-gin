// Package api  SysOrganizationApi 自动生成模板
// @description <TODO description class purpose>
// @author
// @File: sys_organization
// @version 1.0.0
// @create 2023-08-18 14:00:53
package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"manager-gin/src/app/admin/sys/sys_organization/service"
	"manager-gin/src/app/admin/sys/sys_organization/service/view"
	"manager-gin/src/common"
	response "manager-gin/src/common/response"
	"manager-gin/src/global"
	"manager-gin/src/utils"
)

type SysOrganizationApi struct {
}

var sysOrganizationService = service.SysOrganizationServiceApp

// Create 创建SysOrganization
// @Summary 创建SysOrganization
// @Router /sysOrganization/create [post]
func (sysOrganizationApi *SysOrganizationApi) Create(c *gin.Context) {
	var sysOrganizationView view.SysOrganizationView
	_ = c.ShouldBindJSON(&sysOrganizationView)
	sysOrganizationView.Id = utils.GenUID()
	sysOrganizationView.CreateTime = utils.GetCurTimeStr()
	sysOrganizationView.UpdateTime = utils.GetCurTimeStr()
	if err := sysOrganizationService.Create(&sysOrganizationView); err != nil {
		global.Logger.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// Delete 删除SysOrganization
// @Summary 删除SysOrganization
// @Router /sysOrganization/delete [delete]
func (sysOrganizationApi *SysOrganizationApi) Delete(c *gin.Context) {
	var id common.Id
	_ = c.ShouldBindJSON(&id)
	if err := sysOrganizationService.Delete(id.ID); err != nil {
		global.Logger.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteByIds 批量删除SysOrganization
// @Summary 批量删除SysOrganization
// @Router /sysOrganization/deleteByIds [delete]
func (sysOrganizationApi *SysOrganizationApi) DeleteByIds(c *gin.Context) {
	var ids common.Ids
	_ = c.ShouldBindJSON(&ids)
	if err := sysOrganizationService.DeleteByIds(ids.Ids); err != nil {
		global.Logger.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// Update 更新SysOrganization
// @Summary 更新SysOrganization
// @Router /sysOrganization/update [put]
func (sysOrganizationApi *SysOrganizationApi) Update(c *gin.Context) {
	var sysOrganizationView view.SysOrganizationView
	_ = c.ShouldBindJSON(&sysOrganizationView)
	id := sysOrganizationView.Id
	if id == "" {
		response.FailWithMessage("更新失败", c)
	}
	sysOrganizationView.UpdateTime = utils.GetCurTimeStr()
	if err := sysOrganizationService.Update(id, &sysOrganizationView); err != nil {
		global.Logger.Error("更新持久化失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// Get 用id查询SysOrganization
// @Summary 用id查询SysOrganization
// @Router /sysOrganization/get [get]
func (sysOrganizationApi *SysOrganizationApi) Get(c *gin.Context) {
	id := c.Query("id")
	if err, sysOrganizationView := sysOrganizationService.Get(id); err != nil {
		global.Logger.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"sysOrganizationView": sysOrganizationView}, c)
	}
}

// Find 分页获取SysOrganization列表
// @Summary 分页获取SysOrganization列表
// @Router /sysOrganization/find [get]
func (sysOrganizationApi *SysOrganizationApi) Find(c *gin.Context) {
	var pageInfo common.PageInfoV2
	// 绑定查询参数到 pageInfo
	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		response.FailWithMessage("获取分页数据解析失败!", c)
	}
	// 调用 Calculate 方法自动计算 Limit 和 Offset
	pageInfo.Calculate()
	if err := sysOrganizationService.Find(&pageInfo); err != nil {
		global.Logger.Error("获取分页信息失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(pageInfo, "获取成功", c)
	}
}
