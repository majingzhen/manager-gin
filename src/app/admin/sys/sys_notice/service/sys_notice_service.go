// Package service 自动生成模板 SysNoticeService
// @description <TODO description class purpose>
// @author
// @File: sys_notice
// @version 1.0.0
// @create 2023-08-08 10:06:19
package service

import (
	"manager-gin/src/app/admin/sys/sys_notice/model"
	"manager-gin/src/app/admin/sys/sys_notice/service/view"
	"manager-gin/src/common"
)

var sysNoticeDao = model.SysNoticeDaoApp
var viewUtils = view.SysNoticeViewUtilsApp

type SysNoticeService struct{}

// Create 创建SysNotice记录
// Author
func (sysNoticeService *SysNoticeService) Create(sysNoticeView *view.SysNoticeView) (err error) {
	err1, sysNotice := viewUtils.View2Data(sysNoticeView)
	if err1 != nil {
		return err1
	}
	err2 := sysNoticeDao.Create(*sysNotice)
	if err2 != nil {
		return err2
	}
	return nil
}

// Delete 删除SysNotice记录
// Author
func (sysNoticeService *SysNoticeService) Delete(id int) (err error) {
	err = sysNoticeDao.Delete(id)
	return err
}

// DeleteByIds 批量删除SysNotice记录
// Author
func (sysNoticeService *SysNoticeService) DeleteByIds(ids []int) (err error) {
	err = sysNoticeDao.DeleteByIds(ids)
	return err
}

// Update 更新SysNotice记录
// Author
func (sysNoticeService *SysNoticeService) Update(id int, sysNoticeView *view.SysNoticeView) (err error) {
	sysNoticeView.Id = id
	err1, sysNotice := viewUtils.View2Data(sysNoticeView)
	if err1 != nil {
		return err1
	}
	err = sysNoticeDao.Update(*sysNotice)
	return err
}

// Get 根据id获取SysNotice记录
// Author
func (sysNoticeService *SysNoticeService) Get(id int) (err error, sysNoticeView *view.SysNoticeView) {
	err1, sysNotice := sysNoticeDao.Get(id)
	if err1 != nil {
		return err1, nil
	}
	err2, sysNoticeView := viewUtils.Data2View(sysNotice)
	if err2 != nil {
		return err2, nil
	}
	return
}

// Find 分页获取SysNotice记录
// Author
func (sysNoticeService *SysNoticeService) Find(info *common.PageInfoV2) (err error) {
	err1, sysNotices, total := sysNoticeDao.Find(info)
	if err1 != nil {
		return err1
	}
	info.Total = total
	err2, viewList := viewUtils.Data2ViewList(sysNotices)
	if err2 != nil {
		return err2
	}
	info.FormList = viewList
	return err
}
