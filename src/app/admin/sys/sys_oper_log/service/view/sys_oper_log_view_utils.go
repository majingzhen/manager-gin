// Package view
// @description <TODO description class purpose>
// @author
// @File: sys_oper_log
// @version 1.0.0
// @create 2023-08-18 13:41:26
package view

import (
	"fmt"
	"go.uber.org/zap"
	"manager-gin/src/app/admin/sys/sys_oper_log/model"
	"manager-gin/src/global"
	"manager-gin/src/utils"
)

type SysOperLogViewUtils struct{}

func (sysOperLogViewUtils *SysOperLogViewUtils) Data2View(data *model.SysOperLog) (err error, view *SysOperLogView) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysOperLogViewUtils View2Data error: %v", e)
			global.Logger.Error("SysOperLogViewUtils.Data2View:格式转换异常",
				zap.Any("error", e))
		}
	}()
	var tmp SysOperLogView

	tmp.Id = data.Id

	tmp.Title = data.Title

	tmp.BusinessType = data.BusinessType

	tmp.Method = data.Method

	tmp.RequestMethod = data.RequestMethod

	tmp.OperatorType = data.OperatorType

	tmp.OperName = data.OperName

	tmp.DeptName = data.DeptName

	tmp.OperUrl = data.OperUrl

	tmp.OperIp = data.OperIp

	tmp.OperLocation = data.OperLocation

	tmp.OperParam = data.OperParam

	tmp.JsonResult = data.JsonResult

	tmp.Status = data.Status

	tmp.ErrorMsg = data.ErrorMsg

	tmp.OperTime = utils.Time2Str(data.OperTime)

	tmp.CostTime = data.CostTime

	view = &tmp
	return
}
func (sysOperLogViewUtils *SysOperLogViewUtils) View2Data(view *SysOperLogView) (err error, data *model.SysOperLog) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysOperLogViewUtils View2Data error: %v", e)
			global.Logger.Error("SysOperLogViewUtils.View2Data:格式转换异常",
				zap.Any("error", e))
		}
	}()
	var tmp model.SysOperLog

	tmp.Id = view.Id

	tmp.Title = view.Title

	tmp.BusinessType = view.BusinessType

	tmp.Method = view.Method

	tmp.RequestMethod = view.RequestMethod

	tmp.OperatorType = view.OperatorType

	tmp.OperName = view.OperName

	tmp.DeptName = view.DeptName

	tmp.OperUrl = view.OperUrl

	tmp.OperIp = view.OperIp

	tmp.OperLocation = view.OperLocation

	tmp.OperParam = view.OperParam

	tmp.JsonResult = view.JsonResult

	tmp.Status = view.Status

	tmp.ErrorMsg = view.ErrorMsg

	tmp.OperTime = utils.Str2Time(view.OperTime)

	tmp.CostTime = view.CostTime

	data = &tmp
	return
}

func (sysOperLogViewUtils *SysOperLogViewUtils) View2DataList(viewList *[]SysOperLogView) (err error, dataList *[]model.SysOperLog) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysOperLogViewUtils View2DataList error: %v", e)
			global.Logger.Error("SysOperLogViewUtils.View2DataList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if viewList != nil {
		var dataTmpList []model.SysOperLog
		for i := range *dataList {
			view := (*viewList)[i]
			err, data := sysOperLogViewUtils.View2Data(&view)
			if err == nil {
				dataTmpList = append(dataTmpList, *data)
			}
		}
		dataList = &dataTmpList
	}
	return
}

func (sysOperLogViewUtils *SysOperLogViewUtils) Data2ViewList(dataList *[]model.SysOperLog) (err error, viewList *[]SysOperLogView) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysOperLogViewUtils Data2ViewList error: %v", e)
			global.Logger.Error("SysOperLogViewUtils.Data2ViewList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if dataList != nil {
		var viewTmpList []SysOperLogView
		for i := range *dataList {
			data := (*dataList)[i]
			err, view := sysOperLogViewUtils.Data2View(&data)
			if err == nil {
				viewTmpList = append(viewTmpList, *view)
			}
		}
		viewList = &viewTmpList
	}
	return
}
