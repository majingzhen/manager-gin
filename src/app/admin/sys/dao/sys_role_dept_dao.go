// Package model 自动生成模板 SysRoleDeptDao
// @description <TODO description class purpose>
// @author
// @File: sys_role_dept
// @version 1.0.0
// @create 2023-08-18 14:02:24
package dao

import (
	"gorm.io/gorm"
	"manager-gin/src/app/admin/sys/model"
)

// SysRoleDeptDao 结构体

type SysRoleDeptDao struct{}

// DeleteRoleDeptByRoleIds 根据角色id集合删除角色部门关联数据
func (dao *SysRoleDeptDao) DeleteRoleDeptByRoleIds(tx *gorm.DB, ids []string) error {
	return tx.Delete(&[]model.SysRoleDept{}, "role_id in ?", ids).Error
}

// DeleteRoleDeptByRoleId 根据角色id集合删除角色部门关联数据
func (dao *SysRoleDeptDao) DeleteRoleDeptByRoleId(tx *gorm.DB, id string) error {
	return tx.Delete(&[]model.SysRoleDept{}, "role_id = ?", id).Error

}

// CreateBatch 批量创建SysRoleDept记录
func (dao *SysRoleDeptDao) CreateBatch(tx *gorm.DB, depts []model.SysRoleDept) error {
	return tx.Create(&depts).Error
}
