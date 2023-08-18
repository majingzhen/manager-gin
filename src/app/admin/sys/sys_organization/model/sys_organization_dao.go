// Package model 自动生成模板 SysOrganizationDao
// @description <TODO description class purpose>
// @author
// @File: sys_organization
// @version 1.0.0
// @create 2023-08-18 14:00:53
package model

import (
	"manager-gin/src/common"
	"manager-gin/src/global"
)

// SysOrganizationDao 结构体

type SysOrganizationDao struct{}

// Create 创建SysOrganization记录
// Author
func (dao *SysOrganizationDao) Create(sysOrganization SysOrganization) (err error) {
	err = global.GOrmDao.Create(&sysOrganization).Error
	return err
}

// Delete 删除SysOrganization记录
// Author
func (dao *SysOrganizationDao) Delete(id string) (err error) {
	err = global.GOrmDao.Delete(&[]SysOrganization{}, "id = ?", id).Error
	return err
}

// DeleteByIds 批量删除SysOrganization记录
// Author
func (dao *SysOrganizationDao) DeleteByIds(ids []string) (err error) {
	err = global.GOrmDao.Delete(&[]SysOrganization{}, "id in ?", ids).Error
	return err
}

// Update 更新SysOrganization记录
// Author
func (dao *SysOrganizationDao) Update(sysOrganization SysOrganization) (err error) {
	err = global.GOrmDao.Save(&sysOrganization).Error
	return err
}

// Get 根据id获取SysOrganization记录
// Author
func (dao *SysOrganizationDao) Get(id string) (err error, sysOrganization *SysOrganization) {
	err = global.GOrmDao.Where("id = ?", id).First(&sysOrganization).Error
	return
}

// Find 分页获取SysOrganization记录
// Author
func (dao *SysOrganizationDao) Find(info *common.PageInfoV2) (err error, sysOrganizations *[]SysOrganization, total int64) {
	// 创建db
	db := global.GOrmDao.Model(&SysOrganization{})
	// 如果有条件搜索 下方会自动创建搜索语句
	//if info.Id != "" {
	//	db = db.Where("ID = ?", info.Id)
	//}

	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var tmp []SysOrganization
	err = db.Limit(info.Limit).Offset(info.Offset).Find(&tmp).Error
	sysOrganizations = &tmp
	return err, sysOrganizations, total
}
