// Package api  SysMenuApi 自动生成模板
// @description <TODO description class purpose>
// @author
// @File: sys_menu
// @version 1.0.0
// @create 2023-08-20 21:21:34
package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"manager-gin/src/app/admin/sys/service/sys_menu"
	"manager-gin/src/app/admin/sys/service/sys_menu/view"
	response "manager-gin/src/common/response"
	"manager-gin/src/global"
	"manager-gin/src/utils"
)

type SysMenuApi struct {
	BasicApi
	sysMenuService sys_menu.SysMenuService
}

// Create 创建SysMenu
// @Summary 创建SysMenu
// @Router /sysMenu/create [post]
func (api *SysMenuApi) Create(c *gin.Context) {
	var sysMenuView view.SysMenuView
	_ = c.ShouldBindJSON(&sysMenuView)
	sysMenuView.Id = utils.GenUID()
	sysMenuView.CreateTime = utils.GetCurTimeStr()
	sysMenuView.UpdateTime = utils.GetCurTimeStr()
	sysMenuView.CreateBy = api.GetLoginUserName(c)
	if err := api.sysMenuService.Create(&sysMenuView); err != nil {
		global.Logger.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// Delete 删除SysMenu
// @Summary 删除SysMenu
// @Router /sysMenu/delete [delete]
func (api *SysMenuApi) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := api.sysMenuService.Delete(id); err != nil {
		global.Logger.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// Update 更新SysMenu
// @Summary 更新SysMenu
// @Router /sysMenu/update [put]
func (api *SysMenuApi) Update(c *gin.Context) {
	var sysMenuView view.SysMenuView
	_ = c.ShouldBindJSON(&sysMenuView)
	id := sysMenuView.Id
	if id == "" {
		response.FailWithMessage("更新失败", c)
		return
	}
	sysMenuView.UpdateTime = utils.GetCurTimeStr()
	sysMenuView.UpdateBy = api.GetLoginUserName(c)
	if err := api.sysMenuService.Update(id, &sysMenuView); err != nil {
		global.Logger.Error("更新持久化失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// Get 用id查询SysMenu
// @Summary 用id查询SysMenu
// @Router /sysMenu/get [get]
func (api *SysMenuApi) Get(c *gin.Context) {
	id := c.Param("id")
	if err, sysMenuView := api.sysMenuService.Get(id); err != nil {
		global.Logger.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(sysMenuView, c)
	}
}

// List 获取SysMenu列表
// @Summary 获取SysMenu列表
// @Router /sysMenu/list [get]
func (api *SysMenuApi) List(c *gin.Context) {
	var menuView view.SysMenuView
	// 绑定查询参数到 pageInfo
	if err := c.ShouldBindQuery(&menuView); err != nil {
		response.FailWithMessage("获取参数解析失败!", c)
		return
	}
	userId := api.GetLoginUserId(c)
	if err, res := api.sysMenuService.SelectMenuList(&menuView, userId); err != nil {
		global.Logger.Error("获取数据失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(res, "获取成功", c)
	}
}

// SelectMenuTreeByRoleId 加载对应角色菜单列表树
// @Summary 根据角色id查询菜单树信息
// @Router /sysMenu/selectMenuTreeByRoleId/{roleId} [get]
func (api *SysMenuApi) SelectMenuTreeByRoleId(c *gin.Context) {
	roleId := c.Param("roleId")
	if err, menuList := api.sysMenuService.SelectMenuList(&view.SysMenuView{}, api.GetLoginUserId(c)); err != nil {
		global.Logger.Error("获取数据失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	} else {
		menuTree := api.sysMenuService.BuildMenuTreeSelect(menuList)
		if err, menuListByRoleId := api.sysMenuService.SelectMenuListByRoleId(roleId); err != nil {
			global.Logger.Error("获取数据失败!", zap.Error(err))
			response.FailWithMessage("获取失败", c)
		} else {
			response.OkWithData(gin.H{
				"checkedKeys": menuListByRoleId,
				"menus":       menuTree,
			}, c)
		}
	}
}

// SelectMenuTree 加载菜单列表树
// @Summary 加载菜单列表树
// @Router /sysMenu/selectMenuTree [get]
func (api *SysMenuApi) SelectMenuTree(c *gin.Context) {
	var menuView view.SysMenuView
	if err := c.ShouldBindQuery(&menuView); err != nil {
		response.FailWithMessage("获取参数解析失败!", c)
		return
	}
	if err, menuList := api.sysMenuService.SelectMenuList(&menuView, api.GetLoginUserId(c)); err != nil {
		global.Logger.Error("获取数据失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		menuTree := api.sysMenuService.BuildMenuTreeSelect(menuList)
		response.OkWithData(menuTree, c)
	}
}