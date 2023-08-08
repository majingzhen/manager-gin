// Package model 自动生成模板 SysJobDao
// @description <TODO description class purpose>
// @author
// @File: sys_job
// @version 1.0.0
// @create 2023-08-08 10:06:19
package model

import (
	"manager-gin/src/common"
	"manager-gin/src/global"
)

// SysJobDao 结构体

type SysJobDao struct{}

// Create 创建SysJob记录
// Author
func (dao *SysJobDao) Create(sysJob SysJob) (err error) {
	err = global.GOrmDao.Create(&sysJob).Error
	return err
}

// Delete 删除SysJob记录
// Author
func (dao *SysJobDao) Delete(id int) (err error) {
	err = global.GOrmDao.Delete(&[]SysJob{}, "id = ?", id).Error
	return err
}

// DeleteByIds 批量删除SysJob记录
// Author
func (dao *SysJobDao) DeleteByIds(ids []int) (err error) {
	err = global.GOrmDao.Delete(&[]SysJob{}, "id in ?", ids).Error
	return err
}

// Update 更新SysJob记录
// Author
func (dao *SysJobDao) Update(sysJob SysJob) (err error) {
	err = global.GOrmDao.Save(&sysJob).Error
	return err
}

// Get 根据id获取SysJob记录
// Author
func (dao *SysJobDao) Get(id int) (err error, sysJob *SysJob) {
	err = global.GOrmDao.Where("id = ?", id).First(&sysJob).Error
	return
}

// Find 分页获取SysJob记录
// Author
func (dao *SysJobDao) Find(info *common.PageInfoV2) (err error, sysJobs *[]SysJob, total int64) {
	// 创建db
	db := global.GOrmDao.Model(&SysJob{})
	// 如果有条件搜索 下方会自动创建搜索语句
	//if info.Id != "" {
	//	db = db.Where("ID = ?", info.Id)
	//}

	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var tmp []SysJob
	err = db.Limit(info.Limit).Offset(info.Offset).Find(&tmp).Error
	sysJobs = &tmp
	return err, sysJobs, total
}
