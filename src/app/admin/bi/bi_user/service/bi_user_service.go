// Package service 自动生成模板 BiUserService
// @description <TODO description class purpose>
// @author Administrator
// @File: bi_user
// @version 1.0.0
// @create 2023-08-16 14:53:36
package service

import (
	"manager-gin/src/app/admin/bi/bi_user/model"
	"manager-gin/src/app/admin/bi/bi_user/service/view"
	"manager-gin/src/common"
)

var biUserDao = model.BiUserDaoApp
var viewUtils = view.BiUserViewUtilsApp

type BiUserService struct{}

// Create 创建BiUser记录
// Author Administrator
func (biUserService *BiUserService) Create(biUserView *view.BiUserView) (err error) {
	err1, biUser := viewUtils.View2Data(biUserView)
	if err1 != nil {
		return err1
	}
	err2 := biUserDao.Create(*biUser)
	if err2 != nil {
		return err2
	}
	return nil
}

// Delete 删除BiUser记录
// Author Administrator
func (biUserService *BiUserService) Delete(id string) (err error) {
	err = biUserDao.Delete(id)
	return err
}

// DeleteByIds 批量删除BiUser记录
// Author Administrator
func (biUserService *BiUserService) DeleteByIds(ids []string) (err error) {
	err = biUserDao.DeleteByIds(ids)
	return err
}

// Update 更新BiUser记录
// Author Administrator
func (biUserService *BiUserService) Update(id string, biUserView *view.BiUserView) (err error) {
	biUserView.Id = id
	err1, biUser := viewUtils.View2Data(biUserView)
	if err1 != nil {
		return err1
	}
	err = biUserDao.Update(*biUser)
	return err
}

// Get 根据id获取BiUser记录
// Author Administrator
func (biUserService *BiUserService) Get(id string) (err error, biUserView *view.BiUserView) {
	err1, biUser := biUserDao.Get(id)
	if err1 != nil {
		return err1, nil
	}
	err2, biUserView := viewUtils.Data2View(biUser)
	if err2 != nil {
		return err2, nil
	}
	return
}

// Find 分页获取BiUser记录
// Author Administrator
func (biUserService *BiUserService) Find(info *common.PageInfoV2) (err error) {
	err1, biUsers, total := biUserDao.Find(info)
	if err1 != nil {
		return err1
	}
	info.Total = total
	err2, viewList := viewUtils.Data2ViewList(biUsers)
	if err2 != nil {
		return err2
	}
	info.FormList = viewList
	return err
}
