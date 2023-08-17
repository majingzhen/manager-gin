// Package service 自动生成模板 BiRoleService
// @description <TODO description class purpose>
// @author Majz
// @File: bi_role
// @version 1.0.0
// @create 2023-08-16 14:53:36
package service

import (
	"manager-gin/src/app/admin/bi/bi_role/model"
	"manager-gin/src/app/admin/bi/bi_role/service/view"
	"manager-gin/src/common"
	"strconv"
)

var biRoleDao = model.BiRoleDaoApp
var viewUtils = view.BiRoleViewUtilsApp

type BiRoleService struct{}

// Create 创建BiRole记录
// Author Majz
func (biRoleService *BiRoleService) Create(biRoleView *view.BiRoleView) (err error) {
	err1, biRole := viewUtils.View2Data(biRoleView)
	if err1 != nil {
		return err1
	}
	err2 := biRoleDao.Create(*biRole)
	if err2 != nil {
		return err2
	}
	return nil
}

// Delete 删除BiRole记录
// Author Majz
func (biRoleService *BiRoleService) Delete(id int) (err error) {
	err = biRoleDao.Delete(id)
	return err
}

// DeleteByIds 批量删除BiRole记录
// Author Majz
func (biRoleService *BiRoleService) DeleteByIds(ids []int) (err error) {
	err = biRoleDao.DeleteByIds(ids)
	return err
}

// Update 更新BiRole记录
// Author Majz
func (biRoleService *BiRoleService) Update(id int, biRoleView *view.BiRoleView) (err error) {
	biRoleView.Id = strconv.Itoa(id)
	err1, biRole := viewUtils.View2Data(biRoleView)
	if err1 != nil {
		return err1
	}
	err = biRoleDao.Update(*biRole)
	return err
}

// Get 根据id获取BiRole记录
// Author Majz
func (biRoleService *BiRoleService) Get(id int) (err error, biRoleView *view.BiRoleView) {
	err1, biRole := biRoleDao.Get(id)
	if err1 != nil {
		return err1, nil
	}
	err2, biRoleView := viewUtils.Data2View(biRole)
	if err2 != nil {
		return err2, nil
	}
	return
}

// Find 分页获取BiRole记录
// Author Majz
func (biRoleService *BiRoleService) Find(info *common.PageInfoV2) (err error) {
	err1, biRoles, total := biRoleDao.Find(info)
	if err1 != nil {
		return err1
	}
	info.Total = total
	err2, viewList := viewUtils.Data2ViewList(biRoles)
	if err2 != nil {
		return err2
	}
	info.FormList = viewList
	return err
}
