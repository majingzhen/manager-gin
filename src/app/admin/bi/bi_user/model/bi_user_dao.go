// Package model 自动生成模板 BiUserDao
// @description <TODO description class purpose>
// @author Administrator
// @File: bi_user
// @version 1.0.0
// @create 2023-08-16 14:53:36
package model

import (
	"manager-gin/src/common"
	"manager-gin/src/global"
)

// BiUserDao 结构体

type BiUserDao struct{}

// Create 创建BiUser记录
// Author Administrator
func (dao *BiUserDao) Create(biUser BiUser) (err error) {
	err = global.GOrmDao.Create(&biUser).Error
	return err
}

// Delete 删除BiUser记录
// Author Administrator
func (dao *BiUserDao) Delete(id string) (err error) {
	err = global.GOrmDao.Delete(&[]BiUser{}, "id = ?", id).Error
	return err
}

// DeleteByIds 批量删除BiUser记录
// Author Administrator
func (dao *BiUserDao) DeleteByIds(ids []string) (err error) {
	err = global.GOrmDao.Delete(&[]BiUser{}, "id in ?", ids).Error
	return err
}

// Update 更新BiUser记录
// Author Administrator
func (dao *BiUserDao) Update(biUser BiUser) (err error) {
	err = global.GOrmDao.Save(&biUser).Error
	return err
}

// Get 根据id获取BiUser记录
// Author Administrator
func (dao *BiUserDao) Get(id string) (err error, biUser *BiUser) {
	err = global.GOrmDao.Where("id = ?", id).First(&biUser).Error
	return
}

// Find 分页获取BiUser记录
// Author Administrator
func (dao *BiUserDao) Find(info *common.PageInfoV2) (err error, biUsers *[]BiUser, total int64) {
	// 创建db
	db := global.GOrmDao.Model(&BiUser{})
	// 如果有条件搜索 下方会自动创建搜索语句
	//if info.Id != "" {
	//	db = db.Where("ID = ?", info.Id)
	//}

	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var tmp []BiUser
	err = db.Limit(info.Limit).Offset(info.Offset).Find(&tmp).Error
	biUsers = &tmp
	return err, biUsers, total
}
