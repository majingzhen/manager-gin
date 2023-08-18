// Package view 自动生成模板 SysRole
// @description <TODO description class purpose>
// @author
// @File: sys_role
// @version 1.0.0
// @create 2023-08-18 14:00:53
package view

// SysRoleView 结构体

type SysRoleView struct {
	Id string `json:"id"`

	RoleName string `json:"roleName"`

	RoleKey string `json:"roleKey"`

	RoleSort int `json:"roleSort"`

	DataScope string `json:"dataScope"`

	MenuCheckStrictly interface{} `json:"menuCheckStrictly"`

	DeptCheckStrictly interface{} `json:"deptCheckStrictly"`

	Status string `json:"status"`

	DeletedAt string `json:"deletedAt"`

	CreateBy string `json:"createBy"`

	CreateTime string `json:"createTime"`

	UpdateBy string `json:"updateBy"`

	UpdateTime string `json:"updateTime"`

	Remark string `json:"remark"`
}
