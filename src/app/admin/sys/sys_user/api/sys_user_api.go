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
	postSer "manager-gin/src/app/admin/sys/sys_post/service"
	roleSer "manager-gin/src/app/admin/sys/sys_role/service"
	"manager-gin/src/app/admin/sys/sys_user/service"
	"manager-gin/src/app/admin/sys/sys_user/service/view"
	"manager-gin/src/common"
	response "manager-gin/src/common/response"
	"manager-gin/src/framework"
	"manager-gin/src/global"
	"manager-gin/src/utils"
	"strings"
)

type SysUserApi struct {
}

var sysUserService = service.SysUserServiceApp
var roleService = roleSer.SysRoleServiceApp
var postService = postSer.SysPostServiceApp

// Create 创建SysUser
// @Summary 创建SysUser
// @Router /sysUser/create [post]
func (api *SysUserApi) Create(c *gin.Context) {
	var sysUserView view.SysUserView
	_ = c.ShouldBindJSON(&sysUserView)
	// 校验参数
	if err := sysUserService.CheckFieldUnique("user_name", sysUserView.UserName); err == nil {
		response.FailWithMessage("登录账号已存在", c)
		return
	}
	if err := sysUserService.CheckFieldUnique("phone_number", sysUserView.PhoneNumber); err == nil {
		response.FailWithMessage("手机号码已存在", c)
		return
	}
	if err := sysUserService.CheckFieldUnique("email", sysUserView.Email); err == nil {
		response.FailWithMessage("邮箱账号已存在", c)
		return
	}
	sysUserView.Id = utils.GenUID()
	sysUserView.CreateTime = utils.GetCurTimeStr()
	sysUserView.UpdateTime = utils.GetCurTimeStr()
	sysUserView.Salt = utils.GenUID()
	sysUserView.Password = utils.EncryptionPassword(sysUserView.Password, sysUserView.Salt)
	sysUserView.CreateBy = framework.GetLoginUserName(c)
	if err := sysUserService.Create(&sysUserView); err != nil {
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
	if err := sysUserService.DeleteByIds(ids); err != nil {
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
	err := sysUserService.CheckUserDataScope(id)
	if err != nil {
		response.FailWithMessage("没有权限访问用户数据", c)
		return
	}
	// 校验参数
	if err := sysUserService.CheckFieldUnique("user_name", sysUserView.UserName); err == nil {
		response.FailWithMessage("登录账号已存在", c)
		return
	}
	if err := sysUserService.CheckFieldUnique("phone_number", sysUserView.PhoneNumber); err == nil {
		response.FailWithMessage("手机号码已存在", c)
		return
	}
	if err := sysUserService.CheckFieldUnique("email", sysUserView.Email); err == nil {
		response.FailWithMessage("邮箱账号已存在", c)
		return
	}
	sysUserView.UpdateTime = utils.GetCurTimeStr()
	sysUserView.UpdateBy = framework.GetLoginUserName(c)
	if err := sysUserService.Update(id, &sysUserView); err != nil {
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
	err, roles := roleService.SelectRoleAll()
	if err == nil {
		// 剔除超级管理员角色
		for i := 0; i < len(*roles); i++ {
			if (*roles)[i].Id == common.SYSTEM_ROLE_ADMIN_ID {
				*roles = append((*roles)[:i], (*roles)[i+1:]...)
				break
			}
		}
		sysUserInfoView.Roles = roles
	}
	err, views := postService.SelectPostAll()
	if err == nil {
		sysUserInfoView.Posts = views
	}
	id := c.Param("id")
	if id != "" {
		if err1, sysUserView := sysUserService.Get(id); err1 != nil {
			global.Logger.Error("查询失败!", zap.Error(err1))
			response.OkWithMessage(err1.Error(), c)
			return
		} else {
			sysUserInfoView.SysUserView = *sysUserView
			err2, postIds := postService.SelectPostListByUserId(id)
			if err2 != nil {
				global.Logger.Error("查询失败!", zap.Error(err1))
				response.OkWithMessage(err1.Error(), c)
				return
			}
			sysUserInfoView.PostIds = postIds
			var roleIds []string
			for _, roleView := range *sysUserView.Roles {
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
	user := framework.GetLoginUser(c)
	if err, res := sysUserService.Page(&pageInfo, user); err != nil {
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
	// userId := framework.GetLoginUserId(c)
	if err, res := sysUserService.List(&view); err != nil {
		global.Logger.Error("获取数据失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(res, "获取成功", c)
	}
}
