// Package service 自动生成模板 SysOrganizationService
// @description <TODO description class purpose>
// @author
// @File: sys_organization
// @version 1.0.0
// @create 2023-08-18 14:00:53
package service

import (
	"manager-gin/src/app/admin/sys/sys_organization/model"
	"manager-gin/src/app/admin/sys/sys_organization/service/view"
	"manager-gin/src/common"
)

var sysOrganizationDao = model.SysOrganizationDaoApp
var viewUtils = view.SysOrganizationViewUtilsApp

type SysOrganizationService struct{}

// Create 创建SysOrganization记录
// Author
func (sysOrganizationService *SysOrganizationService) Create(sysOrganizationView *view.SysOrganizationView) (err error) {
	err1, sysOrganization := viewUtils.View2Data(sysOrganizationView)
	if err1 != nil {
		return err1
	}
	err2 := sysOrganizationDao.Create(*sysOrganization)
	if err2 != nil {
		return err2
	}
	return nil
}

// Delete 删除SysOrganization记录
// Author
func (sysOrganizationService *SysOrganizationService) Delete(id string) (err error) {
	err = sysOrganizationDao.Delete(id)
	return err
}

// DeleteByIds 批量删除SysOrganization记录
// Author
func (sysOrganizationService *SysOrganizationService) DeleteByIds(ids []string) (err error) {
	err = sysOrganizationDao.DeleteByIds(ids)
	return err
}

// Update 更新SysOrganization记录
// Author
func (sysOrganizationService *SysOrganizationService) Update(id string, sysOrganizationView *view.SysOrganizationView) (err error) {
	sysOrganizationView.Id = id
	err1, sysOrganization := viewUtils.View2Data(sysOrganizationView)
	if err1 != nil {
		return err1
	}
	err = sysOrganizationDao.Update(*sysOrganization)
	return err
}

// Get 根据id获取SysOrganization记录
// Author
func (sysOrganizationService *SysOrganizationService) Get(id string) (err error, sysOrganizationView *view.SysOrganizationView) {
	err1, sysOrganization := sysOrganizationDao.Get(id)
	if err1 != nil {
		return err1, nil
	}
	err2, sysOrganizationView := viewUtils.Data2View(sysOrganization)
	if err2 != nil {
		return err2, nil
	}
	return
}

// Find 分页获取SysOrganization记录
// Author
func (sysOrganizationService *SysOrganizationService) Find(info *common.PageInfoV2) (err error) {
	err1, sysOrganizations, total := sysOrganizationDao.Find(info)
	if err1 != nil {
		return err1
	}
	info.Total = total
	err2, viewList := viewUtils.Data2ViewList(sysOrganizations)
	if err2 != nil {
		return err2
	}
	info.FormList = viewList
	return err
}
