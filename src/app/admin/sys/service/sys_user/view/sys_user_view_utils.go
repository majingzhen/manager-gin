// Package view
// @description <TODO description class purpose>
// @author
// @File: sys_user
// @version 1.0.0
// @create 2023-08-21 14:20:37
package view

import (
	"fmt"
	"go.uber.org/zap"
	"manager-gin/src/app/admin/sys/model"
	"manager-gin/src/common"
	"manager-gin/src/global"
	"manager-gin/src/utils"
)

type SysUserViewUtils struct{}

func (viewUtils *SysUserViewUtils) Data2View(data *model.SysUser) (err error, view *SysUserView) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysUserViewUtils View2Data error: %v", e)
			global.Logger.Error("SysUserViewUtils.Data2View:格式转换异常",
				zap.Any("error", e))
		}
	}()
	var tmp SysUserView

	tmp.Id = data.Id

	tmp.DeptId = data.DeptId

	tmp.UserName = data.UserName

	tmp.NickName = data.NickName

	tmp.UserType = data.UserType

	tmp.Email = data.Email

	tmp.PhoneNumber = data.PhoneNumber

	tmp.Sex = data.Sex

	tmp.Avatar = data.Avatar

	tmp.Password = data.Password

	tmp.Salt = data.Salt

	tmp.Status = data.Status

	tmp.LoginIp = data.LoginIp

	tmp.LoginDate = utils.Time2Str(data.LoginDate)

	tmp.CreateBy = data.CreateBy

	tmp.CreateTime = utils.Time2Str(data.CreateTime)

	tmp.UpdateBy = data.UpdateBy

	tmp.UpdateTime = utils.Time2Str(data.UpdateTime)

	tmp.Remark = data.Remark

	view = &tmp
	return
}
func (viewUtils *SysUserViewUtils) View2Data(view *SysUserView) (err error, data *model.SysUser) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysUserViewUtils View2Data error: %v", e)
			global.Logger.Error("SysUserViewUtils.View2Data:格式转换异常",
				zap.Any("error", e))
		}
	}()
	var tmp model.SysUser

	tmp.Id = view.Id

	tmp.DeptId = view.DeptId

	tmp.UserName = view.UserName

	tmp.NickName = view.NickName

	tmp.UserType = view.UserType

	tmp.Email = view.Email

	tmp.PhoneNumber = view.PhoneNumber

	tmp.Sex = view.Sex

	tmp.Avatar = view.Avatar

	tmp.Password = view.Password

	tmp.Salt = view.Salt

	tmp.Status = view.Status

	tmp.LoginIp = view.LoginIp

	tmp.LoginDate = utils.Str2Time(view.LoginDate)

	tmp.CreateBy = view.CreateBy

	tmp.CreateTime = utils.Str2Time(view.CreateTime)

	tmp.UpdateBy = view.UpdateBy

	tmp.UpdateTime = utils.Str2Time(view.UpdateTime)

	tmp.Remark = view.Remark

	data = &tmp
	return
}

func (viewUtils *SysUserViewUtils) Page2Data(pageInfo *SysUserPageView) (err error, data *model.SysUser, page *common.PageInfo) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysUserViewUtils View2Data error: %v", e)
			global.Logger.Error("SysUserViewUtils.View2Data:格式转换异常",
				zap.Any("error", e))
		}
	}()
	// TODO 按需修改
	var tmp model.SysUser

	tmp.Id = pageInfo.Id

	tmp.DeptId = pageInfo.DeptId

	tmp.UserName = pageInfo.UserName

	tmp.NickName = pageInfo.NickName

	tmp.UserType = pageInfo.UserType

	tmp.Email = pageInfo.Email

	tmp.PhoneNumber = pageInfo.PhoneNumber

	tmp.Sex = pageInfo.Sex

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

func (viewUtils *SysUserViewUtils) View2DataList(viewList []*SysUserView) (err error, dataList []*model.SysUser) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysUserViewUtils View2DataList error: %v", e)
			global.Logger.Error("SysUserViewUtils.View2DataList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if viewList != nil {
		var dataTmpList []*model.SysUser
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

func (viewUtils *SysUserViewUtils) Data2ViewList(dataList []*model.SysUser) (err error, viewList []*SysUserView) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysUserViewUtils Data2ViewList error: %v", e)
			global.Logger.Error("SysUserViewUtils.Data2ViewList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if dataList != nil {
		var viewTmpList []*SysUserView
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
