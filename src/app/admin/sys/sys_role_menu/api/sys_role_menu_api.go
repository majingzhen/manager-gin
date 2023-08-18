// Package api  SysRoleMenuApi 自动生成模板
// @description <TODO description class purpose>
// @author
// @File: sys_role_menu
// @version 1.0.0
// @create 2023-08-18 13:41:26
package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"manager-gin/src/app/admin/sys/sys_role_menu/service"
	"manager-gin/src/app/admin/sys/sys_role_menu/service/view"
	"manager-gin/src/common"
	response "manager-gin/src/common/response"
	"manager-gin/src/global"
	"manager-gin/src/utils"
	"strconv"
)

type SysRoleMenuApi struct {
}

var sysRoleMenuService = service.SysRoleMenuServiceApp

// Create 创建SysRoleMenu
// @Summary 创建SysRoleMenu
// @Router /sysRoleMenu/create [post]
func (sysRoleMenuApi *SysRoleMenuApi) Create(c *gin.Context) {
	var sysRoleMenuView view.SysRoleMenuView
	_ = c.ShouldBindJSON(&sysRoleMenuView)
	sysRoleMenuView.Id = strings.ReplaceAll(uuid.NewV4().String(), "-", "")
	sysRoleMenuView.CreateTime = utils.GetCurTimeStr()
	sysRoleMenuView.UpdateTime = utils.GetCurTimeStr()
	if err := sysRoleMenuService.Create(&sysRoleMenuView); err != nil {
		global.Logger.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// Delete 删除SysRoleMenu
// @Summary 删除SysRoleMenu
// @Router /sysRoleMenu/delete [delete]
func (sysRoleMenuApi *SysRoleMenuApi) Delete(c *gin.Context) {
	var id common.Id
	_ = c.ShouldBindJSON(&id)
	if err := sysRoleMenuService.Delete(id.ID); err != nil {
		global.Logger.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteByIds 批量删除SysRoleMenu
// @Summary 批量删除SysRoleMenu
// @Router /sysRoleMenu/deleteByIds [delete]
func (sysRoleMenuApi *SysRoleMenuApi) DeleteByIds(c *gin.Context) {
	var ids common.Ids
	_ = c.ShouldBindJSON(&ids)
	if err := sysRoleMenuService.DeleteByIds(ids.Ids); err != nil {
		global.Logger.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// Update 更新SysRoleMenu
// @Summary 更新SysRoleMenu
// @Router /sysRoleMenu/update [put]
func (sysRoleMenuApi *SysRoleMenuApi) Update(c *gin.Context) {
	var sysRoleMenuView view.SysRoleMenuView
	_ = c.ShouldBindJSON(&sysRoleMenuView)
	id := sysRoleMenuView.Id
	if id == "" {
		response.FailWithMessage("更新失败", c)
	}
	sysRoleMenuView.UpdateTime = utils.GetCurTimeStr()
	if err := sysRoleMenuService.Update(id, &sysRoleMenuView); err != nil {
		global.Logger.Error("更新持久化失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// Get 用id查询SysRoleMenu
// @Summary 用id查询SysRoleMenu
// @Router /sysRoleMenu/get [get]
func (sysRoleMenuApi *SysRoleMenuApi) Get(c *gin.Context) {
	id := c.Query("id")
	if err, sysRoleMenuView := sysRoleMenuService.Get(id); err != nil {
		global.Logger.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"sysRoleMenuView": sysRoleMenuView}, c)
	}
}

// Find 分页获取SysRoleMenu列表
// @Summary 分页获取SysRoleMenu列表
// @Router /sysRoleMenu/find [get]
func (sysRoleMenuApi *SysRoleMenuApi) Find(c *gin.Context) {
	var pageInfo common.PageInfoV2
	// 绑定查询参数到 pageInfo
	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		response.FailWithMessage("获取分页数据解析失败!", c)
	}
	// 调用 Calculate 方法自动计算 Limit 和 Offset
	pageInfo.Calculate()
	if err := sysRoleMenuService.Find(&pageInfo); err != nil {
		global.Logger.Error("获取分页信息失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(pageInfo, "获取成功", c)
	}
}
