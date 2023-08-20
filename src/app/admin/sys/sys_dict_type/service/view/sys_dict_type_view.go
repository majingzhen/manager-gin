// Package view 自动生成模板 SysDictType
// @description <TODO description class purpose>
// @author
// @File: sys_dict_type
// @version 1.0.0
// @create 2023-08-18 13:41:26
package view

// SysDictTypeView 结构体

type SysDictTypeView struct {
	Id         string `json:"id"`
	DictName   string `json:"dictName"`
	DictType   string `json:"dictType"`
	Status     string `json:"status"`
	CreateBy   string `json:"createBy"`
	CreateTime string `json:"createTime"`
	UpdateBy   string `json:"updateBy"`
	UpdateTime string `json:"updateTime"`
	Remark     string `json:"remark"`
}

type SysDictTypePageView struct {
	DictName      string `json:"dictName" form:"dictName"`
	DictType      string `json:"dictType" form:"dictType"`
	Status        string `json:"status" form:"status"`
	CreateTime    string `json:"createTime" form:"createTime"`
	OrderByColumn string `form:"orderByColumn" form:"orderByColumn"` //排序字段
	IsAsc         string `form:"isAsc" form:"isAsc"`                 //排序方式
	PageNum       int    `form:"pageNum" form:"pageNum"`             //当前页码
	PageSize      int    `form:"pageSize" form:"pageSize"`           //每页数
}
