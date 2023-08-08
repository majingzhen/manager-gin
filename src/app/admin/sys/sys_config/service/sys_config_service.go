// Package service 自动生成模板 SysConfigService
// @description <TODO description class purpose>
// @author
// @File: sys_config
// @version 1.0.0
// @create 2023-08-08 10:06:19
package service

import (
	"manager-gin/src/app/admin/sys/sys_config/model"
	"manager-gin/src/app/admin/sys/sys_config/service/view"
	"manager-gin/src/common"
)

var sysConfigDao = model.SysConfigDaoApp
var viewUtils = view.SysConfigViewUtilsApp

type SysConfigService struct{}

// Create 创建SysConfig记录
// Author
func (sysConfigService *SysConfigService) Create(sysConfigView *view.SysConfigView) (err error) {
	err1, sysConfig := viewUtils.View2Data(sysConfigView)
	if err1 != nil {
		return err1
	}
	err2 := sysConfigDao.Create(*sysConfig)
	if err2 != nil {
		return err2
	}
	return nil
}

// Delete 删除SysConfig记录
// Author
func (sysConfigService *SysConfigService) Delete(id int) (err error) {
	err = sysConfigDao.Delete(id)
	return err
}

// DeleteByIds 批量删除SysConfig记录
// Author
func (sysConfigService *SysConfigService) DeleteByIds(ids []int) (err error) {
	err = sysConfigDao.DeleteByIds(ids)
	return err
}

// Update 更新SysConfig记录
// Author
func (sysConfigService *SysConfigService) Update(id int, sysConfigView *view.SysConfigView) (err error) {
	sysConfigView.Id = id
	err1, sysConfig := viewUtils.View2Data(sysConfigView)
	if err1 != nil {
		return err1
	}
	err = sysConfigDao.Update(*sysConfig)
	return err
}

// Get 根据id获取SysConfig记录
// Author
func (sysConfigService *SysConfigService) Get(id int) (err error, sysConfigView *view.SysConfigView) {
	err1, sysConfig := sysConfigDao.Get(id)
	if err1 != nil {
		return err1, nil
	}
	err2, sysConfigView := viewUtils.Data2View(sysConfig)
	if err2 != nil {
		return err2, nil
	}
	return
}

// Find 分页获取SysConfig记录
// Author
func (sysConfigService *SysConfigService) Find(info *common.PageInfoV2) (err error) {
	err1, sysConfigs, total := sysConfigDao.Find(info)
	if err1 != nil {
		return err1
	}
	info.Total = total
	err2, viewList := viewUtils.Data2ViewList(sysConfigs)
	if err2 != nil {
		return err2
	}
	info.FormList = viewList
	return err
}
