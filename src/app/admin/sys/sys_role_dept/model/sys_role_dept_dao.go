// Package model 自动生成模板 SysRoleDeptDao
// @description <TODO description class purpose>
// @author
// @File: sys_role_dept
// @version 1.0.0
// @create 2023-08-08 10:06:19
package model

import (
	"manager-gin/src/common"
	"manager-gin/src/global"
)

// SysRoleDeptDao 结构体

type SysRoleDeptDao struct{}

// Create 创建SysRoleDept记录
// Author
func (dao *SysRoleDeptDao) Create(sysRoleDept SysRoleDept) (err error) {
	err = global.GOrmDao.Create(&sysRoleDept).Error
	return err
}

// Delete 删除SysRoleDept记录
// Author
func (dao *SysRoleDeptDao) Delete(id int) (err error) {
	err = global.GOrmDao.Delete(&[]SysRoleDept{}, "id = ?", id).Error
	return err
}

// DeleteByIds 批量删除SysRoleDept记录
// Author
func (dao *SysRoleDeptDao) DeleteByIds(ids []int) (err error) {
	err = global.GOrmDao.Delete(&[]SysRoleDept{}, "id in ?", ids).Error
	return err
}

// Update 更新SysRoleDept记录
// Author
func (dao *SysRoleDeptDao) Update(sysRoleDept SysRoleDept) (err error) {
	err = global.GOrmDao.Save(&sysRoleDept).Error
	return err
}

// Get 根据id获取SysRoleDept记录
// Author
func (dao *SysRoleDeptDao) Get(id int) (err error, sysRoleDept *SysRoleDept) {
	err = global.GOrmDao.Where("id = ?", id).First(&sysRoleDept).Error
	return
}

// Find 分页获取SysRoleDept记录
// Author
func (dao *SysRoleDeptDao) Find(info *common.PageInfoV2) (err error, sysRoleDepts *[]SysRoleDept, total int64) {
	// 创建db
	db := global.GOrmDao.Model(&SysRoleDept{})
	// 如果有条件搜索 下方会自动创建搜索语句
	//if info.Id != "" {
	//	db = db.Where("ID = ?", info.Id)
	//}

	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var tmp []SysRoleDept
	err = db.Limit(info.Limit).Offset(info.Offset).Find(&tmp).Error
	sysRoleDepts = &tmp
	return err, sysRoleDepts, total
}
