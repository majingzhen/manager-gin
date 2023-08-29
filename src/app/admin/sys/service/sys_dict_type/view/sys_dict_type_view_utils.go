// Package view
// @description <TODO description class purpose>
// @author
// @File: sys_dict_type
// @version 1.0.0
// @create 2023-08-18 13:41:26
package view

import (
	"fmt"
	"go.uber.org/zap"
	"manager-gin/src/app/admin/sys/model"
	"manager-gin/src/common"
	"manager-gin/src/global"
	"manager-gin/src/utils"
)

type SysDictTypeViewUtils struct{}

func (viewUtils *SysDictTypeViewUtils) Data2View(data *model.SysDictType) (err error, view *SysDictTypeView) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysDictTypeViewUtils View2Data error: %v", e)
			global.Logger.Error("SysDictTypeViewUtils.Data2View:格式转换异常",
				zap.Any("error", e))
		}
	}()
	var tmp SysDictTypeView
	tmp.Id = data.Id
	tmp.DictName = data.DictName
	tmp.DictType = data.DictType
	tmp.Status = data.Status
	tmp.CreateBy = data.CreateBy
	tmp.CreateTime = utils.Time2Str(data.CreateTime)
	tmp.UpdateBy = data.UpdateBy
	tmp.UpdateTime = utils.Time2Str(data.UpdateTime)
	tmp.Remark = data.Remark
	view = &tmp
	return
}

func (viewUtils *SysDictTypeViewUtils) Page2Data(pageInfo *SysDictTypePageView) (err error, data *model.SysDictType, page *common.PageInfo) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysDictTypeViewUtils View2Data error: %v", e)
			global.Logger.Error("SysDictTypeViewUtils.Data2View:格式转换异常",
				zap.Any("error", e))
		}
	}()

	var tmp model.SysDictType
	tmp.DictName = pageInfo.DictName
	tmp.DictType = pageInfo.DictType
	tmp.Status = pageInfo.Status
	data = &tmp
	page = &common.PageInfo{
		PageSize:      pageInfo.PageSize,
		PageNum:       pageInfo.PageNum,
		OrderByColumn: pageInfo.OrderByColumn,
		IsAsc:         pageInfo.IsAsc,
	}
	return
}

func (viewUtils *SysDictTypeViewUtils) View2Data(view *SysDictTypeView) (err error, data *model.SysDictType) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysDictTypeViewUtils View2Data error: %v", e)
			global.Logger.Error("SysDictTypeViewUtils.View2Data:格式转换异常",
				zap.Any("error", e))
		}
	}()
	var tmp model.SysDictType
	tmp.Id = view.Id
	tmp.DictName = view.DictName
	tmp.DictType = view.DictType
	tmp.Status = view.Status
	tmp.CreateBy = view.CreateBy
	tmp.CreateTime = utils.Str2Time(view.CreateTime)
	tmp.UpdateBy = view.UpdateBy
	tmp.UpdateTime = utils.Str2Time(view.UpdateTime)
	tmp.Remark = view.Remark
	data = &tmp
	return
}

func (viewUtils *SysDictTypeViewUtils) PageData2ViewList(pageInfo *common.PageInfo) (err error, res *common.PageInfo) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysDictTypeViewUtils PageData2ViewList error: %v", e)
			global.Logger.Error("SysDictTypeViewUtils.PageData2ViewList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if pageInfo != nil && pageInfo.Rows != nil {
		if p, ok := pageInfo.Rows.([]*model.SysDictType); ok {
			if err, viewList := viewUtils.Data2ViewList(p); err == nil {
				pageInfo.Rows = viewList
			}
		}
	}
	res = pageInfo
	return
}

func (viewUtils *SysDictTypeViewUtils) View2DataList(viewList []*SysDictTypeView) (err error, dataList []*model.SysDictType) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysDictTypeViewUtils View2DataList error: %v", e)
			global.Logger.Error("SysDictTypeViewUtils.View2DataList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if viewList != nil {
		var dataTmpList []*model.SysDictType
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

func (viewUtils *SysDictTypeViewUtils) Data2ViewList(dataList []*model.SysDictType) (err error, viewList []*SysDictTypeView) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysDictTypeViewUtils Data2ViewList error: %v", e)
			global.Logger.Error("SysDictTypeViewUtils.Data2ViewList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if dataList != nil {
		var viewTmpList []*SysDictTypeView
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
