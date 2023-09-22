// Package api  TableApi 自动生成模板
// @description <TODO description class purpose>
// @author
// @File: gen_table
// @version 1.0.0
// @create 2023-08-31 09:09:53
package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"manager-gin/src/app/admin/gen/service/table"
	"manager-gin/src/app/admin/gen/service/table/view"
	"manager-gin/src/app/admin/gen/service/table_column"
	"manager-gin/src/common/basic"
	response "manager-gin/src/common/response"
	"manager-gin/src/global"
	"manager-gin/src/utils"
	"strings"
)

type TableApi struct {
	basic.BasicApi
	tableService  table.Service
	columnService table_column.TableColumnService
}

// Create 创建Table
// @Summary 创建Table
// @Router /table/create [post]
func (api *TableApi) Create(c *gin.Context) {
	var tableView view.TableView
	_ = c.ShouldBindJSON(&tableView)
	tableView.Id = utils.GenUID()
	tableView.CreateTime = utils.GetCurTimeStr()
	tableView.UpdateTime = utils.GetCurTimeStr()
	tableView.CreateBy = api.GetLoginUserName(c)
	if err := api.tableService.Create(&tableView); err != nil {
		global.Logger.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// Delete 删除Table
// @Summary 删除Table
// @Router /table/delete [delete]
func (api *TableApi) Delete(c *gin.Context) {
	idStr := c.Param("ids")
	if idStr == "" {
		response.FailWithMessage("参数解析错误", c)
		return
	}
	ids := strings.Split(idStr, ",")
	if err := api.tableService.DeleteByIds(ids); err != nil {
		global.Logger.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// Update 更新Table
// @Summary 更新Table
// @Router /table/update [put]
func (api *TableApi) Update(c *gin.Context) {
	var tableView view.TableView
	err := c.ShouldBindJSON(&tableView)
	if err != nil {
		response.FailWithMessage("参数解析错误", c)
		return
	}
	// validateEdit
	if err := api.tableService.ValidateEdit(&tableView); err != nil {
		response.FailWithMessage("参数解析错误"+err.Error(), c)
		return
	}
	tableView.UpdateTime = utils.GetCurTimeStr()
	tableView.UpdateBy = api.GetLoginUserName(c)
	if err := api.tableService.Update(&tableView); err != nil {
		global.Logger.Error("更新持久化失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// Get 用id查询Table
// @Summary 用id查询Table
// @Router /table/get [get]
func (api *TableApi) Get(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.FailWithMessage("参数解析错误", c)
		return
	}
	if err, tableView := api.tableService.Get(id); err != nil {
		global.Logger.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(tableView, c)
	}
}

// Page 分页获取Table列表
// @Summary 分页获取Table列表
// @Router /table/page [get]
func (api *TableApi) Page(c *gin.Context) {
	var pageInfo view.TablePageView
	// 绑定查询参数到 pageInfo
	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		response.FailWithMessage("获取分页数据解析失败!", c)
		return
	}
	if err, res := api.tableService.Page(&pageInfo); err != nil {
		global.Logger.Error("获取分页信息失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(res, "获取成功", c)
	}
}

// List 获取Table列表
// @Summary 获取Table列表
// @Router /table/list [get]
func (api *TableApi) List(c *gin.Context) {
	var view view.TableQueryView
	// 绑定查询参数到 view对象
	if err := c.ShouldBindQuery(&view); err != nil {
		response.FailWithMessage("获取参数解析失败!", c)
		return
	}
	// 判断是否需要根据用户获取数据
	// userId := framework.GetLoginUserId(c)
	if err, res := api.tableService.List(&view); err != nil {
		global.Logger.Error("获取数据失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(res, "获取成功", c)
	}
}

// SelectDbTablePage 获取数据库表列表
// @Summary 获取数据库表列表
// @Router /table/SelectDbTablePage [get]
func (api *TableApi) SelectDbTablePage(c *gin.Context) {
	var view view.TablePageView
	// 绑定查询参数到 view对象
	if err := c.ShouldBindQuery(&view); err != nil {
		response.FailWithMessage("获取参数解析失败!", c)
		return
	}
	// 判断是否需要根据用户获取数据
	// userId := framework.GetLoginUserId(c)
	if err, res := api.tableService.SelectDbTablePage(&view); err != nil {
		global.Logger.Error("获取数据失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(res, "获取成功", c)
	}
}

// ImportTable 导入Table
// @Summary 导入Table
// @Router /table/importTable [get]
func (api *TableApi) ImportTable(c *gin.Context) {
	tables := c.Query("tables")
	if tables == "" {
		global.Logger.Error("参数解析失败!")
		response.FailWithMessage("参数解析失败", c)
		return
	}
	if err := api.tableService.ImportTable(tables, api.GetLoginUserName(c)); err != nil {
		global.Logger.Error("导入失败!", zap.Error(err))
		response.FailWithMessage("导入失败", c)
	} else {
		response.OkWithMessage("导入成功", c)
	}
}

// Preview 预览代码
// @Summary 预览代码
// @Router /table/preview [get]
func (api *TableApi) Preview(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		global.Logger.Error("参数解析失败!")
		response.FailWithMessage("参数解析失败", c)
		return
	}
	if err, code := api.tableService.PreviewCode(id); err != nil {
		global.Logger.Error("预览失败!", zap.Error(err))
		response.FailWithMessage("预览失败", c)
	} else {
		response.OkWithDetailed(code, "预览成功", c)
	}
}

// SyncDb 同步数据库
func (api *TableApi) SyncDb(c *gin.Context) {
	dbName := c.Param("db_name")
	if dbName == "" {
		global.Logger.Error("参数解析失败！")
		response.FailWithMessage("参数解析失败", c)
		return
	}
	if err := api.tableService.SyncDb(dbName); err != nil {
		global.Logger.Error("同步数据库失败！", zap.Error(err))
		response.FailWithMessage("同步数据库失败", c)
	} else {
		response.OkWithMessage("同步数据库成功", c)
	}
}
