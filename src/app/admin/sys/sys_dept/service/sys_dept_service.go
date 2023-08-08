// Package service 自动生成模板 SysDeptService
// @description <TODO description class purpose>
// @author
// @File: sys_dept
// @version 1.0.0
// @create 2023-08-08 10:06:19
package service

import (
	"manager-gin/src/app/admin/sys/sys_dept/model"
	"manager-gin/src/app/admin/sys/sys_dept/service/view"
	"manager-gin/src/common"
)

var sysDeptDao = model.SysDeptDaoApp
var viewUtils = view.SysDeptViewUtilsApp

type SysDeptService struct{}

// Create 创建SysDept记录
// Author
func (sysDeptService *SysDeptService) Create(sysDeptView *view.SysDeptView) (err error) {
	err1, sysDept := viewUtils.View2Data(sysDeptView)
	if err1 != nil {
		return err1
	}
	err2 := sysDeptDao.Create(*sysDept)
	if err2 != nil {
		return err2
	}
	return nil
}

// Delete 删除SysDept记录
// Author
func (sysDeptService *SysDeptService) Delete(id int) (err error) {
	err = sysDeptDao.Delete(id)
	return err
}

// DeleteByIds 批量删除SysDept记录
// Author
func (sysDeptService *SysDeptService) DeleteByIds(ids []int) (err error) {
	err = sysDeptDao.DeleteByIds(ids)
	return err
}

// Update 更新SysDept记录
// Author
func (sysDeptService *SysDeptService) Update(id int, sysDeptView *view.SysDeptView) (err error) {
	sysDeptView.Id = id
	err1, sysDept := viewUtils.View2Data(sysDeptView)
	if err1 != nil {
		return err1
	}
	err = sysDeptDao.Update(*sysDept)
	return err
}

// Get 根据id获取SysDept记录
// Author
func (sysDeptService *SysDeptService) Get(id int) (err error, sysDeptView *view.SysDeptView) {
	err1, sysDept := sysDeptDao.Get(id)
	if err1 != nil {
		return err1, nil
	}
	err2, sysDeptView := viewUtils.Data2View(sysDept)
	if err2 != nil {
		return err2, nil
	}
	return
}

// Find 分页获取SysDept记录
// Author
func (sysDeptService *SysDeptService) Find(info *common.PageInfoV2) (err error) {
	err1, sysDepts, total := sysDeptDao.Find(info)
	if err1 != nil {
		return err1
	}
	info.Total = total
	err2, viewList := viewUtils.Data2ViewList(sysDepts)
	if err2 != nil {
		return err2
	}
	info.FormList = viewList
	return err
}
