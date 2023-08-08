// Package view
// @description <TODO description class purpose>
// @author
// @File: sys_user
// @version 1.0.0
// @create 2023-08-08 10:06:19
package view

import (
	"fmt"
	"go.uber.org/zap"
	"manager-gin/src/app/admin/sys/sys_user/model"
	"manager-gin/src/global"
	"manager-gin/src/utils"
)

type SysUserViewUtils struct{}

func (sysUserViewUtils *SysUserViewUtils) Data2View(data *model.SysUser) (err error, view *SysUserView) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysUserViewUtils View2Data error: %v", e)
			global.Logger.Error("SysUserViewUtils.Data2View:格式转换异常",
				zap.Any("error", e))
		}
	}()
	var tmp SysUserView

	tmp.Avatar = data.Avatar

	tmp.CreateBy = data.CreateBy

	tmp.CreateTime = utils.Time2Str(data.CreateTime)

	tmp.DelFlag = data.DelFlag

	tmp.DeptId = data.DeptId

	tmp.Email = data.Email

	tmp.Id = data.Id

	tmp.LoginDate = utils.Time2Str(data.LoginDate)

	tmp.LoginIp = data.LoginIp

	tmp.NickName = data.NickName

	tmp.Password = data.Password

	tmp.Phonenumber = data.Phonenumber

	tmp.Remark = data.Remark

	tmp.Sex = data.Sex

	tmp.Status = data.Status

	tmp.UpdateBy = data.UpdateBy

	tmp.UpdateTime = utils.Time2Str(data.UpdateTime)

	tmp.UserName = data.UserName

	tmp.UserType = data.UserType

	view = &tmp
	return
}
func (sysUserViewUtils *SysUserViewUtils) View2Data(view *SysUserView) (err error, data *model.SysUser) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysUserViewUtils View2Data error: %v", e)
			global.Logger.Error("SysUserViewUtils.View2Data:格式转换异常",
				zap.Any("error", e))
		}
	}()
	var tmp model.SysUser

	tmp.Avatar = view.Avatar

	tmp.CreateBy = view.CreateBy

	tmp.CreateTime = utils.Str2Time(view.CreateTime)

	tmp.DelFlag = view.DelFlag

	tmp.DeptId = view.DeptId

	tmp.Email = view.Email

	tmp.Id = view.Id

	tmp.LoginDate = utils.Str2Time(view.LoginDate)

	tmp.LoginIp = view.LoginIp

	tmp.NickName = view.NickName

	tmp.Password = view.Password

	tmp.Phonenumber = view.Phonenumber

	tmp.Remark = view.Remark

	tmp.Sex = view.Sex

	tmp.Status = view.Status

	tmp.UpdateBy = view.UpdateBy

	tmp.UpdateTime = utils.Str2Time(view.UpdateTime)

	tmp.UserName = view.UserName

	tmp.UserType = view.UserType

	data = &tmp
	return
}

func (sysUserViewUtils *SysUserViewUtils) View2DataList(viewList *[]SysUserView) (err error, dataList *[]model.SysUser) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysUserViewUtils View2DataList error: %v", e)
			global.Logger.Error("SysUserViewUtils.View2DataList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if viewList != nil {
		var dataTmpList []model.SysUser
		for i := range *dataList {
			view := (*viewList)[i]
			err, data := sysUserViewUtils.View2Data(&view)
			if err == nil {
				dataTmpList = append(dataTmpList, *data)
			}
		}
		dataList = &dataTmpList
	}
	return
}

func (sysUserViewUtils *SysUserViewUtils) Data2ViewList(dataList *[]model.SysUser) (err error, viewList *[]SysUserView) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysUserViewUtils Data2ViewList error: %v", e)
			global.Logger.Error("SysUserViewUtils.Data2ViewList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if dataList != nil {
		var viewTmpList []SysUserView
		for i := range *dataList {
			data := (*dataList)[i]
			err, view := sysUserViewUtils.Data2View(&data)
			if err == nil {
				viewTmpList = append(viewTmpList, *view)
			}
		}
		viewList = &viewTmpList
	}
	return
}
