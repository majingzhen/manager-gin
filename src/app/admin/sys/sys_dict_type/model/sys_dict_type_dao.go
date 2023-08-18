// Package model 自动生成模板 SysDictTypeDao
// @description <TODO description class purpose>
// @author
// @File: sys_dict_type
// @version 1.0.0
// @create 2023-08-18 13:41:26
package model

import (
	"manager-gin/src/common"
	"manager-gin/src/global"
)

// SysDictTypeDao 结构体

type SysDictTypeDao struct{}

// Create 创建SysDictType记录
// Author
func (dao *SysDictTypeDao) Create(sysDictType SysDictType) (err error) {
	err = global.GOrmDao.Create(&sysDictType).Error
	return err
}

// Delete 删除SysDictType记录
// Author
func (dao *SysDictTypeDao) Delete(id string) (err error) {
	err = global.GOrmDao.Delete(&[]SysDictType{}, "id = ?", id).Error
	return err
}

// DeleteByIds 批量删除SysDictType记录
// Author
func (dao *SysDictTypeDao) DeleteByIds(ids []string) (err error) {
	err = global.GOrmDao.Delete(&[]SysDictType{}, "id in ?", ids).Error
	return err
}

// Update 更新SysDictType记录
// Author
func (dao *SysDictTypeDao) Update(sysDictType SysDictType) (err error) {
	err = global.GOrmDao.Save(&sysDictType).Error
	return err
}

// Get 根据id获取SysDictType记录
// Author
func (dao *SysDictTypeDao) Get(id string) (err error, sysDictType *SysDictType) {
	err = global.GOrmDao.Where("id = ?", id).First(&sysDictType).Error
	return
}

// Find 分页获取SysDictType记录
// Author
func (dao *SysDictTypeDao) Find(info *common.PageInfoV2) (err error, sysDictTypes *[]SysDictType, total int64) {
	// 创建db
	db := global.GOrmDao.Model(&SysDictType{})
	// 如果有条件搜索 下方会自动创建搜索语句
	//if info.Id != "" {
	//	db = db.Where("ID = ?", info.Id)
	//}

	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var tmp []SysDictType
	err = db.Limit(info.Limit).Offset(info.Offset).Find(&tmp).Error
	sysDictTypes = &tmp
	return err, sysDictTypes, total
}
