// Package api  SysMenuApi 自动生成模板
// @description <TODO description class purpose>
// @author
// @File: sys_menu
// @version 1.0.0
// @create 2023-08-08 10:06:19
package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"manager-gin/src/app/admin/sys/sys_menu/service"
	"manager-gin/src/app/admin/sys/sys_menu/service/view"
	"manager-gin/src/common"
	response "manager-gin/src/common/response"
	"manager-gin/src/global"
	"manager-gin/src/utils"
	"strconv"
)

type SysMenuApi struct {
}

var sysMenuService = service.SysMenuServiceApp

// Create 创建SysMenu
// @Summary 创建SysMenu
// @Router /sysMenu/create [post]
func (sysMenuApi *SysMenuApi) Create(c *gin.Context) {
	var sysMenuView view.SysMenuView
	_ = c.ShouldBindJSON(&sysMenuView)
	sysMenuView.CreateTime = utils.GetCurTimeStr()
	sysMenuView.UpdateTime = utils.GetCurTimeStr()
	if err := sysMenuService.Create(&sysMenuView); err != nil {
		global.Logger.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// Delete 删除SysMenu
// @Summary 删除SysMenu
// @Router /sysMenu/delete [delete]
func (sysMenuApi *SysMenuApi) Delete(c *gin.Context) {
	var id int
	_ = c.ShouldBindJSON(&id)
	if err := sysMenuService.Delete(id); err != nil {
		global.Logger.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteByIds 批量删除SysMenu
// @Summary 批量删除SysMenu
// @Router /sysMenu/deleteByIds [delete]
func (sysMenuApi *SysMenuApi) DeleteByIds(c *gin.Context) {
	var ids []int
	_ = c.ShouldBindJSON(&ids)
	if err := sysMenuService.DeleteByIds(ids); err != nil {
		global.Logger.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// Update 更新SysMenu
// @Summary 更新SysMenu
// @Router /sysMenu/update [put]
func (sysMenuApi *SysMenuApi) Update(c *gin.Context) {
	id := c.Query("id")
	atoi, err := strconv.Atoi(id)
	if err != nil {
		global.Logger.Error("更新解析上报数据失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	}
	sysMenuViewJson := c.Query("sysMenuView")
	var sysMenuView view.SysMenuView
	err = json.Unmarshal([]byte(sysMenuViewJson), &sysMenuView)
	sysMenuView.UpdateTime = utils.GetCurTimeStr()
	if err := sysMenuService.Update(atoi, &sysMenuView); err != nil {
		global.Logger.Error("更新解析上报数据失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	}
	if err = sysMenuService.Update(atoi, &sysMenuView); err != nil {
		global.Logger.Error("更新持久化失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// Get 用id查询SysMenu
// @Summary 用id查询SysMenu
// @Router /sysMenu/get [get]
func (sysMenuApi *SysMenuApi) Get(c *gin.Context) {
	id := c.Query("id")
	atoi, err := strconv.Atoi(id)
	if err != nil {
		global.Logger.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	}
	if err, sysMenuView := sysMenuService.Get(atoi); err != nil {
		global.Logger.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"sysMenuView": sysMenuView}, c)
	}
}

// Find 分页获取SysMenu列表
// @Summary 分页获取SysMenu列表
// @Router /sysMenu/find [get]
func (sysMenuApi *SysMenuApi) Find(c *gin.Context) {
	var pageInfo common.PageInfoV2
	_ = c.ShouldBindQuery(&pageInfo)
	if err := sysMenuService.Find(&pageInfo); err != nil {
		global.Logger.Error("获取分页信息失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(pageInfo, "获取成功", c)
	}
}
