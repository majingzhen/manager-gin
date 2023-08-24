// Package api  SysRoleApi 自动生成模板
// @description <TODO description class purpose>
// @author
// @File: sys_role
// @version 1.0.0
// @create 2023-08-21 17:37:56
package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"manager-gin/src/app/admin/sys/sys_role/service"
	"manager-gin/src/app/admin/sys/sys_role/service/view"
	"manager-gin/src/common"
	response "manager-gin/src/common/response"
	"manager-gin/src/framework"
	"manager-gin/src/global"
	"manager-gin/src/utils"
	"strings"
)

type SysRoleApi struct {
}

var sysRoleService = service.SysRoleServiceApp

// Create 创建SysRole
// @Summary 创建SysRole
// @Router /sysRole/create [post]
func (api *SysRoleApi) Create(c *gin.Context) {
	var sysRoleView view.SysRoleView
	_ = c.ShouldBindJSON(&sysRoleView)
	if err := sysRoleService.CheckRoleNameUnique(sysRoleView.RoleName); err == nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := sysRoleService.CheckRoleKeyUnique(sysRoleView.RoleKey); err == nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	sysRoleView.Id = utils.GenUID()
	sysRoleView.CreateTime = utils.GetCurTimeStr()
	sysRoleView.UpdateTime = utils.GetCurTimeStr()
	sysRoleView.CreateBy = framework.GetLoginUserName(c)
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
func (api *SysRoleApi) Delete(c *gin.Context) {
	idStr := c.Param("ids")
	ids := strings.Split(idStr, ",")
	if err := sysRoleService.DeleteByIds(ids, framework.GetLoginUser(c)); err != nil {
		global.Logger.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// Update 更新SysRole
// @Summary 更新SysRole
// @Router /sysRole/update [put]
func (api *SysRoleApi) Update(c *gin.Context) {
	var sysRoleView view.SysRoleView
	_ = c.ShouldBindJSON(&sysRoleView)
	id := sysRoleView.Id
	if id == "" {
		response.FailWithMessage("更新失败", c)
		return
	}
	// 校验参数
	if id == common.SYSTEM_ROLE_ADMIN_ID {
		response.FailWithMessage("超级管理员不允许修改", c)
		return
	}
	// 校验数据权限
	if err := sysRoleService.CheckRoleDataScope(id, framework.GetLoginUser(c)); err != nil {
		response.FailWithMessage("没有权限访问角色数据", c)
		return
	}
	if err := sysRoleService.CheckRoleNameUnique(sysRoleView.RoleName); err == nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := sysRoleService.CheckRoleKeyUnique(sysRoleView.RoleKey); err == nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	sysRoleView.UpdateTime = utils.GetCurTimeStr()
	sysRoleView.UpdateBy = framework.GetLoginUserName(c)
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
func (api *SysRoleApi) Get(c *gin.Context) {
	id := c.Param("id")
	if err, sysRoleView := sysRoleService.Get(id); err != nil {
		global.Logger.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(sysRoleView, c)
	}
}

// Page 分页获取SysRole列表
// @Summary 分页获取SysRole列表
// @Router /sysRole/page [get]
func (api *SysRoleApi) Page(c *gin.Context) {
	var pageInfo view.SysRolePageView
	// 绑定查询参数到 pageInfo
	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		response.FailWithMessage("获取分页数据解析失败!", c)
		return
	}
	user := framework.GetLoginUser(c)
	if err, res := sysRoleService.Page(&pageInfo, user); err != nil {
		global.Logger.Error("获取分页信息失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(res, "获取成功", c)
	}
}

// List 获取SysRole列表
// @Summary 获取SysRole列表
// @Router /sysRole/list [get]
func (api *SysRoleApi) List(c *gin.Context) {
	var view view.SysRoleView
	// 绑定查询参数到 view对象
	if err := c.ShouldBindQuery(&view); err != nil {
		response.FailWithMessage("获取参数解析失败!", c)
		return
	}
	// 判断是否需要根据用户获取数据
	// userId := framework.GetLoginUserId(c)
	if err, res := sysRoleService.List(&view); err != nil {
		global.Logger.Error("获取数据失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(res, "获取成功", c)
	}
}
