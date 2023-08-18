// Package view
// @description <TODO description class purpose>
// @author
// @File: sys_user_role
// @version 1.0.0
// @create 2023-08-18 13:41:26
package view

import (
	"fmt"
	"go.uber.org/zap"
	"manager-gin/src/app/admin/sys/sys_user_role/model"
	"manager-gin/src/global"
	"manager-gin/src/utils"
)

type SysUserRoleViewUtils struct{}

func (sysUserRoleViewUtils *SysUserRoleViewUtils) Data2View(data *model.SysUserRole) (err error, view *SysUserRoleView) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysUserRoleViewUtils View2Data error: %v", e)
			global.Logger.Error("SysUserRoleViewUtils.Data2View:格式转换异常",
				zap.Any("error", e))
		}
	}()
	var tmp SysUserRoleView

	tmp.UserId = data.UserId

	tmp.RoleId = data.RoleId

	view = &tmp
	return
}
func (sysUserRoleViewUtils *SysUserRoleViewUtils) View2Data(view *SysUserRoleView) (err error, data *model.SysUserRole) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysUserRoleViewUtils View2Data error: %v", e)
			global.Logger.Error("SysUserRoleViewUtils.View2Data:格式转换异常",
				zap.Any("error", e))
		}
	}()
	var tmp model.SysUserRole

	tmp.UserId = view.UserId

	tmp.RoleId = view.RoleId

	data = &tmp
	return
}

func (sysUserRoleViewUtils *SysUserRoleViewUtils) View2DataList(viewList *[]SysUserRoleView) (err error, dataList *[]model.SysUserRole) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysUserRoleViewUtils View2DataList error: %v", e)
			global.Logger.Error("SysUserRoleViewUtils.View2DataList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if viewList != nil {
		var dataTmpList []model.SysUserRole
		for i := range *dataList {
			view := (*viewList)[i]
			err, data := sysUserRoleViewUtils.View2Data(&view)
			if err == nil {
				dataTmpList = append(dataTmpList, *data)
			}
		}
		dataList = &dataTmpList
	}
	return
}

func (sysUserRoleViewUtils *SysUserRoleViewUtils) Data2ViewList(dataList *[]model.SysUserRole) (err error, viewList *[]SysUserRoleView) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysUserRoleViewUtils Data2ViewList error: %v", e)
			global.Logger.Error("SysUserRoleViewUtils.Data2ViewList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if dataList != nil {
		var viewTmpList []SysUserRoleView
		for i := range *dataList {
			data := (*dataList)[i]
			err, view := sysUserRoleViewUtils.Data2View(&data)
			if err == nil {
				viewTmpList = append(viewTmpList, *view)
			}
		}
		viewList = &viewTmpList
	}
	return
}
