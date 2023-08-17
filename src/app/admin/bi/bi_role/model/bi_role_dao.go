// Package model 自动生成模板 BiRoleDao
// @description <TODO description class purpose>
// @author Majz
// @File: bi_role
// @version 1.0.0
// @create 2023-08-16 14:53:36
package model

import (
	"manager-gin/src/common"
	"manager-gin/src/global"
)

// BiRoleDao 结构体

type BiRoleDao struct{}

// Create 创建BiRole记录
// Author Majz
func (dao *BiRoleDao) Create(biRole BiRole) (err error) {
	err = global.GOrmDao.Create(&biRole).Error
	return err
}

// Delete 删除BiRole记录
// Author Majz
func (dao *BiRoleDao) Delete(id int) (err error) {
	err = global.GOrmDao.Delete(&[]BiRole{}, "id = ?", id).Error
	return err
}

// DeleteByIds 批量删除BiRole记录
// Author Majz
func (dao *BiRoleDao) DeleteByIds(ids []int) (err error) {
	err = global.GOrmDao.Delete(&[]BiRole{}, "id in ?", ids).Error
	return err
}

// Update 更新BiRole记录
// Author Majz
func (dao *BiRoleDao) Update(biRole BiRole) (err error) {
	err = global.GOrmDao.Save(&biRole).Error
	return err
}

// Get 根据id获取BiRole记录
// Author Majz
func (dao *BiRoleDao) Get(id int) (err error, biRole *BiRole) {
	err = global.GOrmDao.Where("id = ?", id).First(&biRole).Error
	return
}

// Find 分页获取BiRole记录
// Author Majz
func (dao *BiRoleDao) Find(info *common.PageInfoV2) (err error, biRoles *[]BiRole, total int64) {
	// 创建db
	db := global.GOrmDao.Model(&BiRole{})
	// 如果有条件搜索 下方会自动创建搜索语句
	//if info.Id != "" {
	//	db = db.Where("ID = ?", info.Id)
	//}

	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var tmp []BiRole
	err = db.Limit(info.Limit).Offset(info.Offset).Find(&tmp).Error
	biRoles = &tmp
	return err, biRoles, total
}
