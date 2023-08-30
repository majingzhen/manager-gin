// Package model 自动生成模板 RoleMenu
// @description <TODO description class purpose>
// @author
// @File: role_menu
// @version 1.0.0
// @create 2023-08-18 14:02:24
package model

// RoleMenu 结构体

type RoleMenu struct {
	RoleId string `json:"roleId" form:"roleId" gorm:"column:role_id;comment:角色ID;"`

	MenuId string `json:"menuId" form:"menuId" gorm:"column:menu_id;comment:菜单ID;"`
}

// TableName RoleMenu 表名
func (RoleMenu) TableName() string {
	return "sys_role_menu"
}
