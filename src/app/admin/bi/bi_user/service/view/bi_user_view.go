// Package view 自动生成模板 BiUser
// @description <TODO description class purpose>
// @author Administrator
// @File: bi_user
// @version 1.0.0
// @create 2023-08-16 14:53:36
package view

// BiUserView 结构体

type BiUserView struct {
	Id string `json:"id" form:"id"`

	TestName string `json:"testName" form:"testName"`

	Gender int `json:"gender" form:"gender"`

	Remark string `json:"remark" form:"remark"`

	Birth string `json:"birth" form:"birth"`

	CreateTime string `json:"createTime" form:"createTime"`

	UpdateTime string `json:"updateTime" form:"updateTime"`
}
