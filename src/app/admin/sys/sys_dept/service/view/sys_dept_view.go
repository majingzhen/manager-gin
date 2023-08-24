// Package view 自动生成模板 SysDept
// @description <TODO description class purpose>
// @author
// @File: sys_dept
// @version 1.0.0
// @create 2023-08-21 10:27:01
package view

// SysDeptView 结构体

type SysDeptView struct {
	Id         string         `json:"id" form:"id"`
	ParentId   string         `json:"parentId" form:"parentId"`
	Ancestors  string         `json:"ancestors" form:"ancestors"`
	DeptName   string         `json:"deptName" form:"deptName"`
	OrderNum   int            `json:"orderNum" form:"orderNum"`
	Leader     string         `json:"leader" form:"leader"`
	Phone      string         `json:"phone" form:"phone"`
	Email      string         `json:"email" form:"email"`
	Status     string         `json:"status" form:"status"`
	CreateBy   string         `json:"createBy" form:"createBy"`
	CreateTime string         `json:"createTime" form:"createTime"`
	UpdateBy   string         `json:"updateBy" form:"updateBy"`
	UpdateTime string         `json:"updateTime" form:"updateTime"`
	Children   *[]SysDeptView `json:"children" form:"children"`
}

type SysDeptPageView struct {
	// TODO 按需修改
	Id         string `json:"id" form:"id"`
	ParentId   string `json:"parentId" form:"parentId"`
	Ancestors  string `json:"ancestors" form:"ancestors"`
	DeptName   string `json:"deptName" form:"deptName"`
	OrderNum   int    `json:"orderNum" form:"orderNum"`
	Leader     string `json:"leader" form:"leader"`
	Phone      string `json:"phone" form:"phone"`
	Email      string `json:"email" form:"email"`
	Status     string `json:"status" form:"status"`
	CreateBy   string `json:"createBy" form:"createBy"`
	CreateTime string `json:"createTime" form:"createTime"`
	UpdateBy   string `json:"updateBy" form:"updateBy"`
	UpdateTime string `json:"updateTime" form:"updateTime"`

	OrderByColumn string `json:"orderByColumn" form:"orderByColumn"` //排序字段
	IsAsc         string `json:"isAsc" form:"isAsc"`                 //排序方式
	PageNum       int    `json:"pageNum" form:"pageNum"`             //当前页码
	PageSize      int    `json:"pageSize" form:"pageSize"`           //每页数
}

type SysDeptTreeView struct {
	Id       string             `json:"id" form:"id"`
	Label    string             `json:"label" form:"label"`
	ParentId string             `json:"parentId" form:"parentId"`
	Children []*SysDeptTreeView `json:"children" form:"children"`
}
