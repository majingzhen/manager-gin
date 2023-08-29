// Package sys_dept Package service 自动生成模板 SysDeptService
// @description <TODO description class purpose>
// @author
// @File: sys_dept
// @version 1.0.0
// @create 2023-08-21 10:27:01
package sys_dept

import (
	"errors"
	"gorm.io/gorm"
	"manager-gin/src/app/admin/sys/dao"
	"manager-gin/src/app/admin/sys/model"
	"manager-gin/src/app/admin/sys/service/sys_dept/view"
	"manager-gin/src/app/admin/sys/service/sys_role"
	"manager-gin/src/app/admin/sys/service/sys_user/extend"
	userView "manager-gin/src/app/admin/sys/service/sys_user/view"
	"manager-gin/src/common"
	"manager-gin/src/framework/aspect"
	"manager-gin/src/global"
	"strings"
)

type SysDeptService struct {
	sysDeptDao  dao.SysDeptDao
	viewUtils   view.SysDeptViewUtils
	userService extend.SysUserExtendService
	roleService sys_role.SysRoleService
}

// Create 创建SysDept记录
// Author
func (s *SysDeptService) Create(sysDeptView *view.SysDeptView) (err error) {
	// 判断名称是否重复
	err, exist := s.sysDeptDao.CheckDeptNameALL(sysDeptView.DeptName, sysDeptView.ParentId)
	if exist {
		return errors.New("部门名称已经存在")
	}
	// 根据前端传入的数据创建model
	if err1, sysDept := s.viewUtils.View2Data(sysDeptView); err1 != nil {
		return err1
	} else {
		if err1, deptView := s.Get(sysDept.ParentId); err1 != nil {
			return errors.New("父级部门不存在")
		} else {
			if deptView.Status == common.DEPT_DISABLE {
				return errors.New("父级部门已经被禁用, 禁止新增")
			}
			sysDept.Ancestors = deptView.Ancestors + "," + sysDept.ParentId
		}
		return s.sysDeptDao.Create(*sysDept)
	}
}

// Delete 删除SysDept记录
// Author
func (s *SysDeptService) Delete(id string) error {
	// 判断是否存在子集
	err, children := s.sysDeptDao.SelectChildrenDeptById(id)
	if err != nil {
		return err
	} else {
		if len(*children) > 0 {
			return errors.New("存在子部门, 不允许删除")
		}
	}
	if err, userView := s.userService.GetByDeptId(id); err != nil {
		return err
	} else {
		if userView != nil {
			return errors.New("部门存在用户, 不允许删除")
		}
	}
	err = s.sysDeptDao.Delete(id)
	return err
}

// Update 更新SysDept记录
// Author
func (s *SysDeptService) Update(id string, sysDeptView *view.SysDeptView) (err error) {
	var oldDept *view.SysDeptView
	// 判断是否存在
	if err, oldDept = s.Get(id); err != nil {
		return errors.New("部门不存在")
	}
	// 判断名称是否重复
	err, exist := s.sysDeptDao.CheckDeptNameALL(sysDeptView.DeptName, sysDeptView.ParentId)
	if exist {
		return errors.New("部门名称已经存在")
	}
	if id == sysDeptView.ParentId {
		return errors.New("上级部门不能是自己")
	}
	sysDeptView.Id = id
	if err1, sysDept := s.viewUtils.View2Data(sysDeptView); err1 != nil {
		return err1
	} else {
		if err1, newParentDept := s.Get(sysDept.ParentId); err1 != nil {
			return errors.New("父级部门不存在")
		} else {
			newAncestors := newParentDept.Ancestors + "," + newParentDept.Id
			oldAncestors := oldDept.Ancestors
			sysDept.Ancestors = newAncestors
			// 更新子部门的祖级列表
			tx := global.GOrmDao.Begin()
			err = s.updateDeptChildren(tx, sysDept.Id, newAncestors, oldAncestors)
			if err != nil {
				tx.Rollback()
				return errors.New("数据更新失败")
			} else {
				if err = s.sysDeptDao.Update(tx, *sysDept); err != nil {
					tx.Rollback()
					return err
				} else {
					tx.Commit()
					return nil
				}
			}
		}
	}
}

// updateDeptChildren 更新子部门的祖级列表
func (s *SysDeptService) updateDeptChildren(tx *gorm.DB, id, newAncestors, oldAncestors string) (err error) {
	var children *[]model.SysDept
	err, children = s.sysDeptDao.SelectChildrenDeptById(id)
	if err != nil {
		return err
	} else {
		for _, child := range *children {
			child.Ancestors = strings.Replace(child.Ancestors, oldAncestors, newAncestors, 1)
			err = s.sysDeptDao.Update(tx, child)
			if err != nil {
				return err
			}
		}
	}
	return err
}

// Get 根据id获取SysDept记录
// Author
func (s *SysDeptService) Get(id string) (err error, sysDeptView *view.SysDeptView) {
	if id == "" {
		return nil, nil
	}
	err1, sysDept := s.sysDeptDao.Get(id)
	if err1 != nil {
		return err1, nil
	}
	err, sysDeptView = s.viewUtils.Data2View(sysDept)
	return err, sysDeptView
}

// Page 分页获取SysDept记录
// Author
func (s *SysDeptService) Page(pageInfo *view.SysDeptPageView) (err error, res *common.PageInfo) {
	if err, res = s.sysDeptDao.Page(pageInfo); err != nil {
		return err, nil
	}
	return s.viewUtils.PageData2ViewList(res)
}

// List 获取SysDept列表
func (s *SysDeptService) List(v *view.SysDeptView, userView *userView.SysUserView) (err error, views []*view.SysDeptView) {
	err, data := s.viewUtils.View2Data(v)
	if err != nil {
		return err, nil
	}
	data.DataScopeSql = aspect.DataScopeFilter(userView, "d", "u", "")
	var datas []*model.SysDept
	if err, datas = s.sysDeptDao.List(data); err != nil {
		return err, nil
	} else {
		err, views = s.viewUtils.Data2ViewList(datas)
		return
	}
}

// SelectDeptTree 获取部门树
func (s *SysDeptService) SelectDeptTree(v *view.SysDeptView, sysUserView *userView.SysUserView) (error, []*view.SysDeptTreeView) {
	err, data := s.viewUtils.View2Data(v)
	if err != nil {
		return err, nil
	}
	data.DataScopeSql = aspect.DataScopeFilter(sysUserView, "d", "u", "")
	var datas []*model.SysDept
	if err, datas = s.sysDeptDao.List(data); err != nil {
		return err, nil
	} else {
		var trees []*view.SysDeptTreeView
		if err, trees = s.viewUtils.Data2TreeList(datas); err != nil {
			return err, nil
		} else {
			deptTrees := getDeptTree(trees)
			return nil, deptTrees
		}

	}
}

// SelectDeptTreeByRole 获取部门树（排除下级）
func (s *SysDeptService) SelectDeptTreeByRole(id string) (error, []string) {
	err, roleView := s.roleService.Get(id)
	if err != nil {
		return err, nil
	}
	return s.sysDeptDao.SelectDeptListByRoleId(id, roleView.DeptCheckStrictly)
}

// getDeptTree 获取部门树
func getDeptTree(departments []*view.SysDeptTreeView) []*view.SysDeptTreeView {
	var rootDepts []*view.SysDeptTreeView

	// 遍历所有部门，找到根节点
	for _, dept := range departments {
		if dept.ParentId == "0" {
			rootDepts = append(rootDepts, dept)
		}
	}
	// 递归获取每个根节点的子部门
	for i := range rootDepts {
		rootDepts[i].Children = getChildren(rootDepts[i].Id, departments)
	}
	return rootDepts
}

// getChildren 获取所有子集部门
func getChildren(parentId string, departments []*view.SysDeptTreeView) []*view.SysDeptTreeView {
	var children []*view.SysDeptTreeView
	// 遍历所有部门，找到指定父节点的子部门
	for _, dept := range departments {
		if dept.ParentId == parentId {
			// 递归获取子部门的子部门
			dept.Children = getChildren(dept.Id, departments)
			children = append(children, dept)
		}
	}
	return children
}
