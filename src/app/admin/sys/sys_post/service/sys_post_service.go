// Package service 自动生成模板 SysPostService
// @description <TODO description class purpose>
// @author
// @File: sys_post
// @version 1.0.0
// @create 2023-08-21 17:37:56
package service

import (
	"errors"
	"fmt"
	"manager-gin/src/app/admin/sys/sys_post/model"
	"manager-gin/src/app/admin/sys/sys_post/service/view"
	"manager-gin/src/common"
)

var sysPostDao = model.SysPostDaoApp
var viewUtils = view.SysPostViewUtilsApp

type SysPostService struct{}

// Create 创建SysPost记录
// Author
func (service *SysPostService) Create(sysPostView *view.SysPostView) error {
	// 校验是否重复
	if err := service.CheckPostCodeUnique(sysPostView.PostCode); err != nil {
		return err
	}
	if err := service.CheckPostNameUnique(sysPostView.PostName); err != nil {
		return err
	}
	if err, sysPost := viewUtils.View2Data(sysPostView); err != nil {
		return err
	} else {
		return sysPostDao.Create(*sysPost)
	}
}

// DeleteByIds 批量删除SysPost记录
// Author
func (service *SysPostService) DeleteByIds(ids []string) error {
	for _, id := range ids {
		err, postView := service.Get(id)
		if err != nil {
			return err
		}
		if err1, count := sysPostDao.CheckPostExistUser(id); err1 != nil {
			return err1
		} else {
			if count > 0 {
				return errors.New(fmt.Sprintf("%s 已分配,不能删除", postView.PostName))
			}
		}
	}
	return sysPostDao.DeleteByIds(ids)
}

// Update 更新SysPost记录
// Author
func (service *SysPostService) Update(id string, sysPostView *view.SysPostView) error {
	// 校验是否重复
	if err := service.CheckPostCodeUnique(sysPostView.PostCode); err != nil {
		return err
	}
	if err := service.CheckPostNameUnique(sysPostView.PostName); err != nil {
		return err
	}
	sysPostView.Id = id
	if err, sysPost := viewUtils.View2Data(sysPostView); err != nil {
		return err
	} else {
		return sysPostDao.Update(*sysPost)
	}
}

// Get 根据id获取SysPost记录
// Author
func (service *SysPostService) Get(id string) (err error, sysPostView *view.SysPostView) {
	err1, sysPost := sysPostDao.Get(id)
	if err1 != nil {
		return err1, nil
	}
	err, sysPostView = viewUtils.Data2View(sysPost)
	return
}

// Page 分页获取SysPost记录
// Author
func (service *SysPostService) Page(pageInfo *view.SysPostPageView) (err error, res *common.PageInfo) {
	err, param, page := viewUtils.Page2Data(pageInfo)
	if err != nil {
		return err, nil
	}
	err1, datas, total := sysPostDao.Page(param, page)
	if err1 != nil {
		return err1, res
	}
	if err2, viewList := viewUtils.Data2ViewList(datas); err2 != nil {
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
func (service *SysPostService) List(v *view.SysPostView) (err error, views *[]view.SysPostView) {
	err, data := viewUtils.View2Data(v)
	if err != nil {
		return err, nil
	}
	var datas *[]model.SysPost
	if err, datas = sysPostDao.List(data); err != nil {
		return err, nil
	} else {
		err, views = viewUtils.Data2ViewList(datas)
		return
	}
}

// CheckPostCodeUnique 校验岗位编码是否唯一
// Author
func (service *SysPostService) CheckPostCodeUnique(postCode string) error {
	if err, count := sysPostDao.CheckPostCodeUnique(postCode); err != nil {
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
func (service *SysPostService) CheckPostNameUnique(postName string) error {
	if err, count := sysPostDao.CheckPostNameUnique(postName); err != nil {
		return err
	} else {
		if count > 0 {
			return errors.New("岗位名称已存在")
		}
	}
	return nil
}

func (service *SysPostService) SelectPostAll() (err error, views *[]view.SysPostView) {
	err, views = service.List(&view.SysPostView{})
	return
}

func (service *SysPostService) SelectPostIdListByUserId(userId string) (err error, ids []string) {
	err, dataList := sysPostDao.SelectPostListByUserId(userId)
	if err != nil {
		return err, nil
	}
	for _, data := range *dataList {
		ids = append(ids, data.Id)
	}
	return
}

// SelectPostListByUserId 根据用户ID查询岗位
func (service *SysPostService) SelectPostListByUserId(userId string) (err error, views *[]view.SysPostView) {
	err, dataList := sysPostDao.SelectPostListByUserId(userId)
	if err != nil {
		return err, nil
	}
	err, views = viewUtils.Data2ViewList(dataList)
	return
}
