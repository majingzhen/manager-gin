// Package model 自动生成模板 SysRoleMenuDao
// @description <TODO description class purpose>
// @author
// @File: sys_role_menu
// @version 1.0.0
// @create 2023-08-18 14:02:24
package dao

import (
	"manager-gin/src/app/admin/sys/model"
	"manager-gin/src/global"
)

// SysRoleMenuDao 结构体

type SysRoleMenuDao struct{}

// DeleteRoleMenuByRoleId 根据角色id删除角色菜单关联数据
func (dao *SysRoleMenuDao) DeleteRoleMenuByRoleId(id string) error {
	return global.GOrmDao.Delete(&[]model.SysRoleMenu{}, "role_id = ?", id).Error
}

// CreateBatch 批量插入
func (dao *SysRoleMenuDao) CreateBatch(menus []model.SysRoleMenu) error {
	return global.GOrmDao.Create(&menus).Error
}

// DeleteRoleMenuByRoleIds 根据角色id集合删除角色菜单关联数据
func (dao *SysRoleMenuDao) DeleteRoleMenuByRoleIds(ids []string) error {
	return global.GOrmDao.Delete(&[]model.SysRoleMenu{}, "role_id in ?", ids).Error
}
