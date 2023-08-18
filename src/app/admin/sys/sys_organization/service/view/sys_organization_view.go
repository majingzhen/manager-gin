// Package view 自动生成模板 SysOrganization
// @description <TODO description class purpose>
// @author
// @File: sys_organization
// @version 1.0.0
// @create 2023-08-18 14:00:53
package view

// SysOrganizationView 结构体

type SysOrganizationView struct {
	Id string `json:"id"`

	ParentId string `json:"parentId"`

	Ancestors string `json:"ancestors"`

	DeptName string `json:"deptName"`

	OrderNum int `json:"orderNum"`

	Leader string `json:"leader"`

	Phone string `json:"phone"`

	Email string `json:"email"`

	Status string `json:"status"`

	DeletedAt string `json:"deletedAt"`

	CreateBy string `json:"createBy"`

	CreateTime string `json:"createTime"`

	UpdateBy string `json:"updateBy"`

	UpdateTime string `json:"updateTime"`
}
