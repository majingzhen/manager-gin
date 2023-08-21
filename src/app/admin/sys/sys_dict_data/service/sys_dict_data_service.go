// Package service 自动生成模板 SysDictDataService
// @description <TODO description class purpose>
// @author
// @File: sys_dict_data
// @version 1.0.0
// @create 2023-08-20 19:08:42
package service

import (
	"manager-gin/src/app/admin/sys/sys_dict_data/model"
	"manager-gin/src/app/admin/sys/sys_dict_data/service/view"
	"manager-gin/src/common"
)

var sysDictDataDao = model.SysDictDataDaoApp
var viewUtils = view.SysDictDataViewUtilsApp

type SysDictDataService struct{}

// Create 创建SysDictData记录
// Author
func (service *SysDictDataService) Create(sysDictDataView *view.SysDictDataView) (err error) {
	err1, sysDictData := viewUtils.View2Data(sysDictDataView)
	if err1 != nil {
		return err1
	}
	err2 := sysDictDataDao.Create(*sysDictData)
	if err2 != nil {
		return err2
	}
	return nil
}

// DeleteByIds 批量删除SysDictData记录
// Author
func (service *SysDictDataService) DeleteByIds(ids []string) (err error) {
	err = sysDictDataDao.DeleteByIds(ids)
	return err
}

// Update 更新SysDictData记录
// Author
func (service *SysDictDataService) Update(id string, sysDictDataView *view.SysDictDataView) (err error) {
	sysDictDataView.Id = id
	err1, sysDictData := viewUtils.View2Data(sysDictDataView)
	if err1 != nil {
		return err1
	}
	err = sysDictDataDao.Update(*sysDictData)
	return err
}

// Get 根据id获取SysDictData记录
// Author
func (service *SysDictDataService) Get(id string) (err error, sysDictDataView *view.SysDictDataView) {
	err1, sysDictData := sysDictDataDao.Get(id)
	if err1 != nil {
		return err1, nil
	}
	err2, sysDictDataView := viewUtils.Data2View(sysDictData)
	if err2 != nil {
		return err2, nil
	}
	return
}

// Page 分页获取SysDictData记录
// Author
func (service *SysDictDataService) Page(pageInfo *view.SysDictDataPageView) (err error, res *common.PageInfo) {
	err, param, page := viewUtils.Page2Data(pageInfo)
	if err != nil {
		return err, nil
	}
	err1, datas, total := sysDictDataDao.Page(param, page)
	if err1 != nil {
		return err1, res
	}
	err2, viewList := viewUtils.Data2ViewList(datas)
	if err2 != nil {
		return err2, res
	}
	res = &common.PageInfo{
		Total: total,
		Rows:  viewList,
	}
	return err, res
}

func (service *SysDictDataService) GetByType(dictType string) (err error, views *[]view.SysDictDataView) {
	err1, datas := sysDictDataDao.GetByType(dictType)
	if err1 != nil {
		return err1, nil
	}
	err2, views := viewUtils.Data2ViewList(datas)
	if err2 != nil {
		return err2, nil
	}
	return
}
