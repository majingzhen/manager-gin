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

// Page 分页获取SysUser记录
// Author
func (dao *SysUserDao) Page(param *SysUser, page *common.PageInfo) (err error, datas *[]SysUser, total int64) {
	// 创建model
	model := global.GOrmDao.Table("sys_user u")
	model.Select("u.id, u.dept_id, u.nick_name, u.user_name, u.email, u.avatar, u.phone_number, u.sex, u.status, u.login_ip, u.login_date, u.create_by, u.create_time, u.remark")
	model.Joins("left join sys_dept d on u.dept_id = d.id")
	// 如果有条件搜索 下方会自动创建搜索语句
	if param.Id != "" {
		model = model.Where("ID = ?", param.Id)
	}
	if param.UserName != "" {
		model = model.Where("u.user_name like ?", "%"+param.UserName+"%")
	}
	if param.PhoneNumber != "" {
		model = model.Where("u.phone_number like ?", "%"+param.PhoneNumber+"%")
	}
	if param.Status != "" {
		model = model.Where("u.status = ?", param.Status)
	}
	if param.DeptId != "" {
		model = model.Where("u.dept_id = ?", param.DeptId)
	}
	if param.DataScopeSql != "" {
		model = model.Where(param.DataScopeSql)
	}
	if err = model.Count(&total).Error; err != nil {
		return
	}
	// 计算分页信息
	page.Calculate()
	// 生成排序信息
	if page.OrderByColumn != "" {
		model.Order(page.OrderByColumn + " " + page.IsAsc + " ")
	}
	var tmp []SysUser
	err = model.Limit(page.Limit).Offset(page.Offset).Find(&tmp).Error
	datas = &tmp
	return err, datas, total
}

// List 获取SysUser记录
// Author
func (dao *SysUserDao) List(data *SysUser) (err error, datas *[]SysUser) {
	var rows []SysUser
	model := global.GOrmDao.Model(&SysUser{})
	if data.Id != "" {
		model = model.Where("ID = ?", data.Id)
	}
	if data.UserName != "" {
		model = model.Where("user_name like ?", "%"+data.UserName+"%")
	}
	if data.PhoneNumber != "" {
		model = model.Where("phone_number like ?", "%"+data.PhoneNumber+"%")
	}
	if data.Status != "" {
		model = model.Where("status = ?", data.Status)
	}
	if data.DeptId != "" {
		model = model.Where("dept_id = ?", data.DeptId)
	}
	if data.DataScopeSql != "" {
		model = model.Where(" ?", data.DataScopeSql)
	}
	model.Order("create_time desc")
	err = model.Find(&rows).Error
	datas = &rows
	return err, datas
}

// GetByUserName 根据用户名获取SysUser记录
func (dao *SysUserDao) GetByUserName(name string) (err error, sysUser *SysUser) {
	err = global.GOrmDao.Where("user_name = ?", name).First(&sysUser).Error
	return
}

// GetByDeptId 根据部门id获取SysUser记录
func (dao *SysUserDao) GetByDeptId(deptId string) (err error, sysUser *[]SysUser) {
	err = global.GOrmDao.Where("dept_id = ?", deptId).Find(&sysUser).Error
	return
}

func (dao *SysUserDao) CheckFieldUnique(fieldName, value string) (error, int64) {
	var count int64
	err := global.GOrmDao.Model(&SysUser{}).Where(fieldName+" = ?", value).Count(&count).Error
	return err, count
}
