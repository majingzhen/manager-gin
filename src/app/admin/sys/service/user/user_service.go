// Package user 自动生成模板 UserService
// @description <TODO description class purpose>
// @author
// @File: user
// @version 1.0.0
// @create 2023-08-18 14:02:24
package user

import (
	"errors"
	"gorm.io/gorm"
	"manager-gin/src/app/admin/sys/dao"
	"manager-gin/src/app/admin/sys/model"
	"manager-gin/src/app/admin/sys/service/dept"
	"manager-gin/src/app/admin/sys/service/role"
	"manager-gin/src/app/admin/sys/service/user/view"
	"manager-gin/src/common"
	"manager-gin/src/common/constants"
	"manager-gin/src/framework/aspect"
	"manager-gin/src/global"
	"manager-gin/src/utils"
)

type UserService struct {
	userDao     dao.UserDao
	viewUtils   view.UserViewUtils
	deptService dept.DeptService
	roleService role.RoleService
	userRoleDao dao.UserRoleDao
	userPostDao dao.UserPostDao
}

// Create 创建User记录
// Author
func (s *UserService) Create(userView *view.UserView) (err error) {
	err1, user := s.viewUtils.View2Data(userView)
	if err1 != nil {
		return err1
	}
	tx := global.GOrmDao.Begin()
	if err = s.userDao.Create(tx, *user); err != nil {
		return err
	} else {
		if userView.RoleIds != nil && len(userView.RoleIds) > 0 {
			if err2 := s.insertUserRole(tx, user.Id, userView.RoleIds); err2 != nil {
				tx.Rollback()
				return err2
			}
		}
		if userView.PostIds != nil && len(userView.PostIds) > 0 {
			if err3 := s.insertUserPost(tx, user.Id, userView.PostIds); err3 != nil {
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
func (s *UserService) insertUserPost(tx *gorm.DB, id string, ids []string) error {
	var userPosts []model.UserPost
	for _, postId := range ids {
		userPosts = append(userPosts, model.UserPost{
			UserId: id,
			PostId: postId,
		})
	}
	return s.userPostDao.CreateBatch(tx, userPosts)
}

// insertUserRole 插入用户角色关联数据
func (s *UserService) insertUserRole(tx *gorm.DB, userId string, roleIds []string) error {
	var userRoles []model.UserRole
	for _, roleId := range roleIds {
		userRoles = append(userRoles, model.UserRole{
			UserId: userId,
			RoleId: roleId,
		})
	}
	return s.userRoleDao.CreateBatch(tx, userRoles)
}

// DeleteByIds 批量删除User记录
// Author
func (s *UserService) DeleteByIds(ids []string, loginUserId string) (err error) {
	for _, id := range ids {
		if constants.SYSTEM_ADMIN_ID == id {
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
	err = s.userDao.DeleteByIds(tx, ids)
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()
	return err
}

// Update 更新User记录
// Author
func (s *UserService) Update(id string, userView *view.UserView) (err error) {
	userView.Id = id
	err1, user := s.viewUtils.View2Data(userView)
	if err1 != nil {
		return err1
	}
	err = s.userDao.Update(*user)
	return err
}

// Get 根据id获取User记录
// Author
func (s *UserService) Get(id string) (err error, userView *view.UserView) {
	if id == "" {
		return nil, nil
	}
	err1, user := s.userDao.Get(id)
	if err1 != nil {
		return err1, nil
	}
	if err, userView = s.viewUtils.Data2View(user); err != nil {
		return err, nil
	} else {
		if err2, deptView := s.deptService.Get(userView.DeptId); err2 != nil {
			return err2, nil
		} else {
			userView.Dept = deptView
		}
		// 组装角色信息
		if err3, roles := s.roleService.AssembleRolesByUserId(id); err3 != nil {
			return err3, nil
		} else {
			userView.Roles = roles
		}
		return
	}
}

// Page 分页获取User记录
// Author
func (s *UserService) Page(pageInfo *view.UserPageView, user *view.UserView) (err error, res *common.PageInfo) {
	pageInfo.DataScopeSql = aspect.DataScopeFilter(user, "d", "u", "")
	if err, res = s.userDao.Page(pageInfo); err != nil {
		return err, res
	}
	if err, res = s.viewUtils.PageData2ViewList(res); err != nil {
		return err, res
	} else {
		if o, ok := res.Rows.([]*view.UserView); ok {
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

// List 获取User记录
func (s *UserService) List(v *view.UserView) (err error, views []*view.UserView) {
	err, data := s.viewUtils.View2Data(v)
	if err != nil {
		return err, nil
	}
	var datas []*model.User
	if err, datas = s.userDao.List(data); err != nil {
		return err, nil
	} else {
		err, views = s.viewUtils.Data2ViewList(datas)
		return
	}
}

// GetByUserName 根据userName获取User记录
// Author
func (s *UserService) GetByUserName(userName string) (err error, userView *view.UserView) {
	err1, user := s.userDao.GetByUserName(userName)
	if err1 != nil {
		return err1, nil
	}
	err2, userView := s.viewUtils.Data2View(user)
	if err2 != nil {
		return err2, nil
	}
	return
}

// CheckFieldUnique 校验字段是否唯一
// Author
func (s *UserService) CheckFieldUnique(fieldName, value, id string) error {
	if fieldName == "" || value == "" {
		return nil
	}
	if err, data := s.userDao.SelectByField(fieldName, value); err != nil {
		return err
	} else {
		if data != nil && data.Id != id {
			return errors.New("数据重复")
		}
		return nil
	}
}

// CheckUserDataScope 校验数据权限
func (s *UserService) CheckUserDataScope(userId, loginUserId string) error {
	if constants.SYSTEM_ADMIN_ID != loginUserId {
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
		param := &model.User{}
		param.Id = userId
		param.DataScopeSql = filter
		// data.DataScopeSql = filter
		err, _ = s.userDao.List(param)
		if err != nil {
			return err
		}
	}
	return nil
}

// ResetPwd 重置密码
func (s *UserService) ResetPwd(v *view.UserView) error {
	err, user := s.viewUtils.View2Data(v)
	if err != nil {
		return err
	}
	salt := utils.GenUID()
	user.Password = utils.EncryptionPassword(user.Password, salt)
	user.Salt = salt
	return s.userDao.Update(*user)
}

// ChangeStatus 更新状态
func (s *UserService) ChangeStatus(v *view.UserView) error {
	err, user := s.viewUtils.View2Data(v)
	if err != nil {
		return err
	}
	return s.userDao.Update(*user)
}

// AuthRole	角色授权
func (s *UserService) AuthRole(v *view.UserView) error {
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
func (s *UserService) SelectAllocatedList(pageInfo *view.UserPageView, user *view.UserView) (err error, res *common.PageInfo) {
	pageInfo.DataScopeSql = aspect.DataScopeFilter(user, "d", "u", "")
	if err, res = s.userDao.SelectAllocatedList(pageInfo); err != nil {
		return err, nil
	}
	if err, res = s.viewUtils.PageData2ViewList(res); err != nil {
		return err, res
	} else {
		if o, ok := res.Rows.([]*view.UserView); ok {
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
func (s *UserService) SelectUnallocatedList(pageInfo *view.UserPageView, user *view.UserView) (err error, res *common.PageInfo) {
	pageInfo.DataScopeSql = aspect.DataScopeFilter(user, "d", "u", "")
	if err, res = s.userDao.SelectUnallocatedList(pageInfo); err != nil {
		return err, nil
	}
	if err, res = s.viewUtils.PageData2ViewList(res); err != nil {
		return err, res
	} else {
		if o, ok := res.Rows.([]*view.UserView); ok {
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

// GetByDeptId 根据部门id获取User记录
func (s *UserService) GetByDeptId(deptId string) (err error, userView []*view.UserView) {
	err1, user := s.userDao.GetByDeptId(deptId)
	if err1 != nil {
		return err1, nil
	}
	err2, userView := s.viewUtils.Data2ViewList(user)
	if err2 != nil {
		return err2, nil
	}
	return
}
