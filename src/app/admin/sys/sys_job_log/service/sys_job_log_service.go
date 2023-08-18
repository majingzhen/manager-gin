// Package service 自动生成模板 SysJobLogService
// @description <TODO description class purpose>
// @author
// @File: sys_job_log
// @version 1.0.0
// @create 2023-08-18 13:41:26
package service

import (
	"manager-gin/src/app/admin/sys/sys_job_log/model"
	"manager-gin/src/app/admin/sys/sys_job_log/service/view"
	"manager-gin/src/common"
)

var sysJobLogDao = model.SysJobLogDaoApp
var viewUtils = view.SysJobLogViewUtilsApp

type SysJobLogService struct{}

// Create 创建SysJobLog记录
// Author
func (sysJobLogService *SysJobLogService) Create(sysJobLogView *view.SysJobLogView) (err error) {
	err1, sysJobLog := viewUtils.View2Data(sysJobLogView)
	if err1 != nil {
		return err1
	}
	err2 := sysJobLogDao.Create(*sysJobLog)
	if err2 != nil {
		return err2
	}
	return nil
}

// Delete 删除SysJobLog记录
// Author
func (sysJobLogService *SysJobLogService) Delete(id string) (err error) {
	err = sysJobLogDao.Delete(id)
	return err
}

// DeleteByIds 批量删除SysJobLog记录
// Author
func (sysJobLogService *SysJobLogService) DeleteByIds(ids []string) (err error) {
	err = sysJobLogDao.DeleteByIds(ids)
	return err
}

// Update 更新SysJobLog记录
// Author
func (sysJobLogService *SysJobLogService) Update(id string, sysJobLogView *view.SysJobLogView) (err error) {
	sysJobLogView.Id = id
	err1, sysJobLog := viewUtils.View2Data(sysJobLogView)
	if err1 != nil {
		return err1
	}
	err = sysJobLogDao.Update(*sysJobLog)
	return err
}

// Get 根据id获取SysJobLog记录
// Author
func (sysJobLogService *SysJobLogService) Get(id string) (err error, sysJobLogView *view.SysJobLogView) {
	err1, sysJobLog := sysJobLogDao.Get(id)
	if err1 != nil {
		return err1, nil
	}
	err2, sysJobLogView := viewUtils.Data2View(sysJobLog)
	if err2 != nil {
		return err2, nil
	}
	return
}

// Find 分页获取SysJobLog记录
// Author
func (sysJobLogService *SysJobLogService) Find(info *common.PageInfoV2) (err error) {
	err1, sysJobLogs, total := sysJobLogDao.Find(info)
	if err1 != nil {
		return err1
	}
	info.Total = total
	err2, viewList := viewUtils.Data2ViewList(sysJobLogs)
	if err2 != nil {
		return err2
	}
	info.FormList = viewList
	return err
}
