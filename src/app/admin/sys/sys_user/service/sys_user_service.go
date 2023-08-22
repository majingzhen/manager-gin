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
		if err2 := insertUserRole(sysUser.Id, sysUserView.RoleIds); err2 != nil {
			return err2
		}
		if err3 := insertUserPost(sysUser.Id, sysUserView.PostIds); err3 != nil {
			return err3
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
	if err, sysUserView = viewUtils.Data2View(sysUser); err != nil {
		return err, nil
	} else {
		if err2, deptView := deptService.Get(sysUserView.DeptId); err2 != nil {
			return err2, nil
		} else {
			sysUserView.Dept = deptView
		}
		// 组装角色信息
		if err3, roles := roleService.SelectRolesByUserId(id); err3 != nil {
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
	param.DataScope = aspect.DataScopeFilter(user, "d", "u", "")
	err1, datas, total := sysUserDao.Page(param, page)
	if err1 != nil {
		return err1, res
	}
	if err2, viewList := viewUtils.Data2ViewList(datas); err2 != nil {
		return err2, res
	} else {
		// 组装部门数据
		for i := 0; i < len(*viewList); i++ {
			deptId := (*viewList)[i].DeptId
			if err3, deptView := deptService.Get(deptId); err3 != nil {
				return err3, nil
			} else {
				(*viewList)[i].Dept = deptView
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

// CheckFieldUnique 校验字段是否唯一
// Author
func (service *SysUserService) CheckFieldUnique(fieldName, value string) error {
	if err, count := sysUserDao.CheckFieldUnique(fieldName, value); err != nil {
		return err
	} else {
		if count > 0 {
			return errors.New("数据重复")
		}
		return nil
	}
}

func (service *SysUserService) CheckUserDataScope(userId string) error {
	if common.SYSTEM_ADMIN_ID != userId {
		err, userView := service.Get(userId)
		if err != nil {
			return err
		}
		// 数据权限控制
		err, data := viewUtils.View2Data(userView)
		if err != nil {
			return err
		}
		filter := aspect.DataScopeFilter(userView, "d", "u", "")
		data.DataScope = filter
		err, _ = sysUserDao.List(data)
		if err != nil {
			return err
		}
	}
	return nil
}
