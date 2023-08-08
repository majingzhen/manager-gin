// Package service 自动生成模板 SysRoleMenuService
// @description <TODO description class purpose>
// @author
// @File: sys_role_menu
// @version 1.0.0
// @create 2023-08-08 10:06:19
package service

import (
	"manager-gin/src/app/admin/sys/sys_role_menu/model"
	"manager-gin/src/app/admin/sys/sys_role_menu/service/view"
	"manager-gin/src/common"
)

var sysRoleMenuDao = model.SysRoleMenuDaoApp
var viewUtils = view.SysRoleMenuViewUtilsApp

type SysRoleMenuService struct{}

// Create 创建SysRoleMenu记录
// Author
func (sysRoleMenuService *SysRoleMenuService) Create(sysRoleMenuView *view.SysRoleMenuView) (err error) {
	err1, sysRoleMenu := viewUtils.View2Data(sysRoleMenuView)
	if err1 != nil {
		return err1
	}
	err2 := sysRoleMenuDao.Create(*sysRoleMenu)
	if err2 != nil {
		return err2
	}
	return nil
}

// Delete 删除SysRoleMenu记录
// Author
func (sysRoleMenuService *SysRoleMenuService) Delete(id int) (err error) {
	err = sysRoleMenuDao.Delete(id)
	return err
}

// DeleteByIds 批量删除SysRoleMenu记录
// Author
func (sysRoleMenuService *SysRoleMenuService) DeleteByIds(ids []int) (err error) {
	err = sysRoleMenuDao.DeleteByIds(ids)
	return err
}

// Update 更新SysRoleMenu记录
// Author
func (sysRoleMenuService *SysRoleMenuService) Update(id int, sysRoleMenuView *view.SysRoleMenuView) (err error) {
	sysRoleMenuView.Id = id
	err1, sysRoleMenu := viewUtils.View2Data(sysRoleMenuView)
	if err1 != nil {
		return err1
	}
	err = sysRoleMenuDao.Update(*sysRoleMenu)
	return err
}

// Get 根据id获取SysRoleMenu记录
// Author
func (sysRoleMenuService *SysRoleMenuService) Get(id int) (err error, sysRoleMenuView *view.SysRoleMenuView) {
	err1, sysRoleMenu := sysRoleMenuDao.Get(id)
	if err1 != nil {
		return err1, nil
	}
	err2, sysRoleMenuView := viewUtils.Data2View(sysRoleMenu)
	if err2 != nil {
		return err2, nil
	}
	return
}

// Find 分页获取SysRoleMenu记录
// Author
func (sysRoleMenuService *SysRoleMenuService) Find(info *common.PageInfoV2) (err error) {
	err1, sysRoleMenus, total := sysRoleMenuDao.Find(info)
	if err1 != nil {
		return err1
	}
	info.Total = total
	err2, viewList := viewUtils.Data2ViewList(sysRoleMenus)
	if err2 != nil {
		return err2
	}
	info.FormList = viewList
	return err
}
