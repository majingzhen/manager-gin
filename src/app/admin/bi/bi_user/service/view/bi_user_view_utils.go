// Package view
// @description <TODO description class purpose>
// @author Administrator
// @File: bi_user
// @version 1.0.0
// @create 2023-08-16 14:53:36
package view

import (
	"fmt"
	"go.uber.org/zap"
	"manager-gin/src/app/admin/bi/bi_user/model"
	"manager-gin/src/global"
	"manager-gin/src/utils"
)

type BiUserViewUtils struct{}

func (biUserViewUtils *BiUserViewUtils) Data2View(data *model.BiUser) (err error, view *BiUserView) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("BiUserViewUtils View2Data error: %v", e)
			global.Logger.Error("BiUserViewUtils.Data2View:格式转换异常",
				zap.Any("error", e))
		}
	}()
	var tmp BiUserView
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
func (biUserViewUtils *BiUserViewUtils) View2Data(view *BiUserView) (err error, data *model.BiUser) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("BiUserViewUtils View2Data error: %v", e)
			global.Logger.Error("BiUserViewUtils.View2Data:格式转换异常",
				zap.Any("error", e))
		}
	}()
	var tmp model.BiUser

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

func (biUserViewUtils *BiUserViewUtils) View2DataList(viewList *[]BiUserView) (err error, dataList *[]model.BiUser) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("BiUserViewUtils View2DataList error: %v", e)
			global.Logger.Error("BiUserViewUtils.View2DataList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if viewList != nil {
		var dataTmpList []model.BiUser
		for i := range *dataList {
			view := (*viewList)[i]
			err, data := biUserViewUtils.View2Data(&view)
			if err == nil {
				dataTmpList = append(dataTmpList, *data)
			}
		}
		dataList = &dataTmpList
	}
	return
}

func (biUserViewUtils *BiUserViewUtils) Data2ViewList(dataList *[]model.BiUser) (err error, viewList *[]BiUserView) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("BiUserViewUtils Data2ViewList error: %v", e)
			global.Logger.Error("BiUserViewUtils.Data2ViewList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if dataList != nil {
		var viewTmpList []BiUserView
		for i := range *dataList {
			data := (*dataList)[i]
			err, view := biUserViewUtils.Data2View(&data)
			if err == nil {
				viewTmpList = append(viewTmpList, *view)
			}
		}
		viewList = &viewTmpList
	}
	return
}
