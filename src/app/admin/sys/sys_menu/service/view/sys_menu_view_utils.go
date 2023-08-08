// Package view
// @description <TODO description class purpose>
// @author
// @File: sys_menu
// @version 1.0.0
// @create 2023-08-08 10:06:19
package view

import (
	"fmt"
	"go.uber.org/zap"
	"manager-gin/src/app/admin/sys/sys_menu/model"
	"manager-gin/src/global"
	"manager-gin/src/utils"
)

type SysMenuViewUtils struct{}

func (sysMenuViewUtils *SysMenuViewUtils) Data2View(data *model.SysMenu) (err error, view *SysMenuView) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysMenuViewUtils View2Data error: %v", e)
			global.Logger.Error("SysMenuViewUtils.Data2View:格式转换异常",
				zap.Any("error", e))
		}
	}()
	var tmp SysMenuView

	tmp.Component = data.Component

	tmp.CreateBy = data.CreateBy

	tmp.CreateTime = utils.Time2Str(data.CreateTime)

	tmp.Icon = data.Icon

	tmp.Id = data.Id

	tmp.IsCache = data.IsCache

	tmp.IsFrame = data.IsFrame

	tmp.MenuName = data.MenuName

	tmp.MenuType = data.MenuType

	tmp.OrderNum = data.OrderNum

	tmp.ParentId = data.ParentId

	tmp.Path = data.Path

	tmp.Perms = data.Perms

	tmp.Query = data.Query

	tmp.Remark = data.Remark

	tmp.Status = data.Status

	tmp.UpdateBy = data.UpdateBy

	tmp.UpdateTime = utils.Time2Str(data.UpdateTime)

	tmp.Visible = data.Visible

	view = &tmp
	return
}
func (sysMenuViewUtils *SysMenuViewUtils) View2Data(view *SysMenuView) (err error, data *model.SysMenu) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysMenuViewUtils View2Data error: %v", e)
			global.Logger.Error("SysMenuViewUtils.View2Data:格式转换异常",
				zap.Any("error", e))
		}
	}()
	var tmp model.SysMenu

	tmp.Component = view.Component

	tmp.CreateBy = view.CreateBy

	tmp.CreateTime = utils.Str2Time(view.CreateTime)

	tmp.Icon = view.Icon

	tmp.Id = view.Id

	tmp.IsCache = view.IsCache

	tmp.IsFrame = view.IsFrame

	tmp.MenuName = view.MenuName

	tmp.MenuType = view.MenuType

	tmp.OrderNum = view.OrderNum

	tmp.ParentId = view.ParentId

	tmp.Path = view.Path

	tmp.Perms = view.Perms

	tmp.Query = view.Query

	tmp.Remark = view.Remark

	tmp.Status = view.Status

	tmp.UpdateBy = view.UpdateBy

	tmp.UpdateTime = utils.Str2Time(view.UpdateTime)

	tmp.Visible = view.Visible

	data = &tmp
	return
}

func (sysMenuViewUtils *SysMenuViewUtils) View2DataList(viewList *[]SysMenuView) (err error, dataList *[]model.SysMenu) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysMenuViewUtils View2DataList error: %v", e)
			global.Logger.Error("SysMenuViewUtils.View2DataList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if viewList != nil {
		var dataTmpList []model.SysMenu
		for i := range *dataList {
			view := (*viewList)[i]
			err, data := sysMenuViewUtils.View2Data(&view)
			if err == nil {
				dataTmpList = append(dataTmpList, *data)
			}
		}
		dataList = &dataTmpList
	}
	return
}

func (sysMenuViewUtils *SysMenuViewUtils) Data2ViewList(dataList *[]model.SysMenu) (err error, viewList *[]SysMenuView) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysMenuViewUtils Data2ViewList error: %v", e)
			global.Logger.Error("SysMenuViewUtils.Data2ViewList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if dataList != nil {
		var viewTmpList []SysMenuView
		for i := range *dataList {
			data := (*dataList)[i]
			err, view := sysMenuViewUtils.Data2View(&data)
			if err == nil {
				viewTmpList = append(viewTmpList, *view)
			}
		}
		viewList = &viewTmpList
	}
	return
}
