// Package model 自动生成模板 SysRoleDept
// @description <TODO description class purpose>
// @author
// @File: sys_role_dept
// @version 1.0.0
// @create 2023-08-18 14:02:24
package model

// SysRoleDept 结构体

type SysRoleDept struct {
	RoleId string `json:"roleId" form:"roleId" gorm:"column:role_id;comment:角色ID;"`

	DeptId string `json:"deptId" form:"deptId" gorm:"column:dept_id;comment:部门ID;"`
}

// TableName SysRoleDept 表名
func (SysRoleDept) TableName() string {
	return "sys_role_dept"
}
