// Package model 自动生成模板 SysConfigDao
// @description <TODO description class purpose>
// @author
// @File: sys_config
// @version 1.0.0
// @create 2023-08-08 10:06:19
package model

import (
	"manager-gin/src/common"
	"manager-gin/src/global"
)

// SysConfigDao 结构体

type SysConfigDao struct{}

// Create 创建SysConfig记录
// Author
func (dao *SysConfigDao) Create(sysConfig SysConfig) (err error) {
	err = global.GOrmDao.Create(&sysConfig).Error
	return err
}

// Delete 删除SysConfig记录
// Author
func (dao *SysConfigDao) Delete(id int) (err error) {
	err = global.GOrmDao.Delete(&[]SysConfig{}, "id = ?", id).Error
	return err
}

// DeleteByIds 批量删除SysConfig记录
// Author
func (dao *SysConfigDao) DeleteByIds(ids []int) (err error) {
	err = global.GOrmDao.Delete(&[]SysConfig{}, "id in ?", ids).Error
	return err
}

// Update 更新SysConfig记录
// Author
func (dao *SysConfigDao) Update(sysConfig SysConfig) (err error) {
	err = global.GOrmDao.Save(&sysConfig).Error
	return err
}

// Get 根据id获取SysConfig记录
// Author
func (dao *SysConfigDao) Get(id int) (err error, sysConfig *SysConfig) {
	err = global.GOrmDao.Where("id = ?", id).First(&sysConfig).Error
	return
}

// Find 分页获取SysConfig记录
// Author
func (dao *SysConfigDao) Find(info *common.PageInfoV2) (err error, sysConfigs *[]SysConfig, total int64) {
	// 创建db
	db := global.GOrmDao.Model(&SysConfig{})
	// 如果有条件搜索 下方会自动创建搜索语句
	//if info.Id != "" {
	//	db = db.Where("ID = ?", info.Id)
	//}

	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var tmp []SysConfig
	err = db.Limit(info.Limit).Offset(info.Offset).Find(&tmp).Error
	sysConfigs = &tmp
	return err, sysConfigs, total
}
