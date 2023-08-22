// Package model 自动生成模板 SysUserPostDao
// @description <TODO description class purpose>
// @author
// @File: sys_user_post
// @version 1.0.0
// @create 2023-08-21 17:37:56
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

// Page 分页获取SysUserPost记录
// Author
func (dao *SysUserPostDao) Page(param *SysUserPost, page *common.PageInfo) (err error, datas *[]SysUserPost, total int64) {
	// 创建model
	model := global.GOrmDao.Model(&SysUserPost{})
	// 如果有条件搜索 下方会自动创建搜索语句
	//if param.Id != "" {
	//	model = model.Where("ID = ?", info.Id)
	//}
	if err = model.Count(&total).Error; err != nil {
		return
	}
	// 计算分页信息
	page.Calculate()
	// 生成排序信息
	if page.OrderByColumn != "" {
		model.Order(page.OrderByColumn + " " + page.IsAsc + " ")
	}
	var tmp []SysUserPost
	err = model.Limit(page.Limit).Offset(page.Offset).Find(&tmp).Error
	datas = &tmp
	return err, datas, total
}

// List 获取SysUserPost记录
// Author
func (dao *SysUserPostDao) List(data *SysUserPost) (err error, datas *[]SysUserPost) {
	var rows []SysUserPost
	db := global.GOrmDao.Model(&SysUserPost{})
	// TODO 输入查询条件
	//if data.Id != "" {
	//    db.Where("id = ?", data.Id)
	//}
	db.Order("create_time desc")
	err = db.Find(&rows).Error
	datas = &rows
	return err, datas
}

// CreateBatch 批量创建SysUserPost记录
func (dao *SysUserPostDao) CreateBatch(posts []SysUserPost) error {
	return global.GOrmDao.Create(&posts).Error
}
