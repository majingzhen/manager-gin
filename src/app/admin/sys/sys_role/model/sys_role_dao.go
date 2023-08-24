// Package model 自动生成模板 SysRoleDao
// @description <TODO description class purpose>
// @author
// @File: sys_role
// @version 1.0.0
// @create 2023-08-18 14:00:53
package model

import (
	"manager-gin/src/common"
	"manager-gin/src/global"
)

// SysRoleDao 结构体

type SysRoleDao struct{}

// Create 创建SysRole记录
// Author
func (dao *SysRoleDao) Create(sysRole SysRole) (err error) {
	err = global.GOrmDao.Create(&sysRole).Error
	return err
}

// Delete 删除SysRole记录
// Author
func (dao *SysRoleDao) Delete(id string) (err error) {
	err = global.GOrmDao.Delete(&[]SysRole{}, "id = ?", id).Error
	return err
}

// DeleteByIds 批量删除SysRole记录
// Author
func (dao *SysRoleDao) DeleteByIds(ids []string) (err error) {
	err = global.GOrmDao.Delete(&[]SysRole{}, "id in ?", ids).Error
	return err
}

// Update 更新SysRole记录
// Author
func (dao *SysRoleDao) Update(sysRole *SysRole) (err error) {
	err = global.GOrmDao.Updates(sysRole).Error
	return err
}

// Get 根据id获取SysRole记录
// Author
func (dao *SysRoleDao) Get(id string) (err error, sysRole *SysRole) {
	err = global.GOrmDao.Where("id = ?", id).First(&sysRole).Error
	return
}

// Page 分页获取SysRole记录
// Author
func (dao *SysRoleDao) Page(param *SysRole, page *common.PageInfo) (err error, datas []*SysRole, total int64) {
	// 创建model
	model := global.GOrmDao.Table("sys_role r")
	model.Select("distinct r.id, r.role_name, r.role_key, r.role_sort, r.data_scope, r.menu_check_strictly, r.dept_check_strictly,r.status, r.create_time, r.remark ")
	model.Joins("left join sys_user_role ur on ur.role_id = r.id")
	model.Joins("left join sys_user u on u.id = ur.user_id")
	model.Joins("left join sys_dept d on u.dept_id = d.id")
	// 如果有条件搜索 下方会自动创建搜索语句
	if param.Id != "" {
		model = model.Where("ID = ?", param.Id)
	}
	if param.RoleName != "" {
		model = model.Where("role_name = ?", "%"+param.RoleName+"%")
	}
	if param.RoleKey != "" {
		model = model.Where("role_key = ?", "%"+param.RoleKey+"%")
	}
	if param.Status != "" {
		model = model.Where("status = ?", param.Status)
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
	var tmp []*SysRole
	err = model.Limit(page.Limit).Offset(page.Offset).Find(&tmp).Error
	datas = tmp
	return err, datas, total
}

// List 获取SysRole记录
// Author
func (dao *SysRoleDao) List(data *SysRole) (err error, datas []*SysRole) {
	var rows []*SysRole
	model := global.GOrmDao.Table("sys_role r")
	model.Select("distinct r.id, r.role_name, r.role_key, r.role_sort, r.data_scope, r.menu_check_strictly, r.dept_check_strictly,r.status, r.create_time, r.remark ")
	model.Joins("left join sys_user_role ur on ur.role_id = r.id")
	model.Joins("left join sys_user u on u.id = ur.user_id")
	model.Joins("left join sys_dept d on u.dept_id = d.id")
	if data.Id != "" {
		model = model.Where("ID = ?", data.Id)
	}
	if data.RoleName != "" {
		model = model.Where("role_name = ?", "%"+data.RoleName+"%")
	}
	if data.RoleKey != "" {
		model = model.Where("role_key = ?", "%"+data.RoleKey+"%")
	}
	if data.Status != "" {
		model = model.Where("status = ?", data.Status)
	}
	if data.DataScopeSql != "" {
		model = model.Where(data.DataScopeSql)
	}
	model.Order("create_time desc")
	err = model.Find(&rows).Error
	datas = rows
	return err, datas
}

// GetRoleByUserId 根据用户获取角色集合
func (dao *SysRoleDao) GetRoleByUserId(userId string) (err error, roles []*SysRole) {
	var tmp []*SysRole
	model := global.GOrmDao.Table("sys_role r")
	model.Select("distinct r.id, r.role_name, r.role_key, r.role_sort, r.data_scope, r.menu_check_strictly, r.dept_check_strictly,r.status, r.create_time, r.remark ")
	model.Joins("left join sys_user_role ur on ur.role_id = r.id")
	model.Where("ur.user_id = ? and r.status = ?", userId, common.STATUS_NORMAL)
	err = model.Find(&tmp).Error
	roles = tmp
	return err, roles
}

// CheckRoleNameUnique 校验角色名称是否唯一
func (dao *SysRoleDao) CheckRoleNameUnique(name string) (error, *SysRole) {
	var data []*SysRole
	if err := global.GOrmDao.Table("sys_role").Where("role_name = ?", name).Find(&data).Error; err != nil {
		return err, nil
	} else {
		if data != nil && len(data) > 0 {
			return nil, data[0]
		} else {
			return nil, nil
		}
	}
}

// CheckRoleKeyUnique 校验角色权限是否唯一
func (dao *SysRoleDao) CheckRoleKeyUnique(key string) (error, *SysRole) {
	var data []*SysRole
	if err := global.GOrmDao.Table("sys_role").Where("role_key = ?", key).Find(&data).Error; err != nil {
		return err, nil
	} else {
		if data != nil && len(data) > 0 {
			return nil, data[0]
		} else {
			return nil, nil
		}
	}
}
