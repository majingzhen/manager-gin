// Package sys_config 自动生成模板 SysConfigService
// @description <TODO description class purpose>
// @author
// @File: sys_config
// @version 1.0.0
// @create 2023-08-21 14:20:37
package sys_config

import (
	"errors"
	"manager-gin/src/app/admin/sys/dao"
	model2 "manager-gin/src/app/admin/sys/model"
	"manager-gin/src/app/admin/sys/service/sys_config/view"
	"manager-gin/src/common"
)

type SysConfigService struct {
	sysConfigDao dao.SysConfigDao
	viewUtils    view.SysConfigViewUtils
}

// Create 创建SysConfig记录
// Author
func (s *SysConfigService) Create(sysConfigView *view.SysConfigView) error {
	if err, sysConfig := s.viewUtils.View2Data(sysConfigView); err != nil {
		return err
	} else {
		return s.sysConfigDao.Create(*sysConfig)
	}
}

// DeleteByIds 批量删除SysConfig记录
// Author
func (s *SysConfigService) DeleteByIds(ids []string) (err error) {
	// 判断是否为系统配置
	for _, id := range ids {
		if err1, sysConfig := s.sysConfigDao.Get(id); err1 != nil {
			return err1
		} else {
			if sysConfig.ConfigType == common.YES {
				return errors.New("系统内置，不可删除")
			}
		}
	}
	err = s.sysConfigDao.DeleteByIds(ids)
	return err
}

// Update 更新SysConfig记录
// Author
func (s *SysConfigService) Update(id string, sysConfigView *view.SysConfigView) (err error) {
	sysConfigView.Id = id
	if err1, sysConfig := s.viewUtils.View2Data(sysConfigView); err1 != nil {
		return err1
	} else {
		return s.sysConfigDao.Update(*sysConfig)
	}
}

// Get 根据id获取SysConfig记录
// Author
func (s *SysConfigService) Get(id string) (err error, sysConfigView *view.SysConfigView) {
	if id == "" {
		return nil, nil
	}
	err1, sysConfig := s.sysConfigDao.Get(id)
	if err1 != nil {
		return err1, nil
	}
	err, sysConfigView = s.viewUtils.Data2View(sysConfig)
	return
}

// Page 分页获取SysConfig记录
// Author
func (s *SysConfigService) Page(pageInfo *view.SysConfigPageView) (err error, res *common.PageInfo) {
	//err, param, page := viewUtils.Page2Data(pageInfo)
	//if err != nil {
	//	return err, nil
	//}
	err, res = s.sysConfigDao.Page(pageInfo)
	if err != nil {
		return err, nil
	}
	return nil, res
	//if err2, viewList := viewUtils.Data2ViewList(datas); err2 != nil {
	//	return err2, res
	//} else {
	//	res = &common.PageInfo{
	//		Total: total,
	//		Rows:  viewList,
	//	}
	//	return err, res
	//}

}

func (s *SysConfigService) List(v *view.SysConfigView) (err error, views []*view.SysConfigView) {
	err, data := s.viewUtils.View2Data(v)
	if err != nil {
		return err, nil
	}
	var datas []*model2.SysConfig
	if err, datas = s.sysConfigDao.List(data); err != nil {
		return err, nil
	} else {
		err, views = s.viewUtils.Data2ViewList(datas)
		return
	}
}

// SelectConfigByKey 根据key查询SysConfig记录
func (s *SysConfigService) SelectConfigByKey(key string) (error, *view.SysConfigView) {
	if err, sysConfig := s.sysConfigDao.SelectConfigByKey(key); err != nil {
		return err, nil
	} else {
		if err, configView := s.viewUtils.Data2View(sysConfig); err != nil {
			return err, nil
		} else {
			return nil, configView
		}
	}
}
