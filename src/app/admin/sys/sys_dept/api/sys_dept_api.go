// Package api  SysDeptApi 自动生成模板
// @description <TODO description class purpose>
// @author
// @File: sys_dept
// @version 1.0.0
// @create 2023-08-08 10:06:19
package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"manager-gin/src/app/admin/sys/sys_dept/service"
	"manager-gin/src/app/admin/sys/sys_dept/service/view"
	"manager-gin/src/common"
	"manager-gin/src/common/response"
	"manager-gin/src/global"
	"manager-gin/src/utils"
	"strconv"
)

type SysDeptApi struct {
}

var sysDeptService = service.SysDeptServiceApp

// Create 创建SysDept
// @Summary 创建SysDept
// @Router /sysDept/create [post]
func (sysDeptApi *SysDeptApi) Create(c *gin.Context) {
	var sysDeptView view.SysDeptView
	_ = c.ShouldBindJSON(&sysDeptView)
	sysDeptView.CreateTime = utils.GetCurTimeStr()
	sysDeptView.UpdateTime = utils.GetCurTimeStr()
	if err := sysDeptService.Create(&sysDeptView); err != nil {
		global.Logger.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// Delete 删除SysDept
// @Summary 删除SysDept
// @Router /sysDept/delete [delete]
func (sysDeptApi *SysDeptApi) Delete(c *gin.Context) {
	var id int
	_ = c.ShouldBindJSON(&id)
	if err := sysDeptService.Delete(id); err != nil {
		global.Logger.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteByIds 批量删除SysDept
// @Summary 批量删除SysDept
// @Router /sysDept/deleteByIds [delete]
func (sysDeptApi *SysDeptApi) DeleteByIds(c *gin.Context) {
	var ids []int
	_ = c.ShouldBindJSON(&ids)
	if err := sysDeptService.DeleteByIds(ids); err != nil {
		global.Logger.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// Update 更新SysDept
// @Summary 更新SysDept
// @Router /sysDept/update [put]
func (sysDeptApi *SysDeptApi) Update(c *gin.Context) {
	id := c.Query("id")
	atoi, err := strconv.Atoi(id)
	if err != nil {
		global.Logger.Error("更新解析上报数据失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	}
	sysDeptViewJson := c.Query("sysDeptView")
	var sysDeptView view.SysDeptView
	err = json.Unmarshal([]byte(sysDeptViewJson), &sysDeptView)
	sysDeptView.UpdateTime = utils.GetCurTimeStr()
	if err := sysDeptService.Update(atoi, &sysDeptView); err != nil {
		global.Logger.Error("更新解析上报数据失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	}
	if err = sysDeptService.Update(atoi, &sysDeptView); err != nil {
		global.Logger.Error("更新持久化失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// Get 用id查询SysDept
// @Summary 用id查询SysDept
// @Router /sysDept/get [get]
func (sysDeptApi *SysDeptApi) Get(c *gin.Context) {
	id := c.Query("id")
	atoi, err := strconv.Atoi(id)
	if err != nil {
		global.Logger.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	}
	if err, sysDeptView := sysDeptService.Get(atoi); err != nil {
		global.Logger.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"sysDeptView": sysDeptView}, c)
	}
}

// Find 分页获取SysDept列表
// @Summary 分页获取SysDept列表
// @Router /sysDept/find [get]
func (sysDeptApi *SysDeptApi) Find(c *gin.Context) {
	var pageInfo common.PageInfoV2
	_ = c.ShouldBindQuery(&pageInfo)
	if err := sysDeptService.Find(&pageInfo); err != nil {
		global.Logger.Error("获取分页信息失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(pageInfo, "获取成功", c)
	}
}
