// Package api  SysUserRoleApi 自动生成模板
// @description <TODO description class purpose>
// @author
// @File: sys_user_role
// @version 1.0.0
// @create 2023-08-18 13:41:26
package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"manager-gin/src/app/admin/sys/sys_user_role/service"
	"manager-gin/src/app/admin/sys/sys_user_role/service/view"
	"manager-gin/src/common"
	response "manager-gin/src/common/response"
	"manager-gin/src/global"
	"manager-gin/src/utils"
	"strconv"
)

type SysUserRoleApi struct {
}

var sysUserRoleService = service.SysUserRoleServiceApp

// Create 创建SysUserRole
// @Summary 创建SysUserRole
// @Router /sysUserRole/create [post]
func (sysUserRoleApi *SysUserRoleApi) Create(c *gin.Context) {
	var sysUserRoleView view.SysUserRoleView
	_ = c.ShouldBindJSON(&sysUserRoleView)
	sysUserRoleView.Id = strings.ReplaceAll(uuid.NewV4().String(), "-", "")
	sysUserRoleView.CreateTime = utils.GetCurTimeStr()
	sysUserRoleView.UpdateTime = utils.GetCurTimeStr()
	if err := sysUserRoleService.Create(&sysUserRoleView); err != nil {
		global.Logger.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// Delete 删除SysUserRole
// @Summary 删除SysUserRole
// @Router /sysUserRole/delete [delete]
func (sysUserRoleApi *SysUserRoleApi) Delete(c *gin.Context) {
	var id common.Id
	_ = c.ShouldBindJSON(&id)
	if err := sysUserRoleService.Delete(id.ID); err != nil {
		global.Logger.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteByIds 批量删除SysUserRole
// @Summary 批量删除SysUserRole
// @Router /sysUserRole/deleteByIds [delete]
func (sysUserRoleApi *SysUserRoleApi) DeleteByIds(c *gin.Context) {
	var ids common.Ids
	_ = c.ShouldBindJSON(&ids)
	if err := sysUserRoleService.DeleteByIds(ids.Ids); err != nil {
		global.Logger.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// Update 更新SysUserRole
// @Summary 更新SysUserRole
// @Router /sysUserRole/update [put]
func (sysUserRoleApi *SysUserRoleApi) Update(c *gin.Context) {
	var sysUserRoleView view.SysUserRoleView
	_ = c.ShouldBindJSON(&sysUserRoleView)
	id := sysUserRoleView.Id
	if id == "" {
		response.FailWithMessage("更新失败", c)
	}
	sysUserRoleView.UpdateTime = utils.GetCurTimeStr()
	if err := sysUserRoleService.Update(id, &sysUserRoleView); err != nil {
		global.Logger.Error("更新持久化失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// Get 用id查询SysUserRole
// @Summary 用id查询SysUserRole
// @Router /sysUserRole/get [get]
func (sysUserRoleApi *SysUserRoleApi) Get(c *gin.Context) {
	id := c.Query("id")
	if err, sysUserRoleView := sysUserRoleService.Get(id); err != nil {
		global.Logger.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"sysUserRoleView": sysUserRoleView}, c)
	}
}

// Find 分页获取SysUserRole列表
// @Summary 分页获取SysUserRole列表
// @Router /sysUserRole/find [get]
func (sysUserRoleApi *SysUserRoleApi) Find(c *gin.Context) {
	var pageInfo common.PageInfoV2
	// 绑定查询参数到 pageInfo
	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		response.FailWithMessage("获取分页数据解析失败!", c)
	}
	// 调用 Calculate 方法自动计算 Limit 和 Offset
	pageInfo.Calculate()
	if err := sysUserRoleService.Find(&pageInfo); err != nil {
		global.Logger.Error("获取分页信息失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(pageInfo, "获取成功", c)
	}
}
