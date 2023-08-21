// Package service 自动生成模板 SysConfigService
// @description <TODO description class purpose>
// @author
// @File: sys_config
// @version 1.0.0
// @create 2023-08-21 14:20:37
package service

import (
	"errors"
	"manager-gin/src/app/admin/sys/sys_config/model"
	"manager-gin/src/app/admin/sys/sys_config/service/view"
	"manager-gin/src/common"
)

var sysConfigDao = model.SysConfigDaoApp
var viewUtils = view.SysConfigViewUtilsApp

type SysConfigService struct{}

// Create 创建SysConfig记录
// Author
func (service *SysConfigService) Create(sysConfigView *view.SysConfigView) (err error) {
	if err1, sysConfig := viewUtils.View2Data(sysConfigView); err1 != nil {
		return errors.New("数据解析失败")
	} else {
		return sysConfigDao.Create(*sysConfig)
	}
}

// DeleteByIds 批量删除SysConfig记录
// Author
func (service *SysConfigService) DeleteByIds(ids []string) (err error) {
	err = sysConfigDao.DeleteByIds(ids)
	return err
}

// Update 更新SysConfig记录
// Author
func (service *SysConfigService) Update(id string, sysConfigView *view.SysConfigView) (err error) {
	sysConfigView.Id = id
	if err1, sysConfig := viewUtils.View2Data(sysConfigView); err1 != nil {
		return err1
	} else {
		return sysConfigDao.Update(*sysConfig)
	}
}

// Get 根据id获取SysConfig记录
// Author
func (service *SysConfigService) Get(id string) (err error, sysConfigView *view.SysConfigView) {
	err1, sysConfig := sysConfigDao.Get(id)
	if err1 != nil {
		return err1, nil
	}
	err, sysConfigView = viewUtils.Data2View(sysConfig)
	return
}

// Page 分页获取SysConfig记录
// Author
func (service *SysConfigService) Page(pageInfo *view.SysConfigPageView) (err error, res *common.PageInfo) {
	err, param, page := viewUtils.Page2Data(pageInfo)
	if err != nil {
		return err, nil
	}
	err1, datas, total := sysConfigDao.Page(param, page)
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

func (service *SysConfigService) List(v *view.SysConfigView) (err error, views *[]view.SysConfigView) {
	err, data := viewUtils.View2Data(v)
	if err != nil {
		return err, nil
	}
	var datas *[]model.SysConfig
	if err, datas = sysConfigDao.List(data); err != nil {
		return err, nil
	} else {
		err, views = viewUtils.Data2ViewList(datas)
		return
	}
}
