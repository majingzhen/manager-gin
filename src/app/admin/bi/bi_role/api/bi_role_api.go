// Package api  BiRoleApi 自动生成模板
// @description <TODO description class purpose>
// @author Majz
// @File: bi_role
// @version 1.0.0
// @create 2023-08-16 14:53:36
package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"manager-gin/src/app/admin/bi/bi_role/service"
	"manager-gin/src/app/admin/bi/bi_role/service/view"
	"manager-gin/src/common"
	response "manager-gin/src/common/response"
	"manager-gin/src/global"
	"manager-gin/src/utils"
	"net/http"
	"strconv"
)

type BiRoleApi struct {
}

var biRoleService = service.BiRoleServiceApp

// Create 创建BiRole
// @Summary 创建BiRole
// @Router /biRole/create [post]
func (biRoleApi *BiRoleApi) Create(c *gin.Context) {
	var biRoleView view.BiRoleView
	_ = c.ShouldBindJSON(&biRoleView)
	biRoleView.CreateTime = utils.GetCurTimeStr()
	biRoleView.UpdateTime = utils.GetCurTimeStr()
	if err := biRoleService.Create(&biRoleView); err != nil {
		global.Logger.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// Delete 删除BiRole
// @Summary 删除BiRole
// @Router /biRole/delete [delete]
func (biRoleApi *BiRoleApi) Delete(c *gin.Context) {
	var id int
	_ = c.ShouldBindJSON(&id)
	if err := biRoleService.Delete(id); err != nil {
		global.Logger.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteByIds 批量删除BiRole
// @Summary 批量删除BiRole
// @Router /biRole/deleteByIds [delete]
func (biRoleApi *BiRoleApi) DeleteByIds(c *gin.Context) {
	var ids []int
	_ = c.ShouldBindJSON(&ids)
	if err := biRoleService.DeleteByIds(ids); err != nil {
		global.Logger.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// Update 更新BiRole
// @Summary 更新BiRole
// @Router /biRole/update [put]
func (biRoleApi *BiRoleApi) Update(c *gin.Context) {
	id := c.Query("id")
	atoi, err := strconv.Atoi(id)
	if err != nil {
		global.Logger.Error("更新解析上报数据失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	}
	biRoleViewJson := c.Query("biRoleView")
	var biRoleView view.BiRoleView
	err = json.Unmarshal([]byte(biRoleViewJson), &biRoleView)
	biRoleView.UpdateTime = utils.GetCurTimeStr()
	if err := biRoleService.Update(atoi, &biRoleView); err != nil {
		global.Logger.Error("更新解析上报数据失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	}
	if err = biRoleService.Update(atoi, &biRoleView); err != nil {
		global.Logger.Error("更新持久化失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// Get 用id查询BiRole
// @Summary 用id查询BiRole
// @Router /biRole/get [get]
func (biRoleApi *BiRoleApi) Get(c *gin.Context) {
	id := c.Query("id")
	atoi, err := strconv.Atoi(id)
	if err != nil {
		global.Logger.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	}
	if err, biRoleView := biRoleService.Get(atoi); err != nil {
		global.Logger.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"biRoleView": biRoleView}, c)
	}
}

// Find 分页获取BiRole列表
// @Summary 分页获取BiRole列表
// @Router /biRole/find [get]
func (biRoleApi *BiRoleApi) Find(c *gin.Context) {
	var pageInfo common.PageInfoV2
	// 绑定查询参数到 pageInfo
	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 调用 Calculate 方法自动计算 Limit 和 Offset
	pageInfo.Calculate()
	if err := biRoleService.Find(&pageInfo); err != nil {
		global.Logger.Error("获取分页信息失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(pageInfo, "获取成功", c)
	}
}
