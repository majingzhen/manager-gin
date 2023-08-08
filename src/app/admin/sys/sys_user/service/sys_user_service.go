// Package service 自动生成模板 SysUserService
// @description <TODO description class purpose>
// @author
// @File: sys_user
// @version 1.0.0
// @create 2023-08-08 10:06:19
package service

import (
	"manager-gin/src/app/admin/sys/sys_user/model"
	"manager-gin/src/app/admin/sys/sys_user/service/view"
	"manager-gin/src/common"
)

var sysUserDao = model.SysUserDaoApp
var viewUtils = view.SysUserViewUtilsApp

type SysUserService struct{}

// Create 创建SysUser记录
// Author
func (sysUserService *SysUserService) Create(sysUserView *view.SysUserView) (err error) {
	err1, sysUser := viewUtils.View2Data(sysUserView)
	if err1 != nil {
		return err1
	}
	err2 := sysUserDao.Create(*sysUser)
	if err2 != nil {
		return err2
	}
	return nil
}

// Delete 删除SysUser记录
// Author
func (sysUserService *SysUserService) Delete(id int) (err error) {
	err = sysUserDao.Delete(id)
	return err
}

// DeleteByIds 批量删除SysUser记录
// Author
func (sysUserService *SysUserService) DeleteByIds(ids []int) (err error) {
	err = sysUserDao.DeleteByIds(ids)
	return err
}

// Update 更新SysUser记录
// Author
func (sysUserService *SysUserService) Update(id int, sysUserView *view.SysUserView) (err error) {
	sysUserView.Id = id
	err1, sysUser := viewUtils.View2Data(sysUserView)
	if err1 != nil {
		return err1
	}
	err = sysUserDao.Update(*sysUser)
	return err
}

// Get 根据id获取SysUser记录
// Author
func (sysUserService *SysUserService) Get(id int) (err error, sysUserView *view.SysUserView) {
	err1, sysUser := sysUserDao.Get(id)
	if err1 != nil {
		return err1, nil
	}
	err2, sysUserView := viewUtils.Data2View(sysUser)
	if err2 != nil {
		return err2, nil
	}
	return
}

// Find 分页获取SysUser记录
// Author
func (sysUserService *SysUserService) Find(info *common.PageInfoV2) (err error) {
	err1, sysUsers, total := sysUserDao.Find(info)
	if err1 != nil {
		return err1
	}
	info.Total = total
	err2, viewList := viewUtils.Data2ViewList(sysUsers)
	if err2 != nil {
		return err2
	}
	info.FormList = viewList
	return err
}
