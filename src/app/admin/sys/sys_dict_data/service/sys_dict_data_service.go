// Package service 自动生成模板 SysDictDataService
// @description <TODO description class purpose>
// @author
// @File: sys_dict_data
// @version 1.0.0
// @create 2023-08-18 13:41:26
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
func (sysDictDataService *SysDictDataService) Create(sysDictDataView *view.SysDictDataView) (err error) {
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

// Delete 删除SysDictData记录
// Author
func (sysDictDataService *SysDictDataService) Delete(id string) (err error) {
	err = sysDictDataDao.Delete(id)
	return err
}

// DeleteByIds 批量删除SysDictData记录
// Author
func (sysDictDataService *SysDictDataService) DeleteByIds(ids []string) (err error) {
	err = sysDictDataDao.DeleteByIds(ids)
	return err
}

// Update 更新SysDictData记录
// Author
func (sysDictDataService *SysDictDataService) Update(id string, sysDictDataView *view.SysDictDataView) (err error) {
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
func (sysDictDataService *SysDictDataService) Get(id string) (err error, sysDictDataView *view.SysDictDataView) {
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

// Find 分页获取SysDictData记录
// Author
func (sysDictDataService *SysDictDataService) Find(info *common.PageInfoV2) (err error) {
	err1, sysDictDatas, total := sysDictDataDao.Find(info)
	if err1 != nil {
		return err1
	}
	info.Total = total
	err2, viewList := viewUtils.Data2ViewList(sysDictDatas)
	if err2 != nil {
		return err2
	}
	info.FormList = viewList
	return err
}
