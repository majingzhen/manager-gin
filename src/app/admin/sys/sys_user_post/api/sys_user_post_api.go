// Package api  SysUserPostApi 自动生成模板
// @description <TODO description class purpose>
// @author
// @File: sys_user_post
// @version 1.0.0
// @create 2023-08-08 10:06:19
package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"manager-gin/src/app/admin/sys/sys_user_post/service"
	"manager-gin/src/app/admin/sys/sys_user_post/service/view"
	"manager-gin/src/common"
	"manager-gin/src/global"
	"manager-gin/src/utils"
)

type SysUserPostApi struct {
}

var sysUserPostService = service.SysUserPostServiceApp

// Create 创建SysUserPost
// @Summary 创建SysUserPost
// @Router /sysUserPost/create [post]
func (sysUserPostApi *SysUserPostApi) Create(c *gin.Context) {
	var sysUserPostView view.SysUserPostView
	_ = c.ShouldBindJSON(&sysUserPostView)
	sysUserPostView.CreateTime = utils.GetCurTimeStr()
	sysUserPostView.UpdateTime = utils.GetCurTimeStr()
	if err := sysUserPostService.Create(&sysUserPostView); err != nil {
		global.Logger.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// Delete 删除SysUserPost
// @Summary 删除SysUserPost
// @Router /sysUserPost/delete [delete]
func (sysUserPostApi *SysUserPostApi) Delete(c *gin.Context) {
	var id int
	_ = c.ShouldBindJSON(&id)
	if err := sysUserPostService.Delete(id); err != nil {
		global.Logger.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteByIds 批量删除SysUserPost
// @Summary 批量删除SysUserPost
// @Router /sysUserPost/deleteByIds [delete]
func (sysUserPostApi *SysUserPostApi) DeleteByIds(c *gin.Context) {
	var ids []int
	_ = c.ShouldBindJSON(&ids)
	if err := sysUserPostService.DeleteByIds(ids); err != nil {
		global.Logger.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// Update 更新SysUserPost
// @Summary 更新SysUserPost
// @Router /sysUserPost/update [put]
func (sysUserPostApi *SysUserPostApi) Update(c *gin.Context) {
	id := c.Query("id")
	atoi, err := strconv.Atoi(id)
	if err != nil {
		global.Logger.Error("更新解析上报数据失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	}
	sysUserPostViewJson := c.Query("sysUserPostView")
	var sysUserPostView view.SysUserPostView
	err = json.Unmarshal([]byte(sysUserPostViewJson), &sysUserPostView)
	sysUserPostView.UpdateTime = utils.GetCurTimeStr()
	if err := sysUserPostService.Update(atoi, &sysUserPostView); err != nil {
		global.Logger.Error("更新解析上报数据失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	}
	if err = sysUserPostService.Update(atoi, &sysUserPostView); err != nil {
		global.Logger.Error("更新持久化失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// Get 用id查询SysUserPost
// @Summary 用id查询SysUserPost
// @Router /sysUserPost/get [get]
func (sysUserPostApi *SysUserPostApi) Get(c *gin.Context) {
	id := c.Query("id")
	atoi, err := strconv.Atoi(id)
	if err != nil {
		global.Logger.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	}
	if err, sysUserPostView := sysUserPostService.Get(atoi); err != nil {
		global.Logger.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"sysUserPostView": sysUserPostView}, c)
	}
}

// Find 分页获取SysUserPost列表
// @Summary 分页获取SysUserPost列表
// @Router /sysUserPost/find [get]
func (sysUserPostApi *SysUserPostApi) Find(c *gin.Context) {
	var pageInfo common.PageInfoV2
	_ = c.ShouldBindQuery(&pageInfo)
	if err := sysUserPostService.Find(&pageInfo); err != nil {
		global.Logger.Error("获取分页信息失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(pageInfo, "获取成功", c)
	}
}
