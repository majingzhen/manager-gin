// Package view 自动生成模板 SysPost
// @description <TODO description class purpose>
// @author
// @File: sys_post
// @version 1.0.0
// @create 2023-08-08 10:06:19
package view

// SysPostView 结构体

type SysPostView struct {
	CreateBy string `json:"createBy"`

	CreateTime string `json:"createTime"`

	Id int `json:"id"`

	PostCode string `json:"postCode"`

	PostName string `json:"postName"`

	PostSort int `json:"postSort"`

	Remark string `json:"remark"`

	Status interface{} `json:"status"`

	UpdateBy string `json:"updateBy"`

	UpdateTime string `json:"updateTime"`
}
