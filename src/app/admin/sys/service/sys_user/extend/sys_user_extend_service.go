package extend

import (
	"manager-gin/src/app/admin/sys/dao"
	"manager-gin/src/app/admin/sys/service/sys_user/view"
)

type SysUserExtendService struct {
	sysUserDao dao.SysUserDao
	viewUtils  view.SysUserViewUtils
}

// GetByDeptId 根据部门id获取SysUser记录
func (s *SysUserExtendService) GetByDeptId(deptId string) (err error, sysUserView []*view.SysUserView) {
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
