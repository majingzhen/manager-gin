// Package service 自动生成模板 SysOperLogService
// @description <TODO description class purpose>
// @author
// @File: sys_oper_log
// @version 1.0.0
// @create 2023-08-08 10:06:19
package service

import (
	"manager-gin/src/app/admin/sys/sys_oper_log/model"
	"manager-gin/src/app/admin/sys/sys_oper_log/service/view"
	"manager-gin/src/common"
)

var sysOperLogDao = model.SysOperLogDaoApp
var viewUtils = view.SysOperLogViewUtilsApp

type SysOperLogService struct{}

// Create 创建SysOperLog记录
// Author
func (sysOperLogService *SysOperLogService) Create(sysOperLogView *view.SysOperLogView) (err error) {
	err1, sysOperLog := viewUtils.View2Data(sysOperLogView)
	if err1 != nil {
		return err1
	}
	err2 := sysOperLogDao.Create(*sysOperLog)
	if err2 != nil {
		return err2
	}
	return nil
}

// Delete 删除SysOperLog记录
// Author
func (sysOperLogService *SysOperLogService) Delete(id int) (err error) {
	err = sysOperLogDao.Delete(id)
	return err
}

// DeleteByIds 批量删除SysOperLog记录
// Author
func (sysOperLogService *SysOperLogService) DeleteByIds(ids []int) (err error) {
	err = sysOperLogDao.DeleteByIds(ids)
	return err
}

// Update 更新SysOperLog记录
// Author
func (sysOperLogService *SysOperLogService) Update(id int, sysOperLogView *view.SysOperLogView) (err error) {
	sysOperLogView.Id = id
	err1, sysOperLog := viewUtils.View2Data(sysOperLogView)
	if err1 != nil {
		return err1
	}
	err = sysOperLogDao.Update(*sysOperLog)
	return err
}

// Get 根据id获取SysOperLog记录
// Author
func (sysOperLogService *SysOperLogService) Get(id int) (err error, sysOperLogView *view.SysOperLogView) {
	err1, sysOperLog := sysOperLogDao.Get(id)
	if err1 != nil {
		return err1, nil
	}
	err2, sysOperLogView := viewUtils.Data2View(sysOperLog)
	if err2 != nil {
		return err2, nil
	}
	return
}

// Find 分页获取SysOperLog记录
// Author
func (sysOperLogService *SysOperLogService) Find(info *common.PageInfoV2) (err error) {
	err1, sysOperLogs, total := sysOperLogDao.Find(info)
	if err1 != nil {
		return err1
	}
	info.Total = total
	err2, viewList := viewUtils.Data2ViewList(sysOperLogs)
	if err2 != nil {
		return err2
	}
	info.FormList = viewList
	return err
}
