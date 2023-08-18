// Package model 自动生成模板 SysRoleMenuDao
// @description <TODO description class purpose>
// @author
// @File: sys_role_menu
// @version 1.0.0
// @create 2023-08-18 13:41:26
package model

import (
	"manager-gin/src/common"
	"manager-gin/src/global"
)

// SysRoleMenuDao 结构体

type SysRoleMenuDao struct{}

// Create 创建SysRoleMenu记录
// Author
func (dao *SysRoleMenuDao) Create(sysRoleMenu SysRoleMenu) (err error) {
	err = global.GOrmDao.Create(&sysRoleMenu).Error
	return err
}

// Delete 删除SysRoleMenu记录
// Author
func (dao *SysRoleMenuDao) Delete(id string) (err error) {
	err = global.GOrmDao.Delete(&[]SysRoleMenu{}, "id = ?", id).Error
	return err
}

// DeleteByIds 批量删除SysRoleMenu记录
// Author
func (dao *SysRoleMenuDao) DeleteByIds(ids []string) (err error) {
	err = global.GOrmDao.Delete(&[]SysRoleMenu{}, "id in ?", ids).Error
	return err
}

// Update 更新SysRoleMenu记录
// Author
func (dao *SysRoleMenuDao) Update(sysRoleMenu SysRoleMenu) (err error) {
	err = global.GOrmDao.Save(&sysRoleMenu).Error
	return err
}

// Get 根据id获取SysRoleMenu记录
// Author
func (dao *SysRoleMenuDao) Get(id string) (err error, sysRoleMenu *SysRoleMenu) {
	err = global.GOrmDao.Where("id = ?", id).First(&sysRoleMenu).Error
	return
}

// Find 分页获取SysRoleMenu记录
// Author
func (dao *SysRoleMenuDao) Find(info *common.PageInfoV2) (err error, sysRoleMenus *[]SysRoleMenu, total int64) {
	// 创建db
	db := global.GOrmDao.Model(&SysRoleMenu{})
	// 如果有条件搜索 下方会自动创建搜索语句
	//if info.Id != "" {
	//	db = db.Where("ID = ?", info.Id)
	//}

	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var tmp []SysRoleMenu
	err = db.Limit(info.Limit).Offset(info.Offset).Find(&tmp).Error
	sysRoleMenus = &tmp
	return err, sysRoleMenus, total
}
