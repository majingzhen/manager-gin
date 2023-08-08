// Package view
// @description <TODO description class purpose>
// @author
// @File: sys_job_log
// @version 1.0.0
// @create 2023-08-08 10:06:19
package view

import (
	"fmt"
	"go.uber.org/zap"
	"manager-gin/src/app/admin/sys/sys_job_log/model"
	"manager-gin/src/global"
	"manager-gin/src/utils"
)

type SysJobLogViewUtils struct{}

func (sysJobLogViewUtils *SysJobLogViewUtils) Data2View(data *model.SysJobLog) (err error, view *SysJobLogView) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysJobLogViewUtils View2Data error: %v", e)
			global.Logger.Error("SysJobLogViewUtils.Data2View:格式转换异常",
				zap.Any("error", e))
		}
	}()
	var tmp SysJobLogView

	tmp.CreateTime = utils.Time2Str(data.CreateTime)

	tmp.ExceptionInfo = data.ExceptionInfo

	tmp.Id = data.Id

	tmp.InvokeTarget = data.InvokeTarget

	tmp.JobGroup = data.JobGroup

	tmp.JobMessage = data.JobMessage

	tmp.JobName = data.JobName

	tmp.Status = data.Status

	view = &tmp
	return
}
func (sysJobLogViewUtils *SysJobLogViewUtils) View2Data(view *SysJobLogView) (err error, data *model.SysJobLog) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysJobLogViewUtils View2Data error: %v", e)
			global.Logger.Error("SysJobLogViewUtils.View2Data:格式转换异常",
				zap.Any("error", e))
		}
	}()
	var tmp model.SysJobLog

	tmp.CreateTime = utils.Str2Time(view.CreateTime)

	tmp.ExceptionInfo = view.ExceptionInfo

	tmp.Id = view.Id

	tmp.InvokeTarget = view.InvokeTarget

	tmp.JobGroup = view.JobGroup

	tmp.JobMessage = view.JobMessage

	tmp.JobName = view.JobName

	tmp.Status = view.Status

	data = &tmp
	return
}

func (sysJobLogViewUtils *SysJobLogViewUtils) View2DataList(viewList *[]SysJobLogView) (err error, dataList *[]model.SysJobLog) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysJobLogViewUtils View2DataList error: %v", e)
			global.Logger.Error("SysJobLogViewUtils.View2DataList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if viewList != nil {
		var dataTmpList []model.SysJobLog
		for i := range *dataList {
			view := (*viewList)[i]
			err, data := sysJobLogViewUtils.View2Data(&view)
			if err == nil {
				dataTmpList = append(dataTmpList, *data)
			}
		}
		dataList = &dataTmpList
	}
	return
}

func (sysJobLogViewUtils *SysJobLogViewUtils) Data2ViewList(dataList *[]model.SysJobLog) (err error, viewList *[]SysJobLogView) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysJobLogViewUtils Data2ViewList error: %v", e)
			global.Logger.Error("SysJobLogViewUtils.Data2ViewList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if dataList != nil {
		var viewTmpList []SysJobLogView
		for i := range *dataList {
			data := (*dataList)[i]
			err, view := sysJobLogViewUtils.Data2View(&data)
			if err == nil {
				viewTmpList = append(viewTmpList, *view)
			}
		}
		viewList = &viewTmpList
	}
	return
}
