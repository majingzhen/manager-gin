// Package service 自动生成模板 SysRoleDeptService
// @description <TODO description class purpose>
// @author
// @File: sys_role_dept
// @version 1.0.0
// @create 2023-08-08 10:06:19
package service

import (
	"manager-gin/src/app/admin/sys/sys_role_dept/model"
	"manager-gin/src/app/admin/sys/sys_role_dept/service/view"
	"manager-gin/src/common"
)

var sysRoleDeptDao = model.SysRoleDeptDaoApp
var viewUtils = view.SysRoleDeptViewUtilsApp

type SysRoleDeptService struct{}

// Create 创建SysRoleDept记录
// Author
func (sysRoleDeptService *SysRoleDeptService) Create(sysRoleDeptView *view.SysRoleDeptView) (err error) {
	err1, sysRoleDept := viewUtils.View2Data(sysRoleDeptView)
	if err1 != nil {
		return err1
	}
	err2 := sysRoleDeptDao.Create(*sysRoleDept)
	if err2 != nil {
		return err2
	}
	return nil
}

// Delete 删除SysRoleDept记录
// Author
func (sysRoleDeptService *SysRoleDeptService) Delete(id int) (err error) {
	err = sysRoleDeptDao.Delete(id)
	return err
}

// DeleteByIds 批量删除SysRoleDept记录
// Author
func (sysRoleDeptService *SysRoleDeptService) DeleteByIds(ids []int) (err error) {
	err = sysRoleDeptDao.DeleteByIds(ids)
	return err
}

// Update 更新SysRoleDept记录
// Author
func (sysRoleDeptService *SysRoleDeptService) Update(id int, sysRoleDeptView *view.SysRoleDeptView) (err error) {
	sysRoleDeptView.Id = id
	err1, sysRoleDept := viewUtils.View2Data(sysRoleDeptView)
	if err1 != nil {
		return err1
	}
	err = sysRoleDeptDao.Update(*sysRoleDept)
	return err
}

// Get 根据id获取SysRoleDept记录
// Author
func (sysRoleDeptService *SysRoleDeptService) Get(id int) (err error, sysRoleDeptView *view.SysRoleDeptView) {
	err1, sysRoleDept := sysRoleDeptDao.Get(id)
	if err1 != nil {
		return err1, nil
	}
	err2, sysRoleDeptView := viewUtils.Data2View(sysRoleDept)
	if err2 != nil {
		return err2, nil
	}
	return
}

// Find 分页获取SysRoleDept记录
// Author
func (sysRoleDeptService *SysRoleDeptService) Find(info *common.PageInfoV2) (err error) {
	err1, sysRoleDepts, total := sysRoleDeptDao.Find(info)
	if err1 != nil {
		return err1
	}
	info.Total = total
	err2, viewList := viewUtils.Data2ViewList(sysRoleDepts)
	if err2 != nil {
		return err2
	}
	info.FormList = viewList
	return err
}
