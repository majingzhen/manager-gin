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

// SelectMenuNormalAll 获取全部数据
func (dao *SysMenuDao) SelectMenuNormalAll() {

}

func (dao *SysMenuDao) GetMenuPermissionByRoleId(roleId string) (err error, perms []string) {
	var rows []SysMenu
	db := global.GOrmDao.Table("sys_menu m")
	db.Joins("join sys_role_menu rm", "m.id = rm.menu_id")
	db.Select("distinct m.perms")
	db.Where("rm.role_id = ? and r.status = ?", roleId, common.STATUS_NORMAL)
	err = db.Scan(&rows).Error
	if err != nil {
		return
	}
	for _, menu := range rows {
		perms = append(perms, menu.Perms)
	}
	return err, perms
}

func (dao *SysMenuDao) GetMenuPermissionByUserId(userId string) (err error, perms []string) {
	var rows []SysMenu
	db := global.GOrmDao.Table("sys_menu m")
	db.Joins("join sys_role_menu rm", "m.id = rm.menu_id")
	db.Joins("join sys_user_role ur", "rm.role_id = ur.role_id")
	db.Select("distinct m.perms")
	db.Where("ur.role_id = ? and r.status = ? and m.status = ?", userId, common.STATUS_NORMAL, common.STATUS_NORMAL)
	err = db.Scan(&rows).Error
	if err != nil {
		return
	}
	for _, menu := range rows {
		perms = append(perms, menu.Perms)
	}
	return err, perms
}

func (dao *SysMenuDao) SelectMenuAll() (err error, menus *[]SysMenu) {
	db := global.GOrmDao.Model(&[]SysMenu{})
	db.Where("status = ? and menu_type in (?, ?)", common.STATUS_NORMAL, common.MENU_TYPE_DIR, common.MENU_TYPE_MENU)
	var tmp []SysMenu
	err1 := db.Find(&tmp).Error
	if err1 != nil {
		return err1, nil
	}
	menus = &tmp
	return err, menus
}

func (dao *SysMenuDao) SelectMenuByUserId(userId string) (err error, menus *[]SysMenu) {
	var rows []SysMenu
	db := global.GOrmDao.Table("sys_menu m")
	db.Joins("join sys_role_menu rm", "m.id = rm.menu_id")
	db.Joins("join sys_user_role ur", "rm.role_id = ur.role_id")
	db.Select("distinct m.perms")
	db.Where("ur.role_id = ? and r.status = ? and m.status = ?", userId, common.STATUS_NORMAL, common.STATUS_NORMAL)
	err = db.Scan(&rows).Error
	if err != nil {
		return err, nil
	}
	menus = &rows
	return err, menus
}
