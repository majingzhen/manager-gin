// Package model 自动生成模板 SysUserDao
// @description <TODO description class purpose>
// @author
// @File: sys_user
// @version 1.0.0
// @create 2023-08-18 14:02:24
package dao

import (
	"manager-gin/src/app/admin/sys/model"
	"manager-gin/src/common"
	"manager-gin/src/global"
)

// SysUserDao 结构体

type SysUserDao struct{}

// Create 创建SysUser记录
// Author
func (dao *SysUserDao) Create(sysUser model.SysUser) (err error) {
	err = global.GOrmDao.Create(&sysUser).Error
	return err
}

// Delete 删除SysUser记录
// Author
func (dao *SysUserDao) Delete(id string) (err error) {
	err = global.GOrmDao.Delete(&[]model.SysUser{}, "id = ?", id).Error
	return err
}

// DeleteByIds 批量删除SysUser记录
// Author
func (dao *SysUserDao) DeleteByIds(ids []string) (err error) {
	err = global.GOrmDao.Delete(&[]model.SysUser{}, "id in ?", ids).Error
	return err
}

// Update 更新SysUser记录
// Author
func (dao *SysUserDao) Update(sysUser model.SysUser) (err error) {
	err = global.GOrmDao.Updates(&sysUser).Error
	return err
}

// Get 根据id获取SysUser记录
// Author
func (dao *SysUserDao) Get(id string) (err error, sysUser *model.SysUser) {
	err = global.GOrmDao.Where("id = ?", id).First(&sysUser).Error
	return
}

// Page 分页获取SysUser记录
// Author
func (dao *SysUserDao) Page(param *model.SysUser, page *common.PageInfo) (err error, datas []*model.SysUser, total int64) {
	// 创建model
	db := global.GOrmDao.Table("sys_user u")
	db.Select("distinct u.id, u.dept_id, u.nick_name, u.user_name, u.email, u.avatar, u.phone_number, u.sex, u.status, u.login_ip, u.login_date, u.create_by, u.create_time, u.remark")
	db.Joins("left join sys_dept d on u.dept_id = d.id")
	// 如果有条件搜索 下方会自动创建搜索语句
	if param.Id != "" {
		db.Where("ID = ?", param.Id)
	}
	if param.UserName != "" {
		db.Where("u.user_name like ?", "%"+param.UserName+"%")
	}
	if param.PhoneNumber != "" {
		db.Where("u.phone_number like ?", "%"+param.PhoneNumber+"%")
	}
	if param.Status != "" {
		db.Where("u.status = ?", param.Status)
	}
	if param.DeptId != "" {
		db.Where("u.dept_id = ?", param.DeptId)
	}
	if param.DataScopeSql != "" {
		db.Where(param.DataScopeSql)
	}
	if err = db.Count(&total).Error; err != nil {
		return
	}
	// 计算分页信息
	page.Calculate()
	// 生成排序信息
	if page.OrderByColumn != "" {
		db.Order(page.OrderByColumn + " " + page.IsAsc + " ")
	}
	var tmp []*model.SysUser
	err = db.Limit(page.Limit).Offset(page.Offset).Find(&tmp).Error
	datas = tmp
	return err, datas, total
}

// List 获取SysUser记录
// Author
func (dao *SysUserDao) List(data *model.SysUser) (err error, datas []*model.SysUser) {
	var rows []*model.SysUser
	model := global.GOrmDao.Table("sys_user u")
	model.Select("distinct u.id, u.dept_id, u.nick_name, u.user_name, u.email, u.avatar, u.phone_number, u.sex, u.status, u.login_ip, u.login_date, u.create_by, u.create_time, u.remark")
	model.Joins("left join sys_dept d on u.dept_id = d.id")
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
	datas = rows
	return err, datas
}

// GetByUserName 根据用户名获取SysUser记录
func (dao *SysUserDao) GetByUserName(name string) (err error, sysUser *model.SysUser) {
	err = global.GOrmDao.Where("user_name = ?", name).First(&sysUser).Error
	return
}

// GetByDeptId 根据部门id获取SysUser记录
func (dao *SysUserDao) GetByDeptId(deptId string) (err error, sysUser []*model.SysUser) {
	err = global.GOrmDao.Where("dept_id = ?", deptId).Find(&sysUser).Error
	return
}

func (dao *SysUserDao) CheckFieldUnique(fieldName, value string) (error, int64) {
	var count int64
	err := global.GOrmDao.Model(&model.SysUser{}).Where(fieldName+" = ?", value).Count(&count).Error
	return err, count
}

func (dao *SysUserDao) SelectByField(fieldName string, value string) (error, *model.SysUser) {
	var users []*model.SysUser
	if err := global.GOrmDao.Model(&model.SysUser{}).Where(fieldName+" = ?", value).Find(&users).Error; err != nil {
		return err, nil
	} else {
		if users != nil && len(users) > 0 {
			return nil, users[0]
		}
		return nil, nil
	}

}

func (dao *SysUserDao) SelectAllocatedList(param *model.SysUser, page *common.PageInfo, roleId string) (err error, datas []*model.SysUser, total int64) {
	// 创建model
	db := global.GOrmDao.Table("sys_user u")
	db.Select("distinct u.id, u.dept_id, u.user_name, u.nick_name, u.email, u.phone_number, u.status, u.create_time")
	db.Joins("left join sys_dept d on u.dept_id = d.id")
	db.Joins("left join sys_user_role ur on u.id = ur.user_id")
	db.Joins("left join sys_role r on ur.role_id = r.id")
	db.Where("r.id = ?", roleId)
	// 如果有条件搜索 下方会自动创建搜索语句
	if param.UserName != "" {
		db.Where("u.user_name like ?", "%"+param.UserName+"%")
	}
	if param.PhoneNumber != "" {
		db.Where("u.phone_number like ?", "%"+param.PhoneNumber+"%")
	}
	if param.DataScopeSql != "" {
		db.Where(param.DataScopeSql)
	}
	if err = db.Count(&total).Error; err != nil {
		return
	}
	// 计算分页信息
	page.Calculate()
	// 生成排序信息
	if page.OrderByColumn != "" {
		db.Order(page.OrderByColumn + " " + page.IsAsc + " ")
	}
	var tmp []*model.SysUser
	err = db.Limit(page.Limit).Offset(page.Offset).Find(&tmp).Error
	datas = tmp
	return err, datas, total
}

func (dao *SysUserDao) SelectUnallocatedList(param *model.SysUser, page *common.PageInfo, roleId string) (err error, datas []*model.SysUser, total int64) {
	// 创建model
	db := global.GOrmDao.Table("sys_user u")
	db.Select("distinct u.id, u.dept_id, u.user_name, u.nick_name, u.email, u.phone_number, u.status, u.create_time")
	db.Joins("left join sys_dept d on u.dept_id = d.id")
	db.Joins("left join sys_user_role ur on u.id = ur.user_id")
	db.Joins("left join sys_role r on ur.role_id = r.id")
	db.Where("(r.id != ? or r.id is null)", roleId)
	db.Where("u.id not in (select u.id from sys_user u inner join sys_user_role ur on u.id = ur.user_id and ur.role_id = ?)", roleId)
	// 如果有条件搜索 下方会自动创建搜索语句
	if param.UserName != "" {
		db.Where("u.user_name like ?", "%"+param.UserName+"%")
	}
	if param.PhoneNumber != "" {
		db.Where("u.phone_number like ?", "%"+param.PhoneNumber+"%")
	}
	if param.DataScopeSql != "" {
		db.Where(param.DataScopeSql)
	}
	if err = db.Count(&total).Error; err != nil {
		return
	}
	// 计算分页信息
	page.Calculate()
	// 生成排序信息
	if page.OrderByColumn != "" {
		db.Order(page.OrderByColumn + " " + page.IsAsc + " ")
	}
	var tmp []*model.SysUser
	err = db.Limit(page.Limit).Offset(page.Offset).Find(&tmp).Error
	datas = tmp
	return err, datas, total
}
