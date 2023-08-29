// Package dao 自动生成模板 SysDictDataDao
// @description <TODO description class purpose>
// @author
// @File: sys_dict_data
// @version 1.0.0
// @create 2023-08-20 19:08:42
package dao

import (
	"manager-gin/src/app/admin/sys/model"
	"manager-gin/src/app/admin/sys/service/sys_dict_data/view"
	"manager-gin/src/common"
	"manager-gin/src/global"
)

// SysDictDataDao 结构体

type SysDictDataDao struct{}

// Create 创建SysDictData记录
// Author
func (dao *SysDictDataDao) Create(sysDictData model.SysDictData) (err error) {
	err = global.GOrmDao.Create(&sysDictData).Error
	return err
}

// DeleteByIds 批量删除SysDictData记录
// Author
func (dao *SysDictDataDao) DeleteByIds(ids []string) (err error) {
	err = global.GOrmDao.Delete(&[]model.SysDictData{}, "id in ?", ids).Error
	return err
}

// Update 更新SysDictData记录
// Author
func (dao *SysDictDataDao) Update(sysDictData model.SysDictData) (err error) {
	err = global.GOrmDao.Updates(&sysDictData).Error
	return err
}

// Get 根据id获取SysDictData记录
// Author
func (dao *SysDictDataDao) Get(id string) (err error, sysDictData *model.SysDictData) {
	err = global.GOrmDao.Where("id = ?", id).First(&sysDictData).Error
	return
}

// Page 分页获取SysDictData记录
// Author
func (dao *SysDictDataDao) Page(param *view.SysDictDataPageView) (err error, page *common.PageInfo) {
	// 创建model
	db := global.GOrmDao.Model(&model.SysDictData{})
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
	var dataList []*model.SysDictData
	err = db.Limit(page.Limit).Offset(page.Offset).Find(&dataList).Error
	page.Rows = dataList
	return err, page
}

func (dao *SysDictDataDao) GetByType(dictType string) (err error, sysDictData []*model.SysDictData) {
	err = global.GOrmDao.Where("dict_type = ?", dictType).Find(&sysDictData).Error
	return
}
