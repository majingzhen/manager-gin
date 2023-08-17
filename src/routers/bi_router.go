package routers

import (
	"github.com/gin-gonic/gin"
	role "manager-gin/src/app/admin/bi/bi_role/router"
	user "manager-gin/src/app/admin/bi/bi_user/router"
)

type BiRouter struct{}

// InitBiRouter 初始化 InitBi 路由信息
func (r *BiRouter) InitBiRouter(Router *gin.RouterGroup) {

	var userRouter user.BiUserRouter
	var roleRouter role.BiRoleRouter

	sys := Router.Group("bi")
	{
		userRouter.InitBiUserRouter(sys)
		roleRouter.InitBiRoleRouter(sys)
	}
}
