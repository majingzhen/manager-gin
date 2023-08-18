// Package view
// @description <TODO description class purpose>
// @author
// @File: sys_job
// @version 1.0.0
// @create 2023-08-18 13:41:26
package view

import (
	"fmt"
	"go.uber.org/zap"
	"manager-gin/src/app/admin/sys/sys_job/model"
	"manager-gin/src/global"
	"manager-gin/src/utils"
)

type SysJobViewUtils struct{}

func (sysJobViewUtils *SysJobViewUtils) Data2View(data *model.SysJob) (err error, view *SysJobView) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysJobViewUtils View2Data error: %v", e)
			global.Logger.Error("SysJobViewUtils.Data2View:格式转换异常",
				zap.Any("error", e))
		}
	}()
	var tmp SysJobView

	tmp.Id = data.Id

	tmp.JobName = data.JobName

	tmp.JobGroup = data.JobGroup

	tmp.InvokeTarget = data.InvokeTarget

	tmp.CronExpression = data.CronExpression

	tmp.MisfirePolicy = data.MisfirePolicy

	tmp.Concurrent = data.Concurrent

	tmp.Status = data.Status

	tmp.CreateBy = data.CreateBy

	tmp.CreateTime = utils.Time2Str(data.CreateTime)

	tmp.UpdateBy = data.UpdateBy

	tmp.UpdateTime = utils.Time2Str(data.UpdateTime)

	tmp.Remark = data.Remark

	view = &tmp
	return
}
func (sysJobViewUtils *SysJobViewUtils) View2Data(view *SysJobView) (err error, data *model.SysJob) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysJobViewUtils View2Data error: %v", e)
			global.Logger.Error("SysJobViewUtils.View2Data:格式转换异常",
				zap.Any("error", e))
		}
	}()
	var tmp model.SysJob

	tmp.Id = view.Id

	tmp.JobName = view.JobName

	tmp.JobGroup = view.JobGroup

	tmp.InvokeTarget = view.InvokeTarget

	tmp.CronExpression = view.CronExpression

	tmp.MisfirePolicy = view.MisfirePolicy

	tmp.Concurrent = view.Concurrent

	tmp.Status = view.Status

	tmp.CreateBy = view.CreateBy

	tmp.CreateTime = utils.Str2Time(view.CreateTime)

	tmp.UpdateBy = view.UpdateBy

	tmp.UpdateTime = utils.Str2Time(view.UpdateTime)

	tmp.Remark = view.Remark

	data = &tmp
	return
}

func (sysJobViewUtils *SysJobViewUtils) View2DataList(viewList *[]SysJobView) (err error, dataList *[]model.SysJob) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysJobViewUtils View2DataList error: %v", e)
			global.Logger.Error("SysJobViewUtils.View2DataList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if viewList != nil {
		var dataTmpList []model.SysJob
		for i := range *dataList {
			view := (*viewList)[i]
			err, data := sysJobViewUtils.View2Data(&view)
			if err == nil {
				dataTmpList = append(dataTmpList, *data)
			}
		}
		dataList = &dataTmpList
	}
	return
}

func (sysJobViewUtils *SysJobViewUtils) Data2ViewList(dataList *[]model.SysJob) (err error, viewList *[]SysJobView) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysJobViewUtils Data2ViewList error: %v", e)
			global.Logger.Error("SysJobViewUtils.Data2ViewList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if dataList != nil {
		var viewTmpList []SysJobView
		for i := range *dataList {
			data := (*dataList)[i]
			err, view := sysJobViewUtils.Data2View(&data)
			if err == nil {
				viewTmpList = append(viewTmpList, *view)
			}
		}
		viewList = &viewTmpList
	}
	return
}
