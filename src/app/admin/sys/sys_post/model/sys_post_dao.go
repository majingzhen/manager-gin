// Package model 自动生成模板 SysPostDao
// @description <TODO description class purpose>
// @author
// @File: sys_post
// @version 1.0.0
// @create 2023-08-18 13:41:26
package model

import (
	"manager-gin/src/common"
	"manager-gin/src/global"
)

// SysPostDao 结构体

type SysPostDao struct{}

// Create 创建SysPost记录
// Author
func (dao *SysPostDao) Create(sysPost SysPost) (err error) {
	err = global.GOrmDao.Create(&sysPost).Error
	return err
}

// Delete 删除SysPost记录
// Author
func (dao *SysPostDao) Delete(id string) (err error) {
	err = global.GOrmDao.Delete(&[]SysPost{}, "id = ?", id).Error
	return err
}

// DeleteByIds 批量删除SysPost记录
// Author
func (dao *SysPostDao) DeleteByIds(ids []string) (err error) {
	err = global.GOrmDao.Delete(&[]SysPost{}, "id in ?", ids).Error
	return err
}

// Update 更新SysPost记录
// Author
func (dao *SysPostDao) Update(sysPost SysPost) (err error) {
	err = global.GOrmDao.Save(&sysPost).Error
	return err
}

// Get 根据id获取SysPost记录
// Author
func (dao *SysPostDao) Get(id string) (err error, sysPost *SysPost) {
	err = global.GOrmDao.Where("id = ?", id).First(&sysPost).Error
	return
}

// Find 分页获取SysPost记录
// Author
func (dao *SysPostDao) Find(info *common.PageInfoV2) (err error, sysPosts *[]SysPost, total int64) {
	// 创建db
	db := global.GOrmDao.Model(&SysPost{})
	// 如果有条件搜索 下方会自动创建搜索语句
	//if info.Id != "" {
	//	db = db.Where("ID = ?", info.Id)
	//}

	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var tmp []SysPost
	err = db.Limit(info.Limit).Offset(info.Offset).Find(&tmp).Error
	sysPosts = &tmp
	return err, sysPosts, total
}
