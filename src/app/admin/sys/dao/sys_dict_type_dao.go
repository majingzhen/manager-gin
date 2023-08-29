// Package dao 自动生成模板 SysDictTypeDao
// @description <TODO description class purpose>
// @author
// @File: sys_dict_type
// @version 1.0.0
// @create 2023-08-18 13:41:26
package dao

import (
	"manager-gin/src/app/admin/sys/model"
	"manager-gin/src/app/admin/sys/service/sys_dict_type/view"
	"manager-gin/src/common"
	"manager-gin/src/global"
)

// SysDictTypeDao 结构体

type SysDictTypeDao struct{}

// Create 创建SysDictType记录
// Author
func (dao *SysDictTypeDao) Create(sysDictType model.SysDictType) (err error) {
	err = global.GOrmDao.Create(&sysDictType).Error
	return err
}

// Delete 删除SysDictType记录
// Author
func (dao *SysDictTypeDao) Delete(id string) (err error) {
	err = global.GOrmDao.Delete(&[]model.SysDictType{}, "id = ?", id).Error
	return err
}

// DeleteByIds 批量删除SysDictType记录
// Author
func (dao *SysDictTypeDao) DeleteByIds(ids []string) (err error) {
	err = global.GOrmDao.Delete(&[]model.SysDictType{}, "id in ?", ids).Error
	return err
}

// Update 更新SysDictType记录
// Author
func (dao *SysDictTypeDao) Update(sysDictType model.SysDictType) (err error) {
	err = global.GOrmDao.Updates(&sysDictType).Error
	return err
}

// Get 根据id获取SysDictType记录
// Author
func (dao *SysDictTypeDao) Get(id string) (err error, sysDictType *model.SysDictType) {
	err = global.GOrmDao.Where("id = ?", id).First(&sysDictType).Error
	return
}

// Page 分页获取SysDictType记录
// Author
func (dao *SysDictTypeDao) Page(param *view.SysDictTypePageView) (err error, page *common.PageInfo) {
	// 创建db
	db := global.GOrmDao.Model(&model.SysDictType{})
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
	page = common.CreatePageInfo(param.PageNum, param.PageSize)
	err = db.Count(&page.Total).Error
	if err != nil {
		return
	}
	if param.OrderByColumn != "" {
		db.Order(param.OrderByColumn + " " + param.IsAsc + " ")
	}
	var dataList []*model.SysDictType
	err = db.Limit(page.Limit).Offset(page.Offset).Find(&dataList).Error
	page.Rows = dataList
	return err, page
}

func (dao *SysDictTypeDao) SelectDictTypeAll() (err error, datas []*model.SysDictType) {
	err = global.GOrmDao.Find(&datas).Error
	return
}
