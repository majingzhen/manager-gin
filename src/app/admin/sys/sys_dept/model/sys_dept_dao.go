// Package model 自动生成模板 SysDeptDao
// @description <TODO description class purpose>
// @author
// @File: sys_dept
// @version 1.0.0
// @create 2023-08-21 10:27:01
package model

import (
	"manager-gin/src/common"
	"manager-gin/src/global"
)

// SysDeptDao 结构体

type SysDeptDao struct{}

// Create 创建SysDept记录
// Author
func (dao *SysDeptDao) Create(sysDept SysDept) (err error) {
	err = global.GOrmDao.Create(&sysDept).Error
	return err
}

// DeleteByIds 批量删除SysDept记录
// Author
func (dao *SysDeptDao) DeleteByIds(ids []string) (err error) {
	err = global.GOrmDao.Delete(&[]SysDept{}, "id in ?", ids).Error
	return err
}

// Update 更新SysDept记录
// Author
func (dao *SysDeptDao) Update(sysDept SysDept) (err error) {
	err = global.GOrmDao.Save(&sysDept).Error
	return err
}

// Get 根据id获取SysDept记录
// Author
func (dao *SysDeptDao) Get(id string) (err error, sysDept *SysDept) {
	err = global.GOrmDao.Where("id = ?", id).First(&sysDept).Error
	return
}

// Page 分页获取SysDept记录
// Author
func (dao *SysDeptDao) Page(param *SysDept, page *common.PageInfo) (err error, datas *[]SysDept, total int64) {
	// 创建model
	model := global.GOrmDao.Model(&SysDept{})
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
	var tmp []SysDept
	err = model.Limit(page.Limit).Offset(page.Offset).Find(&tmp).Error
	datas = &tmp
	return err, datas, total
}

// List 获取SysDept记录
// Author
func (dao *SysDeptDao) List(data *SysDept) (err error, datas *[]SysDept) {
	var rows []SysDept
	db := global.GOrmDao.Model(&SysDept{}).Table("sys_dept d")
	// TODO 输入查询条件
	if data.DeptName != "" {
		db.Where("d.dept_name like ?", "?"+data.DeptName+"?")
	}
	if data.Status != "" {
		db.Where("d.status = ?", data.Status)
	}
	if data.ParentId != "" {
		db.Where("d.parent_id = ?", data.ParentId)
	}
	if data.Id != "" {
		db.Where("d.id = ?", data.Id)
	}
	if data.DataScope != "" {
		db.Where(data.DataScope)
	}
	db.Order("d.parent_id, d.order_num")
	err = db.Find(&rows).Error
	datas = &rows
	return err, datas
}

// CheckDeptNameALL 检查部门名称是否存在
// Author
func (dao *SysDeptDao) CheckDeptNameALL(deptName, parentId string) (err error, res bool) {
	var count int64
	err = global.GOrmDao.Model(&SysDept{}).Where("dept_name = ?,and parent_id = ?", deptName, parentId).Count(&count).Error
	if err != nil {
		return err, false
	}
	return err, count > 0
}

// SelectChildrenDeptById 根据id查询所有下级部门
func (dao *SysDeptDao) SelectChildrenDeptById(id string) (err error, res *[]SysDept) {
	var rows []SysDept
	err = global.GOrmDao.Model(&SysDept{}).Where(" find_in_set(?, ancestors)", id).Find(&rows).Error
	res = &rows
	return
}

// Delete 删除SysDept记录
func (dao *SysDeptDao) Delete(id string) error {
	return global.GOrmDao.Delete(&SysDept{}, "id = ?", id).Error
}
