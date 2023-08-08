// Package api  SysPostApi 自动生成模板
// @description <TODO description class purpose>
// @author
// @File: sys_post
// @version 1.0.0
// @create 2023-08-08 10:06:19
package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"manager-gin/src/app/admin/sys/sys_post/service"
	"manager-gin/src/app/admin/sys/sys_post/service/view"
	"manager-gin/src/common"
	response "manager-gin/src/common/response"
	"manager-gin/src/global"
	"manager-gin/src/utils"
	"strconv"
)

type SysPostApi struct {
}

var sysPostService = service.SysPostServiceApp

// Create 创建SysPost
// @Summary 创建SysPost
// @Router /sysPost/create [post]
func (sysPostApi *SysPostApi) Create(c *gin.Context) {
	var sysPostView view.SysPostView
	_ = c.ShouldBindJSON(&sysPostView)
	sysPostView.CreateTime = utils.GetCurTimeStr()
	sysPostView.UpdateTime = utils.GetCurTimeStr()
	if err := sysPostService.Create(&sysPostView); err != nil {
		global.Logger.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// Delete 删除SysPost
// @Summary 删除SysPost
// @Router /sysPost/delete [delete]
func (sysPostApi *SysPostApi) Delete(c *gin.Context) {
	var id int
	_ = c.ShouldBindJSON(&id)
	if err := sysPostService.Delete(id); err != nil {
		global.Logger.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteByIds 批量删除SysPost
// @Summary 批量删除SysPost
// @Router /sysPost/deleteByIds [delete]
func (sysPostApi *SysPostApi) DeleteByIds(c *gin.Context) {
	var ids []int
	_ = c.ShouldBindJSON(&ids)
	if err := sysPostService.DeleteByIds(ids); err != nil {
		global.Logger.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// Update 更新SysPost
// @Summary 更新SysPost
// @Router /sysPost/update [put]
func (sysPostApi *SysPostApi) Update(c *gin.Context) {
	id := c.Query("id")
	atoi, err := strconv.Atoi(id)
	if err != nil {
		global.Logger.Error("更新解析上报数据失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	}
	sysPostViewJson := c.Query("sysPostView")
	var sysPostView view.SysPostView
	err = json.Unmarshal([]byte(sysPostViewJson), &sysPostView)
	sysPostView.UpdateTime = utils.GetCurTimeStr()
	if err := sysPostService.Update(atoi, &sysPostView); err != nil {
		global.Logger.Error("更新解析上报数据失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	}
	if err = sysPostService.Update(atoi, &sysPostView); err != nil {
		global.Logger.Error("更新持久化失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// Get 用id查询SysPost
// @Summary 用id查询SysPost
// @Router /sysPost/get [get]
func (sysPostApi *SysPostApi) Get(c *gin.Context) {
	id := c.Query("id")
	atoi, err := strconv.Atoi(id)
	if err != nil {
		global.Logger.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	}
	if err, sysPostView := sysPostService.Get(atoi); err != nil {
		global.Logger.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"sysPostView": sysPostView}, c)
	}
}

// Find 分页获取SysPost列表
// @Summary 分页获取SysPost列表
// @Router /sysPost/find [get]
func (sysPostApi *SysPostApi) Find(c *gin.Context) {
	var pageInfo common.PageInfoV2
	_ = c.ShouldBindQuery(&pageInfo)
	if err := sysPostService.Find(&pageInfo); err != nil {
		global.Logger.Error("获取分页信息失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(pageInfo, "获取成功", c)
	}
}
