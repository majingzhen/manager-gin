// Package api  TableColumnApi 自动生成模板
// @description <TODO description class purpose>
// @author
// @File: gen_table_column
// @version 1.0.0
// @create 2023-08-31 09:09:53
package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"manager-gin/src/app/admin/gen/service/table_column"
	"manager-gin/src/app/admin/gen/service/table_column/view"
	"manager-gin/src/common/basic"
	response "manager-gin/src/common/response"
	"manager-gin/src/global"
	"manager-gin/src/utils"
	"strings"
)

type TableColumnApi struct {
	basic.BasicApi
	tableColumnService table_column.TableColumnService
}

// Delete 删除TableColumn
// @Summary 删除TableColumn
// @Router /tableColumn/delete [delete]
func (api *TableColumnApi) Delete(c *gin.Context) {
	idStr := c.Param("ids")
	if idStr == "" {
		response.FailWithMessage("参数解析错误", c)
		return
	}
	ids := strings.Split(idStr, ",")
	if err := api.tableColumnService.DeleteByIds(ids); err != nil {
		global.Logger.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// Update 更新TableColumn
// @Summary 更新TableColumn
// @Router /tableColumn/update [put]
func (api *TableColumnApi) Update(c *gin.Context) {
	var tableColumnView view.TableColumnView
	err := c.ShouldBindJSON(&tableColumnView)
	if err != nil {
		response.FailWithMessage("参数解析错误", c)
		return
	}
	id := tableColumnView.Id
	if id == "" {
		response.FailWithMessage("参数解析错误", c)
		return
	}
	tableColumnView.UpdateTime = utils.GetCurTimeStr()
	tableColumnView.UpdateBy = api.GetLoginUserName(c)
	if err := api.tableColumnService.Update(id, &tableColumnView); err != nil {
		global.Logger.Error("更新持久化失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// Get 用id查询TableColumn
// @Summary 用id查询TableColumn
// @Router /tableColumn/get [get]
func (api *TableColumnApi) Get(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.FailWithMessage("参数解析错误", c)
		return
	}
	if err, tableColumnView := api.tableColumnService.Get(id); err != nil {
		global.Logger.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(tableColumnView, c)
	}
}

// Page 分页获取TableColumn列表
// @Summary 分页获取TableColumn列表
// @Router /tableColumn/page [get]
func (api *TableColumnApi) Page(c *gin.Context) {
	var pageInfo view.TableColumnPageView
	// 绑定查询参数到 pageInfo
	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		response.FailWithMessage("获取分页数据解析失败!", c)
		return
	}
	if err, res := api.tableColumnService.Page(&pageInfo); err != nil {
		global.Logger.Error("获取分页信息失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(res, "获取成功", c)
	}
}

// List 获取TableColumn列表
// @Summary 获取TableColumn列表
// @Router /tableColumn/list [get]
func (api *TableColumnApi) List(c *gin.Context) {
	var view view.TableColumnQueryView
	// 绑定查询参数到 view对象
	if err := c.ShouldBindQuery(&view); err != nil {
		response.FailWithMessage("获取参数解析失败!", c)
		return
	}
	// 判断是否需要根据用户获取数据
	// userId := framework.GetLoginUserId(c)
	if err, res := api.tableColumnService.List(&view); err != nil {
		global.Logger.Error("获取数据失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(res, "获取成功", c)
	}
}
