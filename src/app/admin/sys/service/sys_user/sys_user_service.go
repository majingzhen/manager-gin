// Package service 自动生成模板 SysUserService
// @description <TODO description class purpose>
// @author
// @File: sys_user
// @version 1.0.0
// @create 2023-08-18 14:02:24
package sys_user

import (
	"errors"
	"manager-gin/src/app/admin/sys/dao"
	"manager-gin/src/app/admin/sys/model"
	"manager-gin/src/app/admin/sys/service/sys_dept"
	"manager-gin/src/app/admin/sys/service/sys_role"
	"manager-gin/src/app/admin/sys/service/sys_user/view"
	"manager-gin/src/common"
	"manager-gin/src/framework/aspect"
	"manager-gin/src/utils"
)

var SysUserServiceApp = new(SysUserService)

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
	if err = s.sysUserDao.Create(*sysUser); err != nil {
		return err
	} else {
		if sysUserView.RoleIds != nil && len(sysUserView.RoleIds) > 0 {
			if err2 := s.insertUserRole(sysUser.Id, sysUserView.RoleIds); err2 != nil {
				return err2
			}
		}
		if sysUserView.PostIds != nil && len(sysUserView.PostIds) > 0 {
			if err3 := s.insertUserPost(sysUser.Id, sysUserView.PostIds); err3 != nil {
				return err3
			}
		}
	}
	return nil
}

// insertUserPost 插入用户岗位关联数据
func (s *SysUserService) insertUserPost(id string, ids []string) error {
	var userPosts []model.SysUserPost
	for _, postId := range ids {
		userPosts = append(userPosts, model.SysUserPost{
			UserId: id,
			PostId: postId,
		})
	}
	return s.userPostDao.CreateBatch(userPosts)
}

// insertUserRole 插入用户角色关联数据
func (s *SysUserService) insertUserRole(userId string, roleIds []string) error {
	var userRoles []model.SysUserRole
	for _, roleId := range roleIds {
		userRoles = append(userRoles, model.SysUserRole{
			UserId: userId,
			RoleId: roleId,
		})
	}
	return s.userRoleDao.CreateBatch(userRoles)
}

// Delete 删除SysUser记录
// Author
func (s *SysUserService) Delete(id string) (err error) {
	err = s.sysUserDao.Delete(id)
	return err
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
	// 删除用户角色关联数据
	if err = s.userRoleDao.DeleteByUserIds(ids); err != nil {
		return err
	}
	// 删除用户岗位关联数据
	if err = s.userPostDao.DeleteByUserIds(ids); err != nil {
		return err
	}
	err = s.sysUserDao.DeleteByIds(ids)
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
	err, param, page := s.viewUtils.Page2Data(pageInfo)
	if err != nil {
		return err, nil
	}
	param.DataScopeSql = aspect.DataScopeFilter(user, "d", "u", "")
	err1, datas, total := s.sysUserDao.Page(param, page)
	if err1 != nil {
		return err1, res
	}
	if err2, viewList := s.viewUtils.Data2ViewList(datas); err2 != nil {
		return err2, res
	} else {
		// 组装部门数据
		for i := 0; i < len(viewList); i++ {
			deptId := viewList[i].DeptId
			if err3, deptView := s.deptService.Get(deptId); err3 != nil {
				return err3, nil
			} else {
				viewList[i].Dept = deptView
			}
		}
		res = &common.PageInfo{
			Total: total,
			Rows:  viewList,
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
	// 删除用户角色关联数据
	if err := s.userRoleDao.DeleteByUserIds([]string{v.Id}); err != nil {
		return err
	}
	// 插入用户角色关联数据
	if err := s.insertUserRole(v.Id, v.RoleIds); err != nil {
		return err
	}
	return nil
}

// SelectAllocatedList 获取已分配用户角色的用户列表
func (s *SysUserService) SelectAllocatedList(pageInfo *view.SysUserPageView, user *view.SysUserView) (error, *common.PageInfo) {
	err, param, page := s.viewUtils.Page2Data(pageInfo)
	if err != nil {
		return err, nil
	}
	roleId := pageInfo.RoleId
	param.DataScopeSql = aspect.DataScopeFilter(user, "d", "u", "")
	err1, datas, total := s.sysUserDao.SelectAllocatedList(param, page, roleId)
	if err1 != nil {
		return err1, nil
	}
	if err2, viewList := s.viewUtils.Data2ViewList(datas); err2 != nil {
		return err2, nil
	} else {
		// 组装部门数据
		for i := 0; i < len(viewList); i++ {
			deptId := viewList[i].DeptId
			if err3, deptView := s.deptService.Get(deptId); err3 != nil {
				return err3, nil
			} else {
				viewList[i].Dept = deptView
			}
		}
		res := &common.PageInfo{
			Total: total,
			Rows:  viewList,
		}
		return err, res
	}
}

// SelectUnallocatedList 获取未分配用户角色的用户列表
func (s *SysUserService) SelectUnallocatedList(pageInfo *view.SysUserPageView, user *view.SysUserView) (error, *common.PageInfo) {
	err, param, page := s.viewUtils.Page2Data(pageInfo)
	if err != nil {
		return err, nil
	}
	roleId := pageInfo.RoleId
	param.DataScopeSql = aspect.DataScopeFilter(user, "d", "u", "")
	err1, datas, total := s.sysUserDao.SelectUnallocatedList(param, page, roleId)
	if err1 != nil {
		return err1, nil
	}
	if err2, viewList := s.viewUtils.Data2ViewList(datas); err2 != nil {
		return err2, nil
	} else {
		// 组装部门数据
		for i := 0; i < len(viewList); i++ {
			deptId := viewList[i].DeptId
			if err3, deptView := s.deptService.Get(deptId); err3 != nil {
				return err3, nil
			} else {
				viewList[i].Dept = deptView
			}
		}
		res := &common.PageInfo{
			Total: total,
			Rows:  viewList,
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
