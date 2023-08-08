// Package model 自动生成模板 SysUserRole
// @description <TODO description class purpose>
// @author
// @File: sys_user_role
// @version 1.0.0
// @create 2023-08-08 10:06:19
package model

// SysUserRole 结构体

type SysUserRole struct {
	RoleId int `json:"roleId" form:"roleId" gorm:"column:role_id;comment:;"`

	UserId int `json:"userId" form:"userId" gorm:"column:user_id;comment:;"`
}

// TableName SysUserRole 表名
func (SysUserRole) TableName() string {
	return "sys_user_role"
}
