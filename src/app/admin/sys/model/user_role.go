// Package model 自动生成模板 UserRole
// @description <TODO description class purpose>
// @author
// @File: user_role
// @version 1.0.0
// @create 2023-08-21 17:37:56
package model

// UserRole 结构体

type UserRole struct {
	UserId string `json:"userId" form:"userId" gorm:"column:user_id;comment:用户ID;"`

	RoleId string `json:"roleId" form:"roleId" gorm:"column:role_id;comment:角色ID;"`
}

// TableName UserRole 表名
func (UserRole) TableName() string {
	return "sys_user_role"
}
