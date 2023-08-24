// Package model 自动生成模板 SysPostDao
// @description <TODO description class purpose>
// @author
// @File: sys_post
// @version 1.0.0
// @create 2023-08-21 17:37:56
package model

import (
	"manager-gin/src/common"
	"manager-gin/src/global"
)

// SysPostDao 结构体

type SysPostDao struct{}

// Create 创建SysPost记录
// Author
func (dao *SysPostDao) Create(sysPost SysPost) (err error) {
	err = global.GOrmDao.Create(&sysPost).Error
	return err
}

// DeleteByIds 批量删除SysPost记录
// Author
func (dao *SysPostDao) DeleteByIds(ids []string) (err error) {
	err = global.GOrmDao.Delete(&[]SysPost{}, "id in ?", ids).Error
	return err
}

// Update 更新SysPost记录
// Author
func (dao *SysPostDao) Update(sysPost SysPost) (err error) {
	err = global.GOrmDao.Updates(&sysPost).Error
	return err
}

// Get 根据id获取SysPost记录
// Author
func (dao *SysPostDao) Get(id string) (err error, sysPost *SysPost) {
	err = global.GOrmDao.Where("id = ?", id).First(&sysPost).Error
	return
}

// Page 分页获取SysPost记录
// Author
func (dao *SysPostDao) Page(param *SysPost, page *common.PageInfo) (err error, datas []*SysPost, total int64) {
	// 创建model
	model := global.GOrmDao.Model(&SysPost{})
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
	err = model.Limit(page.Limit).Offset(page.Offset).Find(datas).Error
	return err, datas, total
}

// List 获取SysPost记录
// Author
func (dao *SysPostDao) List(data *SysPost) (err error, datas []*SysPost) {
	db := global.GOrmDao.Model(&SysPost{})
	if data.PostName != "" {
		db.Where("post_name like ?", "%"+data.PostName+"%")
	}
	if data.PostCode != "" {
		db.Where("post_code like ?", "%"+data.PostCode+"%")
	}
	if data.Status != "" {
		db.Where("status = ?", data.Status)
	}
	db.Order("create_time desc")
	err = db.Find(&datas).Error
	return err, datas
}

// CheckPostNameUnique 校验岗位名称是否唯一
// Author
func (dao *SysPostDao) CheckPostNameUnique(postName string) (err error, count int64) {
	var sysPost SysPost
	err = global.GOrmDao.Model(&SysPost{}).Where("post_name = ?", postName).First(&sysPost).Count(&count).Error
	return
}

// CheckPostCodeUnique 校验岗位编码是否唯一
// Author
func (dao *SysPostDao) CheckPostCodeUnique(postCode string) (err error, count int64) {
	var sysPost SysPost
	err = global.GOrmDao.Model(&SysPost{}).Where("post_code = ?", postCode).First(&sysPost).Count(&count).Error
	return
}

func (dao *SysPostDao) CheckPostExistUser(postId string) (err error, count int64) {
	err = global.GOrmDao.Table("sys_user_post").Where("post_id = ?", postId).Count(&count).Error
	return
}

// SelectPostListByUserId 根据用户ID查询岗位
func (dao *SysPostDao) SelectPostListByUserId(userId string) (error, []*SysPost) {
	db := global.GOrmDao.Table("sys_post p")
	db.Joins("left join sys_user_post up on p.id = up.post_id")
	db.Joins("left join sys_user u on u.id = up.user_id")
	db.Where("u.id = ?", userId)
	db.Select("distinct p.id, p.post_code, p.post_name, p.post_sort, p.status, p.create_by, p.create_time, p.remark ")
	var res []*SysPost
	err := db.Find(&res).Error
	return err, res
}
