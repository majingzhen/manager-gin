// Package view 自动生成模板 SysPost
// @description <TODO description class purpose>
// @author
// @File: sys_post
// @version 1.0.0
// @create 2023-08-18 13:41:26
package view

// SysPostView 结构体

type SysPostView struct {
	Id string `json:"id"`

	PostCode string `json:"postCode"`

	PostName string `json:"postName"`

	PostSort int `json:"postSort"`

	Status string `json:"status"`

	CreateBy string `json:"createBy"`

	CreateTime string `json:"createTime"`

	UpdateBy string `json:"updateBy"`

	UpdateTime string `json:"updateTime"`

	Remark string `json:"remark"`
}
