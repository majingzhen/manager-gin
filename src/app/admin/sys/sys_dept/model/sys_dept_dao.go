// Package model 自动生成模板 SysDeptDao
// @description <TODO description class purpose>
// @author
// @File: sys_dept
// @version 1.0.0
// @create 2023-08-08 10:06:19
package model

import (
	"manager-gin/src/common"
	"manager-gin/src/global"
)

// SysDeptDao 结构体

type SysDeptDao struct{}

// Create 创建SysDept记录
// Author
func (dao *SysDeptDao) Create(sysDept SysDept) (err error) {
	err = global.GOrmDao.Create(&sysDept).Error
	return err
}

// Delete 删除SysDept记录
// Author
func (dao *SysDeptDao) Delete(id int) (err error) {
	err = global.GOrmDao.Delete(&[]SysDept{}, "id = ?", id).Error
	return err
}

// DeleteByIds 批量删除SysDept记录
// Author
func (dao *SysDeptDao) DeleteByIds(ids []int) (err error) {
	err = global.GOrmDao.Delete(&[]SysDept{}, "id in ?", ids).Error
	return err
}

// Update 更新SysDept记录
// Author
func (dao *SysDeptDao) Update(sysDept SysDept) (err error) {
	err = global.GOrmDao.Save(&sysDept).Error
	return err
}

// Get 根据id获取SysDept记录
// Author
func (dao *SysDeptDao) Get(id int) (err error, sysDept *SysDept) {
	err = global.GOrmDao.Where("id = ?", id).First(&sysDept).Error
	return
}

// Find 分页获取SysDept记录
// Author
func (dao *SysDeptDao) Find(info *common.PageInfoV2) (err error, sysDepts *[]SysDept, total int64) {
	// 创建db
	db := global.GOrmDao.Model(&SysDept{})
	// 如果有条件搜索 下方会自动创建搜索语句
	//if info.Id != "" {
	//	db = db.Where("ID = ?", info.Id)
	//}

	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var tmp []SysDept
	err = db.Limit(info.Limit).Offset(info.Offset).Find(&tmp).Error
	sysDepts = &tmp
	return err, sysDepts, total
}
