// Package view
// @description <TODO description class purpose>
// @author Majz
// @File: bi_role
// @version 1.0.0
// @create 2023-08-16 14:53:36
package view

import (
	"fmt"
	"go.uber.org/zap"
	"manager-gin/src/app/admin/bi/bi_role/model"
	"manager-gin/src/global"
	"manager-gin/src/utils"
)

type BiRoleViewUtils struct{}

func (biRoleViewUtils *BiRoleViewUtils) Data2View(data *model.BiRole) (err error, view *BiRoleView) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("BiRoleViewUtils View2Data error: %v", e)
			global.Logger.Error("BiRoleViewUtils.Data2View:格式转换异常",
				zap.Any("error", e))
		}
	}()
	var tmp BiRoleView

	tmp.Id = data.Id

	tmp.TestName = data.TestName

	tmp.Gender = data.Gender

	tmp.Remark = data.Remark

	tmp.Birth = utils.Time2Str(data.Birth)

	tmp.CreateTime = utils.Time2Str(data.CreateTime)

	tmp.UpdateTime = utils.Time2Str(data.UpdateTime)

	view = &tmp
	return
}
func (biRoleViewUtils *BiRoleViewUtils) View2Data(view *BiRoleView) (err error, data *model.BiRole) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("BiRoleViewUtils View2Data error: %v", e)
			global.Logger.Error("BiRoleViewUtils.View2Data:格式转换异常",
				zap.Any("error", e))
		}
	}()
	var tmp model.BiRole

	tmp.Id = view.Id

	tmp.TestName = view.TestName

	tmp.Gender = view.Gender

	tmp.Remark = view.Remark

	tmp.Birth = utils.Str2Time(view.Birth)

	tmp.CreateTime = utils.Str2Time(view.CreateTime)

	tmp.UpdateTime = utils.Str2Time(view.UpdateTime)

	data = &tmp
	return
}

func (biRoleViewUtils *BiRoleViewUtils) View2DataList(viewList *[]BiRoleView) (err error, dataList *[]model.BiRole) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("BiRoleViewUtils View2DataList error: %v", e)
			global.Logger.Error("BiRoleViewUtils.View2DataList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if viewList != nil {
		var dataTmpList []model.BiRole
		for i := range *dataList {
			view := (*viewList)[i]
			err, data := biRoleViewUtils.View2Data(&view)
			if err == nil {
				dataTmpList = append(dataTmpList, *data)
			}
		}
		dataList = &dataTmpList
	}
	return
}

func (biRoleViewUtils *BiRoleViewUtils) Data2ViewList(dataList *[]model.BiRole) (err error, viewList *[]BiRoleView) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("BiRoleViewUtils Data2ViewList error: %v", e)
			global.Logger.Error("BiRoleViewUtils.Data2ViewList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if dataList != nil {
		var viewTmpList []BiRoleView
		for i := range *dataList {
			data := (*dataList)[i]
			err, view := biRoleViewUtils.Data2View(&data)
			if err == nil {
				viewTmpList = append(viewTmpList, *view)
			}
		}
		viewList = &viewTmpList
	}
	return
}
