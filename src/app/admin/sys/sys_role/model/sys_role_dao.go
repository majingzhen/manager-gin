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

// Page 分页获取SysRole记录
// Author
func (dao *SysRoleDao) Page(param *SysRole, page *common.PageInfo) (err error, datas *[]SysRole, total int64) {
	// 创建model
	model := global.GOrmDao.Model(&SysRole{})
	// 如果有条件搜索 下方会自动创建搜索语句
	//if param.Id != "" {
	//	model = model.Where("ID = ?", info.Id)
	//}
	if err = model.Count(&total).Error; err != nil {
		return
	}
	// 计算分页信息
	page.Calculate()
	// 生成排序信息
	if page.OrderByColumn != "" {
		model.Order(page.OrderByColumn + " " + page.IsAsc + " ")
	}
	var tmp []SysRole
	err = model.Limit(page.Limit).Offset(page.Offset).Find(&tmp).Error
	datas = &tmp
	return err, datas, total
}

// List 获取SysRole记录
// Author
func (dao *SysRoleDao) List(data *SysRole) (err error, datas *[]SysRole) {
	var rows []SysRole
	db := global.GOrmDao.Model(&SysRole{})
	// TODO 输入查询条件
	//if data.Id != "" {
	//    db.Where("id = ?", data.Id)
	//}
	db.Order("create_time desc")
	err = db.Find(&rows).Error
	datas = &rows
	return err, datas
}

// GetRoleByUserId 根据用户获取角色集合
func (dao *SysRoleDao) GetRoleByUserId(userId string) (err error, roles *[]SysRole) {
	var tmp []SysRole
	db := global.GOrmDao.Table("sys_role r")
	db.Joins("join sys_user_role ur", "r.id = ur.role_id")
	db.Where("ur.user_id = ? and r.status = ?", userId, common.STATUS_NORMAL)
	err = db.Find(&tmp).Error
	roles = &tmp
	return err, roles
}
