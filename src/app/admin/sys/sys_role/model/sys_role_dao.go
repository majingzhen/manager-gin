// Package model 自动生成模板 SysRoleDao
// @description <TODO description class purpose>
// @author
// @File: sys_role
// @version 1.0.0
// @create 2023-08-18 14:00:53
package model

import (
	"manager-gin/src/common"
	"manager-gin/src/global"
)

// SysRoleDao 结构体

type SysRoleDao struct{}

// Create 创建SysRole记录
// Author
func (dao *SysRoleDao) Create(sysRole SysRole) (err error) {
	err = global.GOrmDao.Create(&sysRole).Error
	return err
}

// Delete 删除SysRole记录
// Author
func (dao *SysRoleDao) Delete(id string) (err error) {
	err = global.GOrmDao.Delete(&[]SysRole{}, "id = ?", id).Error
	return err
}

// DeleteByIds 批量删除SysRole记录
// Author
func (dao *SysRoleDao) DeleteByIds(ids []string) (err error) {
	err = global.GOrmDao.Delete(&[]SysRole{}, "id in ?", ids).Error
	return err
}

// Update 更新SysRole记录
// Author
func (dao *SysRoleDao) Update(sysRole SysRole) (err error) {
	err = global.GOrmDao.Save(&sysRole).Error
	return err
}

// Get 根据id获取SysRole记录
// Author
func (dao *SysRoleDao) Get(id string) (err error, sysRole *SysRole) {
	err = global.GOrmDao.Where("id = ?", id).First(&sysRole).Error
	return
}

// Find 分页获取SysRole记录
// Author
func (dao *SysRoleDao) Find(info *common.PageInfoV2) (err error, sysRoles *[]SysRole, total int64) {
	// 创建db
	db := global.GOrmDao.Model(&SysRole{})
	// 如果有条件搜索 下方会自动创建搜索语句
	//if info.Id != "" {
	//	db = db.Where("ID = ?", info.Id)
	//}

	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var tmp []SysRole
	err = db.Limit(info.Limit).Offset(info.Offset).Find(&tmp).Error
	sysRoles = &tmp
	return err, sysRoles, total
}
