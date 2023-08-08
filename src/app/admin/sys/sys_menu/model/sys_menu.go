// Package model 自动生成模板 SysMenu
// @description <TODO description class purpose>
// @author
// @File: sys_menu
// @version 1.0.0
// @create 2023-08-08 10:06:19
package model

import (
	"time"
)

// SysMenu 结构体

type SysMenu struct {
	Component string `json:"component" form:"component" gorm:"column:component;comment:;"`

	CreateBy string `json:"createBy" form:"createBy" gorm:"column:create_by;comment:;"`

	CreateTime time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:;"`

	Icon string `json:"icon" form:"icon" gorm:"column:icon;comment:;"`

	Id int `json:"id" form:"id" gorm:"column:id;comment:;"`

	IsCache int `json:"isCache" form:"isCache" gorm:"column:is_cache;comment:;"`

	IsFrame int `json:"isFrame" form:"isFrame" gorm:"column:is_frame;comment:;"`

	MenuName string `json:"menuName" form:"menuName" gorm:"column:menu_name;comment:;"`

	MenuType interface{} `json:"menuType" form:"menuType" gorm:"column:menu_type;comment:;"`

	OrderNum int `json:"orderNum" form:"orderNum" gorm:"column:order_num;comment:;"`

	ParentId int `json:"parentId" form:"parentId" gorm:"column:parent_id;comment:;"`

	Path string `json:"path" form:"path" gorm:"column:path;comment:;"`

	Perms string `json:"perms" form:"perms" gorm:"column:perms;comment:;"`

	Query string `json:"query" form:"query" gorm:"column:query;comment:;"`

	Remark string `json:"remark" form:"remark" gorm:"column:remark;comment:;"`

	Status interface{} `json:"status" form:"status" gorm:"column:status;comment:;"`

	UpdateBy string `json:"updateBy" form:"updateBy" gorm:"column:update_by;comment:;"`

	UpdateTime time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:;"`

	Visible interface{} `json:"visible" form:"visible" gorm:"column:visible;comment:;"`
}

// TableName SysMenu 表名
func (SysMenu) TableName() string {
	return "sys_menu"
}
