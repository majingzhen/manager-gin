// Package sys_user 自动生成模板 SysUserService
// @description <TODO description class purpose>
// @author
// @File: sys_user
// @version 1.0.0
// @create 2023-08-18 14:02:24
package sys_user

import (
	"errors"
	"gorm.io/gorm"
	"manager-gin/src/app/admin/sys/dao"
	"manager-gin/src/app/admin/sys/model"
	"manager-gin/src/app/admin/sys/service/sys_dept"
	"manager-gin/src/app/admin/sys/service/sys_role"
	"manager-gin/src/app/admin/sys/service/sys_user/view"
	"manager-gin/src/common"
	"manager-gin/src/framework/aspect"
	"manager-gin/src/global"
	"manager-gin/src/utils"
)

type SysUserService struct {
	sysUserDao  dao.SysUserDao
	viewUtils   view.SysUserViewUtils
	deptService sys_dept.SysDeptService
	roleService sys_role.SysRoleService
	userRoleDao dao.SysUserRoleDao
	userPostDao dao.SysUserPostDao
}

// Create 创建SysUser记录
// Author
func (s *SysUserService) Create(sysUserView *view.SysUserView) (err error) {
	err1, sysUser := s.viewUtils.View2Data(sysUserView)
	if err1 != nil {
		return err1
	}
	tx := global.GOrmDao.Begin()
	if err = s.sysUserDao.Create(tx, *sysUser); err != nil {
		return err
	} else {
		if sysUserView.RoleIds != nil && len(sysUserView.RoleIds) > 0 {
			if err2 := s.insertUserRole(tx, sysUser.Id, sysUserView.RoleIds); err2 != nil {
				tx.Rollback()
				return err2
			}
		}
		if sysUserView.PostIds != nil && len(sysUserView.PostIds) > 0 {
			if err3 := s.insertUserPost(tx, sysUser.Id, sysUserView.PostIds); err3 != nil {
				tx.Rollback()
				return err3
			}
		}
		// 提交事务
		tx.Commit()
		return nil
	}
}

// insertUserPost 插入用户岗位关联数据
func (s *SysUserService) insertUserPost(tx *gorm.DB, id string, ids []string) error {
	var userPosts []model.SysUserPost
	for _, postId := range ids {
		userPosts = append(userPosts, model.SysUserPost{
			UserId: id,
			PostId: postId,
		})
	}
	return s.userPostDao.CreateBatch(tx, userPosts)
}

// insertUserRole 插入用户角色关联数据
func (s *SysUserService) insertUserRole(tx *gorm.DB, userId string, roleIds []string) error {
	var userRoles []model.SysUserRole
	for _, roleId := range roleIds {
		userRoles = append(userRoles, model.SysUserRole{
			UserId: userId,
			RoleId: roleId,
		})
	}
	return s.userRoleDao.CreateBatch(tx, userRoles)
}

// DeleteByIds 批量删除SysUser记录
// Author
func (s *SysUserService) DeleteByIds(ids []string, loginUserId string) (err error) {
	for _, id := range ids {
		if common.SYSTEM_ADMIN_ID == id {
			return errors.New("不允许操作超级管理员用户")
		}
		if err = s.CheckUserDataScope(id, loginUserId); err != nil {
			return err
		}
	}
	tx := global.GOrmDao.Begin()
	// 删除用户角色关联数据
	if err = s.userRoleDao.DeleteByUserIds(tx, ids); err != nil {
		tx.Rollback()
		return err
	}
	// 删除用户岗位关联数据
	if err = s.userPostDao.DeleteByUserIds(tx, ids); err != nil {
		tx.Rollback()
		return err
	}
	err = s.sysUserDao.DeleteByIds(tx, ids)
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()
	return err
}

// Update 更新SysUser记录
// Author
func (s *SysUserService) Update(id string, sysUserView *view.SysUserView) (err error) {
	sysUserView.Id = id
	err1, sysUser := s.viewUtils.View2Data(sysUserView)
	if err1 != nil {
		return err1
	}
	err = s.sysUserDao.Update(*sysUser)
	return err
}

// Get 根据id获取SysUser记录
// Author
func (s *SysUserService) Get(id string) (err error, sysUserView *view.SysUserView) {
	if id == "" {
		return nil, nil
	}
	err1, sysUser := s.sysUserDao.Get(id)
	if err1 != nil {
		return err1, nil
	}
	if err, sysUserView = s.viewUtils.Data2View(sysUser); err != nil {
		return err, nil
	} else {
		if err2, deptView := s.deptService.Get(sysUserView.DeptId); err2 != nil {
			return err2, nil
		} else {
			sysUserView.Dept = deptView
		}
		// 组装角色信息
		if err3, roles := s.roleService.AssembleRolesByUserId(id); err3 != nil {
			return err3, nil
		} else {
			sysUserView.Roles = roles
		}
		return
	}
}

// Page 分页获取SysUser记录
// Author
func (s *SysUserService) Page(pageInfo *view.SysUserPageView, user *view.SysUserView) (err error, res *common.PageInfo) {
	pageInfo.DataScopeSql = aspect.DataScopeFilter(user, "d", "u", "")
	if err, res = s.sysUserDao.Page(pageInfo); err != nil {
		return err, res
	}
	if err, res = s.viewUtils.PageData2ViewList(res); err != nil {
		return err, res
	} else {
		if o, ok := res.Rows.([]*view.SysUserView); ok {
			// 组装部门数据
			for i := 0; i < len(o); i++ {
				deptId := o[i].DeptId
				if err3, deptView := s.deptService.Get(deptId); err3 != nil {
					return err3, nil
				} else {
					o[i].Dept = deptView
				}
			}
		}
		return err, res
	}
}

// List 获取SysUser记录
func (s *SysUserService) List(v *view.SysUserView) (err error, views []*view.SysUserView) {
	err, data := s.viewUtils.View2Data(v)
	if err != nil {
		return err, nil
	}
	var datas []*model.SysUser
	if err, datas = s.sysUserDao.List(data); err != nil {
		return err, nil
	} else {
		err, views = s.viewUtils.Data2ViewList(datas)
		return
	}
}

// GetByUserName 根据userName获取SysUser记录
// Author
func (s *SysUserService) GetByUserName(userName string) (err error, sysUserView *view.SysUserView) {
	err1, sysUser := s.sysUserDao.GetByUserName(userName)
	if err1 != nil {
		return err1, nil
	}
	err2, sysUserView := s.viewUtils.Data2View(sysUser)
	if err2 != nil {
		return err2, nil
	}
	return
}

// CheckFieldUnique 校验字段是否唯一
// Author
func (s *SysUserService) CheckFieldUnique(fieldName, value, id string) error {
	if fieldName == "" || value == "" {
		return nil
	}
	if err, data := s.sysUserDao.SelectByField(fieldName, value); err != nil {
		return err
	} else {
		if data != nil && data.Id != id {
			return errors.New("数据重复")
		}
		return nil
	}
}

// CheckUserDataScope 校验数据权限
func (s *SysUserService) CheckUserDataScope(userId, loginUserId string) error {
	if common.SYSTEM_ADMIN_ID != loginUserId {
		err, userView := s.Get(userId)
		if err != nil {
			return err
		}
		// 数据权限控制
		// err, data := s.viewUtils.View2Data(userView)
		if err != nil {
			return err
		}
		filter := aspect.DataScopeFilter(userView, "d", "u", "")
		param := &model.SysUser{}
		param.Id = userId
		param.DataScopeSql = filter
		// data.DataScopeSql = filter
		err, _ = s.sysUserDao.List(param)
		if err != nil {
			return err
		}
	}
	return nil
}

// ResetPwd 重置密码
func (s *SysUserService) ResetPwd(v *view.SysUserView) error {
	err, sysUser := s.viewUtils.View2Data(v)
	if err != nil {
		return err
	}
	salt := utils.GenUID()
	sysUser.Password = utils.EncryptionPassword(sysUser.Password, salt)
	sysUser.Salt = salt
	return s.sysUserDao.Update(*sysUser)
}

// ChangeStatus 更新状态
func (s *SysUserService) ChangeStatus(v *view.SysUserView) error {
	err, sysUser := s.viewUtils.View2Data(v)
	if err != nil {
		return err
	}
	return s.sysUserDao.Update(*sysUser)
}

// AuthRole	角色授权
func (s *SysUserService) AuthRole(v *view.SysUserView) error {
	tx := global.GOrmDao.Begin()
	// 删除用户角色关联数据
	if err := s.userRoleDao.DeleteByUserIds(tx, []string{v.Id}); err != nil {
		tx.Rollback()
		return err
	}
	// 插入用户角色关联数据
	if err := s.insertUserRole(tx, v.Id, v.RoleIds); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

// SelectAllocatedList 获取已分配用户角色的用户列表
func (s *SysUserService) SelectAllocatedList(pageInfo *view.SysUserPageView, user *view.SysUserView) (err error, res *common.PageInfo) {
	pageInfo.DataScopeSql = aspect.DataScopeFilter(user, "d", "u", "")
	if err, res = s.sysUserDao.SelectAllocatedList(pageInfo); err != nil {
		return err, nil
	}
	if err, res = s.viewUtils.PageData2ViewList(res); err != nil {
		return err, res
	} else {
		if o, ok := res.Rows.([]*view.SysUserView); ok {
			// 组装部门数据
			for i := 0; i < len(o); i++ {
				deptId := o[i].DeptId
				if err3, deptView := s.deptService.Get(deptId); err3 != nil {
					return err3, nil
				} else {
					o[i].Dept = deptView
				}
			}
		}
		return err, res
	}
}

// SelectUnallocatedList 获取未分配用户角色的用户列表
func (s *SysUserService) SelectUnallocatedList(pageInfo *view.SysUserPageView, user *view.SysUserView) (err error, res *common.PageInfo) {
	pageInfo.DataScopeSql = aspect.DataScopeFilter(user, "d", "u", "")
	if err, res = s.sysUserDao.SelectUnallocatedList(pageInfo); err != nil {
		return err, nil
	}
	if err, res = s.viewUtils.PageData2ViewList(res); err != nil {
		return err, res
	} else {
		if o, ok := res.Rows.([]*view.SysUserView); ok {
			// 组装部门数据
			for i := 0; i < len(o); i++ {
				deptId := o[i].DeptId
				if err3, deptView := s.deptService.Get(deptId); err3 != nil {
					return err3, nil
				} else {
					o[i].Dept = deptView
				}
			}
		}
		return err, res
	}
}

// GetByDeptId 根据部门id获取SysUser记录
func (s *SysUserService) GetByDeptId(deptId string) (err error, sysUserView []*view.SysUserView) {
	err1, sysUser := s.sysUserDao.GetByDeptId(deptId)
	if err1 != nil {
		return err1, nil
	}
	err2, sysUserView := s.viewUtils.Data2ViewList(sysUser)
	if err2 != nil {
		return err2, nil
	}
	return
}
