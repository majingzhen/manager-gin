// Package api  SysLogininforApi 自动生成模板
// @description <TODO description class purpose>
// @author
// @File: sys_logininfor
// @version 1.0.0
// @create 2023-08-08 10:06:19
package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"manager-gin/src/app/admin/sys/sys_logininfor/service"
	"manager-gin/src/app/admin/sys/sys_logininfor/service/view"
	"manager-gin/src/common"
	response "manager-gin/src/common/response"
	"manager-gin/src/global"
	"strconv"
)

type SysLogininforApi struct {
}

var sysLogininforService = service.SysLogininforServiceApp

// Create 创建SysLogininfor
// @Summary 创建SysLogininfor
// @Router /sysLogininfor/create [post]
func (sysLogininforApi *SysLogininforApi) Create(c *gin.Context) {
	var sysLogininforView view.SysLogininforView
	_ = c.ShouldBindJSON(&sysLogininforView)
	if err := sysLogininforService.Create(&sysLogininforView); err != nil {
		global.Logger.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// Delete 删除SysLogininfor
// @Summary 删除SysLogininfor
// @Router /sysLogininfor/delete [delete]
func (sysLogininforApi *SysLogininforApi) Delete(c *gin.Context) {
	var id int
	_ = c.ShouldBindJSON(&id)
	if err := sysLogininforService.Delete(id); err != nil {
		global.Logger.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteByIds 批量删除SysLogininfor
// @Summary 批量删除SysLogininfor
// @Router /sysLogininfor/deleteByIds [delete]
func (sysLogininforApi *SysLogininforApi) DeleteByIds(c *gin.Context) {
	var ids []int
	_ = c.ShouldBindJSON(&ids)
	if err := sysLogininforService.DeleteByIds(ids); err != nil {
		global.Logger.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// Update 更新SysLogininfor
// @Summary 更新SysLogininfor
// @Router /sysLogininfor/update [put]
func (sysLogininforApi *SysLogininforApi) Update(c *gin.Context) {
	id := c.Query("id")
	atoi, err := strconv.Atoi(id)
	if err != nil {
		global.Logger.Error("更新解析上报数据失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	}
	sysLogininforViewJson := c.Query("sysLogininforView")
	var sysLogininforView view.SysLogininforView
	err = json.Unmarshal([]byte(sysLogininforViewJson), &sysLogininforView)
	if err := sysLogininforService.Update(atoi, &sysLogininforView); err != nil {
		global.Logger.Error("更新解析上报数据失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	}
	if err = sysLogininforService.Update(atoi, &sysLogininforView); err != nil {
		global.Logger.Error("更新持久化失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// Get 用id查询SysLogininfor
// @Summary 用id查询SysLogininfor
// @Router /sysLogininfor/get [get]
func (sysLogininforApi *SysLogininforApi) Get(c *gin.Context) {
	id := c.Query("id")
	atoi, err := strconv.Atoi(id)
	if err != nil {
		global.Logger.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	}
	if err, sysLogininforView := sysLogininforService.Get(atoi); err != nil {
		global.Logger.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"sysLogininforView": sysLogininforView}, c)
	}
}

// Find 分页获取SysLogininfor列表
// @Summary 分页获取SysLogininfor列表
// @Router /sysLogininfor/find [get]
func (sysLogininforApi *SysLogininforApi) Find(c *gin.Context) {
	var pageInfo common.PageInfoV2
	_ = c.ShouldBindQuery(&pageInfo)
	if err := sysLogininforService.Find(&pageInfo); err != nil {
		global.Logger.Error("获取分页信息失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(pageInfo, "获取成功", c)
	}
}
