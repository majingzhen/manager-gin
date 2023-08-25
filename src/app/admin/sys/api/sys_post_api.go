// Package api  SysPostApi 自动生成模板
// @description <TODO description class purpose>
// @author
// @File: sys_post
// @version 1.0.0
// @create 2023-08-21 17:37:56
package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"manager-gin/src/app/admin/sys/service/sys_post"
	"manager-gin/src/app/admin/sys/service/sys_post/view"
	response "manager-gin/src/common/response"
	"manager-gin/src/framework"
	"manager-gin/src/global"
	"manager-gin/src/utils"
	"strings"
)

type SysPostApi struct {
	sysPostService sys_post.SysPostService
}

// Create 创建SysPost
// @Summary 创建SysPost
// @Router /sysPost/create [post]
func (api *SysPostApi) Create(c *gin.Context) {
	var sysPostView view.SysPostView
	_ = c.ShouldBindJSON(&sysPostView)
	sysPostView.Id = utils.GenUID()
	sysPostView.CreateTime = utils.GetCurTimeStr()
	sysPostView.UpdateTime = utils.GetCurTimeStr()
	sysPostView.CreateBy = framework.GetLoginUserName(c)
	if err := api.sysPostService.Create(&sysPostView); err != nil {
		global.Logger.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// Delete 删除SysPost
// @Summary 删除SysPost
// @Router /sysPost/delete [delete]
func (api *SysPostApi) Delete(c *gin.Context) {
	idStr := c.Param("ids")
	ids := strings.Split(idStr, ",")
	if err := api.sysPostService.DeleteByIds(ids); err != nil {
		global.Logger.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// Update 更新SysPost
// @Summary 更新SysPost
// @Router /sysPost/update [put]
func (api *SysPostApi) Update(c *gin.Context) {
	var sysPostView view.SysPostView
	_ = c.ShouldBindJSON(&sysPostView)
	id := sysPostView.Id
	if id == "" {
		response.FailWithMessage("更新失败", c)
		return
	}
	sysPostView.UpdateTime = utils.GetCurTimeStr()
	sysPostView.UpdateBy = framework.GetLoginUserName(c)
	if err := api.sysPostService.Update(id, &sysPostView); err != nil {
		global.Logger.Error("更新持久化失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// Get 用id查询SysPost
// @Summary 用id查询SysPost
// @Router /sysPost/get [get]
func (api *SysPostApi) Get(c *gin.Context) {
	id := c.Param("id")
	if err, sysPostView := api.sysPostService.Get(id); err != nil {
		global.Logger.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(sysPostView, c)
	}
}

// Page 分页获取SysPost列表
// @Summary 分页获取SysPost列表
// @Router /sysPost/page [get]
func (api *SysPostApi) Page(c *gin.Context) {
	var pageInfo view.SysPostPageView
	// 绑定查询参数到 pageInfo
	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		response.FailWithMessage("获取分页数据解析失败!", c)
		return
	}

	if err, res := api.sysPostService.Page(&pageInfo); err != nil {
		global.Logger.Error("获取分页信息失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(res, "获取成功", c)
	}
}

// List 获取SysPost列表
// @Summary 获取SysPost列表
// @Router /sysPost/list [get]
func (api *SysPostApi) List(c *gin.Context) {
	var view view.SysPostView
	// 绑定查询参数到 view对象
	if err := c.ShouldBindQuery(&view); err != nil {
		response.FailWithMessage("获取参数解析失败!", c)
		return
	}
	// 判断是否需要根据用户获取数据
	// userId := framework.GetLoginUserId(c)
	if err, res := api.sysPostService.List(&view); err != nil {
		global.Logger.Error("获取数据失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(res, "获取成功", c)
	}
}
