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

// List 分页获取SysDictType记录
// Author
func (dao *SysDictTypeDao) List(param *SysDictType, page *common.PageInfo) (err error, sysDictTypes *[]SysDictType, total int64) {
	// 创建db
	db := global.GOrmDao.Model(&SysDictType{})
	// 如果有条件搜索 下方会自动创建搜索语句
	if param.DictName != "" {
		db = db.Where("dict_name like ?", "%"+param.DictName+"%")
	}
	if param.DictType != "" {
		db = db.Where("dict_type like ?", "%"+param.DictType+"%")
	}
	if param.Status != "" {
		db = db.Where("status = ?", param.Status)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	// 计算分页信息
	page.Calculate()
	if page.OrderByColumn != "" {
		db.Order(page.OrderByColumn + " " + page.IsAsc + " ")
	}
	var tmp []SysDictType
	err = db.Limit(page.Limit).Offset(page.Offset).Find(&tmp).Error
	sysDictTypes = &tmp
	return err, sysDictTypes, total
}

func (dao *SysDictTypeDao) SelectDictTypeAll() (err error, datas *[]SysDictType) {
	err = global.GOrmDao.Find(&datas).Error
	return
}
