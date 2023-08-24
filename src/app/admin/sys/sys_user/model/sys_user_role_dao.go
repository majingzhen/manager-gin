// Package model 自动生成模板 SysUserRoleDao
// @description <TODO description class purpose>
// @author
// @File: sys_user_role
// @version 1.0.0
// @create 2023-08-21 17:37:56
package model

import (
	"manager-gin/src/global"
)

// SysUserRoleDao 结构体

type SysUserRoleDao struct{}

// CreateBatch 批量创建SysUserRole记录
func (dao *SysUserRoleDao) CreateBatch(roles []SysUserRole) error {
	return global.GOrmDao.Create(&roles).Error
}

// DeleteByUserIds 根据用户id删除用户角色关联数据
func (dao *SysUserRoleDao) DeleteByUserIds(ids []string) error {
	return global.GOrmDao.Delete(&[]SysUserRole{}, "user_id in ?", ids).Error
}

// CountUserRoleByRoleId 根据角色id查询用户数量
func (dao *SysUserRoleDao) CountUserRoleByRoleId(id string) (error, int64) {
	var total int64
	err := global.GOrmDao.Model(&SysUserRole{}).Where("role_id = ?", id).Count(&total).Error
	return err, total

}
