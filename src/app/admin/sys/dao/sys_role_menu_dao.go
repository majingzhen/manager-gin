// Package dao 自动生成模板 SysRoleMenuDao
// @description <TODO description class purpose>
// @author
// @File: sys_role_menu
// @version 1.0.0
// @create 2023-08-18 14:02:24
package dao

import (
	"gorm.io/gorm"
	"manager-gin/src/app/admin/sys/model"
)

// SysRoleMenuDao 结构体

type SysRoleMenuDao struct{}

// DeleteRoleMenuByRoleId 根据角色id删除角色菜单关联数据
func (dao *SysRoleMenuDao) DeleteRoleMenuByRoleId(tx *gorm.DB, id string) error {
	return tx.Delete(&[]model.SysRoleMenu{}, "role_id = ?", id).Error
}

// CreateBatch 批量插入
func (dao *SysRoleMenuDao) CreateBatch(tx *gorm.DB, menus []model.SysRoleMenu) error {
	return tx.Create(&menus).Error
}

// DeleteRoleMenuByRoleIds 根据角色id集合删除角色菜单关联数据
func (dao *SysRoleMenuDao) DeleteRoleMenuByRoleIds(tx *gorm.DB, ids []string) error {
	return tx.Delete(&[]model.SysRoleMenu{}, "role_id in ?", ids).Error
}
