// Package view
// @description <TODO description class purpose>
// @author
// @File: sys_user_post
// @version 1.0.0
// @create 2023-08-18 13:41:26
package view

import (
	"fmt"
	"go.uber.org/zap"
	"manager-gin/src/app/admin/sys/sys_user_post/model"
	"manager-gin/src/global"
	"manager-gin/src/utils"
)

type SysUserPostViewUtils struct{}

func (sysUserPostViewUtils *SysUserPostViewUtils) Data2View(data *model.SysUserPost) (err error, view *SysUserPostView) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysUserPostViewUtils View2Data error: %v", e)
			global.Logger.Error("SysUserPostViewUtils.Data2View:格式转换异常",
				zap.Any("error", e))
		}
	}()
	var tmp SysUserPostView

	tmp.UserId = data.UserId

	tmp.PostId = data.PostId

	view = &tmp
	return
}
func (sysUserPostViewUtils *SysUserPostViewUtils) View2Data(view *SysUserPostView) (err error, data *model.SysUserPost) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysUserPostViewUtils View2Data error: %v", e)
			global.Logger.Error("SysUserPostViewUtils.View2Data:格式转换异常",
				zap.Any("error", e))
		}
	}()
	var tmp model.SysUserPost

	tmp.UserId = view.UserId

	tmp.PostId = view.PostId

	data = &tmp
	return
}

func (sysUserPostViewUtils *SysUserPostViewUtils) View2DataList(viewList *[]SysUserPostView) (err error, dataList *[]model.SysUserPost) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysUserPostViewUtils View2DataList error: %v", e)
			global.Logger.Error("SysUserPostViewUtils.View2DataList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if viewList != nil {
		var dataTmpList []model.SysUserPost
		for i := range *dataList {
			view := (*viewList)[i]
			err, data := sysUserPostViewUtils.View2Data(&view)
			if err == nil {
				dataTmpList = append(dataTmpList, *data)
			}
		}
		dataList = &dataTmpList
	}
	return
}

func (sysUserPostViewUtils *SysUserPostViewUtils) Data2ViewList(dataList *[]model.SysUserPost) (err error, viewList *[]SysUserPostView) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysUserPostViewUtils Data2ViewList error: %v", e)
			global.Logger.Error("SysUserPostViewUtils.Data2ViewList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if dataList != nil {
		var viewTmpList []SysUserPostView
		for i := range *dataList {
			data := (*dataList)[i]
			err, view := sysUserPostViewUtils.Data2View(&data)
			if err == nil {
				viewTmpList = append(viewTmpList, *view)
			}
		}
		viewList = &viewTmpList
	}
	return
}
