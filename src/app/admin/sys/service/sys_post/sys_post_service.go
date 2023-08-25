// Package service 自动生成模板 SysPostService
// @description <TODO description class purpose>
// @author
// @File: sys_post
// @version 1.0.0
// @create 2023-08-21 17:37:56
package sys_post

import (
	"errors"
	"fmt"
	"manager-gin/src/app/admin/sys/dao"
	"manager-gin/src/app/admin/sys/model"
	"manager-gin/src/app/admin/sys/service/sys_post/view"
	"manager-gin/src/common"
)

type SysPostService struct {
	sysPostDao dao.SysPostDao
	viewUtils  view.SysPostViewUtils
}

// Create 创建SysPost记录
// Author
func (s *SysPostService) Create(sysPostView *view.SysPostView) error {
	// 校验是否重复
	if err := s.CheckPostCodeUnique(sysPostView.PostCode); err != nil {
		return err
	}
	if err := s.CheckPostNameUnique(sysPostView.PostName); err != nil {
		return err
	}
	if err, sysPost := s.viewUtils.View2Data(sysPostView); err != nil {
		return err
	} else {
		return s.sysPostDao.Create(*sysPost)
	}
}

// DeleteByIds 批量删除SysPost记录
// Author
func (s *SysPostService) DeleteByIds(ids []string) error {
	for _, id := range ids {
		err, postView := s.Get(id)
		if err != nil {
			return err
		}
		if err1, count := s.sysPostDao.CheckPostExistUser(id); err1 != nil {
			return err1
		} else {
			if count > 0 {
				return errors.New(fmt.Sprintf("%s 已分配,不能删除", postView.PostName))
			}
		}
	}
	return s.sysPostDao.DeleteByIds(ids)
}

// Update 更新SysPost记录
// Author
func (s *SysPostService) Update(id string, sysPostView *view.SysPostView) error {
	// 校验是否重复
	if err := s.CheckPostCodeUnique(sysPostView.PostCode); err != nil {
		return err
	}
	if err := s.CheckPostNameUnique(sysPostView.PostName); err != nil {
		return err
	}
	sysPostView.Id = id
	if err, sysPost := s.viewUtils.View2Data(sysPostView); err != nil {
		return err
	} else {
		return s.sysPostDao.Update(*sysPost)
	}
}

// Get 根据id获取SysPost记录
// Author
func (s *SysPostService) Get(id string) (err error, sysPostView *view.SysPostView) {
	if id == "" {
		return nil, nil
	}
	err1, sysPost := s.sysPostDao.Get(id)
	if err1 != nil {
		return err1, nil
	}
	err, sysPostView = s.viewUtils.Data2View(sysPost)
	return
}

// Page 分页获取SysPost记录
// Author
func (s *SysPostService) Page(pageInfo *view.SysPostPageView) (err error, res *common.PageInfo) {
	err, param, page := s.viewUtils.Page2Data(pageInfo)
	if err != nil {
		return err, nil
	}
	err1, datas, total := s.sysPostDao.Page(param, page)
	if err1 != nil {
		return err1, res
	}
	if err2, viewList := s.viewUtils.Data2ViewList(datas); err2 != nil {
		return err2, res
	} else {
		res = &common.PageInfo{
			Total: total,
			Rows:  viewList,
		}
		return err, res
	}

}

// List 获取SysPost列表
// Author
func (s *SysPostService) List(v *view.SysPostView) (err error, views []*view.SysPostView) {
	err, data := s.viewUtils.View2Data(v)
	if err != nil {
		return err, nil
	}
	var datas []*model.SysPost
	if err, datas = s.sysPostDao.List(data); err != nil {
		return err, nil
	} else {
		err, views = s.viewUtils.Data2ViewList(datas)
		return
	}
}

// CheckPostCodeUnique 校验岗位编码是否唯一
// Author
func (s *SysPostService) CheckPostCodeUnique(postCode string) error {
	if err, count := s.sysPostDao.CheckPostCodeUnique(postCode); err != nil {
		return err
	} else {
		if count > 0 {
			return errors.New("岗位编码已存在")
		}
	}
	return nil
}

// CheckPostNameUnique 校验岗位名称是否唯一
// Author
func (s *SysPostService) CheckPostNameUnique(postName string) error {
	if err, count := s.sysPostDao.CheckPostNameUnique(postName); err != nil {
		return err
	} else {
		if count > 0 {
			return errors.New("岗位名称已存在")
		}
	}
	return nil
}

func (s *SysPostService) SelectPostAll() (err error, views []*view.SysPostView) {
	err, views = s.List(&view.SysPostView{})
	return
}

func (s *SysPostService) SelectPostIdListByUserId(userId string) (err error, ids []string) {
	err, dataList := s.sysPostDao.SelectPostListByUserId(userId)
	if err != nil {
		return err, nil
	}
	for _, data := range dataList {
		ids = append(ids, data.Id)
	}
	return
}

// SelectPostListByUserId 根据用户ID查询岗位
func (s *SysPostService) SelectPostListByUserId(userId string) (err error, views []*view.SysPostView) {
	err, dataList := s.sysPostDao.SelectPostListByUserId(userId)
	if err != nil {
		return err, nil
	}
	err, views = s.viewUtils.Data2ViewList(dataList)
	return
}
