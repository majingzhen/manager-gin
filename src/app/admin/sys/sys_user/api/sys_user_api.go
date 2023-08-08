// Package api  SysUserApi 自动生成模板
// @description <TODO description class purpose>
// @author
// @File: sys_user
// @version 1.0.0
// @create 2023-08-08 10:06:19
package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"manager-gin/src/app/admin/sys/sys_user/service"
	"manager-gin/src/app/admin/sys/sys_user/service/view"
	"manager-gin/src/common"
	response "manager-gin/src/common/response"
	"manager-gin/src/global"
	"manager-gin/src/utils"
	"strconv"
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
	var id int
	_ = c.ShouldBindJSON(&id)
	if err := sysUserService.Delete(id); err != nil {
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
	var ids []int
	_ = c.ShouldBindJSON(&ids)
	if err := sysUserService.DeleteByIds(ids); err != nil {
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
	id := c.Query("id")
	atoi, err := strconv.Atoi(id)
	if err != nil {
		global.Logger.Error("更新解析上报数据失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	}
	sysUserViewJson := c.Query("sysUserView")
	var sysUserView view.SysUserView
	err = json.Unmarshal([]byte(sysUserViewJson), &sysUserView)
	sysUserView.UpdateTime = utils.GetCurTimeStr()
	if err := sysUserService.Update(atoi, &sysUserView); err != nil {
		global.Logger.Error("更新解析上报数据失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	}
	if err = sysUserService.Update(atoi, &sysUserView); err != nil {
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
	atoi, err := strconv.Atoi(id)
	if err != nil {
		global.Logger.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	}
	if err, sysUserView := sysUserService.Get(atoi); err != nil {
		global.Logger.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"sysUserView": sysUserView}, c)
	}
}

// Find 分页获取SysUser列表
// @Summary 分页获取SysUser列表
// @Router /sysUser/find [get]
func (sysUserApi *SysUserApi) Find(c *gin.Context) {
	var pageInfo common.PageInfoV2
	_ = c.ShouldBindQuery(&pageInfo)
	if err := sysUserService.Find(&pageInfo); err != nil {
		global.Logger.Error("获取分页信息失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(pageInfo, "获取成功", c)
	}
}
