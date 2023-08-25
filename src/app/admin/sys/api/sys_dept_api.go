// Package api  SysDeptApi 自动生成模板
// @description <TODO description class purpose>
// @author
// @File: sys_dept
// @version 1.0.0
// @create 2023-08-21 10:27:01
package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"manager-gin/src/app/admin/sys/service/sys_dept"
	"manager-gin/src/app/admin/sys/service/sys_dept/view"
	response "manager-gin/src/common/response"
	"manager-gin/src/framework"
	"manager-gin/src/global"
	"manager-gin/src/utils"
	"strings"
)

type SysDeptApi struct {
	sysDeptService sys_dept.SysDeptService
}

// Create 创建SysDept
// @Summary 创建SysDept
// @Router /sysDept/create [post]
func (api *SysDeptApi) Create(c *gin.Context) {
	var sysDeptView view.SysDeptView
	_ = c.ShouldBindJSON(&sysDeptView)
	sysDeptView.Id = utils.GenUID()
	sysDeptView.CreateTime = utils.GetCurTimeStr()
	sysDeptView.UpdateTime = utils.GetCurTimeStr()
	sysDeptView.CreateBy = framework.GetLoginUserName(c)
	if err := api.sysDeptService.Create(&sysDeptView); err != nil {
		global.Logger.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// Delete 删除SysDept
// @Summary 删除SysDept
// @Router /sysDept/delete [delete]
func (api *SysDeptApi) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := api.sysDeptService.Delete(id); err != nil {
		global.Logger.Error("删除失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// Update 更新SysDept
// @Summary 更新SysDept
// @Router /sysDept/update [put]
func (api *SysDeptApi) Update(c *gin.Context) {
	var sysDeptView view.SysDeptView
	_ = c.ShouldBindJSON(&sysDeptView)
	id := sysDeptView.Id
	if id == "" {
		response.FailWithMessage("更新失败", c)
		return
	}
	sysDeptView.UpdateTime = utils.GetCurTimeStr()
	sysDeptView.UpdateBy = framework.GetLoginUser(c).UserName
	if err := api.sysDeptService.Update(id, &sysDeptView); err != nil {
		global.Logger.Error("更新持久化失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// Get 用id查询SysDept
// @Summary 用id查询SysDept
// @Router /sysDept/get [get]
func (api *SysDeptApi) Get(c *gin.Context) {
	id := c.Param("id")
	if err, sysDeptView := api.sysDeptService.Get(id); err != nil {
		global.Logger.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(sysDeptView, c)
	}
}

// Page 分页获取SysDept列表
// @Summary 分页获取SysDept列表
// @Router /sysDept/page [get]
func (api *SysDeptApi) Page(c *gin.Context) {
	var pageInfo view.SysDeptPageView
	// 绑定查询参数到 pageInfo
	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		response.FailWithMessage("获取分页数据解析失败!", c)
		return
	}

	if err, res := api.sysDeptService.Page(&pageInfo); err != nil {
		global.Logger.Error("获取分页信息失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(res, "获取成功", c)
	}
}

// List 获取SysDept列表
// @Summary 获取SysDept列表
// @Router /sysDept/list [get]
func (api *SysDeptApi) List(c *gin.Context) {
	var view view.SysDeptView
	// 绑定查询参数到 view对象
	if err := c.ShouldBindQuery(&view); err != nil {
		response.FailWithMessage("获取参数解析失败!", c)
		return
	}
	// 判断是否需要根据用户获取数据
	// userId := framework.GetLoginUserId(c)
	user := framework.GetLoginUser(c)
	if err, res := api.sysDeptService.List(&view, user); err != nil {
		global.Logger.Error("获取数据失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(res, "获取成功", c)
	}
}

// ListExclude 查询部门列表（排除节点）
// @Summary 查询部门列表（排除节点）
// @Router /sysDept/listExclude [get]
func (api *SysDeptApi) ListExclude(c *gin.Context) {
	id := c.Param("id")
	user := framework.GetLoginUser(c)
	if err, sysDeptView := api.sysDeptService.List(&view.SysDeptView{}, user); err != nil {
		global.Logger.Error("查询失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		var deleteIndexes []int // 用于记录要删除的元素索引
		for i, deptView := range sysDeptView {
			parentIds := strings.Split(deptView.Ancestors, ",")
			// 要把查询的这个节点排除掉
			if deptView.Id == id || utils.Contains(parentIds, id) {
				deleteIndexes = append(deleteIndexes, i)
			}
		}
		// 根据记录的索引删除元素
		for i := len(deleteIndexes) - 1; i >= 0; i-- {
			sysDeptView = append(sysDeptView[:deleteIndexes[i]], sysDeptView[deleteIndexes[i]+1:]...)
		}
		response.OkWithData(sysDeptView, c)
	}
}

// SelectDeptTree 查询部门树列表
// @Summary 查询部门树列表
// @Router /sysDept/selectDeptTree [get]
func (api *SysDeptApi) SelectDeptTree(c *gin.Context) {
	var view view.SysDeptView
	// 绑定查询参数到 view对象
	if err := c.ShouldBindQuery(&view); err != nil {
		response.FailWithMessage("获取参数解析失败!", c)
		return
	}
	user := framework.GetLoginUser(c)
	if err, sysDeptView := api.sysDeptService.SelectDeptTree(&view, user); err != nil {
		global.Logger.Error("查询失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(sysDeptView, c)
	}
}

// SelectDeptTreeByRole 查询部门树列表
// @Summary 根据角色查询部门树列表
// @Router /sysDept/selectDeptTreeByRole [get]
func (api *SysDeptApi) SelectDeptTreeByRole(c *gin.Context) {
	roleId := c.Param("roleId")
	if roleId == "" {
		response.FailWithMessage("获取参数解析失败!", c)
		return
	}
	if err, checkedKeys := api.sysDeptService.SelectDeptTreeByRole(roleId); err != nil {
		global.Logger.Error("查询失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		err1, deptTree := api.sysDeptService.SelectDeptTree(&view.SysDeptView{}, framework.GetLoginUser(c))
		if err1 != nil {
			global.Logger.Error("查询失败!", zap.Error(err))
			response.FailWithMessage(err.Error(), c)
			return
		}
		response.OkWithData(gin.H{
			"checkedKeys": checkedKeys,
			"depts":       deptTree,
		}, c)
	}
}
