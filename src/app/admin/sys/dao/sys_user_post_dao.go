// Package model 自动生成模板 SysUserPostDao
// @description <TODO description class purpose>
// @author
// @File: sys_user_post
// @version 1.0.0
// @create 2023-08-21 17:37:56
package dao

import (
	"manager-gin/src/app/admin/sys/model"
	"manager-gin/src/common"
	"manager-gin/src/global"
)

// SysUserPostDao 结构体

type SysUserPostDao struct{}

// Create 创建SysUserPost记录
// Author
func (dao *SysUserPostDao) Create(sysUserPost model.SysUserPost) (err error) {
	err = global.GOrmDao.Create(&sysUserPost).Error
	return err
}

// DeleteByIds 批量删除SysUserPost记录
// Author
func (dao *SysUserPostDao) DeleteByIds(ids []string) (err error) {
	err = global.GOrmDao.Delete(&[]model.SysUserPost{}, "id in ?", ids).Error
	return err
}

// Update 更新SysUserPost记录
// Author
func (dao *SysUserPostDao) Update(sysUserPost model.SysUserPost) (err error) {
	err = global.GOrmDao.Updates(&sysUserPost).Error
	return err
}

// Get 根据id获取SysUserPost记录
// Author
func (dao *SysUserPostDao) Get(id string) (err error, sysUserPost *model.SysUserPost) {
	err = global.GOrmDao.Where("id = ?", id).First(&sysUserPost).Error
	return
}

// Page 分页获取SysUserPost记录
// Author
func (dao *SysUserPostDao) Page(param *model.SysUserPost, page *common.PageInfo) (err error, datas *[]model.SysUserPost, total int64) {
	// 创建model
	db := global.GOrmDao.Model(&model.SysUserPost{})
	// 如果有条件搜索 下方会自动创建搜索语句
	//if param.Id != "" {
	//	model = model.Where("ID = ?", info.Id)
	//}
	if err = db.Count(&total).Error; err != nil {
		return
	}
	// 计算分页信息
	page.Calculate()
	// 生成排序信息
	if page.OrderByColumn != "" {
		db.Order(page.OrderByColumn + " " + page.IsAsc + " ")
	}
	var tmp []model.SysUserPost
	err = db.Limit(page.Limit).Offset(page.Offset).Find(&tmp).Error
	datas = &tmp
	return err, datas, total
}

// List 获取SysUserPost记录
// Author
func (dao *SysUserPostDao) List(data *model.SysUserPost) (err error, datas *[]model.SysUserPost) {
	var rows []model.SysUserPost
	db := global.GOrmDao.Model(&model.SysUserPost{})
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
func (dao *SysUserPostDao) CreateBatch(posts []model.SysUserPost) error {
	return global.GOrmDao.Create(&posts).Error
}

// DeleteByUserIds 根据用户id批量删除SysUserPost记录
func (dao *SysUserPostDao) DeleteByUserIds(ids []string) error {
	return global.GOrmDao.Delete(&model.SysUserPost{}, "user_id in ?", ids).Error
}
