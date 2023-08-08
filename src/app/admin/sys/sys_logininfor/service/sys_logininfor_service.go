// Package service 自动生成模板 SysLogininforService
// @description <TODO description class purpose>
// @author
// @File: sys_logininfor
// @version 1.0.0
// @create 2023-08-08 10:06:19
package service

import (
	"manager-gin/src/app/admin/sys/sys_logininfor/model"
	"manager-gin/src/app/admin/sys/sys_logininfor/service/view"
	"manager-gin/src/common"
)

var sysLogininforDao = model.SysLogininforDaoApp
var viewUtils = view.SysLogininforViewUtilsApp

type SysLogininforService struct{}

// Create 创建SysLogininfor记录
// Author
func (sysLogininforService *SysLogininforService) Create(sysLogininforView *view.SysLogininforView) (err error) {
	err1, sysLogininfor := viewUtils.View2Data(sysLogininforView)
	if err1 != nil {
		return err1
	}
	err2 := sysLogininforDao.Create(*sysLogininfor)
	if err2 != nil {
		return err2
	}
	return nil
}

// Delete 删除SysLogininfor记录
// Author
func (sysLogininforService *SysLogininforService) Delete(id int) (err error) {
	err = sysLogininforDao.Delete(id)
	return err
}

// DeleteByIds 批量删除SysLogininfor记录
// Author
func (sysLogininforService *SysLogininforService) DeleteByIds(ids []int) (err error) {
	err = sysLogininforDao.DeleteByIds(ids)
	return err
}

// Update 更新SysLogininfor记录
// Author
func (sysLogininforService *SysLogininforService) Update(id int, sysLogininforView *view.SysLogininforView) (err error) {
	sysLogininforView.Id = id
	err1, sysLogininfor := viewUtils.View2Data(sysLogininforView)
	if err1 != nil {
		return err1
	}
	err = sysLogininforDao.Update(*sysLogininfor)
	return err
}

// Get 根据id获取SysLogininfor记录
// Author
func (sysLogininforService *SysLogininforService) Get(id int) (err error, sysLogininforView *view.SysLogininforView) {
	err1, sysLogininfor := sysLogininforDao.Get(id)
	if err1 != nil {
		return err1, nil
	}
	err2, sysLogininforView := viewUtils.Data2View(sysLogininfor)
	if err2 != nil {
		return err2, nil
	}
	return
}

// Find 分页获取SysLogininfor记录
// Author
func (sysLogininforService *SysLogininforService) Find(info *common.PageInfoV2) (err error) {
	err1, sysLogininfors, total := sysLogininforDao.Find(info)
	if err1 != nil {
		return err1
	}
	info.Total = total
	err2, viewList := viewUtils.Data2ViewList(sysLogininfors)
	if err2 != nil {
		return err2
	}
	info.FormList = viewList
	return err
}
