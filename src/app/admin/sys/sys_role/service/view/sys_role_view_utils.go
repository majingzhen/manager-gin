// Package view
// @description <TODO description class purpose>
// @author
// @File: sys_role
// @version 1.0.0
// @create 2023-08-08 10:06:19
package view

import (
	"fmt"
	"go.uber.org/zap"
	"manager-gin/src/app/admin/sys/sys_role/model"
	"manager-gin/src/global"
	"manager-gin/src/utils"
)

type SysRoleViewUtils struct{}

func (sysRoleViewUtils *SysRoleViewUtils) Data2View(data *model.SysRole) (err error, view *SysRoleView) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysRoleViewUtils View2Data error: %v", e)
			global.Logger.Error("SysRoleViewUtils.Data2View:格式转换异常",
				zap.Any("error", e))
		}
	}()
	var tmp SysRoleView

	tmp.CreateBy = data.CreateBy

	tmp.CreateTime = utils.Time2Str(data.CreateTime)

	tmp.DataScope = data.DataScope

	tmp.DelFlag = data.DelFlag

	tmp.DeptCheckStrictly = data.DeptCheckStrictly

	tmp.Id = data.Id

	tmp.MenuCheckStrictly = data.MenuCheckStrictly

	tmp.Remark = data.Remark

	tmp.RoleKey = data.RoleKey

	tmp.RoleName = data.RoleName

	tmp.RoleSort = data.RoleSort

	tmp.Status = data.Status

	tmp.UpdateBy = data.UpdateBy

	tmp.UpdateTime = utils.Time2Str(data.UpdateTime)

	view = &tmp
	return
}
func (sysRoleViewUtils *SysRoleViewUtils) View2Data(view *SysRoleView) (err error, data *model.SysRole) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysRoleViewUtils View2Data error: %v", e)
			global.Logger.Error("SysRoleViewUtils.View2Data:格式转换异常",
				zap.Any("error", e))
		}
	}()
	var tmp model.SysRole

	tmp.CreateBy = view.CreateBy

	tmp.CreateTime = utils.Str2Time(view.CreateTime)

	tmp.DataScope = view.DataScope

	tmp.DelFlag = view.DelFlag

	tmp.DeptCheckStrictly = view.DeptCheckStrictly

	tmp.Id = view.Id

	tmp.MenuCheckStrictly = view.MenuCheckStrictly

	tmp.Remark = view.Remark

	tmp.RoleKey = view.RoleKey

	tmp.RoleName = view.RoleName

	tmp.RoleSort = view.RoleSort

	tmp.Status = view.Status

	tmp.UpdateBy = view.UpdateBy

	tmp.UpdateTime = utils.Str2Time(view.UpdateTime)

	data = &tmp
	return
}

func (sysRoleViewUtils *SysRoleViewUtils) View2DataList(viewList *[]SysRoleView) (err error, dataList *[]model.SysRole) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysRoleViewUtils View2DataList error: %v", e)
			global.Logger.Error("SysRoleViewUtils.View2DataList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if viewList != nil {
		var dataTmpList []model.SysRole
		for i := range *dataList {
			view := (*viewList)[i]
			err, data := sysRoleViewUtils.View2Data(&view)
			if err == nil {
				dataTmpList = append(dataTmpList, *data)
			}
		}
		dataList = &dataTmpList
	}
	return
}

func (sysRoleViewUtils *SysRoleViewUtils) Data2ViewList(dataList *[]model.SysRole) (err error, viewList *[]SysRoleView) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysRoleViewUtils Data2ViewList error: %v", e)
			global.Logger.Error("SysRoleViewUtils.Data2ViewList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if dataList != nil {
		var viewTmpList []SysRoleView
		for i := range *dataList {
			data := (*dataList)[i]
			err, view := sysRoleViewUtils.Data2View(&data)
			if err == nil {
				viewTmpList = append(viewTmpList, *view)
			}
		}
		viewList = &viewTmpList
	}
	return
}
