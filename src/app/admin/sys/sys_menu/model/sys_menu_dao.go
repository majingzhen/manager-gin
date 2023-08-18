// Package model 自动生成模板 SysMenuDao
// @description <TODO description class purpose>
// @author
// @File: sys_menu
// @version 1.0.0
// @create 2023-08-18 13:41:26
package model

import (
	"manager-gin/src/common"
	"manager-gin/src/global"
)

// SysMenuDao 结构体

type SysMenuDao struct{}

// Create 创建SysMenu记录
// Author
func (dao *SysMenuDao) Create(sysMenu SysMenu) (err error) {
	err = global.GOrmDao.Create(&sysMenu).Error
	return err
}

// Delete 删除SysMenu记录
// Author
func (dao *SysMenuDao) Delete(id string) (err error) {
	err = global.GOrmDao.Delete(&[]SysMenu{}, "id = ?", id).Error
	return err
}

// DeleteByIds 批量删除SysMenu记录
// Author
func (dao *SysMenuDao) DeleteByIds(ids []string) (err error) {
	err = global.GOrmDao.Delete(&[]SysMenu{}, "id in ?", ids).Error
	return err
}

// Update 更新SysMenu记录
// Author
func (dao *SysMenuDao) Update(sysMenu SysMenu) (err error) {
	err = global.GOrmDao.Save(&sysMenu).Error
	return err
}

// Get 根据id获取SysMenu记录
// Author
func (dao *SysMenuDao) Get(id string) (err error, sysMenu *SysMenu) {
	err = global.GOrmDao.Where("id = ?", id).First(&sysMenu).Error
	return
}

// Find 分页列表
// Author
func (dao *SysMenuDao) Find(info *common.PageInfoV2) (err error, sysMenus *[]SysMenu, total int64) {
	// 创建db
	db := global.GOrmDao.Model(&SysMenu{})
	// 如果有条件搜索 下方会自动创建搜索语句
	//if info.Id != "" {
	//	db = db.Where("ID = ?", info.Id)
	//}

	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var tmp []SysMenu
	err = db.Limit(info.Limit).Offset(info.Offset).Find(&tmp).Error
	sysMenus = &tmp
	return err, sysMenus, total
}

// MenuTreeDataByRoles 根据角色集合获取菜单
// Author
func (dao *SysMenuDao) MenuTreeDataByRoles(roleIds []string) (err error, sysMenus *[]SysMenu) {
	// 创建db
	model := global.GOrmDao.Table("sys_menu m")
	model.Joins("JOIN sys_role_menu rm", "m.menu_id = rm.menu.id")
	model.Where("rm.role_id in (?)", roleIds)
	model.Order("m.parent_id, m.order_num")
	_, err = model.Select("m.*").Rows()
	if err != nil {
		return
	}

	return err, sysMenus
}
