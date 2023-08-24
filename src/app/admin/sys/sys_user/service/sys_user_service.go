// Package service 自动生成模板 SysUserService
// @description <TODO description class purpose>
// @author
// @File: sys_user
// @version 1.0.0
// @create 2023-08-18 14:02:24
package service

import (
	"errors"
	deptSer "manager-gin/src/app/admin/sys/sys_dept/service"
	roleSer "manager-gin/src/app/admin/sys/sys_role/service"
	"manager-gin/src/app/admin/sys/sys_user/model"
	"manager-gin/src/app/admin/sys/sys_user/service/view"
	"manager-gin/src/common"
	"manager-gin/src/framework/aspect"
	"manager-gin/src/utils"
)

var sysUserDao = model.SysUserDaoApp
var userRoleDao = model.SysUserRoleDaoApp
var userPostDao = model.SysUserPostDaoApp
var viewUtils = view.SysUserViewUtilsApp
var deptService = deptSer.SysDeptServiceApp
var roleService = roleSer.SysRoleServiceApp

type SysUserService struct{}

// Create 创建SysUser记录
// Author
func (service *SysUserService) Create(sysUserView *view.SysUserView) (err error) {
	err1, sysUser := viewUtils.View2Data(sysUserView)
	if err1 != nil {
		return err1
	}
	if err = sysUserDao.Create(*sysUser); err != nil {
		return err
	} else {
		if sysUserView.RoleIds != nil && len(sysUserView.RoleIds) > 0 {
			if err2 := insertUserRole(sysUser.Id, sysUserView.RoleIds); err2 != nil {
				return err2
			}
		}
		if sysUserView.PostIds != nil && len(sysUserView.PostIds) > 0 {
			if err3 := insertUserPost(sysUser.Id, sysUserView.PostIds); err3 != nil {
				return err3
			}
		}
	}
	return nil
}

// insertUserPost 插入用户岗位关联数据
func insertUserPost(id string, ids []string) error {
	var userPosts []model.SysUserPost
	for _, postId := range ids {
		userPosts = append(userPosts, model.SysUserPost{
			UserId: id,
			PostId: postId,
		})
	}
	return userPostDao.CreateBatch(userPosts)
}

// insertUserRole 插入用户角色关联数据
func insertUserRole(userId string, roleIds []string) error {
	var userRoles []model.SysUserRole
	for _, roleId := range roleIds {
		userRoles = append(userRoles, model.SysUserRole{
			UserId: userId,
			RoleId: roleId,
		})
	}
	return userRoleDao.CreateBatch(userRoles)
}

// Delete 删除SysUser记录
// Author
func (service *SysUserService) Delete(id string) (err error) {
	err = sysUserDao.Delete(id)
	return err
}

// DeleteByIds 批量删除SysUser记录
// Author
func (service *SysUserService) DeleteByIds(ids []string, loginUserId string) (err error) {
	for _, id := range ids {
		if common.SYSTEM_ADMIN_ID == id {
			return errors.New("不允许操作超级管理员用户")
		}
		if err = service.CheckUserDataScope(id, loginUserId); err != nil {
			return err
		}
	}
	// 删除用户角色关联数据
	if err = userRoleDao.DeleteByUserIds(ids); err != nil {
		return err
	}
	// 删除用户岗位关联数据
	if err = userPostDao.DeleteByUserIds(ids); err != nil {
		return err
	}
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
	if id == "" {
		return nil, nil
	}
	err1, sysUser := sysUserDao.Get(id)
	if err1 != nil {
		return err1, nil
	}
	if err, sysUserView = viewUtils.Data2View(sysUser); err != nil {
		return err, nil
	} else {
		if err2, deptView := deptService.Get(sysUserView.DeptId); err2 != nil {
			return err2, nil
		} else {
			sysUserView.Dept = deptView
		}
		// 组装角色信息
		if err3, roles := roleService.AssembleRolesByUserId(id); err3 != nil {
			return err3, nil
		} else {
			sysUserView.Roles = roles
		}
		return
	}
}

// Page 分页获取SysUser记录
// Author
func (service *SysUserService) Page(pageInfo *view.SysUserPageView, user *view.SysUserView) (err error, res *common.PageInfo) {
	err, param, page := viewUtils.Page2Data(pageInfo)
	if err != nil {
		return err, nil
	}
	param.DataScopeSql = aspect.DataScopeFilter(user, "d", "u", "")
	err1, datas, total := sysUserDao.Page(param, page)
	if err1 != nil {
		return err1, res
	}
	if err2, viewList := viewUtils.Data2ViewList(datas); err2 != nil {
		return err2, res
	} else {
		// 组装部门数据
		for i := 0; i < len(viewList); i++ {
			deptId := viewList[i].DeptId
			if err3, deptView := deptService.Get(deptId); err3 != nil {
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
func (service *SysUserService) List(v *view.SysUserView) (err error, views []*view.SysUserView) {
	err, data := viewUtils.View2Data(v)
	if err != nil {
		return err, nil
	}
	var datas []*model.SysUser
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

// CheckFieldUnique 校验字段是否唯一
// Author
func (service *SysUserService) CheckFieldUnique(fieldName, value, id string) error {
	if fieldName == "" || value == "" {
		return nil
	}
	if err, data := sysUserDao.SelectByField(fieldName, value); err != nil {
		return err
	} else {
		if data != nil && data.Id != id {
			return errors.New("数据重复")
		}
		return nil
	}
}

// CheckUserDataScope 校验数据权限
func (service *SysUserService) CheckUserDataScope(userId, loginUserId string) error {
	if common.SYSTEM_ADMIN_ID != loginUserId {
		err, userView := service.Get(userId)
		if err != nil {
			return err
		}
		// 数据权限控制
		// err, data := viewUtils.View2Data(userView)
		if err != nil {
			return err
		}
		filter := aspect.DataScopeFilter(userView, "d", "u", "")
		param := &model.SysUser{}
		param.Id = userId
		param.DataScopeSql = filter
		// data.DataScopeSql = filter
		err, _ = sysUserDao.List(param)
		if err != nil {
			return err
		}
	}
	return nil
}

// ResetPwd 重置密码
func (service *SysUserService) ResetPwd(v *view.SysUserView) error {
	err, sysUser := viewUtils.View2Data(v)
	if err != nil {
		return err
	}
	salt := utils.GenUID()
	sysUser.Password = utils.EncryptionPassword(sysUser.Password, salt)
	sysUser.Salt = salt
	return sysUserDao.Update(*sysUser)
}

// ChangeStatus 更新状态
func (service *SysUserService) ChangeStatus(v *view.SysUserView) error {
	err, sysUser := viewUtils.View2Data(v)
	if err != nil {
		return err
	}
	return sysUserDao.Update(*sysUser)
}

func (service *SysUserService) AuthRole(v *view.SysUserView) error {
	// 删除用户角色关联数据
	if err := userRoleDao.DeleteByUserIds([]string{v.Id}); err != nil {
		return err
	}
	// 插入用户角色关联数据
	if err := insertUserRole(v.Id, v.RoleIds); err != nil {
		return err
	}
	return nil
}
