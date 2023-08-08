// Package service 自动生成模板 SysUserRoleService
// @description <TODO description class purpose>
// @author
// @File: sys_user_role
// @version 1.0.0
// @create 2023-08-08 10:06:19
package service

import (
	"manager-gin/src/app/admin/sys/sys_user_role/model"
	"manager-gin/src/app/admin/sys/sys_user_role/service/view"
	"manager-gin/src/common"
)

var sysUserRoleDao = model.SysUserRoleDaoApp
var viewUtils = view.SysUserRoleViewUtilsApp

type SysUserRoleService struct{}

// Create 创建SysUserRole记录
// Author
func (sysUserRoleService *SysUserRoleService) Create(sysUserRoleView *view.SysUserRoleView) (err error) {
	err1, sysUserRole := viewUtils.View2Data(sysUserRoleView)
	if err1 != nil {
		return err1
	}
	err2 := sysUserRoleDao.Create(*sysUserRole)
	if err2 != nil {
		return err2
	}
	return nil
}

// Delete 删除SysUserRole记录
// Author
func (sysUserRoleService *SysUserRoleService) Delete(id int) (err error) {
	err = sysUserRoleDao.Delete(id)
	return err
}

// DeleteByIds 批量删除SysUserRole记录
// Author
func (sysUserRoleService *SysUserRoleService) DeleteByIds(ids []int) (err error) {
	err = sysUserRoleDao.DeleteByIds(ids)
	return err
}

// Update 更新SysUserRole记录
// Author
func (sysUserRoleService *SysUserRoleService) Update(id int, sysUserRoleView *view.SysUserRoleView) (err error) {
	sysUserRoleView.Id = id
	err1, sysUserRole := viewUtils.View2Data(sysUserRoleView)
	if err1 != nil {
		return err1
	}
	err = sysUserRoleDao.Update(*sysUserRole)
	return err
}

// Get 根据id获取SysUserRole记录
// Author
func (sysUserRoleService *SysUserRoleService) Get(id int) (err error, sysUserRoleView *view.SysUserRoleView) {
	err1, sysUserRole := sysUserRoleDao.Get(id)
	if err1 != nil {
		return err1, nil
	}
	err2, sysUserRoleView := viewUtils.Data2View(sysUserRole)
	if err2 != nil {
		return err2, nil
	}
	return
}

// Find 分页获取SysUserRole记录
// Author
func (sysUserRoleService *SysUserRoleService) Find(info *common.PageInfoV2) (err error) {
	err1, sysUserRoles, total := sysUserRoleDao.Find(info)
	if err1 != nil {
		return err1
	}
	info.Total = total
	err2, viewList := viewUtils.Data2ViewList(sysUserRoles)
	if err2 != nil {
		return err2
	}
	info.FormList = viewList
	return err
}
