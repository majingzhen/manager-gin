// Package model 自动生成模板 SysDictDataDao
// @description <TODO description class purpose>
// @author
// @File: sys_dict_data
// @version 1.0.0
// @create 2023-08-18 13:41:26
package model

import (
	"manager-gin/src/common"
	"manager-gin/src/global"
)

// SysDictDataDao 结构体

type SysDictDataDao struct{}

// Create 创建SysDictData记录
// Author
func (dao *SysDictDataDao) Create(sysDictData SysDictData) (err error) {
	err = global.GOrmDao.Create(&sysDictData).Error
	return err
}

// Delete 删除SysDictData记录
// Author
func (dao *SysDictDataDao) Delete(id string) (err error) {
	err = global.GOrmDao.Delete(&[]SysDictData{}, "id = ?", id).Error
	return err
}

// DeleteByIds 批量删除SysDictData记录
// Author
func (dao *SysDictDataDao) DeleteByIds(ids []string) (err error) {
	err = global.GOrmDao.Delete(&[]SysDictData{}, "id in ?", ids).Error
	return err
}

// Update 更新SysDictData记录
// Author
func (dao *SysDictDataDao) Update(sysDictData SysDictData) (err error) {
	err = global.GOrmDao.Save(&sysDictData).Error
	return err
}

// Get 根据id获取SysDictData记录
// Author
func (dao *SysDictDataDao) Get(id string) (err error, sysDictData *SysDictData) {
	err = global.GOrmDao.Where("id = ?", id).First(&sysDictData).Error
	return
}

// Find 分页获取SysDictData记录
// Author
func (dao *SysDictDataDao) Find(info *common.PageInfoV2) (err error, sysDictDatas *[]SysDictData, total int64) {
	// 创建db
	db := global.GOrmDao.Model(&SysDictData{})
	// 如果有条件搜索 下方会自动创建搜索语句
	//if info.Id != "" {
	//	db = db.Where("ID = ?", info.Id)
	//}

	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var tmp []SysDictData
	err = db.Limit(info.Limit).Offset(info.Offset).Find(&tmp).Error
	sysDictDatas = &tmp
	return err, sysDictDatas, total
}
