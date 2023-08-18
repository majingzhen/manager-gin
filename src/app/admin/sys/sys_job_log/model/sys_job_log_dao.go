// Package model 自动生成模板 SysJobLogDao
// @description <TODO description class purpose>
// @author
// @File: sys_job_log
// @version 1.0.0
// @create 2023-08-18 13:41:26
package model

import (
	"manager-gin/src/common"
	"manager-gin/src/global"
)

// SysJobLogDao 结构体

type SysJobLogDao struct{}

// Create 创建SysJobLog记录
// Author
func (dao *SysJobLogDao) Create(sysJobLog SysJobLog) (err error) {
	err = global.GOrmDao.Create(&sysJobLog).Error
	return err
}

// Delete 删除SysJobLog记录
// Author
func (dao *SysJobLogDao) Delete(id string) (err error) {
	err = global.GOrmDao.Delete(&[]SysJobLog{}, "id = ?", id).Error
	return err
}

// DeleteByIds 批量删除SysJobLog记录
// Author
func (dao *SysJobLogDao) DeleteByIds(ids []string) (err error) {
	err = global.GOrmDao.Delete(&[]SysJobLog{}, "id in ?", ids).Error
	return err
}

// Update 更新SysJobLog记录
// Author
func (dao *SysJobLogDao) Update(sysJobLog SysJobLog) (err error) {
	err = global.GOrmDao.Save(&sysJobLog).Error
	return err
}

// Get 根据id获取SysJobLog记录
// Author
func (dao *SysJobLogDao) Get(id string) (err error, sysJobLog *SysJobLog) {
	err = global.GOrmDao.Where("id = ?", id).First(&sysJobLog).Error
	return
}

// Find 分页获取SysJobLog记录
// Author
func (dao *SysJobLogDao) Find(info *common.PageInfoV2) (err error, sysJobLogs *[]SysJobLog, total int64) {
	// 创建db
	db := global.GOrmDao.Model(&SysJobLog{})
	// 如果有条件搜索 下方会自动创建搜索语句
	//if info.Id != "" {
	//	db = db.Where("ID = ?", info.Id)
	//}

	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var tmp []SysJobLog
	err = db.Limit(info.Limit).Offset(info.Offset).Find(&tmp).Error
	sysJobLogs = &tmp
	return err, sysJobLogs, total
}
