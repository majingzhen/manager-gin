// Package model 自动生成模板 SysDictDataDao
// @description <TODO description class purpose>
// @author
// @File: sys_dict_data
// @version 1.0.0
// @create 2023-08-20 19:08:42
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

// List 分页获取SysDictData记录
// Author
func (dao *SysDictDataDao) List(param *SysDictData, page *common.PageInfo) (err error, datas *[]SysDictData, total int64) {
	// 创建model
	model := global.GOrmDao.Model(&SysDictData{})
	// 如果有条件搜索 下方会自动创建搜索语句
	if param.DictType != "" {
		model = model.Where("dict_type = ?", param.DictType)
	}
	if param.DictLabel != "" {
		model = model.Where("dict_label like ?", "%"+param.DictLabel+"%")
	}
	if param.Status != "" {
		model = model.Where("status = ?", param.Status)
	}

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
	var tmp []SysDictData
	err = model.Limit(page.Limit).Offset(page.Offset).Find(&tmp).Error
	datas = &tmp
	return err, datas, total
}

func (dao *SysDictDataDao) GetByType(dictType string) (err error, sysDictData *[]SysDictData) {
	err = global.GOrmDao.Where("dict_type = ?", dictType).Find(&sysDictData).Error
	return
}
