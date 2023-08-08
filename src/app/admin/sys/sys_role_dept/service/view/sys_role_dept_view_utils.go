// Package view
// @description <TODO description class purpose>
// @author
// @File: sys_role_dept
// @version 1.0.0
// @create 2023-08-08 10:06:19
package view

import (
	"fmt"
	"go.uber.org/zap"
	"manager-gin/src/app/admin/sys/sys_role_dept/model"
	"manager-gin/src/global"
	"manager-gin/src/utils"
)

type SysRoleDeptViewUtils struct{}

func (sysRoleDeptViewUtils *SysRoleDeptViewUtils) Data2View(data *model.SysRoleDept) (err error, view *SysRoleDeptView) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysRoleDeptViewUtils View2Data error: %v", e)
			global.Logger.Error("SysRoleDeptViewUtils.Data2View:格式转换异常",
				zap.Any("error", e))
		}
	}()
	var tmp SysRoleDeptView

	tmp.DeptId = data.DeptId

	tmp.RoleId = data.RoleId

	view = &tmp
	return
}
func (sysRoleDeptViewUtils *SysRoleDeptViewUtils) View2Data(view *SysRoleDeptView) (err error, data *model.SysRoleDept) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysRoleDeptViewUtils View2Data error: %v", e)
			global.Logger.Error("SysRoleDeptViewUtils.View2Data:格式转换异常",
				zap.Any("error", e))
		}
	}()
	var tmp model.SysRoleDept

	tmp.DeptId = view.DeptId

	tmp.RoleId = view.RoleId

	data = &tmp
	return
}

func (sysRoleDeptViewUtils *SysRoleDeptViewUtils) View2DataList(viewList *[]SysRoleDeptView) (err error, dataList *[]model.SysRoleDept) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysRoleDeptViewUtils View2DataList error: %v", e)
			global.Logger.Error("SysRoleDeptViewUtils.View2DataList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if viewList != nil {
		var dataTmpList []model.SysRoleDept
		for i := range *dataList {
			view := (*viewList)[i]
			err, data := sysRoleDeptViewUtils.View2Data(&view)
			if err == nil {
				dataTmpList = append(dataTmpList, *data)
			}
		}
		dataList = &dataTmpList
	}
	return
}

func (sysRoleDeptViewUtils *SysRoleDeptViewUtils) Data2ViewList(dataList *[]model.SysRoleDept) (err error, viewList *[]SysRoleDeptView) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysRoleDeptViewUtils Data2ViewList error: %v", e)
			global.Logger.Error("SysRoleDeptViewUtils.Data2ViewList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if dataList != nil {
		var viewTmpList []SysRoleDeptView
		for i := range *dataList {
			data := (*dataList)[i]
			err, view := sysRoleDeptViewUtils.Data2View(&data)
			if err == nil {
				viewTmpList = append(viewTmpList, *view)
			}
		}
		viewList = &viewTmpList
	}
	return
}
