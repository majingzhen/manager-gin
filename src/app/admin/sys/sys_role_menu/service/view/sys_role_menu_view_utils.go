// Package view
// @description <TODO description class purpose>
// @author
// @File: sys_role_menu
// @version 1.0.0
// @create 2023-08-08 10:06:19
package view

import (
	"fmt"
	"go.uber.org/zap"
	"manager-gin/src/app/admin/sys/sys_role_menu/model"
	"manager-gin/src/global"
	"manager-gin/src/utils"
)

type SysRoleMenuViewUtils struct{}

func (sysRoleMenuViewUtils *SysRoleMenuViewUtils) Data2View(data *model.SysRoleMenu) (err error, view *SysRoleMenuView) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysRoleMenuViewUtils View2Data error: %v", e)
			global.Logger.Error("SysRoleMenuViewUtils.Data2View:格式转换异常",
				zap.Any("error", e))
		}
	}()
	var tmp SysRoleMenuView

	tmp.MenuId = data.MenuId

	tmp.RoleId = data.RoleId

	view = &tmp
	return
}
func (sysRoleMenuViewUtils *SysRoleMenuViewUtils) View2Data(view *SysRoleMenuView) (err error, data *model.SysRoleMenu) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysRoleMenuViewUtils View2Data error: %v", e)
			global.Logger.Error("SysRoleMenuViewUtils.View2Data:格式转换异常",
				zap.Any("error", e))
		}
	}()
	var tmp model.SysRoleMenu

	tmp.MenuId = view.MenuId

	tmp.RoleId = view.RoleId

	data = &tmp
	return
}

func (sysRoleMenuViewUtils *SysRoleMenuViewUtils) View2DataList(viewList *[]SysRoleMenuView) (err error, dataList *[]model.SysRoleMenu) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysRoleMenuViewUtils View2DataList error: %v", e)
			global.Logger.Error("SysRoleMenuViewUtils.View2DataList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if viewList != nil {
		var dataTmpList []model.SysRoleMenu
		for i := range *dataList {
			view := (*viewList)[i]
			err, data := sysRoleMenuViewUtils.View2Data(&view)
			if err == nil {
				dataTmpList = append(dataTmpList, *data)
			}
		}
		dataList = &dataTmpList
	}
	return
}

func (sysRoleMenuViewUtils *SysRoleMenuViewUtils) Data2ViewList(dataList *[]model.SysRoleMenu) (err error, viewList *[]SysRoleMenuView) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysRoleMenuViewUtils Data2ViewList error: %v", e)
			global.Logger.Error("SysRoleMenuViewUtils.Data2ViewList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if dataList != nil {
		var viewTmpList []SysRoleMenuView
		for i := range *dataList {
			data := (*dataList)[i]
			err, view := sysRoleMenuViewUtils.Data2View(&data)
			if err == nil {
				viewTmpList = append(viewTmpList, *view)
			}
		}
		viewList = &viewTmpList
	}
	return
}
