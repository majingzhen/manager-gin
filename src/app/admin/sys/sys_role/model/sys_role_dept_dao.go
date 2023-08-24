// Package model 自动生成模板 SysRoleDeptDao
// @description <TODO description class purpose>
// @author
// @File: sys_role_dept
// @version 1.0.0
// @create 2023-08-18 14:02:24
package model

import "manager-gin/src/global"

// SysRoleDeptDao 结构体

type SysRoleDeptDao struct{}

// DeleteRoleDeptByRoleIds 根据角色id集合删除角色部门关联数据
func (dao *SysRoleDeptDao) DeleteRoleDeptByRoleIds(ids []string) error {
	return global.GOrmDao.Delete(&[]SysRoleDept{}, "role_id in ?", ids).Error
}
