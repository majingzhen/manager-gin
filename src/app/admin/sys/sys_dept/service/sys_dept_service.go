// Package service 自动生成模板 SysDeptService
// @description <TODO description class purpose>
// @author
// @File: sys_dept
// @version 1.0.0
// @create 2023-08-21 10:27:01
package service

import (
	"errors"
	"manager-gin/src/app/admin/sys/sys_dept/model"
	"manager-gin/src/app/admin/sys/sys_dept/service/view"
	"manager-gin/src/app/admin/sys/sys_user/service"
	"manager-gin/src/common"
	"strings"
)

var sysDeptDao = model.SysDeptDaoApp
var viewUtils = view.SysDeptViewUtilsApp
var userService = service.SysUserServiceApp

type SysDeptService struct{}

// Create 创建SysDept记录
// Author
func (service *SysDeptService) Create(sysDeptView *view.SysDeptView) (err error) {
	// 判断名称是否重复
	err, exist := sysDeptDao.CheckDeptNameALL(sysDeptView.DeptName, sysDeptView.ParentId)
	if exist {
		return errors.New("部门名称已经存在")
	}
	// 根据前端传入的数据创建model
	if err1, sysDept := viewUtils.View2Data(sysDeptView); err1 != nil {
		return errors.New("数据解析失败")
	} else {
		if err1, deptView := service.Get(sysDept.ParentId); err1 != nil {
			return errors.New("父级部门不存在")
		} else {
			if deptView.Status == common.DEPT_DISABLE {
				return errors.New("父级部门已经被禁用, 禁止新增")
			}
			sysDept.Ancestors = deptView.Ancestors + "," + sysDept.ParentId
		}
		return sysDeptDao.Create(*sysDept)
	}
}

// Delete 删除SysDept记录
// Author
func (service *SysDeptService) Delete(id string) error {
	// 判断是否存在子集
	err, children := sysDeptDao.SelectChildrenDeptById(id)
	if err != nil {
		return err
	} else {
		if len(*children) > 0 {
			return errors.New("存在子部门, 不允许删除")
		}
	}
	//TODO 判断是否存在用户
	if err, userView := userService.GetByDeptId(id); err != nil {
		return err
	} else {
		if userView != nil {
			return errors.New("部门存在用户, 不允许删除")
		}
	}
	err = sysDeptDao.Delete(id)
	return err
}

// Update 更新SysDept记录
// Author
func (service *SysDeptService) Update(id string, sysDeptView *view.SysDeptView) (err error) {
	var oldDept *view.SysDeptView
	// 判断是否存在
	if err, oldDept = service.Get(id); err != nil {
		return errors.New("部门不存在")
	}
	// 判断名称是否重复
	err, exist := sysDeptDao.CheckDeptNameALL(sysDeptView.DeptName, sysDeptView.ParentId)
	if exist {
		return errors.New("部门名称已经存在")
	}
	if id == sysDeptView.ParentId {
		return errors.New("上级部门不能是自己")
	}
	sysDeptView.Id = id
	if err1, sysDept := viewUtils.View2Data(sysDeptView); err1 != nil {
		return errors.New("数据解析失败")
	} else {
		if err1, newParentDept := service.Get(sysDept.ParentId); err1 != nil {
			return errors.New("父级部门不存在")
		} else {
			newAncestors := newParentDept.Ancestors + "," + newParentDept.Id
			oldAncestors := oldDept.Ancestors
			sysDept.Ancestors = newAncestors
			// 更新子部门的祖级列表
			err = updateDeptChildren(sysDept.Id, newAncestors, oldAncestors)
			if err != nil {
				return errors.New("数据更新失败")
			}
		}
		return sysDeptDao.Update(*sysDept)
	}
}

// updateDeptChildren 更新子部门的祖级列表
func updateDeptChildren(id, newAncestors, oldAncestors string) (err error) {
	var children *[]model.SysDept
	err, children = sysDeptDao.SelectChildrenDeptById(id)
	if err != nil {
		return err
	} else {
		for _, child := range *children {
			child.Ancestors = strings.Replace(child.Ancestors, oldAncestors, newAncestors, 1)
			err = sysDeptDao.Update(child)
			if err != nil {
				return err
			}
		}
	}
	return err
}

// Get 根据id获取SysDept记录
// Author
func (service *SysDeptService) Get(id string) (err error, sysDeptView *view.SysDeptView) {
	err1, sysDept := sysDeptDao.Get(id)
	if err1 != nil {
		return err1, nil
	}
	err, sysDeptView = viewUtils.Data2View(sysDept)
	return err, sysDeptView
}

// Page 分页获取SysDept记录
// Author
func (service *SysDeptService) Page(pageInfo *view.SysDeptPageView) (err error, res *common.PageInfo) {
	err, param, page := viewUtils.Page2Data(pageInfo)
	if err != nil {
		return err, nil
	}
	err1, datas, total := sysDeptDao.Page(param, page)
	if err1 != nil {
		return err1, nil
	}
	if err2, viewList := viewUtils.Data2ViewList(datas); err2 != nil {
		return err2, nil
	} else {
		res = &common.PageInfo{
			Total: total,
			Rows:  viewList,
		}
		return err, res
	}
}

// List 获取SysDept列表
func (service *SysDeptService) List(v *view.SysDeptView) (err error, views *[]view.SysDeptView) {
	err, data := viewUtils.View2Data(v)
	if err != nil {
		return err, nil
	}
	var datas *[]model.SysDept
	if err, datas = sysDeptDao.List(data); err != nil {
		return err, nil
	} else {
		err, views = viewUtils.Data2ViewList(datas)
		return
	}
}

// SelectDeptTree 获取部门树
func (service *SysDeptService) SelectDeptTree(v *view.SysDeptView) (error, *[]view.SysDeptView) {
	err, data := viewUtils.View2Data(v)
	if err != nil {
		return err, nil
	}
	var datas *[]model.SysDept
	if err, datas = sysDeptDao.List(data); err != nil {
		return err, nil
	} else {
		var views *[]view.SysDeptView
		if err, views = viewUtils.Data2ViewList(datas); err != nil {
			return err, nil
		} else {
			return buildDeptTree(views)
		}

	}
}

func buildDeptTree(depts *[]view.SysDeptView) (error, *[]view.SysDeptView) {
	var deptTree []view.SysDeptView
	for _, dept := range *depts {
		if dept.ParentId == "0" {
			deptTree = append(deptTree, dept)
		}
	}
	for _, dept := range deptTree {
		dept.Children = getChildren(dept.Id, *depts)
	}
	return nil, &deptTree
}

func getChildren(id string, views []view.SysDeptView) *[]view.SysDeptView {
	var children []view.SysDeptView
	for _, dept := range views {
		if dept.ParentId == id {
			children = append(children, dept)
		}
	}
	for _, child := range children {
		child.Children = getChildren(child.Id, views)
	}
	return &children
}
