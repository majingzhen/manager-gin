// Package service 自动生成模板 SysDictTypeService
// @description <TODO description class purpose>
// @author
// @File: sys_dict_type
// @version 1.0.0
// @create 2023-08-18 13:41:26
package sys_dict_type

import (
	"manager-gin/src/app/admin/sys/dao"
	"manager-gin/src/app/admin/sys/service/sys_dict_type/view"
	"manager-gin/src/common"
)

type SysDictTypeService struct {
	sysDictTypeDao dao.SysDictTypeDao
	viewUtils      view.SysDictTypeViewUtils
}

// Create 创建SysDictType记录
// Author
func (s *SysDictTypeService) Create(sysDictTypeView *view.SysDictTypeView) (err error) {
	err1, sysDictType := s.viewUtils.View2Data(sysDictTypeView)
	if err1 != nil {
		return err1
	}
	err2 := s.sysDictTypeDao.Create(*sysDictType)
	if err2 != nil {
		return err2
	}
	return nil
}

// Delete 删除SysDictType记录
// Author
func (s *SysDictTypeService) Delete(id string) (err error) {
	err = s.sysDictTypeDao.Delete(id)
	return err
}

// DeleteByIds 批量删除SysDictType记录
// Author
func (s *SysDictTypeService) DeleteByIds(ids []string) (err error) {
	err = s.sysDictTypeDao.DeleteByIds(ids)
	return err
}

// Update 更新SysDictType记录
// Author
func (s *SysDictTypeService) Update(id string, sysDictTypeView *view.SysDictTypeView) (err error) {
	sysDictTypeView.Id = id
	err1, sysDictType := s.viewUtils.View2Data(sysDictTypeView)
	if err1 != nil {
		return err1
	}
	err = s.sysDictTypeDao.Update(*sysDictType)
	return err
}

// Get 根据id获取SysDictType记录
// Author
func (s *SysDictTypeService) Get(id string) (err error, sysDictTypeView *view.SysDictTypeView) {
	if id == "" {
		return nil, nil
	}
	err1, sysDictType := s.sysDictTypeDao.Get(id)
	if err1 != nil {
		return err1, nil
	}
	err2, sysDictTypeView := s.viewUtils.Data2View(sysDictType)
	if err2 != nil {
		return err2, nil
	}
	return
}

// Page 分页获取SysDictType记录
// Author
func (s *SysDictTypeService) Page(pageInfo *view.SysDictTypePageView) (err error, res *common.PageInfo) {
	if err, res = s.sysDictTypeDao.Page(pageInfo); err != nil {
		return err, nil
	}
	return s.viewUtils.PageData2ViewList(res)
}

// SelectDictTypeAll 获取全部数据
func (s *SysDictTypeService) SelectDictTypeAll() (err error, views []*view.SysDictTypeView) {
	err, datas := s.sysDictTypeDao.SelectDictTypeAll()
	if err != nil {
		return err, nil
	}
	err, views = s.viewUtils.Data2ViewList(datas)
	if err != nil {
		return err, nil
	}
	return
}
