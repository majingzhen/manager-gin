// Package service 自动生成模板 SysRoleService
// @description <TODO description class purpose>
// @author
// @File: sys_role
// @version 1.0.0
// @create 2023-08-18 14:00:53
package service

import (
	"errors"
	"fmt"
	"manager-gin/src/app/admin/sys/sys_role/model"
	"manager-gin/src/app/admin/sys/sys_role/service/view"
	userModel "manager-gin/src/app/admin/sys/sys_user/model"
	userView "manager-gin/src/app/admin/sys/sys_user/service/view"
	"manager-gin/src/common"
	"manager-gin/src/framework/aspect"
)

var sysRoleDao = model.SysRoleDaoApp
var viewUtils = view.SysRoleViewUtilsApp
var roleMenuDao = model.SysRoleMenuDaoApp
var roleDeptDao = model.SysRoleDeptDaoApp
var userRoleDao = userModel.SysUserRoleDaoApp

type SysRoleService struct{}

// Create 创建SysRole记录
// Author
func (service *SysRoleService) Create(sysRoleView *view.SysRoleView) (err error) {

	err1, sysRole := viewUtils.View2Data(sysRoleView)
	if err1 != nil {
		return err1
	}
	err2 := sysRoleDao.Create(*sysRole)
	if err2 != nil {
		return err2
	}
	return nil
}

// Delete 删除SysRole记录
// Author
func (service *SysRoleService) Delete(id string) (err error) {
	err = sysRoleDao.Delete(id)
	return err
}

// DeleteByIds 批量删除SysRole记录
// Author
func (service *SysRoleService) DeleteByIds(ids []string, loginUser *userView.SysUserView) (err error) {
	for _, id := range ids {
		if id == common.SYSTEM_ROLE_ADMIN_ID {
			return errors.New("不允许删除超级管理员角色")
		}
		if err := service.CheckRoleDataScope(id, loginUser); err != nil {
			return errors.New("角色数据权限不足，不能删除")
		}
		// 根据角色获取用户
		if err, total := userRoleDao.CountUserRoleByRoleId(id); err != nil {
			return err
		} else if total > 0 {
			if err1, roleView := service.Get(id); err != nil {
				return err1
			} else {
				return errors.New(fmt.Sprintf("%s已分配,不能删除", roleView.RoleName))
			}
		}
	}
	// 删除角色与菜单关联
	if err = roleMenuDao.DeleteRoleMenuByRoleIds(ids); err != nil {
		return err
	}
	// 删除角色与部门关联
	if err = roleDeptDao.DeleteRoleDeptByRoleIds(ids); err != nil {
		return err
	}
	err = sysRoleDao.DeleteByIds(ids)
	return err
}

// Update 更新SysRole记录
// Author
func (service *SysRoleService) Update(id string, sysRoleView *view.SysRoleView) (err error) {
	sysRoleView.Id = id
	err1, sysRole := viewUtils.View2Data(sysRoleView)
	if err1 != nil {
		return err1
	}
	if err = sysRoleDao.Update(*sysRole); err != nil {

		return err
	} else {
		// 删除角色与菜单关联
		if err = roleMenuDao.DeleteRoleMenuByRoleId(id); err != nil {
			return err
		}
		// 插入角色与菜单关联
		if err = insertRoleMenu(id, sysRoleView.MenuIds); err != nil {
			return err
		}
	}
	return err
}

// insertRoleMenu 新增角色菜单信息
func insertRoleMenu(id string, ids []string) error {
	var roleMenus []model.SysRoleMenu
	for _, menuId := range ids {
		roleMenus = append(roleMenus, model.SysRoleMenu{
			RoleId: id,
			MenuId: menuId,
		})
	}
	return roleMenuDao.CreateBatch(roleMenus)
}

// Get 根据id获取SysRole记录
// Author
func (service *SysRoleService) Get(id string) (err error, sysRoleView *view.SysRoleView) {
	if id == "" {
		return nil, nil
	}
	err1, sysRole := sysRoleDao.Get(id)
	if err1 != nil {
		return err1, nil
	}
	err2, sysRoleView := viewUtils.Data2View(sysRole)
	if err2 != nil {
		return err2, nil
	}
	return
}

// Page 分页获取SysRole记录
// Author
func (service *SysRoleService) Page(pageInfo *view.SysRolePageView, sysUserView *userView.SysUserView) (err error, res *common.PageInfo) {
	err, param, page := viewUtils.Page2Data(pageInfo)
	if err != nil {
		return err, nil
	}
	param.DataScopeSql = aspect.DataScopeFilter(sysUserView, "d", "u", "")
	err1, datas, total := sysRoleDao.Page(param, page)
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

// List 获取SysRole列表
// Author
func (service *SysRoleService) List(v *view.SysRoleView, loginUser *userView.SysUserView) (err error, views []*view.SysRoleView) {
	err, data := viewUtils.View2Data(v)
	if err != nil {
		return err, nil
	}
	data.DataScopeSql = aspect.DataScopeFilter(loginUser, "d", "u", "")
	var datas []*model.SysRole
	if err, datas = sysRoleDao.List(data); err != nil {
		return err, nil
	} else {
		err, views = viewUtils.Data2ViewList(datas)
		return
	}
}

// GetRoleByUserId 根据用户获取角色集合
func (service *SysRoleService) GetRoleByUserId(user *userView.SysUserView) (err error, roleNames []string) {
	is := user.Id == common.SYSTEM_ADMIN_ID
	if is {
		roleNames = append(roleNames, "admin")
	} else {
		err1, roles := sysRoleDao.GetRoleByUserId(user.Id)
		if err1 != nil {
			return err1, nil
		}
		for _, role := range roles {
			roleNames = append(roleNames, role.RoleKey)
		}
		_, user.Roles = viewUtils.Data2ViewList(roles)
	}
	return nil, roleNames
}

// SelectRoleAll 查询所有角色
func (service *SysRoleService) SelectRoleAll(loginUser *userView.SysUserView) (err error, roles []*view.SysRoleView) {
	err, roles = service.List(&view.SysRoleView{}, loginUser)
	return
}

// SelectRolesByUserId 根据用户ID查询角色
func (service *SysRoleService) SelectRolesByUserId(userId string) (err error, roles []*view.SysRoleView) {
	err1, datas := sysRoleDao.GetRoleByUserId(userId)
	if err1 != nil {
		return err1, nil
	}
	err, roles = viewUtils.Data2ViewList(datas)
	return
}

// AssembleRolesByUserId 根据用户ID查询授权角色
func (service *SysRoleService) AssembleRolesByUserId(userId string) (error, []*view.SysRoleView) {
	if err, roles := service.SelectRoleAll(nil); err != nil {
		return err, nil
	} else {
		if err1, userRoles := service.SelectRolesByUserId(userId); err1 != nil {
			return err1, nil
		} else {
			for i := 0; i < len(roles); i++ {
				for j := 0; j < len(userRoles); j++ {
					if roles[i].Id == userRoles[j].Id {
						roles[i].Flag = true
						break
					}
				}
			}
		}
		return nil, roles
	}
}

// CheckRoleNameUnique 校验角色名称是否唯一
func (service *SysRoleService) CheckRoleNameUnique(roleName string) error {
	if err, count := sysRoleDao.CheckRoleNameUnique(roleName); err != nil {
		return err
	} else {
		if count > 0 {
			return errors.New("角色名称已存在")
		}
		return nil
	}
}

// CheckRoleKeyUnique 校验角色权限是否唯一
func (service *SysRoleService) CheckRoleKeyUnique(roleKey string) error {
	if err, count := sysRoleDao.CheckRoleKeyUnique(roleKey); err != nil {
		return err
	} else {
		if count > 0 {
			return errors.New("角色权限已存在")
		}
		return nil
	}
}

func (service *SysRoleService) CheckRoleDataScope(id string, loginUser *userView.SysUserView) error {
	if loginUser.Id != common.SYSTEM_ADMIN_ID {
		role := &model.SysRole{
			Id:           id,
			DataScopeSql: aspect.DataScopeFilter(loginUser, "d", "u", ""),
		}
		err, _ := sysRoleDao.List(role)
		if err != nil {
			return err
		}
	}
	return nil
}
