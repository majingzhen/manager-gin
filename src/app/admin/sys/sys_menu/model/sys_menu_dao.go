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

// List 分页获取SysMenu记录
// Author
func (dao *SysMenuDao) List(param *SysMenu, page *common.PageInfo) (err error, datas *[]SysMenu, total int64) {
	// 创建model
	model := global.GOrmDao.Model(&SysMenu{})
	// 如果有条件搜索 下方会自动创建搜索语句
	//if param.Id != "" {
	//	model = model.Where("ID = ?", info.Id)
	//}

	err = model.Count(&total).Error
	if err != nil {
		return
	}
	// 计算分页信息
	page.Calculate()
	// 生成排序信息
	if page.OrderByColumn != "" {
		model.Order(page.OrderByColumn + " " + page.IsAsc + " ")
	}
	var tmp []SysMenu
	err = model.Limit(page.Limit).Offset(page.Offset).Find(&tmp).Error
	datas = &tmp
	return err, datas, total
}

func (dao *SysMenuDao) GetMenuPermissionByRoleId(roleId string) (err error, perms []string) {
	var rows []SysMenu
	db := global.GOrmDao.Table("sys_menu m")
	db.Joins("left join sys_role_menu rm on m.id = rm.menu_id")
	db.Joins("left join sys_role r on r.id = rm.role_id")
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
	db.Joins("left join sys_role_menu rm on m.id = rm.menu_id")
	db.Joins("left join sys_user_role ur on rm.role_id = ur.role_id")
	db.Select("distinct m.perms")
	db.Where("ur.user_id = ? and r.status = ? and m.status = ?", userId, common.STATUS_NORMAL, common.STATUS_NORMAL)
	err = db.Scan(&rows).Error
	if err != nil {
		return
	}
	for _, menu := range rows {
		perms = append(perms, menu.Perms)
	}
	return err, perms
}

// SelectMenuAll 查询所有菜单
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

// SelectMenuByUserId 根据用户id查询菜单
func (dao *SysMenuDao) SelectMenuByUserId(userId string) (err error, menus *[]SysMenu) {
	var rows []SysMenu
	db := global.GOrmDao.Table("sys_menu m")
	db.Joins("left join sys_role_menu rm on m.id = rm.menu_id")
	db.Joins("left join sys_user_role ur on rm.role_id = ur.role_id")
	db.Joins("left join sys_role r on r.id = rm.role_id")
	db.Select("distinct m.id, m.parent_id, m.menu_name, m.path, m.component, m.`query`, m.visible, m.status, perms, m.is_frame, m.is_cache, m.menu_type, m.icon, m.order_num, m.create_time")
	db.Where("ur.user_id = ? and r.status = ? and m.status = ? and menu_type in (?, ?)", userId, common.STATUS_NORMAL, common.STATUS_NORMAL, common.MENU_TYPE_DIR, common.MENU_TYPE_MENU)
	db.Order("m.parent_id, m.order_num")
	err = db.Scan(&rows).Error
	if err != nil {
		return err, nil
	}
	menus = &rows
	return err, menus
}

func (dao *SysMenuDao) SelectMenuList(data *SysMenu) (err error, menus *[]SysMenu) {
	var rows []SysMenu
	db := global.GOrmDao.Table(data.TableName())
	db.Select("distinct id, parent_id, menu_name, path, component, `query`, visible, status, perms, is_frame, is_cache, menu_type, icon, order_num, create_time")
	if data.MenuName != "" {
		db.Where("menu_name like ?", "%"+data.MenuName+"%")
	}
	if data.Visible != "" {
		db.Where("visible = ?", data.Visible)
	}
	if data.Status != "" {
		db.Where("status = ?", data.Status)
	}
	db.Order("parent_id, order_num")
	err = db.Scan(&rows).Error
	menus = &rows
	return err, menus
}

// SelectMenuListByUserId 根据用户id查询菜单
func (dao *SysMenuDao) SelectMenuListByUserId(data *SysMenu, userId string) (err error, menus *[]SysMenu) {
	var rows []SysMenu
	db := global.GOrmDao.Table("sys_menu m")
	db.Joins("left join sys_role_menu rm on m.id = rm.menu_id")
	db.Joins("left join sys_user_role ur on rm.role_id = ur.role_id")
	db.Select("distinct m.id, m.parent_id, m.menu_name, m.path, m.component, m.`query`, m.visible, m.status, m.perms, m.is_frame, m.is_cache, m.menu_type, m.icon, m.order_num, m.create_time")
	db.Where("ur.user_id = ? ", userId)
	if data.MenuName != "" {
		db.Where("menu_name like ?", "%"+data.MenuName+"%")
	}
	if data.Visible != "" {
		db.Where("visible = ?", data.Visible)
	}
	if data.Status != "" {
		db.Where("status = ?", data.Status)
	}
	db.Order("m.parent_id, m.order_num")
	err = db.Scan(&rows).Error
	menus = &rows
	return err, menus
}

// CheckMenuNameUniqueAll 判断名称是否唯一
func (dao *SysMenuDao) CheckMenuNameUniqueAll(data *SysMenu) (error, bool) {
	var menu SysMenu
	db := global.GOrmDao.Model(&SysMenu{})
	db.Where("menu_name = ?", data.MenuName)
	if data.ParentId != "" {
		db.Where("parent_id = ?", data.ParentId)
	}
	err := db.First(&menu).Error
	if err != nil {
		return err, false
	}
	return nil, true
}

// SelectMenuByParentId 根据父级id查询菜单
func (dao *SysMenuDao) SelectMenuByParentId(id string) (error, *[]SysMenu) {
	var rows []SysMenu
	db := global.GOrmDao.Model(&SysMenu{})
	db.Where("parent_id = ?", id)
	err := db.Find(&rows).Error
	if err != nil {
		return err, nil
	}
	return nil, &rows
}

// CheckMenuExistRole 判断菜单是否存在角色
func (dao *SysMenuDao) CheckMenuExistRole(roleId string) (error, bool) {
	var count int64
	db := global.GOrmDao.Table("sys_role_menu").Where("menu_id = ?", roleId)
	err := db.Count(&count).Error
	if err != nil {
		return err, false
	}
	if count > 0 {
		return nil, true
	}
	return nil, false
}
