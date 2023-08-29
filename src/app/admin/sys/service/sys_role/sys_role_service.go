// Package service 自动生成模板 SysRoleService
// @description <TODO description class purpose>
// @author
// @File: sys_role
// @version 1.0.0
// @create 2023-08-18 14:00:53
package sys_role

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"manager-gin/src/app/admin/sys/dao"
	"manager-gin/src/app/admin/sys/model"
	"manager-gin/src/app/admin/sys/service/sys_role/view"
	userView "manager-gin/src/app/admin/sys/service/sys_user/view"
	"manager-gin/src/common"
	"manager-gin/src/framework/aspect"
	"manager-gin/src/global"
)

type SysRoleService struct {
	sysRoleDao  dao.SysRoleDao
	viewUtils   view.SysRoleViewUtils
	roleMenuDao dao.SysRoleMenuDao
	roleDeptDao dao.SysRoleDeptDao
	userRoleDao dao.SysUserRoleDao
}

// Create 创建SysRole记录
// Author
func (s *SysRoleService) Create(sysRoleView *view.SysRoleView) (err error) {

	err1, sysRole := s.viewUtils.View2Data(sysRoleView)
	if err1 != nil {
		return err1
	}
	tx := global.GOrmDao.Begin()
	err2 := s.sysRoleDao.Create(tx, *sysRole)
	if err2 != nil {
		tx.Rollback()
		return err2
	}
	// 插入角色与菜单关联
	if err = s.insertRoleMenu(tx, sysRole.Id, sysRoleView.MenuIds); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

// Delete 删除SysRole记录
// Author
func (s *SysRoleService) Delete(id string) (err error) {
	err = s.sysRoleDao.Delete(id)
	return err
}

// DeleteByIds 批量删除SysRole记录
// Author
func (s *SysRoleService) DeleteByIds(ids []string, loginUser *userView.SysUserView) (err error) {
	for _, id := range ids {
		if id == common.SYSTEM_ROLE_ADMIN_ID {
			return errors.New("不允许删除超级管理员角色")
		}
		if err := s.CheckRoleDataScope(id, loginUser); err != nil {
			return errors.New("角色数据权限不足，不能删除")
		}
		// 根据角色获取用户
		if err, total := s.userRoleDao.CountUserRoleByRoleId(id); err != nil {
			return err
		} else if total > 0 {
			if err1, roleView := s.Get(id); err != nil {
				return err1
			} else {
				return errors.New(fmt.Sprintf("%s已分配,不能删除", roleView.RoleName))
			}
		}
	}
	tx := global.GOrmDao.Begin()
	// 删除角色与菜单关联
	if err = s.roleMenuDao.DeleteRoleMenuByRoleIds(tx, ids); err != nil {
		tx.Rollback()
		return err
	}
	// 删除角色与部门关联
	if err = s.roleDeptDao.DeleteRoleDeptByRoleIds(tx, ids); err != nil {
		tx.Rollback()
		return err
	}
	if err = s.sysRoleDao.DeleteByIds(tx, ids); err != nil {
		tx.Rollback()
		return err
	} else {
		tx.Commit()
		return nil
	}
}

// Update 更新SysRole记录
// Author
func (s *SysRoleService) Update(id string, sysRoleView *view.SysRoleView) (err error) {
	sysRoleView.Id = id
	err1, sysRole := s.viewUtils.View2Data(sysRoleView)
	if err1 != nil {
		return err1
	}
	tx := global.GOrmDao.Begin()
	if err = s.sysRoleDao.Update(tx, sysRole); err != nil {
		return err
	} else {
		// 删除角色与菜单关联
		if err = s.roleMenuDao.DeleteRoleMenuByRoleId(tx, id); err != nil {
			tx.Rollback()
			return err
		}
		if sysRoleView.MenuIds != nil && len(sysRoleView.MenuIds) > 0 {
			// 插入角色与菜单关联
			if err = s.insertRoleMenu(tx, id, sysRoleView.MenuIds); err != nil {
				tx.Rollback()
				return err
			}
		}
		tx.Commit()
		return nil
	}
}

// insertRoleMenu 新增角色菜单信息
func (s *SysRoleService) insertRoleMenu(tx *gorm.DB, id string, ids []string) error {
	var roleMenus []model.SysRoleMenu
	for _, menuId := range ids {
		roleMenus = append(roleMenus, model.SysRoleMenu{
			RoleId: id,
			MenuId: menuId,
		})
	}
	return s.roleMenuDao.CreateBatch(tx, roleMenus)
}

// Get 根据id获取SysRole记录
// Author
func (s *SysRoleService) Get(id string) (err error, sysRoleView *view.SysRoleView) {
	if id == "" {
		return nil, nil
	}
	err1, sysRole := s.sysRoleDao.Get(id)
	if err1 != nil {
		return err1, nil
	}
	err2, sysRoleView := s.viewUtils.Data2View(sysRole)
	if err2 != nil {
		return err2, nil
	}
	return
}

// Page 分页获取SysRole记录
// Author
func (s *SysRoleService) Page(pageInfo *view.SysRolePageView, sysUserView *userView.SysUserView) (err error, res *common.PageInfo) {
	pageInfo.DataScopeSql = aspect.DataScopeFilter(sysUserView, "d", "u", "")
	if err, res = s.sysRoleDao.Page(pageInfo); err != nil {
		return err, res
	}
	return s.viewUtils.PageData2ViewList(res)
}

// List 获取SysRole列表
// Author
func (s *SysRoleService) List(v *view.SysRoleView, loginUser *userView.SysUserView) (err error, views []*view.SysRoleView) {
	err, data := s.viewUtils.View2Data(v)
	if err != nil {
		return err, nil
	}
	data.DataScopeSql = aspect.DataScopeFilter(loginUser, "d", "u", "")
	var datas []*model.SysRole
	if err, datas = s.sysRoleDao.List(data); err != nil {
		return err, nil
	} else {
		err, views = s.viewUtils.Data2ViewList(datas)
		return
	}
}

// GetRoleByUserId 根据用户获取角色集合
func (s *SysRoleService) GetRoleByUserId(user *userView.SysUserView) (err error, roleNames []string) {
	is := user.Id == common.SYSTEM_ADMIN_ID
	if is {
		roleNames = append(roleNames, "admin")
	} else {
		err1, roles := s.sysRoleDao.GetRoleByUserId(user.Id)
		if err1 != nil {
			return err1, nil
		}
		for _, role := range roles {
			roleNames = append(roleNames, role.RoleKey)
		}
		_, user.Roles = s.viewUtils.Data2ViewList(roles)
	}
	return nil, roleNames
}

// SelectRoleAll 查询所有角色
func (s *SysRoleService) SelectRoleAll(loginUser *userView.SysUserView) (err error, roles []*view.SysRoleView) {
	err, roles = s.List(&view.SysRoleView{}, loginUser)
	return
}

// SelectRolesByUserId 根据用户ID查询角色
func (s *SysRoleService) SelectRolesByUserId(userId string) (err error, roles []*view.SysRoleView) {
	err1, datas := s.sysRoleDao.GetRoleByUserId(userId)
	if err1 != nil {
		return err1, nil
	}
	err, roles = s.viewUtils.Data2ViewList(datas)
	return
}

// AssembleRolesByUserId 根据用户ID查询授权角色
func (s *SysRoleService) AssembleRolesByUserId(userId string) (error, []*view.SysRoleView) {
	if err, roles := s.SelectRoleAll(nil); err != nil {
		return err, nil
	} else {
		if err1, userRoles := s.SelectRolesByUserId(userId); err1 != nil {
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
func (s *SysRoleService) CheckRoleNameUnique(roleName, id string) error {
	if err, data := s.sysRoleDao.CheckRoleNameUnique(roleName); err != nil {
		return err
	} else {
		if data != nil && data.Id != id {
			return errors.New("角色名称已存在")
		}
		return nil
	}
}

// CheckRoleKeyUnique 校验角色权限是否唯一
func (s *SysRoleService) CheckRoleKeyUnique(roleKey string, id string) error {
	if err, data := s.sysRoleDao.CheckRoleKeyUnique(roleKey); err != nil {
		return err
	} else {
		if data != nil && data.Id != id {
			return errors.New("角色权限已存在")
		}
		return nil
	}
}

// CheckRoleDataScope 校验角色是否允许操作
func (s *SysRoleService) CheckRoleDataScope(id string, loginUser *userView.SysUserView) error {
	if loginUser.Id != common.SYSTEM_ADMIN_ID {
		role := &model.SysRole{
			Id:           id,
			DataScopeSql: aspect.DataScopeFilter(loginUser, "d", "u", ""),
		}
		err, _ := s.sysRoleDao.List(role)
		if err != nil {
			return err
		}
	}
	return nil
}

// UpdateStatus 更新状态
func (s *SysRoleService) UpdateStatus(view *view.SysRoleView) error {
	if err, data := s.viewUtils.View2Data(view); err != nil {
		return err
	} else {
		return s.sysRoleDao.Update(global.GOrmDao, data)
	}
}

// AuthDataScope 数据权限
func (s *SysRoleService) AuthDataScope(v *view.SysRoleView) error {
	if err, data := s.viewUtils.View2Data(v); err != nil {
		return err
	} else {
		tx := global.GOrmDao.Begin()
		if err := s.sysRoleDao.Update(tx, data); err != nil {
			tx.Rollback()
			return err
		} else {
			// 删除角色与部门关联
			if err = s.roleDeptDao.DeleteRoleDeptByRoleId(tx, v.Id); err != nil {
				tx.Rollback()
				return err
			}
			if v.DeptIds != nil && len(v.DeptIds) > 0 {
				// 插入角色与部门关联
				if err = s.insertRoleDept(tx, v.Id, v.DeptIds); err != nil {
					tx.Rollback()
					return err
				}
			}
			tx.Commit()
			return nil
		}
	}
}

// CancelAuthUser 取消用户授权
func (s *SysRoleService) CancelAuthUser(v *view.SysUserRoleView) error {
	return s.userRoleDao.DeleteUserRoleInfo(v.UserId, v.RoleId)
}

// BatchCancelAuthUser 批量取消用户授权
func (s *SysRoleService) BatchCancelAuthUser(roleId string, userIds []string) error {
	return s.userRoleDao.DeleteUsersRoleInfo(roleId, userIds)
}

// BatchSelectAuthUser 批量选择用户授权
func (s *SysRoleService) BatchSelectAuthUser(roleId string, userIds []string) error {
	return s.userRoleDao.InsertUsersRoleInfo(roleId, userIds)
}

// insertRoleDept 新增角色部门信息
func (s *SysRoleService) insertRoleDept(tx *gorm.DB, id string, ids []string) error {
	var roleDepts []model.SysRoleDept
	for _, deptId := range ids {
		roleDepts = append(roleDepts, model.SysRoleDept{
			RoleId: id,
			DeptId: deptId,
		})
	}
	return s.roleDeptDao.CreateBatch(tx, roleDepts)
}
