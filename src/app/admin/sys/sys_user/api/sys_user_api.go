// Package api  SysUserApi 自动生成模板
// @description <TODO description class purpose>
// @author
// @File: sys_user
// @version 1.0.0
// @create 2023-08-18 14:02:24
package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"manager-gin/src/app/admin/sys/sys_user/service"
	"manager-gin/src/app/admin/sys/sys_user/service/view"
	"manager-gin/src/common"
	response "manager-gin/src/common/response"
	"manager-gin/src/global"
	"manager-gin/src/utils"
)

type SysUserApi struct {
}

var sysUserService = service.SysUserServiceApp

// Create 创建SysUser
// @Summary 创建SysUser
// @Router /sysUser/create [post]
func (sysUserApi *SysUserApi) Create(c *gin.Context) {
	var sysUserView view.SysUserView
	_ = c.ShouldBindJSON(&sysUserView)
	sysUserView.Id = utils.GenUID()
	sysUserView.CreateTime = utils.GetCurTimeStr()
	sysUserView.UpdateTime = utils.GetCurTimeStr()
	if err := sysUserService.Create(&sysUserView); err != nil {
		global.Logger.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// Delete 删除SysUser
// @Summary 删除SysUser
// @Router /sysUser/delete [delete]
func (sysUserApi *SysUserApi) Delete(c *gin.Context) {
	var id common.Id
	_ = c.ShouldBindJSON(&id)
	if err := sysUserService.Delete(id.ID); err != nil {
		global.Logger.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteByIds 批量删除SysUser
// @Summary 批量删除SysUser
// @Router /sysUser/deleteByIds [delete]
func (sysUserApi *SysUserApi) DeleteByIds(c *gin.Context) {
	var ids common.Ids
	_ = c.ShouldBindJSON(&ids)
	if err := sysUserService.DeleteByIds(ids.Ids); err != nil {
		global.Logger.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// Update 更新SysUser
// @Summary 更新SysUser
// @Router /sysUser/update [put]
func (sysUserApi *SysUserApi) Update(c *gin.Context) {
	var sysUserView view.SysUserView
	_ = c.ShouldBindJSON(&sysUserView)
	id := sysUserView.Id
	if id == "" {
		response.FailWithMessage("更新失败", c)
	}
	sysUserView.UpdateTime = utils.GetCurTimeStr()
	if err := sysUserService.Update(id, &sysUserView); err != nil {
		global.Logger.Error("更新持久化失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// Get 用id查询SysUser
// @Summary 用id查询SysUser
// @Router /sysUser/get [get]
func (sysUserApi *SysUserApi) Get(c *gin.Context) {
	id := c.Query("id")
	if err, sysUserView := sysUserService.Get(id); err != nil {
		global.Logger.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"sysUserView": sysUserView}, c)
	}
}

// List 分页获取SysUser列表
// @Summary 分页获取SysUser列表
// @Router /sysUser/list [get]
func (sysUserApi *SysUserApi) List(c *gin.Context) {
	var pageInfo common.PageInfo
	// 绑定查询参数到 pageInfo
	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		response.FailWithMessage("获取分页数据解析失败!", c)
	}
	// 调用 Calculate 方法自动计算 Limit 和 Offset
	pageInfo.Calculate()
	if err := sysUserService.List(&pageInfo); err != nil {
		global.Logger.Error("获取分页信息失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(pageInfo, "获取成功", c)
	}
}
