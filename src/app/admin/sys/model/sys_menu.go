// Package model 自动生成模板 SysMenu
// @description <TODO description class purpose>
// @author
// @File: sys_menu
// @version 1.0.0
// @create 2023-08-18 13:41:26
package model

import (
	"time"
)

// SysMenu 结构体

type SysMenu struct {
	Id string `json:"id" form:"id" gorm:"column:id;comment:菜单ID;"`

	MenuName string `json:"menuName" form:"menuName" gorm:"column:menu_name;comment:菜单名称;"`

	ParentId string `json:"parentId" form:"parentId" gorm:"column:parent_id;comment:父菜单ID;"`

	OrderNum int `json:"orderNum" form:"orderNum" gorm:"column:order_num;comment:显示顺序;"`

	Path string `json:"path" form:"path" gorm:"column:path;comment:路由地址;"`

	Component string `json:"component" form:"component" gorm:"column:component;comment:组件路径;"`

	Query string `json:"query" form:"query" gorm:"column:query;comment:路由参数;"`

	IsFrame int `json:"isFrame" form:"isFrame" gorm:"column:is_frame;comment:是否为外链（0是 1否）;"`

	IsCache int `json:"isCache" form:"isCache" gorm:"column:is_cache;comment:是否缓存（0缓存 1不缓存）;"`

	MenuType string `json:"menuType" form:"menuType" gorm:"column:menu_type;comment:菜单类型（M目录 C菜单 F按钮）;"`

	Visible string `json:"visible" form:"visible" gorm:"column:visible;comment:菜单状态（0显示 1隐藏）;"`

	Status string `json:"status" form:"status" gorm:"column:status;comment:菜单状态（0正常 1停用）;"`

	Perms string `json:"perms" form:"perms" gorm:"column:perms;comment:权限标识;"`

	Icon string `json:"icon" form:"icon" gorm:"column:icon;comment:菜单图标;"`

	CreateBy string `json:"createBy" form:"createBy" gorm:"column:create_by;comment:创建者;"`

	CreateTime *time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;"`

	UpdateBy string `json:"updateBy" form:"updateBy" gorm:"column:update_by;comment:更新者;"`

	UpdateTime *time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:更新时间;"`

	Remark string `json:"remark" form:"remark" gorm:"column:remark;comment:备注;"`
}

// TableName SysMenu 表名
func (SysMenu) TableName() string {
	return "sys_menu"
}

// Entity is the golang structure for table sys_menu.
type SysMenuExtend struct {
	SysMenu
	ParentName string          `json:"parentName"` // 父菜单名称
	Children   []SysMenuExtend `json:"children"`   // 子菜单
}
