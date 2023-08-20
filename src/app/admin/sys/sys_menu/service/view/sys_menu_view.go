// Package view 自动生成模板 SysMenu
// @description <TODO description class purpose>
// @author
// @File: sys_menu
// @version 1.0.0
// @create 2023-08-18 13:41:26
package view

// SysMenuView 结构体

type SysMenuView struct {
	Id         string         `json:"id"`
	MenuName   string         `json:"menuName"`
	ParentId   string         `json:"parentId"`
	OrderNum   int            `json:"orderNum"`
	Path       string         `json:"path"`
	Component  string         `json:"component"`
	Query      string         `json:"query"`
	IsFrame    int            `json:"isFrame"`
	IsCache    int            `json:"isCache"`
	MenuType   string         `json:"menuType"`
	Visible    string         `json:"visible"`
	Status     string         `json:"status"`
	Perms      string         `json:"perms"`
	Icon       string         `json:"icon"`
	CreateBy   string         `json:"createBy"`
	CreateTime string         `json:"createTime"`
	UpdateBy   string         `json:"updateBy"`
	UpdateTime string         `json:"updateTime"`
	Remark     string         `json:"remark"`
	Children   []*SysMenuView `json:"children"`
}

type TreeNode struct {
	Id       string      `json:"id"`
	MenuName string      `json:"menuName"`
	Icon     string      `json:"icon"`
	Path     string      `json:"path"`
	Perms    string      `json:"perms"`
	Children []*TreeNode `json:"children"`
}

type RouterView struct {
	Name       string        `json:"name"`
	Path       string        `json:"path"`
	Hidden     bool          `json:"hidden"`
	Redirect   string        `json:"redirect"`
	Component  string        `json:"component"`
	Query      string        `json:"query"`
	AlwaysShow bool          `json:"alwaysShow"`
	Meta       *MetaView     `json:"meta"`
	Children   []*RouterView `json:"children"`
}

type MetaView struct {
	Title   string `json:"title"`
	Icon    string `json:"icon"`
	NoCache bool   `json:"NoCache"`
	Link    string `json:"link"`
}
