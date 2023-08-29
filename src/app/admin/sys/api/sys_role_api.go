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
	"manager-gin/src/app/admin/sys/service/sys_role"
	"manager-gin/src/app/admin/sys/service/sys_role/view"
	"manager-gin/src/common"
	response "manager-gin/src/common/response"
	"manager-gin/src/global"
	"manager-gin/src/utils"
	"strings"
)

type SysRoleApi struct {
	BasicApi
	sysRoleService sys_role.SysRoleService
}

// Create 创建SysRole
// @Summary 创建SysRole
// @Router /sysRole/create [post]
func (api *SysRoleApi) Create(c *gin.Context) {
	var sysRoleView view.SysRoleView
	if err := c.ShouldBindJSON(&sysRoleView); err != nil {
		global.Logger.Error("参数解析失败!", zap.Error(err))
		response.FailWithMessage("参数解析失败", c)
		return
	}
	if err := api.sysRoleService.CheckRoleNameUnique(sysRoleView.RoleName, ""); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := api.sysRoleService.CheckRoleKeyUnique(sysRoleView.RoleKey, ""); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	sysRoleView.Id = utils.GenUID()
	sysRoleView.CreateTime = utils.GetCurTimeStr()
	sysRoleView.UpdateTime = utils.GetCurTimeStr()
	sysRoleView.CreateBy = api.GetLoginUserName(c)
	if err := api.sysRoleService.Create(&sysRoleView); err != nil {
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
	if err := api.sysRoleService.DeleteByIds(ids, api.GetLoginUser(c)); err != nil {
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
	if err := api.sysRoleService.CheckRoleDataScope(id, api.GetLoginUser(c)); err != nil {
		response.FailWithMessage("没有权限访问角色数据", c)
		return
	}
	if err := api.sysRoleService.CheckRoleNameUnique(sysRoleView.RoleName, sysRoleView.Id); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := api.sysRoleService.CheckRoleKeyUnique(sysRoleView.RoleKey, sysRoleView.Id); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	sysRoleView.UpdateTime = utils.GetCurTimeStr()
	sysRoleView.UpdateBy = api.GetLoginUserName(c)
	if err := api.sysRoleService.Update(id, &sysRoleView); err != nil {
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
	if err, sysRoleView := api.sysRoleService.Get(id); err != nil {
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
	user := api.GetLoginUser(c)
	if err, res := api.sysRoleService.Page(&pageInfo, user); err != nil {
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
	// userId := api.GetLoginUserId(c)
	if err, res := api.sysRoleService.List(&view, api.GetLoginUser(c)); err != nil {
		global.Logger.Error("获取数据失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(res, "获取成功", c)
	}
}

// ChangeStatus 更新SysRole状态
// @Summary 更新SysRole状态
// @Router /sysRole/changeStatus [put]
func (api *SysRoleApi) ChangeStatus(c *gin.Context) {
	var view view.SysRoleView
	_ = c.ShouldBindJSON(&view)
	if view.Id == common.SYSTEM_ROLE_ADMIN_ID {
		response.FailWithMessage("超级管理员不允许修改", c)
		return
	}
	if err := api.sysRoleService.CheckRoleDataScope(view.Id, api.GetLoginUser(c)); err != nil {
		response.FailWithMessage("没有权限访问角色数据", c)
		return
	}
	view.UpdateTime = utils.GetCurTimeStr()
	view.UpdateBy = api.GetLoginUserName(c)
	if err := api.sysRoleService.UpdateStatus(&view); err != nil {
		global.Logger.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// DataScope 修改保存数据权限
// @Summary 修改保存数据权限
// @Router /sysRole/dataScope [put]
func (api *SysRoleApi) DataScope(c *gin.Context) {
	var view view.SysRoleView
	_ = c.ShouldBindJSON(&view)
	if view.Id == common.SYSTEM_ROLE_ADMIN_ID {
		response.FailWithMessage("超级管理员不允许修改", c)
		return
	}
	if err := api.sysRoleService.CheckRoleDataScope(view.Id, api.GetLoginUser(c)); err != nil {
		response.FailWithMessage("没有权限访问角色数据", c)
		return
	}
	view.UpdateTime = utils.GetCurTimeStr()
	view.UpdateBy = api.GetLoginUserName(c)
	if err := api.sysRoleService.AuthDataScope(&view); err != nil {
		global.Logger.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// CancelAuthUser 取消授权用户
// @Summary 取消授权用户
// @Router /sysRole/cancelAuthUser [put]
func (api *SysRoleApi) CancelAuthUser(c *gin.Context) {
	var view view.SysUserRoleView
	_ = c.ShouldBindJSON(&view)
	if err := api.sysRoleService.CancelAuthUser(&view); err != nil {
		global.Logger.Error("取消授权失败!", zap.Error(err))
		response.FailWithMessage("取消授权失败", c)
	} else {
		response.OkWithMessage("取消授权成功", c)
	}
}

// BatchCancelAuthUser 批量取消授权用户
// @Summary 批量取消授权用户
// @Router /sysRole/batchCancelAuthUser [put]
func (api *SysRoleApi) BatchCancelAuthUser(c *gin.Context) {
	userId := c.Query("userIds")
	roleId := c.Query("roleId")
	if userId == "" || roleId == "" {
		response.FailWithMessage("参数解析错误", c)
	}
	userIds := strings.Split(userId, ",")
	if err := api.sysRoleService.BatchCancelAuthUser(roleId, userIds); err != nil {
		global.Logger.Error("批量取消授权失败!", zap.Error(err))
		response.FailWithMessage("批量取消授权失败", c)
	} else {
		response.OkWithMessage("批量取消授权成功", c)
	}
}

// BatchSelectAuthUser 批量授权用户
// @Summary 授权用户
// @Router /sysRole/batchSelectAuthUser [put]
func (api *SysRoleApi) BatchSelectAuthUser(c *gin.Context) {
	userId := c.Query("userIds")
	roleId := c.Query("roleId")
	if userId == "" || roleId == "" {
		response.FailWithMessage("参数解析错误", c)
	}
	userIds := strings.Split(userId, ",")
	if err := api.sysRoleService.BatchSelectAuthUser(roleId, userIds); err != nil {
		global.Logger.Error("批量授权失败!", zap.Error(err))
		response.FailWithMessage("批量授权失败", c)
	} else {
		response.OkWithMessage("批量授权成功", c)
	}
}
