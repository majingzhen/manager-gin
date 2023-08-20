// Package service 自动生成模板 SysRoleService
// @description <TODO description class purpose>
// @author
// @File: sys_role
// @version 1.0.0
// @create 2023-08-18 14:00:53
package service

import (
	"manager-gin/src/app/admin/sys/sys_role/model"
	"manager-gin/src/app/admin/sys/sys_role/service/view"
	"manager-gin/src/app/admin/sys/sys_user/service"
	userView "manager-gin/src/app/admin/sys/sys_user/service/view"
	"manager-gin/src/common"
)

var sysRoleDao = model.SysRoleDaoApp
var viewUtils = view.SysRoleViewUtilsApp
var userService = service.SysUserServiceApp

type SysRoleService struct{}

// Create 创建SysRole记录
// Author
func (sysRoleService *SysRoleService) Create(sysRoleView *view.SysRoleView) (err error) {
	err1, sysRole := viewUtils.View2Data(sysRoleView)
	if err1 != nil {
		return err1
	}
	err2 := sysRoleDao.Create(*sysRole)
	if err2 != nil {
		return err2
	}
	return nil
}

// Delete 删除SysRole记录
// Author
func (sysRoleService *SysRoleService) Delete(id string) (err error) {
	err = sysRoleDao.Delete(id)
	return err
}

// DeleteByIds 批量删除SysRole记录
// Author
func (sysRoleService *SysRoleService) DeleteByIds(ids []string) (err error) {
	err = sysRoleDao.DeleteByIds(ids)
	return err
}

// Update 更新SysRole记录
// Author
func (sysRoleService *SysRoleService) Update(id string, sysRoleView *view.SysRoleView) (err error) {
	sysRoleView.Id = id
	err1, sysRole := viewUtils.View2Data(sysRoleView)
	if err1 != nil {
		return err1
	}
	err = sysRoleDao.Update(*sysRole)
	return err
}

// Get 根据id获取SysRole记录
// Author
func (sysRoleService *SysRoleService) Get(id string) (err error, sysRoleView *view.SysRoleView) {
	err1, sysRole := sysRoleDao.Get(id)
	if err1 != nil {
		return err1, nil
	}
	err2, sysRoleView := viewUtils.Data2View(sysRole)
	if err2 != nil {
		return err2, nil
	}
	return
}

// List 分页获取SysRole记录
// Author
func (sysRoleService *SysRoleService) List(info *common.PageInfo) (err error) {
	err1, sysRoles, total := sysRoleDao.List(info)
	if err1 != nil {
		return err1
	}
	info.Total = total
	err2, viewList := viewUtils.Data2ViewList(sysRoles)
	if err2 != nil {
		return err2
	}
	info.Rows = viewList
	return err
}

// GetRoleByUserId 根据用户获取角色集合
func (sysRoleService *SysRoleService) GetRoleByUserId(user *userView.SysUserView) (err error, roleNames []string) {
	is := userService.IsAdmin(user.Id)
	if is {
		roleNames = append(roleNames, "admin")
	}
	err1, roles := sysRoleDao.GetRoleByUserId(user.Id)
	if err1 != nil {
		return err1, nil
	}
	for _, role := range *roles {
		roleNames = append(roleNames, role.RoleName)
	}
	_, user.Roles = viewUtils.Data2ViewList(roles)
	return nil, roleNames
}
