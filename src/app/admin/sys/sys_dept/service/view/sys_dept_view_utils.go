// Package view
// @description <TODO description class purpose>
// @author
// @File: sys_dept
// @version 1.0.0
// @create 2023-08-08 10:06:19
package view

import (
	"fmt"
	"go.uber.org/zap"
	"manager-gin/src/app/admin/sys/sys_dept/model"
	"manager-gin/src/global"
	"manager-gin/src/utils"
)

type SysDeptViewUtils struct{}

func (sysDeptViewUtils *SysDeptViewUtils) Data2View(data *model.SysDept) (err error, view *SysDeptView) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysDeptViewUtils View2Data error: %v", e)
			global.Logger.Error("SysDeptViewUtils.Data2View:格式转换异常",
				zap.Any("error", e))
		}
	}()
	var tmp SysDeptView

	tmp.Ancestors = data.Ancestors

	tmp.CreateBy = data.CreateBy

	tmp.CreateTime = utils.Time2Str(data.CreateTime)

	tmp.DelFlag = data.DelFlag

	tmp.DeptName = data.DeptName

	tmp.Email = data.Email

	tmp.Id = data.Id

	tmp.Leader = data.Leader

	tmp.OrderNum = data.OrderNum

	tmp.ParentId = data.ParentId

	tmp.Phone = data.Phone

	tmp.Status = data.Status

	tmp.UpdateBy = data.UpdateBy

	tmp.UpdateTime = utils.Time2Str(data.UpdateTime)

	view = &tmp
	return
}
func (sysDeptViewUtils *SysDeptViewUtils) View2Data(view *SysDeptView) (err error, data *model.SysDept) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysDeptViewUtils View2Data error: %v", e)
			global.Logger.Error("SysDeptViewUtils.View2Data:格式转换异常",
				zap.Any("error", e))
		}
	}()
	var tmp model.SysDept

	tmp.Ancestors = view.Ancestors

	tmp.CreateBy = view.CreateBy

	tmp.CreateTime = utils.Str2Time(view.CreateTime)

	tmp.DelFlag = view.DelFlag

	tmp.DeptName = view.DeptName

	tmp.Email = view.Email

	tmp.Id = view.Id

	tmp.Leader = view.Leader

	tmp.OrderNum = view.OrderNum

	tmp.ParentId = view.ParentId

	tmp.Phone = view.Phone

	tmp.Status = view.Status

	tmp.UpdateBy = view.UpdateBy

	tmp.UpdateTime = utils.Str2Time(view.UpdateTime)

	data = &tmp
	return
}

func (sysDeptViewUtils *SysDeptViewUtils) View2DataList(viewList *[]SysDeptView) (err error, dataList *[]model.SysDept) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysDeptViewUtils View2DataList error: %v", e)
			global.Logger.Error("SysDeptViewUtils.View2DataList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if viewList != nil {
		var dataTmpList []model.SysDept
		for i := range *dataList {
			view := (*viewList)[i]
			err, data := sysDeptViewUtils.View2Data(&view)
			if err == nil {
				dataTmpList = append(dataTmpList, *data)
			}
		}
		dataList = &dataTmpList
	}
	return
}

func (sysDeptViewUtils *SysDeptViewUtils) Data2ViewList(dataList *[]model.SysDept) (err error, viewList *[]SysDeptView) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysDeptViewUtils Data2ViewList error: %v", e)
			global.Logger.Error("SysDeptViewUtils.Data2ViewList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if dataList != nil {
		var viewTmpList []SysDeptView
		for i := range *dataList {
			data := (*dataList)[i]
			err, view := sysDeptViewUtils.Data2View(&data)
			if err == nil {
				viewTmpList = append(viewTmpList, *view)
			}
		}
		viewList = &viewTmpList
	}
	return
}
