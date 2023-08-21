// Package service 自动生成模板 SysUserService
// @description <TODO description class purpose>
// @author
// @File: sys_user
// @version 1.0.0
// @create 2023-08-18 14:02:24
package service

import (
	"manager-gin/src/app/admin/sys/sys_user/model"
	"manager-gin/src/app/admin/sys/sys_user/service/view"
	"manager-gin/src/common"
)

var sysUserDao = model.SysUserDaoApp
var viewUtils = view.SysUserViewUtilsApp

type SysUserService struct{}

// Create 创建SysUser记录
// Author
func (service *SysUserService) Create(sysUserView *view.SysUserView) (err error) {
	err1, sysUser := viewUtils.View2Data(sysUserView)
	if err1 != nil {
		return err1
	}
	err2 := sysUserDao.Create(*sysUser)
	if err2 != nil {
		return err2
	}
	return nil
}

// Delete 删除SysUser记录
// Author
func (service *SysUserService) Delete(id string) (err error) {
	err = sysUserDao.Delete(id)
	return err
}

// DeleteByIds 批量删除SysUser记录
// Author
func (service *SysUserService) DeleteByIds(ids []string) (err error) {
	err = sysUserDao.DeleteByIds(ids)
	return err
}

// Update 更新SysUser记录
// Author
func (service *SysUserService) Update(id string, sysUserView *view.SysUserView) (err error) {
	sysUserView.Id = id
	err1, sysUser := viewUtils.View2Data(sysUserView)
	if err1 != nil {
		return err1
	}
	err = sysUserDao.Update(*sysUser)
	return err
}

// Get 根据id获取SysUser记录
// Author
func (service *SysUserService) Get(id string) (err error, sysUserView *view.SysUserView) {
	err1, sysUser := sysUserDao.Get(id)
	if err1 != nil {
		return err1, nil
	}
	err2, sysUserView := viewUtils.Data2View(sysUser)
	if err2 != nil {
		return err2, nil
	}
	return
}

// Page 分页获取SysUser记录
// Author
func (service *SysUserService) Page(pageInfo *view.SysUserPageView) (err error, res *common.PageInfo) {
	err, param, page := viewUtils.Page2Data(pageInfo)
	if err != nil {
		return err, nil
	}
	err1, datas, total := sysUserDao.Page(param, page)
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

// List 获取SysUser记录
func (service *SysUserService) List(v *view.SysUserView) (err error, views *[]view.SysUserView) {
	err, data := viewUtils.View2Data(v)
	if err != nil {
		return err, nil
	}
	var datas *[]model.SysUser
	if err, datas = sysUserDao.List(data); err != nil {
		return err, nil
	} else {
		err, views = viewUtils.Data2ViewList(datas)
		return
	}
}

// GetByUserName 根据userName获取SysUser记录
// Author
func (service *SysUserService) GetByUserName(userName string) (err error, sysUserView *view.SysUserView) {
	err1, sysUser := sysUserDao.GetByUserName(userName)
	if err1 != nil {
		return err1, nil
	}
	err2, sysUserView := viewUtils.Data2View(sysUser)
	if err2 != nil {
		return err2, nil
	}
	return
}

// IsAdmin 用户是否管理员
func (service *SysUserService) IsAdmin(userId string) (itIs bool) {
	if common.SYSTEM_ADMIN_ID == userId {
		itIs = true
	}
	return
}

// GetByDeptId 根据部门id获取SysUser记录
func (service *SysUserService) GetByDeptId(deptId string) (err error, sysUserView *[]view.SysUserView) {
	err1, sysUser := sysUserDao.GetByDeptId(deptId)
	if err1 != nil {
		return err1, nil
	}
	err2, sysUserView := viewUtils.Data2ViewList(sysUser)
	if err2 != nil {
		return err2, nil
	}
	return
}
