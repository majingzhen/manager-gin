package framework

import (
	"manager-gin/src/app/admin/sys/sys_user/service"
	"manager-gin/src/app/admin/sys/sys_user/service/view"
	"manager-gin/src/global"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var userService = service.SysUserServiceApp

func GetLoginUser(c *gin.Context) *view.SysUserView {
	err, view := userService.Get(GetLoginUserId(c))
	if err != nil {
		global.Logger.Error("[获取登录用户] is error", zap.Error(err))
		return nil
	}
	return view
}

func GetLoginUserId(c *gin.Context) string {
	userId := c.GetString("user_id")
	return userId
}

func GetLoginUserName(c *gin.Context) string {
	user := GetLoginUser(c)
	return user.UserName
}
