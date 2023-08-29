// Package dao 自动生成模板 SysUserDao
// @description <TODO description class purpose>
// @author
// @File: sys_user
// @version 1.0.0
// @create 2023-08-18 14:02:24
package dao

import (
	"gorm.io/gorm"
	"manager-gin/src/app/admin/sys/model"
	"manager-gin/src/app/admin/sys/service/sys_user/view"
	"manager-gin/src/common"
	"manager-gin/src/global"
)

// SysUserDao 结构体

type SysUserDao struct{}

// Create 创建SysUser记录
// Author
func (dao *SysUserDao) Create(tx *gorm.DB, sysUser model.SysUser) (err error) {
	err = tx.Create(&sysUser).Error
	return err
}

// Delete 删除SysUser记录
// Author
func (dao *SysUserDao) Delete(tx *gorm.DB, id string) (err error) {
	err = tx.Delete(&[]model.SysUser{}, "id = ?", id).Error
	return err
}

// DeleteByIds 批量删除SysUser记录
// Author
func (dao *SysUserDao) DeleteByIds(tx *gorm.DB, ids []string) (err error) {
	err = tx.Delete(&[]model.SysUser{}, "id in ?", ids).Error
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
func (dao *SysUserDao) Page(param *view.SysUserPageView) (err error, page *common.PageInfo) {
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
		db.Where("(u.dept_id = ? or u.dept_id in (SELECT t.id FROM sys_dept t WHERE find_in_set(?, ancestors)))", param.DeptId, param.DeptId)
	}
	if param.DataScopeSql != "" {
		db.Where(param.DataScopeSql)
	}
	page = common.CreatePageInfo(param.PageNum, param.PageSize)
	if err = db.Count(&page.Total).Error; err != nil {
		return
	}
	// 生成排序信息
	if param.OrderByColumn != "" {
		db.Order(param.OrderByColumn + " " + param.IsAsc + " ")
	}
	var tmp []*model.SysUser
	err = db.Limit(page.Limit).Offset(page.Offset).Find(&tmp).Error
	page.Rows = tmp
	return err, page
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

// CheckFieldUnique 判断指定字段是否为空
func (dao *SysUserDao) CheckFieldUnique(fieldName, value string) (error, int64) {
	var count int64
	err := global.GOrmDao.Model(&model.SysUser{}).Where(fieldName+" = ?", value).Count(&count).Error
	return err, count
}

// SelectByField 根据指定字段查询数据
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

// SelectAllocatedList 获取已分配用户角色的用户列表
func (dao *SysUserDao) SelectAllocatedList(param *view.SysUserPageView) (err error, page *common.PageInfo) {
	// 创建model
	db := global.GOrmDao.Table("sys_user u")
	db.Select("distinct u.id, u.dept_id, u.user_name, u.nick_name, u.email, u.phone_number, u.status, u.create_time")
	db.Joins("left join sys_dept d on u.dept_id = d.id")
	db.Joins("left join sys_user_role ur on u.id = ur.user_id")
	db.Joins("left join sys_role r on ur.role_id = r.id")
	db.Where("r.id = ?", param.RoleId)
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
	page = common.CreatePageInfo(param.PageNum, param.PageSize)
	if err = db.Count(&page.Total).Error; err != nil {
		return
	}
	// 生成排序信息
	if param.OrderByColumn != "" {
		db.Order(param.OrderByColumn + " " + param.IsAsc + " ")
	}
	var tmp []*model.SysUser
	err = db.Limit(page.Limit).Offset(page.Offset).Find(&tmp).Error
	page.Rows = tmp
	return err, page
}

// SelectUnallocatedList 获取未分配用户角色的用户列表
func (dao *SysUserDao) SelectUnallocatedList(param *view.SysUserPageView) (err error, page *common.PageInfo) {
	// 创建model
	db := global.GOrmDao.Table("sys_user u")
	db.Select("distinct u.id, u.dept_id, u.user_name, u.nick_name, u.email, u.phone_number, u.status, u.create_time")
	db.Joins("left join sys_dept d on u.dept_id = d.id")
	db.Joins("left join sys_user_role ur on u.id = ur.user_id")
	db.Joins("left join sys_role r on ur.role_id = r.id")
	db.Where("(r.id != ? or r.id is null)", param.RoleId)
	db.Where("u.id not in (select u.id from sys_user u inner join sys_user_role ur on u.id = ur.user_id and ur.role_id = ?)", param.RoleId)
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
	page = common.CreatePageInfo(param.PageNum, param.PageSize)
	if err = db.Count(&page.Total).Error; err != nil {
		return
	}
	// 生成排序信息
	if param.OrderByColumn != "" {
		db.Order(param.OrderByColumn + " " + param.IsAsc + " ")
	}
	var tmp []*model.SysUser
	err = db.Limit(page.Limit).Offset(page.Offset).Find(&tmp).Error
	page.Rows = tmp
	return err, page
}
