// Package view 自动生成模板 SysMenu
// @description <TODO description class purpose>
// @author
// @File: sys_menu
// @version 1.0.0
// @create 2023-08-20 21:21:34
package view

// SysMenuView 结构体

type SysMenuView struct {
	Id         string `json:"id" form:"id"`
	MenuName   string `json:"menuName" form:"menuName"`
	ParentId   string `json:"parentId" form:"parentId"`
	OrderNum   int    `json:"orderNum" form:"orderNum"`
	Path       string `json:"path" form:"path"`
	Component  string `json:"component" form:"component"`
	Query      string `json:"query" form:"query"`
	IsFrame    int    `json:"isFrame" form:"isFrame"`
	IsCache    int    `json:"isCache" form:"isCache"`
	MenuType   string `json:"menuType" form:"menuType"`
	Visible    string `json:"visible" form:"visible"`
	Status     string `json:"status" form:"status"`
	Perms      string `json:"perms" form:"perms"`
	Icon       string `json:"icon" form:"icon"`
	CreateBy   string `json:"createBy" form:"createBy"`
	CreateTime string `json:"createTime" form:"createTime"`
	UpdateBy   string `json:"updateBy" form:"updateBy"`
	UpdateTime string `json:"updateTime" form:"updateTime"`
	Remark     string `json:"remark" form:"remark"`
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

type SysMenuPageView struct {
	// TODO 按需修改
	Id         string `json:"id" form:"id"`
	MenuName   string `json:"menuName" form:"menuName"`
	ParentId   string `json:"parentId" form:"parentId"`
	OrderNum   int    `json:"orderNum" form:"orderNum"`
	Path       string `json:"path" form:"path"`
	Component  string `json:"component" form:"component"`
	Query      string `json:"query" form:"query"`
	IsFrame    int    `json:"isFrame" form:"isFrame"`
	IsCache    int    `json:"isCache" form:"isCache"`
	MenuType   string `json:"menuType" form:"menuType"`
	Visible    string `json:"visible" form:"visible"`
	Status     string `json:"status" form:"status"`
	Perms      string `json:"perms" form:"perms"`
	Icon       string `json:"icon" form:"icon"`
	CreateBy   string `json:"createBy" form:"createBy"`
	CreateTime string `json:"createTime" form:"createTime"`
	UpdateBy   string `json:"updateBy" form:"updateBy"`
	UpdateTime string `json:"updateTime" form:"updateTime"`
	Remark     string `json:"remark" form:"remark"`

	OrderByColumn string `json:"orderByColumn" form:"orderByColumn"` //排序字段
	IsAsc         string `json:"isAsc" form:"isAsc"`                 //排序方式
	PageNum       int    `json:"pageNum" form:"pageNum"`             //当前页码
	PageSize      int    `json:"pageSize" form:"pageSize"`           //每页数
}
