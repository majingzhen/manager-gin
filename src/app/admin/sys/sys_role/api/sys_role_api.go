// Package api  SysRoleApi 自动生成模板
// @description <TODO description class purpose>
// @author
// @File: sys_role
// @version 1.0.0
// @create 2023-08-18 14:00:53
package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"manager-gin/src/app/admin/sys/sys_role/service"
	"manager-gin/src/app/admin/sys/sys_role/service/view"
	"manager-gin/src/common"
	response "manager-gin/src/common/response"
	"manager-gin/src/global"
	"manager-gin/src/utils"
)

type SysRoleApi struct {
}

var sysRoleService = service.SysRoleServiceApp

// Create 创建SysRole
// @Summary 创建SysRole
// @Router /sysRole/create [post]
func (sysRoleApi *SysRoleApi) Create(c *gin.Context) {
	var sysRoleView view.SysRoleView
	_ = c.ShouldBindJSON(&sysRoleView)
	sysRoleView.Id = utils.GenUID()
	sysRoleView.CreateTime = utils.GetCurTimeStr()
	sysRoleView.UpdateTime = utils.GetCurTimeStr()
	if err := sysRoleService.Create(&sysRoleView); err != nil {
		global.Logger.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// Delete 删除SysRole
// @Summary 删除SysRole
// @Router /sysRole/delete [delete]
func (sysRoleApi *SysRoleApi) Delete(c *gin.Context) {
	var id common.Id
	_ = c.ShouldBindJSON(&id)
	if err := sysRoleService.Delete(id.ID); err != nil {
		global.Logger.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteByIds 批量删除SysRole
// @Summary 批量删除SysRole
// @Router /sysRole/deleteByIds [delete]
func (sysRoleApi *SysRoleApi) DeleteByIds(c *gin.Context) {
	var ids common.Ids
	_ = c.ShouldBindJSON(&ids)
	if err := sysRoleService.DeleteByIds(ids.Ids); err != nil {
		global.Logger.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// Update 更新SysRole
// @Summary 更新SysRole
// @Router /sysRole/update [put]
func (sysRoleApi *SysRoleApi) Update(c *gin.Context) {
	var sysRoleView view.SysRoleView
	_ = c.ShouldBindJSON(&sysRoleView)
	id := sysRoleView.Id
	if id == "" {
		response.FailWithMessage("更新失败", c)
	}
	sysRoleView.UpdateTime = utils.GetCurTimeStr()
	if err := sysRoleService.Update(id, &sysRoleView); err != nil {
		global.Logger.Error("更新持久化失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// Get 用id查询SysRole
// @Summary 用id查询SysRole
// @Router /sysRole/get [get]
func (sysRoleApi *SysRoleApi) Get(c *gin.Context) {
	id := c.Query("id")
	if err, sysRoleView := sysRoleService.Get(id); err != nil {
		global.Logger.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"sysRoleView": sysRoleView}, c)
	}
}

// Find 分页获取SysRole列表
// @Summary 分页获取SysRole列表
// @Router /sysRole/find [get]
func (sysRoleApi *SysRoleApi) Find(c *gin.Context) {
	var pageInfo common.PageInfoV2
	// 绑定查询参数到 pageInfo
	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		response.FailWithMessage("获取分页数据解析失败!", c)
	}
	// 调用 Calculate 方法自动计算 Limit 和 Offset
	pageInfo.Calculate()
	if err := sysRoleService.Find(&pageInfo); err != nil {
		global.Logger.Error("获取分页信息失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(pageInfo, "获取成功", c)
	}
}
