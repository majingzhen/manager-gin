// Package role 自动生成模板 RoleService
// @description <TODO description class purpose>
// @author
// @File: sys_role
// @version 1.0.0
// @create 2023-08-18 14:00:53
package role

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"manager-gin/src/app/admin/sys/dao"
	"manager-gin/src/app/admin/sys/model"
	"manager-gin/src/app/admin/sys/service/role/view"
	userView "manager-gin/src/app/admin/sys/service/user/view"
	"manager-gin/src/common"
	"manager-gin/src/common/constants"
	"manager-gin/src/framework/aspect"
	"manager-gin/src/global"
)

type RoleService struct {
	sysRoleDao  dao.RoleDao
	viewUtils   view.RoleViewUtils
	roleMenuDao dao.RoleMenuDao
	roleDeptDao dao.RoleDeptDao
	userRoleDao dao.UserRoleDao
}

// Create 创建Role记录
// Author
func (s *RoleService) Create(sysRoleView *view.RoleView) (err error) {

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

// Delete 删除Role记录
// Author
func (s *RoleService) Delete(id string) (err error) {
	err = s.sysRoleDao.Delete(id)
	return err
}

// DeleteByIds 批量删除Role记录
// Author
func (s *RoleService) DeleteByIds(ids []string, loginUser *userView.UserView) (err error) {
	for _, id := range ids {
		if id == constants.SYSTEM_ROLE_ADMIN_ID {
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

// Update 更新Role记录
// Author
func (s *RoleService) Update(id string, sysRoleView *view.RoleView) (err error) {
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
func (s *RoleService) insertRoleMenu(tx *gorm.DB, id string, ids []string) error {
	var roleMenus []model.RoleMenu
	for _, menuId := range ids {
		roleMenus = append(roleMenus, model.RoleMenu{
			RoleId: id,
			MenuId: menuId,
		})
	}
	return s.roleMenuDao.CreateBatch(tx, roleMenus)
}

// Get 根据id获取Role记录
// Author
func (s *RoleService) Get(id string) (err error, sysRoleView *view.RoleView) {
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

// Page 分页获取Role记录
// Author
func (s *RoleService) Page(pageInfo *view.RolePageView, sysUserView *userView.UserView) (err error, res *common.PageInfo) {
	pageInfo.DataScopeSql = aspect.DataScopeFilter(sysUserView, "d", "u", "")
	if err, res = s.sysRoleDao.Page(pageInfo); err != nil {
		return err, res
	}
	return s.viewUtils.PageData2ViewList(res)
}

// List 获取Role列表
// Author
func (s *RoleService) List(v *view.RoleView, loginUser *userView.UserView) (err error, views []*view.RoleView) {
	err, data := s.viewUtils.View2Data(v)
	if err != nil {
		return err, nil
	}
	data.DataScopeSql = aspect.DataScopeFilter(loginUser, "d", "u", "")
	var datas []*model.Role
	if err, datas = s.sysRoleDao.List(data); err != nil {
		return err, nil
	} else {
		err, views = s.viewUtils.Data2ViewList(datas)
		return
	}
}

// GetRoleByUserId 根据用户获取角色集合
func (s *RoleService) GetRoleByUserId(user *userView.UserView) (err error, roleNames []string) {
	is := user.Id == constants.SYSTEM_ADMIN_ID
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
func (s *RoleService) SelectRoleAll(loginUser *userView.UserView) (err error, roles []*view.RoleView) {
	err, roles = s.List(&view.RoleView{}, loginUser)
	return
}

// SelectRolesByUserId 根据用户ID查询角色
func (s *RoleService) SelectRolesByUserId(userId string) (err error, roles []*view.RoleView) {
	err1, datas := s.sysRoleDao.GetRoleByUserId(userId)
	if err1 != nil {
		return err1, nil
	}
	err, roles = s.viewUtils.Data2ViewList(datas)
	return
}

// AssembleRolesByUserId 根据用户ID查询授权角色
func (s *RoleService) AssembleRolesByUserId(userId string) (error, []*view.RoleView) {
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
func (s *RoleService) CheckRoleNameUnique(roleName, id string) error {
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
func (s *RoleService) CheckRoleKeyUnique(roleKey string, id string) error {
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
func (s *RoleService) CheckRoleDataScope(id string, loginUser *userView.UserView) error {
	if loginUser.Id != constants.SYSTEM_ADMIN_ID {
		role := &model.Role{
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
func (s *RoleService) UpdateStatus(view *view.RoleView) error {
	if err, data := s.viewUtils.View2Data(view); err != nil {
		return err
	} else {
		return s.sysRoleDao.Update(global.GOrmDao, data)
	}
}

// AuthDataScope 数据权限
func (s *RoleService) AuthDataScope(v *view.RoleView) error {
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
func (s *RoleService) CancelAuthUser(v *view.UserRoleView) error {
	return s.userRoleDao.DeleteUserRoleInfo(v.UserId, v.RoleId)
}

// BatchCancelAuthUser 批量取消用户授权
func (s *RoleService) BatchCancelAuthUser(roleId string, userIds []string) error {
	return s.userRoleDao.DeleteUsersRoleInfo(roleId, userIds)
}

// BatchSelectAuthUser 批量选择用户授权
func (s *RoleService) BatchSelectAuthUser(roleId string, userIds []string) error {
	return s.userRoleDao.InsertUsersRoleInfo(roleId, userIds)
}

// insertRoleDept 新增角色部门信息
func (s *RoleService) insertRoleDept(tx *gorm.DB, id string, ids []string) error {
	var roleDepts []model.RoleDept
	for _, deptId := range ids {
		roleDepts = append(roleDepts, model.RoleDept{
			RoleId: id,
			DeptId: deptId,
		})
	}
	return s.roleDeptDao.CreateBatch(tx, roleDepts)
}
