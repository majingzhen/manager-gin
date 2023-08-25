// Package model 自动生成模板 SysRoleDeptDao
// @description <TODO description class purpose>
// @author
// @File: sys_role_dept
// @version 1.0.0
// @create 2023-08-18 14:02:24
package dao

import (
	"manager-gin/src/app/admin/sys/model"
	"manager-gin/src/global"
)

// SysRoleDeptDao 结构体

type SysRoleDeptDao struct{}

// DeleteRoleDeptByRoleIds 根据角色id集合删除角色部门关联数据
func (dao *SysRoleDeptDao) DeleteRoleDeptByRoleIds(ids []string) error {
	return global.GOrmDao.Delete(&[]model.SysRoleDept{}, "role_id in ?", ids).Error
}

// DeleteRoleDeptByDeptIds 根据角色id集合删除角色部门关联数据
func (dao *SysRoleDeptDao) DeleteRoleDeptByRoleId(id string) error {
	return global.GOrmDao.Delete(&[]model.SysRoleDept{}, "role_id = ?", id).Error

}

// CreateBatch 批量创建SysRoleDept记录
func (dao *SysRoleDeptDao) CreateBatch(depts []model.SysRoleDept) error {
	return global.GOrmDao.Create(&depts).Error
}
