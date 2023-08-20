// Package service 自动生成模板 SysDictTypeService
// @description <TODO description class purpose>
// @author
// @File: sys_dict_type
// @version 1.0.0
// @create 2023-08-18 13:41:26
package service

import (
	"manager-gin/src/app/admin/sys/sys_dict_type/model"
	"manager-gin/src/app/admin/sys/sys_dict_type/service/view"
	"manager-gin/src/common"
)

var sysDictTypeDao = model.SysDictTypeDaoApp
var viewUtils = view.SysDictTypeViewUtilsApp

type SysDictTypeService struct{}

// Create 创建SysDictType记录
// Author
func (service *SysDictTypeService) Create(sysDictTypeView *view.SysDictTypeView) (err error) {
	err1, sysDictType := viewUtils.View2Data(sysDictTypeView)
	if err1 != nil {
		return err1
	}
	err2 := sysDictTypeDao.Create(*sysDictType)
	if err2 != nil {
		return err2
	}
	return nil
}

// Delete 删除SysDictType记录
// Author
func (service *SysDictTypeService) Delete(id string) (err error) {
	err = sysDictTypeDao.Delete(id)
	return err
}

// DeleteByIds 批量删除SysDictType记录
// Author
func (service *SysDictTypeService) DeleteByIds(ids []string) (err error) {
	err = sysDictTypeDao.DeleteByIds(ids)
	return err
}

// Update 更新SysDictType记录
// Author
func (service *SysDictTypeService) Update(id string, sysDictTypeView *view.SysDictTypeView) (err error) {
	sysDictTypeView.Id = id
	err1, sysDictType := viewUtils.View2Data(sysDictTypeView)
	if err1 != nil {
		return err1
	}
	err = sysDictTypeDao.Update(*sysDictType)
	return err
}

// Get 根据id获取SysDictType记录
// Author
func (service *SysDictTypeService) Get(id string) (err error, sysDictTypeView *view.SysDictTypeView) {
	err1, sysDictType := sysDictTypeDao.Get(id)
	if err1 != nil {
		return err1, nil
	}
	err2, sysDictTypeView := viewUtils.Data2View(sysDictType)
	if err2 != nil {
		return err2, nil
	}
	return
}

// List 分页获取SysDictType记录
// Author
func (service *SysDictTypeService) List(pageInfo *view.SysDictTypePageView) (err error, res *common.PageInfo) {
	err, param, page := viewUtils.Page2Data(pageInfo)
	if err != nil {
		return err, nil
	}
	err1, sysDictTypes, total := sysDictTypeDao.List(param, page)
	if err1 != nil {
		return err1, nil
	}
	err2, viewList := viewUtils.Data2ViewList(sysDictTypes)
	if err2 != nil {
		return err2, nil
	}
	res = &common.PageInfo{
		Total: total,
		Rows:  viewList,
	}
	return nil, res
}

// SelectDictTypeAll 获取全部数据
func (service *SysDictTypeService) SelectDictTypeAll() (err error, views *[]view.SysDictTypeView) {
	err, datas := sysDictTypeDao.SelectDictTypeAll()
	if err != nil {
		return err, nil
	}
	err, views = viewUtils.Data2ViewList(datas)
	if err != nil {
		return err, nil
	}
	return
}
