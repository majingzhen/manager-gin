// Package view
// @description <TODO description class purpose>
// @author
// @File: sys_role
// @version 1.0.0
// @create 2023-08-21 17:37:56
package view

import (
	"fmt"
	"go.uber.org/zap"
	"manager-gin/src/app/admin/sys/model"
	"manager-gin/src/common"
	"manager-gin/src/global"
	"manager-gin/src/utils"
)

type SysRoleViewUtils struct{}

func (viewUtils *SysRoleViewUtils) Data2View(data *model.SysRole) (err error, view *SysRoleView) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysRoleViewUtils View2Data error: %v", e)
			global.Logger.Error("SysRoleViewUtils.Data2View:格式转换异常",
				zap.Any("error", e))
		}
	}()
	var tmp SysRoleView

	tmp.Id = data.Id

	tmp.RoleName = data.RoleName

	tmp.RoleKey = data.RoleKey

	tmp.RoleSort = data.RoleSort

	tmp.DataScope = data.DataScope

	if data.MenuCheckStrictly == 1 {
		tmp.MenuCheckStrictly = true
	} else {
		tmp.MenuCheckStrictly = false
	}
	if data.DeptCheckStrictly == 1 {
		tmp.DeptCheckStrictly = true
	} else {
		tmp.DeptCheckStrictly = false
	}

	tmp.Status = data.Status

	tmp.CreateBy = data.CreateBy

	tmp.CreateTime = utils.Time2Str(data.CreateTime)

	tmp.UpdateBy = data.UpdateBy

	tmp.UpdateTime = utils.Time2Str(data.UpdateTime)

	tmp.Remark = data.Remark

	view = &tmp
	return
}
func (viewUtils *SysRoleViewUtils) View2Data(view *SysRoleView) (err error, data *model.SysRole) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysRoleViewUtils View2Data error: %v", e)
			global.Logger.Error("SysRoleViewUtils.View2Data:格式转换异常",
				zap.Any("error", e))
		}
	}()
	var tmp model.SysRole

	tmp.Id = view.Id

	tmp.RoleName = view.RoleName

	tmp.RoleKey = view.RoleKey

	tmp.RoleSort = view.RoleSort

	tmp.DataScope = view.DataScope

	if view.MenuCheckStrictly {
		tmp.MenuCheckStrictly = 1
	} else {
		tmp.MenuCheckStrictly = 0
	}

	if view.DeptCheckStrictly {
		tmp.DeptCheckStrictly = 1
	} else {
		tmp.DeptCheckStrictly = 0
	}
	tmp.Status = view.Status

	tmp.CreateBy = view.CreateBy

	tmp.CreateTime = utils.Str2Time(view.CreateTime)

	tmp.UpdateBy = view.UpdateBy

	tmp.UpdateTime = utils.Str2Time(view.UpdateTime)

	tmp.Remark = view.Remark

	data = &tmp
	return
}

func (viewUtils *SysRoleViewUtils) View2DataList(viewList []*SysRoleView) (err error, dataList []*model.SysRole) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysRoleViewUtils View2DataList error: %v", e)
			global.Logger.Error("SysRoleViewUtils.View2DataList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if viewList != nil {
		var dataTmpList []*model.SysRole
		for i := range viewList {
			view := (viewList)[i]
			err, data := viewUtils.View2Data(view)
			if err == nil {
				dataTmpList = append(dataTmpList, data)
			}
		}
		dataList = dataTmpList
	}
	return
}

func (viewUtils *SysRoleViewUtils) Data2ViewList(dataList []*model.SysRole) (err error, viewList []*SysRoleView) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysRoleViewUtils Data2ViewList error: %v", e)
			global.Logger.Error("SysRoleViewUtils.Data2ViewList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if dataList != nil {
		var viewTmpList []*SysRoleView
		for i := range dataList {
			data := (dataList)[i]
			err, view := viewUtils.Data2View(data)
			if err == nil {
				viewTmpList = append(viewTmpList, view)
			}
		}
		viewList = viewTmpList
	}
	return
}

func (viewUtils *SysRoleViewUtils) PageData2ViewList(pageInfo *common.PageInfo) (err error, res *common.PageInfo) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysRoleViewUtils PageData2ViewList error: %v", e)
			global.Logger.Error("SysRoleViewUtils.PageData2ViewList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if pageInfo != nil && pageInfo.Rows != nil {
		if p, ok := pageInfo.Rows.([]*model.SysRole); ok {
			if err, viewList := viewUtils.Data2ViewList(p); err == nil {
				pageInfo.Rows = viewList
			}
		}
	}
	res = pageInfo
	return
}
