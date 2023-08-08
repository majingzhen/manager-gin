// Package model 自动生成模板 SysRole
// @description <TODO description class purpose>
// @author
// @File: sys_role
// @version 1.0.0
// @create 2023-08-08 10:06:19
package model

import (
	"time"
)

// SysRole 结构体

type SysRole struct {
	CreateBy string `json:"createBy" form:"createBy" gorm:"column:create_by;comment:;"`

	CreateTime time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:;"`

	DataScope interface{} `json:"dataScope" form:"dataScope" gorm:"column:data_scope;comment:;"`

	DelFlag interface{} `json:"delFlag" form:"delFlag" gorm:"column:del_flag;comment:;"`

	DeptCheckStrictly int8 `json:"deptCheckStrictly" form:"deptCheckStrictly" gorm:"column:dept_check_strictly;comment:;"`

	Id int `json:"id" form:"id" gorm:"column:id;comment:;"`

	MenuCheckStrictly int8 `json:"menuCheckStrictly" form:"menuCheckStrictly" gorm:"column:menu_check_strictly;comment:;"`

	Remark string `json:"remark" form:"remark" gorm:"column:remark;comment:;"`

	RoleKey string `json:"roleKey" form:"roleKey" gorm:"column:role_key;comment:;"`

	RoleName string `json:"roleName" form:"roleName" gorm:"column:role_name;comment:;"`

	RoleSort int `json:"roleSort" form:"roleSort" gorm:"column:role_sort;comment:;"`

	Status interface{} `json:"status" form:"status" gorm:"column:status;comment:;"`

	UpdateBy string `json:"updateBy" form:"updateBy" gorm:"column:update_by;comment:;"`

	UpdateTime time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:;"`
}

// TableName SysRole 表名
func (SysRole) TableName() string {
	return "sys_role"
}
