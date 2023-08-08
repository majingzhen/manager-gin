// Package view 自动生成模板 SysDept
// @description <TODO description class purpose>
// @author
// @File: sys_dept
// @version 1.0.0
// @create 2023-08-08 10:06:19
package view

// SysDeptView 结构体

type SysDeptView struct {
	Ancestors string `json:"ancestors"`

	CreateBy string `json:"createBy"`

	CreateTime string `json:"createTime"`

	DelFlag interface{} `json:"delFlag"`

	DeptName string `json:"deptName"`

	Email string `json:"email"`

	Id int `json:"id"`

	Leader string `json:"leader"`

	OrderNum int `json:"orderNum"`

	ParentId int `json:"parentId"`

	Phone string `json:"phone"`

	Status interface{} `json:"status"`

	UpdateBy string `json:"updateBy"`

	UpdateTime string `json:"updateTime"`
}
