// Package dao 自动生成模板 SysUserRoleDao
// @description <TODO description class purpose>
// @author
// @File: sys_user_role
// @version 1.0.0
// @create 2023-08-21 17:37:56
package dao

import (
	"gorm.io/gorm"
	"manager-gin/src/app/admin/sys/model"
	"manager-gin/src/global"
)

// SysUserRoleDao 结构体

type SysUserRoleDao struct{}

// CreateBatch 批量创建SysUserRole记录
func (dao *SysUserRoleDao) CreateBatch(tx *gorm.DB, roles []model.SysUserRole) error {
	return tx.Create(&roles).Error
}

// DeleteByUserIds 根据用户id删除用户角色关联数据
func (dao *SysUserRoleDao) DeleteByUserIds(tx *gorm.DB, ids []string) error {
	return tx.Delete(&[]model.SysUserRole{}, "user_id in ?", ids).Error
}

// CountUserRoleByRoleId 根据角色id查询用户数量
func (dao *SysUserRoleDao) CountUserRoleByRoleId(id string) (error, int64) {
	var total int64
	err := global.GOrmDao.Model(&model.SysUserRole{}).Where("role_id = ?", id).Count(&total).Error
	return err, total

}

// DeleteUserRoleInfo 根据用户id和角色id删除用户角色关联数据
func (dao *SysUserRoleDao) DeleteUserRoleInfo(userId, roleId string) error {
	return global.GOrmDao.Delete(&model.SysUserRole{}, "user_id = ? and role_id = ?", userId, roleId).Error
}

// DeleteUsersRoleInfo 根据用户id集合和角色id删除用户角色关联数据
func (dao *SysUserRoleDao) DeleteUsersRoleInfo(roleId string, userIds []string) error {
	return global.GOrmDao.Delete(&model.SysUserRole{}, "user_id in ? and role_id = ?", userIds, roleId).Error
}

// InsertUsersRoleInfo 批量插入用户角色关联数据
func (dao *SysUserRoleDao) InsertUsersRoleInfo(roleId string, userIds []string) error {
	var userRoles []model.SysUserRole
	for _, userId := range userIds {
		userRoles = append(userRoles, model.SysUserRole{RoleId: roleId, UserId: userId})
	}
	return global.GOrmDao.Create(&userRoles).Error
}
