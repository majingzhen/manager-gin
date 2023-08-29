// Package view
// @description <TODO description class purpose>
// @author
// @File: sys_dict_data
// @version 1.0.0
// @create 2023-08-20 19:08:42
package view

import (
	"fmt"
	"go.uber.org/zap"
	"manager-gin/src/app/admin/sys/model"
	"manager-gin/src/common"
	"manager-gin/src/global"
	"manager-gin/src/utils"
)

type SysDictDataViewUtils struct{}

func (viewUtils *SysDictDataViewUtils) Data2View(data *model.SysDictData) (err error, view *SysDictDataView) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysDictDataViewUtils View2Data error: %v", e)
			global.Logger.Error("SysDictDataViewUtils.Data2View:格式转换异常",
				zap.Any("error", e))
		}
	}()
	var tmp SysDictDataView

	tmp.Id = data.Id

	tmp.DictSort = data.DictSort

	tmp.DictLabel = data.DictLabel

	tmp.DictValue = data.DictValue

	tmp.DictType = data.DictType

	tmp.CssClass = data.CssClass

	tmp.ListClass = data.ListClass

	tmp.IsDefault = data.IsDefault

	tmp.Status = data.Status

	tmp.CreateBy = data.CreateBy

	tmp.CreateTime = utils.Time2Str(data.CreateTime)

	tmp.UpdateBy = data.UpdateBy

	tmp.UpdateTime = utils.Time2Str(data.UpdateTime)

	tmp.Remark = data.Remark

	view = &tmp
	return
}
func (viewUtils *SysDictDataViewUtils) View2Data(view *SysDictDataView) (err error, data *model.SysDictData) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysDictDataViewUtils View2Data error: %v", e)
			global.Logger.Error("SysDictDataViewUtils.View2Data:格式转换异常",
				zap.Any("error", e))
		}
	}()
	var tmp model.SysDictData

	tmp.Id = view.Id

	tmp.DictSort = view.DictSort

	tmp.DictLabel = view.DictLabel

	tmp.DictValue = view.DictValue

	tmp.DictType = view.DictType

	tmp.CssClass = view.CssClass

	tmp.ListClass = view.ListClass

	tmp.IsDefault = view.IsDefault

	tmp.Status = view.Status

	tmp.CreateBy = view.CreateBy

	tmp.CreateTime = utils.Str2Time(view.CreateTime)

	tmp.UpdateBy = view.UpdateBy

	tmp.UpdateTime = utils.Str2Time(view.UpdateTime)

	tmp.Remark = view.Remark

	data = &tmp
	return
}

func (viewUtils *SysDictDataViewUtils) PageData2ViewList(pageInfo *common.PageInfo) (err error, res *common.PageInfo) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysDictDataViewUtils PageData2ViewList error: %v", e)
			global.Logger.Error("SysDictDataViewUtils.PageData2ViewList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if pageInfo != nil && pageInfo.Rows != nil {
		if p, ok := pageInfo.Rows.([]*model.SysDictData); ok {
			if err, viewList := viewUtils.Data2ViewList(p); err == nil {
				pageInfo.Rows = viewList
			}
		}
	}
	res = pageInfo
	return
}

func (viewUtils *SysDictDataViewUtils) Page2Data(pageInfo *SysDictDataPageView) (err error, data *model.SysDictData, page *common.PageInfo) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysDictDataViewUtils View2Data error: %v", e)
			global.Logger.Error("SysDictDataViewUtils.View2Data:格式转换异常",
				zap.Any("error", e))
		}
	}()
	// TODO 按需修改
	var tmp model.SysDictData

	tmp.Id = pageInfo.Id

	tmp.DictSort = pageInfo.DictSort

	tmp.DictLabel = pageInfo.DictLabel

	tmp.DictValue = pageInfo.DictValue

	tmp.DictType = pageInfo.DictType

	tmp.CssClass = pageInfo.CssClass

	tmp.ListClass = pageInfo.ListClass

	tmp.IsDefault = pageInfo.IsDefault

	tmp.Status = pageInfo.Status

	tmp.CreateBy = pageInfo.CreateBy

	tmp.UpdateBy = pageInfo.UpdateBy

	tmp.Remark = pageInfo.Remark

	data = &tmp
	page = &common.PageInfo{
		PageSize:      pageInfo.PageSize,
		PageNum:       pageInfo.PageNum,
		OrderByColumn: pageInfo.OrderByColumn,
		IsAsc:         pageInfo.IsAsc,
	}
	return
}

func (viewUtils *SysDictDataViewUtils) View2DataList(viewList []*SysDictDataView) (err error, dataList []*model.SysDictData) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysDictDataViewUtils View2DataList error: %v", e)
			global.Logger.Error("SysDictDataViewUtils.View2DataList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if viewList != nil {
		var dataTmpList []*model.SysDictData
		for i := range viewList {
			view := viewList[i]
			err, data := viewUtils.View2Data(view)
			if err == nil {
				dataTmpList = append(dataTmpList, data)
			}
		}
		dataList = dataTmpList
	}
	return
}

func (viewUtils *SysDictDataViewUtils) Data2ViewList(dataList []*model.SysDictData) (err error, viewList []*SysDictDataView) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysDictDataViewUtils Data2ViewList error: %v", e)
			global.Logger.Error("SysDictDataViewUtils.Data2ViewList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if dataList != nil {
		var viewTmpList []*SysDictDataView
		for i := range dataList {
			data := (dataList)[i]
			err, view := viewUtils.Data2View(data)
			if err == nil {
				viewTmpList = append(viewTmpList, view)
			}
		}
		viewList = viewTmpList
	}
	return
}