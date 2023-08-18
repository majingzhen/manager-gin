// Package api  SysRoleDeptApi 自动生成模板
// @description <TODO description class purpose>
// @author
// @File: sys_role_dept
// @version 1.0.0
// @create 2023-08-18 13:41:26
package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"manager-gin/src/app/admin/sys/sys_role_dept/service"
	"manager-gin/src/app/admin/sys/sys_role_dept/service/view"
	"manager-gin/src/common"
	response "manager-gin/src/common/response"
	"manager-gin/src/global"
	"manager-gin/src/utils"
	"strconv"
)

type SysRoleDeptApi struct {
}

var sysRoleDeptService = service.SysRoleDeptServiceApp

// Create 创建SysRoleDept
// @Summary 创建SysRoleDept
// @Router /sysRoleDept/create [post]
func (sysRoleDeptApi *SysRoleDeptApi) Create(c *gin.Context) {
	var sysRoleDeptView view.SysRoleDeptView
	_ = c.ShouldBindJSON(&sysRoleDeptView)
	sysRoleDeptView.Id = strings.ReplaceAll(uuid.NewV4().String(), "-", "")
	sysRoleDeptView.CreateTime = utils.GetCurTimeStr()
	sysRoleDeptView.UpdateTime = utils.GetCurTimeStr()
	if err := sysRoleDeptService.Create(&sysRoleDeptView); err != nil {
		global.Logger.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// Delete 删除SysRoleDept
// @Summary 删除SysRoleDept
// @Router /sysRoleDept/delete [delete]
func (sysRoleDeptApi *SysRoleDeptApi) Delete(c *gin.Context) {
	var id common.Id
	_ = c.ShouldBindJSON(&id)
	if err := sysRoleDeptService.Delete(id.ID); err != nil {
		global.Logger.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteByIds 批量删除SysRoleDept
// @Summary 批量删除SysRoleDept
// @Router /sysRoleDept/deleteByIds [delete]
func (sysRoleDeptApi *SysRoleDeptApi) DeleteByIds(c *gin.Context) {
	var ids common.Ids
	_ = c.ShouldBindJSON(&ids)
	if err := sysRoleDeptService.DeleteByIds(ids.Ids); err != nil {
		global.Logger.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// Update 更新SysRoleDept
// @Summary 更新SysRoleDept
// @Router /sysRoleDept/update [put]
func (sysRoleDeptApi *SysRoleDeptApi) Update(c *gin.Context) {
	var sysRoleDeptView view.SysRoleDeptView
	_ = c.ShouldBindJSON(&sysRoleDeptView)
	id := sysRoleDeptView.Id
	if id == "" {
		response.FailWithMessage("更新失败", c)
	}
	sysRoleDeptView.UpdateTime = utils.GetCurTimeStr()
	if err := sysRoleDeptService.Update(id, &sysRoleDeptView); err != nil {
		global.Logger.Error("更新持久化失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// Get 用id查询SysRoleDept
// @Summary 用id查询SysRoleDept
// @Router /sysRoleDept/get [get]
func (sysRoleDeptApi *SysRoleDeptApi) Get(c *gin.Context) {
	id := c.Query("id")
	if err, sysRoleDeptView := sysRoleDeptService.Get(id); err != nil {
		global.Logger.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"sysRoleDeptView": sysRoleDeptView}, c)
	}
}

// Find 分页获取SysRoleDept列表
// @Summary 分页获取SysRoleDept列表
// @Router /sysRoleDept/find [get]
func (sysRoleDeptApi *SysRoleDeptApi) Find(c *gin.Context) {
	var pageInfo common.PageInfoV2
	// 绑定查询参数到 pageInfo
	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		response.FailWithMessage("获取分页数据解析失败!", c)
	}
	// 调用 Calculate 方法自动计算 Limit 和 Offset
	pageInfo.Calculate()
	if err := sysRoleDeptService.Find(&pageInfo); err != nil {
		global.Logger.Error("获取分页信息失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(pageInfo, "获取成功", c)
	}
}
