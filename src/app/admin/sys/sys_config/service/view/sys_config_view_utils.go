// Package view
// @description <TODO description class purpose>
// @author
// @File: sys_config
// @version 1.0.0
// @create 2023-08-18 13:41:26
package view

import (
	"fmt"
	"go.uber.org/zap"
	"manager-gin/src/app/admin/sys/sys_config/model"
	"manager-gin/src/global"
	"manager-gin/src/utils"
)

type SysConfigViewUtils struct{}

func (sysConfigViewUtils *SysConfigViewUtils) Data2View(data *model.SysConfig) (err error, view *SysConfigView) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysConfigViewUtils View2Data error: %v", e)
			global.Logger.Error("SysConfigViewUtils.Data2View:格式转换异常",
				zap.Any("error", e))
		}
	}()
	var tmp SysConfigView

	tmp.Id = data.Id

	tmp.ConfigName = data.ConfigName

	tmp.ConfigKey = data.ConfigKey

	tmp.ConfigValue = data.ConfigValue

	tmp.ConfigType = data.ConfigType

	tmp.CreateBy = data.CreateBy

	tmp.CreateTime = utils.Time2Str(data.CreateTime)

	tmp.UpdateBy = data.UpdateBy

	tmp.UpdateTime = utils.Time2Str(data.UpdateTime)

	tmp.Remark = data.Remark

	view = &tmp
	return
}
func (sysConfigViewUtils *SysConfigViewUtils) View2Data(view *SysConfigView) (err error, data *model.SysConfig) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysConfigViewUtils View2Data error: %v", e)
			global.Logger.Error("SysConfigViewUtils.View2Data:格式转换异常",
				zap.Any("error", e))
		}
	}()
	var tmp model.SysConfig

	tmp.Id = view.Id

	tmp.ConfigName = view.ConfigName

	tmp.ConfigKey = view.ConfigKey

	tmp.ConfigValue = view.ConfigValue

	tmp.ConfigType = view.ConfigType

	tmp.CreateBy = view.CreateBy

	tmp.CreateTime = utils.Str2Time(view.CreateTime)

	tmp.UpdateBy = view.UpdateBy

	tmp.UpdateTime = utils.Str2Time(view.UpdateTime)

	tmp.Remark = view.Remark

	data = &tmp
	return
}

func (sysConfigViewUtils *SysConfigViewUtils) View2DataList(viewList *[]SysConfigView) (err error, dataList *[]model.SysConfig) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysConfigViewUtils View2DataList error: %v", e)
			global.Logger.Error("SysConfigViewUtils.View2DataList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if viewList != nil {
		var dataTmpList []model.SysConfig
		for i := range *dataList {
			view := (*viewList)[i]
			err, data := sysConfigViewUtils.View2Data(&view)
			if err == nil {
				dataTmpList = append(dataTmpList, *data)
			}
		}
		dataList = &dataTmpList
	}
	return
}

func (sysConfigViewUtils *SysConfigViewUtils) Data2ViewList(dataList *[]model.SysConfig) (err error, viewList *[]SysConfigView) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysConfigViewUtils Data2ViewList error: %v", e)
			global.Logger.Error("SysConfigViewUtils.Data2ViewList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if dataList != nil {
		var viewTmpList []SysConfigView
		for i := range *dataList {
			data := (*dataList)[i]
			err, view := sysConfigViewUtils.Data2View(&data)
			if err == nil {
				viewTmpList = append(viewTmpList, *view)
			}
		}
		viewList = &viewTmpList
	}
	return
}
