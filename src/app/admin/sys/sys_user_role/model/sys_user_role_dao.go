// Package model 自动生成模板 SysUserRoleDao
// @description <TODO description class purpose>
// @author
// @File: sys_user_role
// @version 1.0.0
// @create 2023-08-18 13:41:26
package model

import (
	"manager-gin/src/common"
	"manager-gin/src/global"
)

// SysUserRoleDao 结构体

type SysUserRoleDao struct{}

// Create 创建SysUserRole记录
// Author
func (dao *SysUserRoleDao) Create(sysUserRole SysUserRole) (err error) {
	err = global.GOrmDao.Create(&sysUserRole).Error
	return err
}

// Delete 删除SysUserRole记录
// Author
func (dao *SysUserRoleDao) Delete(id string) (err error) {
	err = global.GOrmDao.Delete(&[]SysUserRole{}, "id = ?", id).Error
	return err
}

// DeleteByIds 批量删除SysUserRole记录
// Author
func (dao *SysUserRoleDao) DeleteByIds(ids []string) (err error) {
	err = global.GOrmDao.Delete(&[]SysUserRole{}, "id in ?", ids).Error
	return err
}

// Update 更新SysUserRole记录
// Author
func (dao *SysUserRoleDao) Update(sysUserRole SysUserRole) (err error) {
	err = global.GOrmDao.Save(&sysUserRole).Error
	return err
}

// Get 根据id获取SysUserRole记录
// Author
func (dao *SysUserRoleDao) Get(id string) (err error, sysUserRole *SysUserRole) {
	err = global.GOrmDao.Where("id = ?", id).First(&sysUserRole).Error
	return
}

// Find 分页获取SysUserRole记录
// Author
func (dao *SysUserRoleDao) Find(info *common.PageInfoV2) (err error, sysUserRoles *[]SysUserRole, total int64) {
	// 创建db
	db := global.GOrmDao.Model(&SysUserRole{})
	// 如果有条件搜索 下方会自动创建搜索语句
	//if info.Id != "" {
	//	db = db.Where("ID = ?", info.Id)
	//}

	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var tmp []SysUserRole
	err = db.Limit(info.Limit).Offset(info.Offset).Find(&tmp).Error
	sysUserRoles = &tmp
	return err, sysUserRoles, total
}
