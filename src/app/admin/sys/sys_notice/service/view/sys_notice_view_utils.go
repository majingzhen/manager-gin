// Package view
// @description <TODO description class purpose>
// @author
// @File: sys_notice
// @version 1.0.0
// @create 2023-08-08 10:06:19
package view

import (
	"fmt"
	"go.uber.org/zap"
	"manager-gin/src/app/admin/sys/sys_notice/model"
	"manager-gin/src/global"
	"manager-gin/src/utils"
)

type SysNoticeViewUtils struct{}

func (sysNoticeViewUtils *SysNoticeViewUtils) Data2View(data *model.SysNotice) (err error, view *SysNoticeView) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysNoticeViewUtils View2Data error: %v", e)
			global.Logger.Error("SysNoticeViewUtils.Data2View:格式转换异常",
				zap.Any("error", e))
		}
	}()
	var tmp SysNoticeView

	tmp.CreateBy = data.CreateBy

	tmp.CreateTime = utils.Time2Str(data.CreateTime)

	tmp.Id = data.Id

	tmp.NoticeContent = data.NoticeContent

	tmp.NoticeTitle = data.NoticeTitle

	tmp.NoticeType = data.NoticeType

	tmp.Remark = data.Remark

	tmp.Status = data.Status

	tmp.UpdateBy = data.UpdateBy

	tmp.UpdateTime = utils.Time2Str(data.UpdateTime)

	view = &tmp
	return
}
func (sysNoticeViewUtils *SysNoticeViewUtils) View2Data(view *SysNoticeView) (err error, data *model.SysNotice) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysNoticeViewUtils View2Data error: %v", e)
			global.Logger.Error("SysNoticeViewUtils.View2Data:格式转换异常",
				zap.Any("error", e))
		}
	}()
	var tmp model.SysNotice

	tmp.CreateBy = view.CreateBy

	tmp.CreateTime = utils.Str2Time(view.CreateTime)

	tmp.Id = view.Id

	tmp.NoticeContent = view.NoticeContent

	tmp.NoticeTitle = view.NoticeTitle

	tmp.NoticeType = view.NoticeType

	tmp.Remark = view.Remark

	tmp.Status = view.Status

	tmp.UpdateBy = view.UpdateBy

	tmp.UpdateTime = utils.Str2Time(view.UpdateTime)

	data = &tmp
	return
}

func (sysNoticeViewUtils *SysNoticeViewUtils) View2DataList(viewList *[]SysNoticeView) (err error, dataList *[]model.SysNotice) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysNoticeViewUtils View2DataList error: %v", e)
			global.Logger.Error("SysNoticeViewUtils.View2DataList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if viewList != nil {
		var dataTmpList []model.SysNotice
		for i := range *dataList {
			view := (*viewList)[i]
			err, data := sysNoticeViewUtils.View2Data(&view)
			if err == nil {
				dataTmpList = append(dataTmpList, *data)
			}
		}
		dataList = &dataTmpList
	}
	return
}

func (sysNoticeViewUtils *SysNoticeViewUtils) Data2ViewList(dataList *[]model.SysNotice) (err error, viewList *[]SysNoticeView) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysNoticeViewUtils Data2ViewList error: %v", e)
			global.Logger.Error("SysNoticeViewUtils.Data2ViewList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if dataList != nil {
		var viewTmpList []SysNoticeView
		for i := range *dataList {
			data := (*dataList)[i]
			err, view := sysNoticeViewUtils.Data2View(&data)
			if err == nil {
				viewTmpList = append(viewTmpList, *view)
			}
		}
		viewList = &viewTmpList
	}
	return
}
