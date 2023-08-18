// Package model 自动生成模板 SysLogininforDao
// @description <TODO description class purpose>
// @author
// @File: sys_logininfor
// @version 1.0.0
// @create 2023-08-18 13:41:26
package model

import (
	"manager-gin/src/common"
	"manager-gin/src/global"
)

// SysLogininforDao 结构体

type SysLogininforDao struct{}

// Create 创建SysLogininfor记录
// Author
func (dao *SysLogininforDao) Create(sysLogininfor SysLogininfor) (err error) {
	err = global.GOrmDao.Create(&sysLogininfor).Error
	return err
}

// Delete 删除SysLogininfor记录
// Author
func (dao *SysLogininforDao) Delete(id string) (err error) {
	err = global.GOrmDao.Delete(&[]SysLogininfor{}, "id = ?", id).Error
	return err
}

// DeleteByIds 批量删除SysLogininfor记录
// Author
func (dao *SysLogininforDao) DeleteByIds(ids []string) (err error) {
	err = global.GOrmDao.Delete(&[]SysLogininfor{}, "id in ?", ids).Error
	return err
}

// Update 更新SysLogininfor记录
// Author
func (dao *SysLogininforDao) Update(sysLogininfor SysLogininfor) (err error) {
	err = global.GOrmDao.Save(&sysLogininfor).Error
	return err
}

// Get 根据id获取SysLogininfor记录
// Author
func (dao *SysLogininforDao) Get(id string) (err error, sysLogininfor *SysLogininfor) {
	err = global.GOrmDao.Where("id = ?", id).First(&sysLogininfor).Error
	return
}

// Find 分页获取SysLogininfor记录
// Author
func (dao *SysLogininforDao) Find(info *common.PageInfoV2) (err error, sysLogininfors *[]SysLogininfor, total int64) {
	// 创建db
	db := global.GOrmDao.Model(&SysLogininfor{})
	// 如果有条件搜索 下方会自动创建搜索语句
	//if info.Id != "" {
	//	db = db.Where("ID = ?", info.Id)
	//}

	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var tmp []SysLogininfor
	err = db.Limit(info.Limit).Offset(info.Offset).Find(&tmp).Error
	sysLogininfors = &tmp
	return err, sysLogininfors, total
}
