// Package model 自动生成模板 SysNoticeDao
// @description <TODO description class purpose>
// @author
// @File: sys_notice
// @version 1.0.0
// @create 2023-08-08 10:06:19
package model

import (
	"manager-gin/src/common"
	"manager-gin/src/global"
)

// SysNoticeDao 结构体

type SysNoticeDao struct{}

// Create 创建SysNotice记录
// Author
func (dao *SysNoticeDao) Create(sysNotice SysNotice) (err error) {
	err = global.GOrmDao.Create(&sysNotice).Error
	return err
}

// Delete 删除SysNotice记录
// Author
func (dao *SysNoticeDao) Delete(id int) (err error) {
	err = global.GOrmDao.Delete(&[]SysNotice{}, "id = ?", id).Error
	return err
}

// DeleteByIds 批量删除SysNotice记录
// Author
func (dao *SysNoticeDao) DeleteByIds(ids []int) (err error) {
	err = global.GOrmDao.Delete(&[]SysNotice{}, "id in ?", ids).Error
	return err
}

// Update 更新SysNotice记录
// Author
func (dao *SysNoticeDao) Update(sysNotice SysNotice) (err error) {
	err = global.GOrmDao.Save(&sysNotice).Error
	return err
}

// Get 根据id获取SysNotice记录
// Author
func (dao *SysNoticeDao) Get(id int) (err error, sysNotice *SysNotice) {
	err = global.GOrmDao.Where("id = ?", id).First(&sysNotice).Error
	return
}

// Find 分页获取SysNotice记录
// Author
func (dao *SysNoticeDao) Find(info *common.PageInfoV2) (err error, sysNotices *[]SysNotice, total int64) {
	// 创建db
	db := global.GOrmDao.Model(&SysNotice{})
	// 如果有条件搜索 下方会自动创建搜索语句
	//if info.Id != "" {
	//	db = db.Where("ID = ?", info.Id)
	//}

	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var tmp []SysNotice
	err = db.Limit(info.Limit).Offset(info.Offset).Find(&tmp).Error
	sysNotices = &tmp
	return err, sysNotices, total
}
