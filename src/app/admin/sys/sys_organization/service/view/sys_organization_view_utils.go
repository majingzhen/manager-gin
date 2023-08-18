// Package view
// @description <TODO description class purpose>
// @author
// @File: sys_organization
// @version 1.0.0
// @create 2023-08-18 14:00:53
package view

import (
	"fmt"
	"go.uber.org/zap"
	"manager-gin/src/app/admin/sys/sys_organization/model"
	"manager-gin/src/global"
	"manager-gin/src/utils"
)

type SysOrganizationViewUtils struct{}

func (sysOrganizationViewUtils *SysOrganizationViewUtils) Data2View(data *model.SysOrganization) (err error, view *SysOrganizationView) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysOrganizationViewUtils View2Data error: %v", e)
			global.Logger.Error("SysOrganizationViewUtils.Data2View:格式转换异常",
				zap.Any("error", e))
		}
	}()
	var tmp SysOrganizationView

	tmp.Id = data.Id

	tmp.ParentId = data.ParentId

	tmp.Ancestors = data.Ancestors

	tmp.DeptName = data.DeptName

	tmp.OrderNum = data.OrderNum

	tmp.Leader = data.Leader

	tmp.Phone = data.Phone

	tmp.Email = data.Email

	tmp.Status = data.Status

	tmp.DeletedAt = utils.Time2Str(data.DeletedAt)

	tmp.CreateBy = data.CreateBy

	tmp.CreateTime = utils.Time2Str(data.CreateTime)

	tmp.UpdateBy = data.UpdateBy

	tmp.UpdateTime = utils.Time2Str(data.UpdateTime)

	view = &tmp
	return
}
func (sysOrganizationViewUtils *SysOrganizationViewUtils) View2Data(view *SysOrganizationView) (err error, data *model.SysOrganization) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysOrganizationViewUtils View2Data error: %v", e)
			global.Logger.Error("SysOrganizationViewUtils.View2Data:格式转换异常",
				zap.Any("error", e))
		}
	}()
	var tmp model.SysOrganization

	tmp.Id = view.Id

	tmp.ParentId = view.ParentId

	tmp.Ancestors = view.Ancestors

	tmp.DeptName = view.DeptName

	tmp.OrderNum = view.OrderNum

	tmp.Leader = view.Leader

	tmp.Phone = view.Phone

	tmp.Email = view.Email

	tmp.Status = view.Status

	tmp.DeletedAt = utils.Str2Time(view.DeletedAt)

	tmp.CreateBy = view.CreateBy

	tmp.CreateTime = utils.Str2Time(view.CreateTime)

	tmp.UpdateBy = view.UpdateBy

	tmp.UpdateTime = utils.Str2Time(view.UpdateTime)

	data = &tmp
	return
}

func (sysOrganizationViewUtils *SysOrganizationViewUtils) View2DataList(viewList *[]SysOrganizationView) (err error, dataList *[]model.SysOrganization) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysOrganizationViewUtils View2DataList error: %v", e)
			global.Logger.Error("SysOrganizationViewUtils.View2DataList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if viewList != nil {
		var dataTmpList []model.SysOrganization
		for i := range *dataList {
			view := (*viewList)[i]
			err, data := sysOrganizationViewUtils.View2Data(&view)
			if err == nil {
				dataTmpList = append(dataTmpList, *data)
			}
		}
		dataList = &dataTmpList
	}
	return
}

func (sysOrganizationViewUtils *SysOrganizationViewUtils) Data2ViewList(dataList *[]model.SysOrganization) (err error, viewList *[]SysOrganizationView) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysOrganizationViewUtils Data2ViewList error: %v", e)
			global.Logger.Error("SysOrganizationViewUtils.Data2ViewList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if dataList != nil {
		var viewTmpList []SysOrganizationView
		for i := range *dataList {
			data := (*dataList)[i]
			err, view := sysOrganizationViewUtils.Data2View(&data)
			if err == nil {
				viewTmpList = append(viewTmpList, *view)
			}
		}
		viewList = &viewTmpList
	}
	return
}
