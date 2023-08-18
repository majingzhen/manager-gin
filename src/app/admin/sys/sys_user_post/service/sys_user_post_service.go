// Package service 自动生成模板 SysUserPostService
// @description <TODO description class purpose>
// @author
// @File: sys_user_post
// @version 1.0.0
// @create 2023-08-18 13:41:26
package service

import (
	"manager-gin/src/app/admin/sys/sys_user_post/model"
	"manager-gin/src/app/admin/sys/sys_user_post/service/view"
	"manager-gin/src/common"
)

var sysUserPostDao = model.SysUserPostDaoApp
var viewUtils = view.SysUserPostViewUtilsApp

type SysUserPostService struct{}

// Create 创建SysUserPost记录
// Author
func (sysUserPostService *SysUserPostService) Create(sysUserPostView *view.SysUserPostView) (err error) {
	err1, sysUserPost := viewUtils.View2Data(sysUserPostView)
	if err1 != nil {
		return err1
	}
	err2 := sysUserPostDao.Create(*sysUserPost)
	if err2 != nil {
		return err2
	}
	return nil
}

// Delete 删除SysUserPost记录
// Author
func (sysUserPostService *SysUserPostService) Delete(id string) (err error) {
	err = sysUserPostDao.Delete(id)
	return err
}

// DeleteByIds 批量删除SysUserPost记录
// Author
func (sysUserPostService *SysUserPostService) DeleteByIds(ids []string) (err error) {
	err = sysUserPostDao.DeleteByIds(ids)
	return err
}

// Update 更新SysUserPost记录
// Author
func (sysUserPostService *SysUserPostService) Update(id string, sysUserPostView *view.SysUserPostView) (err error) {
	sysUserPostView.Id = id
	err1, sysUserPost := viewUtils.View2Data(sysUserPostView)
	if err1 != nil {
		return err1
	}
	err = sysUserPostDao.Update(*sysUserPost)
	return err
}

// Get 根据id获取SysUserPost记录
// Author
func (sysUserPostService *SysUserPostService) Get(id string) (err error, sysUserPostView *view.SysUserPostView) {
	err1, sysUserPost := sysUserPostDao.Get(id)
	if err1 != nil {
		return err1, nil
	}
	err2, sysUserPostView := viewUtils.Data2View(sysUserPost)
	if err2 != nil {
		return err2, nil
	}
	return
}

// Find 分页获取SysUserPost记录
// Author
func (sysUserPostService *SysUserPostService) Find(info *common.PageInfoV2) (err error) {
	err1, sysUserPosts, total := sysUserPostDao.Find(info)
	if err1 != nil {
		return err1
	}
	info.Total = total
	err2, viewList := viewUtils.Data2ViewList(sysUserPosts)
	if err2 != nil {
		return err2
	}
	info.FormList = viewList
	return err
}
