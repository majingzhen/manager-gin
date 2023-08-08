// Package service 自动生成模板 SysJobService
// @description <TODO description class purpose>
// @author
// @File: sys_job
// @version 1.0.0
// @create 2023-08-08 10:06:19
package service

import (
	"manager-gin/src/app/admin/sys/sys_job/model"
	"manager-gin/src/app/admin/sys/sys_job/service/view"
	"manager-gin/src/common"
)

var sysJobDao = model.SysJobDaoApp
var viewUtils = view.SysJobViewUtilsApp

type SysJobService struct{}

// Create 创建SysJob记录
// Author
func (sysJobService *SysJobService) Create(sysJobView *view.SysJobView) (err error) {
	err1, sysJob := viewUtils.View2Data(sysJobView)
	if err1 != nil {
		return err1
	}
	err2 := sysJobDao.Create(*sysJob)
	if err2 != nil {
		return err2
	}
	return nil
}

// Delete 删除SysJob记录
// Author
func (sysJobService *SysJobService) Delete(id int) (err error) {
	err = sysJobDao.Delete(id)
	return err
}

// DeleteByIds 批量删除SysJob记录
// Author
func (sysJobService *SysJobService) DeleteByIds(ids []int) (err error) {
	err = sysJobDao.DeleteByIds(ids)
	return err
}

// Update 更新SysJob记录
// Author
func (sysJobService *SysJobService) Update(id int, sysJobView *view.SysJobView) (err error) {
	sysJobView.Id = id
	err1, sysJob := viewUtils.View2Data(sysJobView)
	if err1 != nil {
		return err1
	}
	err = sysJobDao.Update(*sysJob)
	return err
}

// Get 根据id获取SysJob记录
// Author
func (sysJobService *SysJobService) Get(id int) (err error, sysJobView *view.SysJobView) {
	err1, sysJob := sysJobDao.Get(id)
	if err1 != nil {
		return err1, nil
	}
	err2, sysJobView := viewUtils.Data2View(sysJob)
	if err2 != nil {
		return err2, nil
	}
	return
}

// Find 分页获取SysJob记录
// Author
func (sysJobService *SysJobService) Find(info *common.PageInfoV2) (err error) {
	err1, sysJobs, total := sysJobDao.Find(info)
	if err1 != nil {
		return err1
	}
	info.Total = total
	err2, viewList := viewUtils.Data2ViewList(sysJobs)
	if err2 != nil {
		return err2
	}
	info.FormList = viewList
	return err
}
