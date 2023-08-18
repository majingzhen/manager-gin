// Package service 自动生成模板 SysMenuService
// @description <TODO description class purpose>
// @author
// @File: sys_menu
// @version 1.0.0
// @create 2023-08-18 13:41:26
package service

import (
	"manager-gin/src/app/admin/sys/sys_menu/model"
	"manager-gin/src/app/admin/sys/sys_menu/service/view"
	"manager-gin/src/common"
)

var sysMenuDao = model.SysMenuDaoApp
var viewUtils = view.SysMenuViewUtilsApp

type SysMenuService struct{}

// Create 创建SysMenu记录
// Author
func (sysMenuService *SysMenuService) Create(sysMenuView *view.SysMenuView) (err error) {
	err1, sysMenu := viewUtils.View2Data(sysMenuView)
	if err1 != nil {
		return err1
	}
	err2 := sysMenuDao.Create(*sysMenu)
	if err2 != nil {
		return err2
	}
	return nil
}

// Delete 删除SysMenu记录
// Author
func (sysMenuService *SysMenuService) Delete(id string) (err error) {
	err = sysMenuDao.Delete(id)
	return err
}

// DeleteByIds 批量删除SysMenu记录
// Author
func (sysMenuService *SysMenuService) DeleteByIds(ids []string) (err error) {
	err = sysMenuDao.DeleteByIds(ids)
	return err
}

// Update 更新SysMenu记录
// Author
func (sysMenuService *SysMenuService) Update(id string, sysMenuView *view.SysMenuView) (err error) {
	sysMenuView.Id = id
	err1, sysMenu := viewUtils.View2Data(sysMenuView)
	if err1 != nil {
		return err1
	}
	err = sysMenuDao.Update(*sysMenu)
	return err
}

// Get 根据id获取SysMenu记录
// Author
func (sysMenuService *SysMenuService) Get(id string) (err error, sysMenuView *view.SysMenuView) {
	err1, sysMenu := sysMenuDao.Get(id)
	if err1 != nil {
		return err1, nil
	}
	err2, sysMenuView := viewUtils.Data2View(sysMenu)
	if err2 != nil {
		return err2, nil
	}
	return
}

// Find 分页获取SysMenu记录
// Author
func (sysMenuService *SysMenuService) Find(info *common.PageInfoV2) (err error) {
	err1, sysMenus, total := sysMenuDao.Find(info)
	if err1 != nil {
		return err1
	}
	info.Total = total
	err2, viewList := viewUtils.Data2ViewList(sysMenus)
	if err2 != nil {
		return err2
	}
	info.FormList = viewList
	return err
}

// RoleMenuTreeData 根据角色查询菜单列表
func (sysMenuService *SysMenuService) RoleMenuTreeData(roleId string) (err error) {

	return err
}
