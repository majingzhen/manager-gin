// Package service 自动生成模板 SysPostService
// @description <TODO description class purpose>
// @author
// @File: sys_post
// @version 1.0.0
// @create 2023-08-08 10:06:19
package service

import (
	"manager-gin/src/app/admin/sys/sys_post/model"
	"manager-gin/src/app/admin/sys/sys_post/service/view"
	"manager-gin/src/common"
)

var sysPostDao = model.SysPostDaoApp
var viewUtils = view.SysPostViewUtilsApp

type SysPostService struct{}

// Create 创建SysPost记录
// Author
func (sysPostService *SysPostService) Create(sysPostView *view.SysPostView) (err error) {
	err1, sysPost := viewUtils.View2Data(sysPostView)
	if err1 != nil {
		return err1
	}
	err2 := sysPostDao.Create(*sysPost)
	if err2 != nil {
		return err2
	}
	return nil
}

// Delete 删除SysPost记录
// Author
func (sysPostService *SysPostService) Delete(id int) (err error) {
	err = sysPostDao.Delete(id)
	return err
}

// DeleteByIds 批量删除SysPost记录
// Author
func (sysPostService *SysPostService) DeleteByIds(ids []int) (err error) {
	err = sysPostDao.DeleteByIds(ids)
	return err
}

// Update 更新SysPost记录
// Author
func (sysPostService *SysPostService) Update(id int, sysPostView *view.SysPostView) (err error) {
	sysPostView.Id = id
	err1, sysPost := viewUtils.View2Data(sysPostView)
	if err1 != nil {
		return err1
	}
	err = sysPostDao.Update(*sysPost)
	return err
}

// Get 根据id获取SysPost记录
// Author
func (sysPostService *SysPostService) Get(id int) (err error, sysPostView *view.SysPostView) {
	err1, sysPost := sysPostDao.Get(id)
	if err1 != nil {
		return err1, nil
	}
	err2, sysPostView := viewUtils.Data2View(sysPost)
	if err2 != nil {
		return err2, nil
	}
	return
}

// Find 分页获取SysPost记录
// Author
func (sysPostService *SysPostService) Find(info *common.PageInfoV2) (err error) {
	err1, sysPosts, total := sysPostDao.Find(info)
	if err1 != nil {
		return err1
	}
	info.Total = total
	err2, viewList := viewUtils.Data2ViewList(sysPosts)
	if err2 != nil {
		return err2
	}
	info.FormList = viewList
	return err
}
