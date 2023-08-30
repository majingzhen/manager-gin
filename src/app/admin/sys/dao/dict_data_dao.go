// Package dao 自动生成模板 DictDataDao
// @description <TODO description class purpose>
// @author
// @File: dict_data
// @version 1.0.0
// @create 2023-08-20 19:08:42
package dao

import (
	"manager-gin/src/app/admin/sys/model"
	"manager-gin/src/app/admin/sys/service/dict_data/view"
	"manager-gin/src/common"
	"manager-gin/src/global"
)

// DictDataDao 结构体

type DictDataDao struct{}

// Create 创建DictData记录
// Author
func (dao *DictDataDao) Create(sysDictData model.DictData) (err error) {
	err = global.GOrmDao.Create(&sysDictData).Error
	return err
}

// DeleteByIds 批量删除DictData记录
// Author
func (dao *DictDataDao) DeleteByIds(ids []string) (err error) {
	err = global.GOrmDao.Delete(&[]model.DictData{}, "id in ?", ids).Error
	return err
}

// Update 更新DictData记录
// Author
func (dao *DictDataDao) Update(sysDictData model.DictData) (err error) {
	err = global.GOrmDao.Updates(&sysDictData).Error
	return err
}

// Get 根据id获取DictData记录
// Author
func (dao *DictDataDao) Get(id string) (err error, sysDictData *model.DictData) {
	err = global.GOrmDao.Where("id = ?", id).First(&sysDictData).Error
	return
}

// Page 分页获取DictData记录
// Author
func (dao *DictDataDao) Page(param *view.DictDataPageView) (err error, page *common.PageInfo) {
	// 创建model
	db := global.GOrmDao.Model(&model.DictData{})
	// 如果有条件搜索 下方会自动创建搜索语句
	if param.DictType != "" {
		db.Where("dict_type = ?", param.DictType)
	}
	if param.DictLabel != "" {
		db.Where("dict_label like ?", "%"+param.DictLabel+"%")
	}
	if param.Status != "" {
		db.Where("status = ?", param.Status)
	}
	page = common.CreatePageInfo(param.PageNum, param.PageSize)
	err = db.Count(&page.Total).Error
	if err != nil {
		return
	}
	// 生成排序信息
	if param.OrderByColumn != "" {
		db.Order(param.OrderByColumn + " " + param.IsAsc + " ")
	}
	var dataList []*model.DictData
	err = db.Limit(page.Limit).Offset(page.Offset).Find(&dataList).Error
	page.Rows = dataList
	return err, page
}

func (dao *DictDataDao) GetByType(dictType string) (err error, sysDictData []*model.DictData) {
	err = global.GOrmDao.Where("dict_type = ?", dictType).Find(&sysDictData).Error
	return
}
