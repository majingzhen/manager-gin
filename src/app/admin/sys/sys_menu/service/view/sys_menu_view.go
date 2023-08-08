// Package view 自动生成模板 SysMenu
// @description <TODO description class purpose>
// @author
// @File: sys_menu
// @version 1.0.0
// @create 2023-08-08 10:06:19
package view

// SysMenuView 结构体

type SysMenuView struct {
	Component string `json:"component"`

	CreateBy string `json:"createBy"`

	CreateTime string `json:"createTime"`

	Icon string `json:"icon"`

	Id int `json:"id"`

	IsCache int `json:"isCache"`

	IsFrame int `json:"isFrame"`

	MenuName string `json:"menuName"`

	MenuType interface{} `json:"menuType"`

	OrderNum int `json:"orderNum"`

	ParentId int `json:"parentId"`

	Path string `json:"path"`

	Perms string `json:"perms"`

	Query string `json:"query"`

	Remark string `json:"remark"`

	Status interface{} `json:"status"`

	UpdateBy string `json:"updateBy"`

	UpdateTime string `json:"updateTime"`

	Visible interface{} `json:"visible"`
}
