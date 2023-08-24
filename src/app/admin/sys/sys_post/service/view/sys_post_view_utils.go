// Package view
// @description <TODO description class purpose>
// @author
// @File: sys_post
// @version 1.0.0
// @create 2023-08-21 17:37:56
package view

import (
	"fmt"
	"go.uber.org/zap"
	"manager-gin/src/app/admin/sys/sys_post/model"
	"manager-gin/src/common"
	"manager-gin/src/global"
	"manager-gin/src/utils"
)

type SysPostViewUtils struct{}

func (viewUtils *SysPostViewUtils) Data2View(data *model.SysPost) (err error, view *SysPostView) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysPostViewUtils View2Data error: %v", e)
			global.Logger.Error("SysPostViewUtils.Data2View:格式转换异常",
				zap.Any("error", e))
		}
	}()
	var tmp SysPostView

	tmp.Id = data.Id

	tmp.PostCode = data.PostCode

	tmp.PostName = data.PostName

	tmp.PostSort = data.PostSort

	tmp.Status = data.Status

	tmp.CreateBy = data.CreateBy

	tmp.CreateTime = utils.Time2Str(data.CreateTime)

	tmp.UpdateBy = data.UpdateBy

	tmp.UpdateTime = utils.Time2Str(data.UpdateTime)

	tmp.Remark = data.Remark

	view = &tmp
	return
}
func (viewUtils *SysPostViewUtils) View2Data(view *SysPostView) (err error, data *model.SysPost) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysPostViewUtils View2Data error: %v", e)
			global.Logger.Error("SysPostViewUtils.View2Data:格式转换异常",
				zap.Any("error", e))
		}
	}()
	var tmp model.SysPost

	tmp.Id = view.Id

	tmp.PostCode = view.PostCode

	tmp.PostName = view.PostName

	tmp.PostSort = view.PostSort

	tmp.Status = view.Status

	tmp.CreateBy = view.CreateBy

	tmp.CreateTime = utils.Str2Time(view.CreateTime)

	tmp.UpdateBy = view.UpdateBy

	tmp.UpdateTime = utils.Str2Time(view.UpdateTime)

	tmp.Remark = view.Remark

	data = &tmp
	return
}

func (viewUtils *SysPostViewUtils) Page2Data(pageInfo *SysPostPageView) (err error, data *model.SysPost, page *common.PageInfo) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysPostViewUtils View2Data error: %v", e)
			global.Logger.Error("SysPostViewUtils.View2Data:格式转换异常",
				zap.Any("error", e))
		}
	}()
	// TODO 按需修改
	var tmp model.SysPost

	tmp.Id = pageInfo.Id

	tmp.PostCode = pageInfo.PostCode

	tmp.PostName = pageInfo.PostName

	tmp.PostSort = pageInfo.PostSort

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

func (viewUtils *SysPostViewUtils) View2DataList(viewList []*SysPostView) (err error, dataList []*model.SysPost) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysPostViewUtils View2DataList error: %v", e)
			global.Logger.Error("SysPostViewUtils.View2DataList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if viewList != nil {
		var dataTmpList []*model.SysPost
		for i := range viewList {
			view := (viewList)[i]
			err, data := viewUtils.View2Data(view)
			if err == nil {
				dataTmpList = append(dataTmpList, data)
			}
		}
		dataList = dataTmpList
	}
	return
}

func (viewUtils *SysPostViewUtils) Data2ViewList(dataList []*model.SysPost) (err error, viewList []*SysPostView) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysPostViewUtils Data2ViewList error: %v", e)
			global.Logger.Error("SysPostViewUtils.Data2ViewList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if dataList != nil {
		var viewTmpList []*SysPostView
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
