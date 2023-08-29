// Package api  SysUserApi 自动生成模板
// @description <TODO description class purpose>
// @author
// @File: sys_user
// @version 1.0.0
// @create 2023-08-21 14:20:37
package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"manager-gin/src/app/admin/sys/service/sys_post"
	"manager-gin/src/app/admin/sys/service/sys_role"
	roleView "manager-gin/src/app/admin/sys/service/sys_role/view"
	"manager-gin/src/app/admin/sys/service/sys_user"
	"manager-gin/src/app/admin/sys/service/sys_user/view"
	"manager-gin/src/common"
	response "manager-gin/src/common/response"
	"manager-gin/src/global"
	"manager-gin/src/utils"
	"strings"
)

type SysUserApi struct {
	BasicApi
	sysUserService sys_user.SysUserService
	roleService    sys_role.SysRoleService
	postService    sys_post.SysPostService
}

// Create 创建SysUser
// @Summary 创建SysUser
// @Router /sysUser/create [post]
func (api *SysUserApi) Create(c *gin.Context) {
	var sysUserView view.SysUserView
	_ = c.ShouldBindJSON(&sysUserView)
	// 校验参数

	if err := api.sysUserService.CheckFieldUnique("user_name", sysUserView.UserName, ""); err != nil {
		response.FailWithMessage("登录账号已存在", c)
		return
	}
	if err := api.sysUserService.CheckFieldUnique("phone_number", sysUserView.PhoneNumber, ""); err != nil {
		response.FailWithMessage("手机号码已存在", c)
		return
	}
	if err := api.sysUserService.CheckFieldUnique("email", sysUserView.Email, ""); err != nil {
		response.FailWithMessage("邮箱账号已存在", c)
		return
	}
	sysUserView.Id = utils.GenUID()
	sysUserView.CreateTime = utils.GetCurTimeStr()
	sysUserView.UpdateTime = utils.GetCurTimeStr()
	sysUserView.Salt = utils.GenUID()
	sysUserView.Password = utils.EncryptionPassword(sysUserView.Password, sysUserView.Salt)
	sysUserView.CreateBy = api.GetLoginUserName(c)
	if err := api.sysUserService.Create(&sysUserView); err != nil {
		global.Logger.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// Delete 删除SysUser
// @Summary 删除SysUser
// @Router /sysUser/delete [delete]
func (api *SysUserApi) Delete(c *gin.Context) {
	idStr := c.Param("ids")
	ids := strings.Split(idStr, ",")
	if utils.Contains(ids, api.GetLoginUserId(c)) {
		response.FailWithMessage("当前用户不能删除", c)
	}
	if err := api.sysUserService.DeleteByIds(ids, api.GetLoginUserId(c)); err != nil {
		global.Logger.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// Update 更新SysUser
// @Summary 更新SysUser
// @Router /sysUser/update [put]
func (api *SysUserApi) Update(c *gin.Context) {
	var sysUserView view.SysUserView
	_ = c.ShouldBindJSON(&sysUserView)
	id := sysUserView.Id
	if id == "" {
		response.FailWithMessage("更新失败", c)
		return
	}
	if common.SYSTEM_ADMIN_ID == id {
		response.FailWithMessage("超级管理员不允许修改", c)
		return
	}
	// 校验用户是否有数据权限
	err := api.sysUserService.CheckUserDataScope(id, api.GetLoginUserId(c))
	if err != nil {
		response.FailWithMessage("没有权限访问用户数据", c)
		return
	}
	// 校验参数
	if err := api.sysUserService.CheckFieldUnique("user_name", sysUserView.UserName, sysUserView.Id); err != nil {
		response.FailWithMessage("登录账号已存在", c)
		return
	}
	if err := api.sysUserService.CheckFieldUnique("phone_number", sysUserView.PhoneNumber, sysUserView.Id); err != nil {
		response.FailWithMessage("手机号码已存在", c)
		return
	}
	if err := api.sysUserService.CheckFieldUnique("email", sysUserView.Email, sysUserView.Id); err != nil {
		response.FailWithMessage("邮箱账号已存在", c)
		return
	}
	sysUserView.UpdateTime = utils.GetCurTimeStr()
	sysUserView.UpdateBy = api.GetLoginUserName(c)
	if err := api.sysUserService.Update(id, &sysUserView); err != nil {
		global.Logger.Error("更新持久化失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// Get 用id查询SysUser
// @Summary 用id查询SysUser
// @Router /sysUser/get [get]
func (api *SysUserApi) Get(c *gin.Context) {
	var sysUserInfoView = new(view.SysUserInfoView)
	err, roles := api.roleService.SelectRoleAll(api.GetLoginUser(c))
	if err == nil {
		removeAdminRole(&roles)
		sysUserInfoView.Roles = roles
	}
	err, views := api.postService.SelectPostAll()
	if err == nil {
		sysUserInfoView.Posts = views
	}
	id := c.Param("id")
	if id != "" {
		if err1, sysUserView := api.sysUserService.Get(id); err1 != nil {
			global.Logger.Error("查询失败!", zap.Error(err1))
			response.OkWithMessage(err1.Error(), c)
			return
		} else {
			sysUserInfoView.SysUserView = *sysUserView
			err2, postIds := api.postService.SelectPostIdListByUserId(id)
			if err2 != nil {
				global.Logger.Error("查询失败!", zap.Error(err1))
				response.OkWithMessage(err1.Error(), c)
				return
			}
			sysUserInfoView.PostIds = postIds
			var roleIds []string
			for _, roleView := range sysUserView.Roles {
				if roleView.Id != common.SYSTEM_ROLE_ADMIN_ID {
					roleIds = append(roleIds, roleView.Id)
				}
			}
			sysUserInfoView.RoleIds = roleIds
		}
	}
	response.OkWithData(sysUserInfoView, c)
}

// Page 分页获取SysUser列表
// @Summary 分页获取SysUser列表
// @Router /sysUser/page [get]
func (api *SysUserApi) Page(c *gin.Context) {
	var pageInfo view.SysUserPageView
	// 绑定查询参数到 pageInfo
	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		response.FailWithMessage("获取分页数据解析失败!", c)
		return
	}
	user := api.GetLoginUser(c)
	if err, res := api.sysUserService.Page(&pageInfo, user); err != nil {
		global.Logger.Error("获取分页信息失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(res, "获取成功", c)
	}
}

// List 获取SysUser列表
// @Summary 获取SysUser列表
// @Router /sysUser/list [get]
func (api *SysUserApi) List(c *gin.Context) {
	var view view.SysUserView
	// 绑定查询参数到 view对象
	if err := c.ShouldBindQuery(&view); err != nil {
		response.FailWithMessage("获取参数解析失败!", c)
		return
	}
	// 判断是否需要根据用户获取数据
	// userId := api.GetLoginUserId(c)
	if err, res := api.sysUserService.List(&view); err != nil {
		global.Logger.Error("获取数据失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(res, "获取成功", c)
	}
}

// ResetPwd 重置密码
// @Summary 重置密码
// @Router /sysUser/resetPwd [put]
func (api *SysUserApi) ResetPwd(c *gin.Context) {
	var req view.SysUserView
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if common.SYSTEM_ADMIN_ID == req.Id {
		response.FailWithMessage("超级管理员不允许修改", c)
		return
	}
	if err := api.sysUserService.CheckUserDataScope(req.Id, api.GetLoginUserId(c)); err != nil {
		response.FailWithMessage("没有权限访问用户数据", c)
		return
	}
	req.UpdateTime = utils.GetCurTimeStr()
	req.UpdateBy = api.GetLoginUserName(c)
	if err := api.sysUserService.ResetPwd(&req); err != nil {
		response.FailWithMessage("重置密码失败", c)
		return
	}
	response.OkWithMessage("重置密码成功", c)
}

// ChangeStatus 更新状态
// @Summary 更新状态
// @Router /sysUser/changeStatus [put]
func (api *SysUserApi) ChangeStatus(c *gin.Context) {
	var req view.SysUserView
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if common.SYSTEM_ADMIN_ID == req.Id {
		response.FailWithMessage("超级管理员不允许修改", c)
		return
	}
	if err := api.sysUserService.CheckUserDataScope(req.Id, api.GetLoginUserId(c)); err != nil {
		response.FailWithMessage("没有权限访问用户数据", c)
		return
	}
	req.UpdateTime = utils.GetCurTimeStr()
	req.UpdateBy = api.GetLoginUserName(c)
	if err := api.sysUserService.ChangeStatus(&req); err != nil {
		response.FailWithMessage("更新状态失败", c)
		return
	}
	response.OkWithMessage("更新状态成功", c)
}

// GetAuthRole 根据用户id获取授权角色
// @Summary 根据用户id获取授权角色
// @Router /sysUser/getAuthRole/{userId} [get]
func (api *SysUserApi) GetAuthRole(c *gin.Context) {
	userId := c.Param("id")
	err, userView := api.sysUserService.Get(userId)
	if err != nil {
		global.Logger.Error("获取数据失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	if err1, res := api.roleService.AssembleRolesByUserId(userId); err1 != nil {
		global.Logger.Error("获取数据失败!", zap.Error(err1))
		response.FailWithMessage("获取失败", c)
	} else {
		removeAdminRole(&res)
		response.OkWithData(gin.H{
			"user":  userView,
			"roles": res,
		}, c)
	}
}

// AuthRole 批量给用户授权角色
// @Summary 批量给用户授权角色
// @Router /sysUser/authRole [put]
func (api *SysUserApi) AuthRole(c *gin.Context) {
	userId := c.Query("userId")
	roleIdStr := c.Query("roleIds")
	if userId == "" || roleIdStr == "" {
		response.FailWithMessage("参数错误", c)
		return
	}
	roleIds := strings.Split(roleIdStr, ",")
	if err := api.sysUserService.CheckUserDataScope(userId, api.GetLoginUserId(c)); err != nil {
		response.FailWithMessage("没有权限访问用户数据", c)
		return
	}
	v := &view.SysUserView{
		Id:         userId,
		RoleIds:    roleIds,
		UpdateTime: utils.GetCurTimeStr(),
		UpdateBy:   api.GetLoginUserName(c),
	}

	if err := api.sysUserService.AuthRole(v); err != nil {
		response.FailWithMessage("授权失败", c)
		return
	}
	response.OkWithMessage("授权成功", c)
}

// 剔除超级管理员
func removeAdminRole(roles *[]*roleView.SysRoleView) {
	for i := 0; i < len(*roles); i++ {
		if (*roles)[i].Id == common.SYSTEM_ROLE_ADMIN_ID {
			*roles = append((*roles)[:i], (*roles)[i+1:]...)
			break
		}
	}
}

// SelectAllocatedList 查询已分配用户角色列表
// @Summary 查询已分配用户角色列表
// @Router /sysUser/selectAllocatedList [get]
func (api *SysUserApi) SelectAllocatedList(c *gin.Context) {
	var pageInfo view.SysUserPageView
	// 绑定查询参数到 pageInfo
	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		response.FailWithMessage("获取分页数据解析失败!", c)
		return
	}
	user := api.GetLoginUser(c)
	if err, res := api.sysUserService.SelectAllocatedList(&pageInfo, user); err != nil {
		global.Logger.Error("获取分页信息失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(res, "获取成功", c)
	}
}

// SelectUnallocatedList 查询未分配用户角色列表
// @Summary 查询未分配用户角色列表
// @Router /sysUser/selectUnallocatedList [get]
func (api *SysUserApi) SelectUnallocatedList(c *gin.Context) {
	var pageInfo view.SysUserPageView
	// 绑定查询参数到 pageInfo
	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		response.FailWithMessage("获取分页数据解析失败!", c)
		return
	}
	user := api.GetLoginUser(c)
	if err, res := api.sysUserService.SelectUnallocatedList(&pageInfo, user); err != nil {
		global.Logger.Error("获取分页信息失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(res, "获取成功", c)
	}
}
