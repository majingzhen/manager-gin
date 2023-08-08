// Package model 自动生成模板 SysRoleMenu
// @description <TODO description class purpose>
// @author
// @File: sys_role_menu
// @version 1.0.0
// @create 2023-08-08 10:06:19
package model

// SysRoleMenu 结构体

type SysRoleMenu struct {
	MenuId int `json:"menuId" form:"menuId" gorm:"column:menu_id;comment:;"`

	RoleId int `json:"roleId" form:"roleId" gorm:"column:role_id;comment:;"`
}

// TableName SysRoleMenu 表名
func (SysRoleMenu) TableName() string {
	return "sys_role_menu"
}
