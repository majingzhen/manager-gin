package extend

import (
	"manager-gin/src/app/admin/sys/sys_user/model"
	"manager-gin/src/app/admin/sys/sys_user/service/view"
)

type SysUserExtendService struct {
}

var sysUserDao = model.SysUserDaoApp
var viewUtils = view.SysUserViewUtilsApp

// GetByDeptId 根据部门id获取SysUser记录
func (service *SysUserExtendService) GetByDeptId(deptId string) (err error, sysUserView []*view.SysUserView) {
	err1, sysUser := sysUserDao.GetByDeptId(deptId)
	if err1 != nil {
		return err1, nil
	}
	err2, sysUserView := viewUtils.Data2ViewList(sysUser)
	if err2 != nil {
		return err2, nil
	}
	return
}
