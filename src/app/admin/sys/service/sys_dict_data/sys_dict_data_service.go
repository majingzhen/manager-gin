// Package sys_dict_data 自动生成模板 SysDictDataService
// @description <TODO description class purpose>
// @author
// @File: sys_dict_data
// @version 1.0.0
// @create 2023-08-20 19:08:42
package sys_dict_data

import (
	"manager-gin/src/app/admin/sys/dao"
	"manager-gin/src/app/admin/sys/service/sys_dict_data/view"
	"manager-gin/src/common"
)

type SysDictDataService struct {
	sysDictDataDao dao.SysDictDataDao
	viewUtils      view.SysDictDataViewUtils
}

// Create 创建SysDictData记录
// Author
func (s *SysDictDataService) Create(sysDictDataView *view.SysDictDataView) (err error) {
	err1, sysDictData := s.viewUtils.View2Data(sysDictDataView)
	if err1 != nil {
		return err1
	}
	err2 := s.sysDictDataDao.Create(*sysDictData)
	if err2 != nil {
		return err2
	}
	return nil
}

// DeleteByIds 批量删除SysDictData记录
// Author
func (s *SysDictDataService) DeleteByIds(ids []string) (err error) {
	err = s.sysDictDataDao.DeleteByIds(ids)
	return err
}

// Update 更新SysDictData记录
// Author
func (s *SysDictDataService) Update(id string, sysDictDataView *view.SysDictDataView) (err error) {
	sysDictDataView.Id = id
	err1, sysDictData := s.viewUtils.View2Data(sysDictDataView)
	if err1 != nil {
		return err1
	}
	err = s.sysDictDataDao.Update(*sysDictData)
	return err
}

// Get 根据id获取SysDictData记录
// Author
func (s *SysDictDataService) Get(id string) (err error, sysDictDataView *view.SysDictDataView) {
	if id == "" {
		return nil, nil
	}
	err1, sysDictData := s.sysDictDataDao.Get(id)
	if err1 != nil {
		return err1, nil
	}
	err2, sysDictDataView := s.viewUtils.Data2View(sysDictData)
	if err2 != nil {
		return err2, nil
	}
	return
}

// Page 分页获取SysDictData记录
// Author
func (s *SysDictDataService) Page(pageInfo *view.SysDictDataPageView) (err error, res *common.PageInfo) {
	err, param, page := s.viewUtils.Page2Data(pageInfo)
	if err != nil {
		return err, nil
	}
	err1, datas, total := s.sysDictDataDao.Page(param, page)
	if err1 != nil {
		return err1, res
	}
	err2, viewList := s.viewUtils.Data2ViewList(datas)
	if err2 != nil {
		return err2, res
	}
	res = &common.PageInfo{
		Total: total,
		Rows:  viewList,
	}
	return err, res
}

func (s *SysDictDataService) GetByType(dictType string) (err error, views []*view.SysDictDataView) {
	err1, datas := s.sysDictDataDao.GetByType(dictType)
	if err1 != nil {
		return err1, nil
	}
	err2, views := s.viewUtils.Data2ViewList(datas)
	if err2 != nil {
		return err2, nil
	}
	return
}
