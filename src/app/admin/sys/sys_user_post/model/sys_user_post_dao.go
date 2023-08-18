// Package model 自动生成模板 SysUserPostDao
// @description <TODO description class purpose>
// @author
// @File: sys_user_post
// @version 1.0.0
// @create 2023-08-18 13:41:26
package model

import (
	"manager-gin/src/common"
	"manager-gin/src/global"
)

// SysUserPostDao 结构体

type SysUserPostDao struct{}

// Create 创建SysUserPost记录
// Author
func (dao *SysUserPostDao) Create(sysUserPost SysUserPost) (err error) {
	err = global.GOrmDao.Create(&sysUserPost).Error
	return err
}

// Delete 删除SysUserPost记录
// Author
func (dao *SysUserPostDao) Delete(id string) (err error) {
	err = global.GOrmDao.Delete(&[]SysUserPost{}, "id = ?", id).Error
	return err
}

// DeleteByIds 批量删除SysUserPost记录
// Author
func (dao *SysUserPostDao) DeleteByIds(ids []string) (err error) {
	err = global.GOrmDao.Delete(&[]SysUserPost{}, "id in ?", ids).Error
	return err
}

// Update 更新SysUserPost记录
// Author
func (dao *SysUserPostDao) Update(sysUserPost SysUserPost) (err error) {
	err = global.GOrmDao.Save(&sysUserPost).Error
	return err
}

// Get 根据id获取SysUserPost记录
// Author
func (dao *SysUserPostDao) Get(id string) (err error, sysUserPost *SysUserPost) {
	err = global.GOrmDao.Where("id = ?", id).First(&sysUserPost).Error
	return
}

// Find 分页获取SysUserPost记录
// Author
func (dao *SysUserPostDao) Find(info *common.PageInfoV2) (err error, sysUserPosts *[]SysUserPost, total int64) {
	// 创建db
	db := global.GOrmDao.Model(&SysUserPost{})
	// 如果有条件搜索 下方会自动创建搜索语句
	//if info.Id != "" {
	//	db = db.Where("ID = ?", info.Id)
	//}

	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var tmp []SysUserPost
	err = db.Limit(info.Limit).Offset(info.Offset).Find(&tmp).Error
	sysUserPosts = &tmp
	return err, sysUserPosts, total
}
