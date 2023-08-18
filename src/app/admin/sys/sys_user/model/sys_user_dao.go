// Package model 自动生成模板 SysUserDao
// @description <TODO description class purpose>
// @author
// @File: sys_user
// @version 1.0.0
// @create 2023-08-18 14:02:24
package model

import (
	"manager-gin/src/common"
	"manager-gin/src/global"
)

// SysUserDao 结构体

type SysUserDao struct{}

// Create 创建SysUser记录
// Author
func (dao *SysUserDao) Create(sysUser SysUser) (err error) {
	err = global.GOrmDao.Create(&sysUser).Error
	return err
}

// Delete 删除SysUser记录
// Author
func (dao *SysUserDao) Delete(id string) (err error) {
	err = global.GOrmDao.Delete(&[]SysUser{}, "id = ?", id).Error
	return err
}

// DeleteByIds 批量删除SysUser记录
// Author
func (dao *SysUserDao) DeleteByIds(ids []string) (err error) {
	err = global.GOrmDao.Delete(&[]SysUser{}, "id in ?", ids).Error
	return err
}

// Update 更新SysUser记录
// Author
func (dao *SysUserDao) Update(sysUser SysUser) (err error) {
	err = global.GOrmDao.Save(&sysUser).Error
	return err
}

// Get 根据id获取SysUser记录
// Author
func (dao *SysUserDao) Get(id string) (err error, sysUser *SysUser) {
	err = global.GOrmDao.Where("id = ?", id).First(&sysUser).Error
	return
}

// Find 分页获取SysUser记录
// Author
func (dao *SysUserDao) Find(info *common.PageInfoV2) (err error, sysUsers *[]SysUser, total int64) {
	// 创建db
	db := global.GOrmDao.Model(&SysUser{})
	// 如果有条件搜索 下方会自动创建搜索语句
	//if info.Id != "" {
	//	db = db.Where("ID = ?", info.Id)
	//}

	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var tmp []SysUser
	err = db.Limit(info.Limit).Offset(info.Offset).Find(&tmp).Error
	sysUsers = &tmp
	return err, sysUsers, total
}

func (dao *SysUserDao) GetByUserName(name string) (err error, sysUser *SysUser) {
	err = global.GOrmDao.Where("user_name = ?", name).First(&sysUser).Error
	return
}
