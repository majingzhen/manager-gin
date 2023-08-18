// Package model 自动生成模板 SysOperLogDao
// @description <TODO description class purpose>
// @author
// @File: sys_oper_log
// @version 1.0.0
// @create 2023-08-18 13:41:26
package model

import (
	"manager-gin/src/common"
	"manager-gin/src/global"
)

// SysOperLogDao 结构体

type SysOperLogDao struct{}

// Create 创建SysOperLog记录
// Author
func (dao *SysOperLogDao) Create(sysOperLog SysOperLog) (err error) {
	err = global.GOrmDao.Create(&sysOperLog).Error
	return err
}

// Delete 删除SysOperLog记录
// Author
func (dao *SysOperLogDao) Delete(id string) (err error) {
	err = global.GOrmDao.Delete(&[]SysOperLog{}, "id = ?", id).Error
	return err
}

// DeleteByIds 批量删除SysOperLog记录
// Author
func (dao *SysOperLogDao) DeleteByIds(ids []string) (err error) {
	err = global.GOrmDao.Delete(&[]SysOperLog{}, "id in ?", ids).Error
	return err
}

// Update 更新SysOperLog记录
// Author
func (dao *SysOperLogDao) Update(sysOperLog SysOperLog) (err error) {
	err = global.GOrmDao.Save(&sysOperLog).Error
	return err
}

// Get 根据id获取SysOperLog记录
// Author
func (dao *SysOperLogDao) Get(id string) (err error, sysOperLog *SysOperLog) {
	err = global.GOrmDao.Where("id = ?", id).First(&sysOperLog).Error
	return
}

// Find 分页获取SysOperLog记录
// Author
func (dao *SysOperLogDao) Find(info *common.PageInfoV2) (err error, sysOperLogs *[]SysOperLog, total int64) {
	// 创建db
	db := global.GOrmDao.Model(&SysOperLog{})
	// 如果有条件搜索 下方会自动创建搜索语句
	//if info.Id != "" {
	//	db = db.Where("ID = ?", info.Id)
	//}

	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var tmp []SysOperLog
	err = db.Limit(info.Limit).Offset(info.Offset).Find(&tmp).Error
	sysOperLogs = &tmp
	return err, sysOperLogs, total
}
