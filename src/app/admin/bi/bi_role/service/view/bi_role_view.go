// Package view 自动生成模板 BiRole
// @description <TODO description class purpose>
// @author Majz
// @File: bi_role
// @version 1.0.0
// @create 2023-08-16 14:53:36
package view

// BiRoleView 结构体

type BiRoleView struct {
	Id string `json:"id"`

	TestName string `json:"testName"`

	Gender int `json:"gender"`

	Remark string `json:"remark"`

	Birth string `json:"birth"`

	CreateTime string `json:"createTime"`

	UpdateTime string `json:"updateTime"`
}
