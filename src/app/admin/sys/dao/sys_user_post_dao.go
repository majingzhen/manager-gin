// Package dao 自动生成模板 SysUserPostDao
// @description <TODO description class purpose>
// @author
// @File: sys_user_post
// @version 1.0.0
// @create 2023-08-21 17:37:56
package dao

import (
	"gorm.io/gorm"
	"manager-gin/src/app/admin/sys/model"
	"manager-gin/src/global"
)

// SysUserPostDao 结构体

type SysUserPostDao struct{}

// CreateBatch 批量创建SysUserPost记录
func (dao *SysUserPostDao) CreateBatch(tx *gorm.DB, posts []model.SysUserPost) error {
	return global.GOrmDao.Create(&posts).Error
}

// DeleteByUserIds 根据用户id批量删除SysUserPost记录
func (dao *SysUserPostDao) DeleteByUserIds(tx *gorm.DB, ids []string) error {
	return global.GOrmDao.Delete(&model.SysUserPost{}, "user_id in ?", ids).Error
}
