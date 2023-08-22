// Package model 自动生成模板 SysUserRoleDao
// @description <TODO description class purpose>
// @author
// @File: sys_user_role
// @version 1.0.0
// @create 2023-08-21 17:37:56
package model

import (
	"manager-gin/src/common"
	"manager-gin/src/global"
)

// SysUserRoleDao 结构体

type SysUserRoleDao struct{}

// Create 创建SysUserRole记录
// Author
func (dao *SysUserRoleDao) Create(sysUserRole SysUserRole) (err error) {
	err = global.GOrmDao.Create(&sysUserRole).Error
	return err
}

// DeleteByIds 批量删除SysUserRole记录
// Author
func (dao *SysUserRoleDao) DeleteByIds(ids []string) (err error) {
	err = global.GOrmDao.Delete(&[]SysUserRole{}, "id in ?", ids).Error
	return err
}

// Update 更新SysUserRole记录
// Author
func (dao *SysUserRoleDao) Update(sysUserRole SysUserRole) (err error) {
	err = global.GOrmDao.Updates(&sysUserRole).Error
	return err
}

// Get 根据id获取SysUserRole记录
// Author
func (dao *SysUserRoleDao) Get(id string) (err error, sysUserRole *SysUserRole) {
	err = global.GOrmDao.Where("id = ?", id).First(&sysUserRole).Error
	return
}

// Page 分页获取SysUserRole记录
// Author
func (dao *SysUserRoleDao) Page(param *SysUserRole, page *common.PageInfo) (err error, datas *[]SysUserRole, total int64) {
	// 创建model
	model := global.GOrmDao.Model(&SysUserRole{})
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
	var tmp []SysUserRole
	err = model.Limit(page.Limit).Offset(page.Offset).Find(&tmp).Error
	datas = &tmp
	return err, datas, total
}

// List 获取SysUserRole记录
// Author
func (dao *SysUserRoleDao) List(data *SysUserRole) (err error, datas *[]SysUserRole) {
	var rows []SysUserRole
	db := global.GOrmDao.Model(&SysUserRole{})
	// TODO 输入查询条件
	//if data.Id != "" {
	//    db.Where("id = ?", data.Id)
	//}
	db.Order("create_time desc")
	err = db.Find(&rows).Error
	datas = &rows
	return err, datas
}

// CreateBatch 批量创建SysUserRole记录
func (dao *SysUserRoleDao) CreateBatch(roles []SysUserRole) error {
	return global.GOrmDao.Create(&roles).Error
}

func (dao *SysUserRoleDao) DeleteByUserIds(ids []string) error {
	return global.GOrmDao.Delete(&[]SysUserRole{}, "user_id in ?", ids).Error
}
