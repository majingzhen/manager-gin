// Package view 自动生成模板 SysRole
// @description <TODO description class purpose>
// @author
// @File: sys_role
// @version 1.0.0
// @create 2023-08-08 10:06:19
package view

// SysRoleView 结构体

type SysRoleView struct {
	CreateBy string `json:"createBy"`

	CreateTime string `json:"createTime"`

	DataScope interface{} `json:"dataScope"`

	DelFlag interface{} `json:"delFlag"`

	DeptCheckStrictly int8 `json:"deptCheckStrictly"`

	Id int `json:"id"`

	MenuCheckStrictly int8 `json:"menuCheckStrictly"`

	Remark string `json:"remark"`

	RoleKey string `json:"roleKey"`

	RoleName string `json:"roleName"`

	RoleSort int `json:"roleSort"`

	Status interface{} `json:"status"`

	UpdateBy string `json:"updateBy"`

	UpdateTime string `json:"updateTime"`
}
