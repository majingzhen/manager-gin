// Package model 自动生成模板 Role
// @description <TODO description class purpose>
// @author
// @File: role
// @version 1.0.0
// @create 2023-08-18 14:00:53
package model

import (
	"time"
)

// Role 结构体

type Role struct {
	Id string `json:"id" form:"id" gorm:"column:id;comment:角色ID;"`

	RoleName string `json:"roleName" form:"roleName" gorm:"column:role_name;comment:角色名称;"`

	RoleKey string `json:"roleKey" form:"roleKey" gorm:"column:role_key;comment:角色权限字符串;"`

	RoleSort int `json:"roleSort" form:"roleSort" gorm:"column:role_sort;comment:显示顺序;"`

	DataScope string `json:"dataScope" form:"dataScope" gorm:"column:data_scope;comment:数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）;"`

	MenuCheckStrictly int8 `json:"menuCheckStrictly" form:"menuCheckStrictly" gorm:"column:menu_check_strictly;comment:菜单树选择项是否关联显示;"`

	DeptCheckStrictly int8 `json:"deptCheckStrictly" form:"deptCheckStrictly" gorm:"column:dept_check_strictly;comment:部门树选择项是否关联显示;"`

	Status string `json:"status" form:"status" gorm:"column:status;comment:角色状态（0正常 1停用）;"`

	DeletedAt *time.Time `json:"deletedAt" form:"deletedAt" gorm:"column:deleted_at;comment:删除标志;"`

	CreateBy string `json:"createBy" form:"createBy" gorm:"column:create_by;comment:创建者;"`

	CreateTime *time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;"`

	UpdateBy string `json:"updateBy" form:"updateBy" gorm:"column:update_by;comment:更新者;"`

	UpdateTime *time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:更新时间;"`

	Remark string `json:"remark" form:"remark" gorm:"column:remark;comment:备注;"`

	DataScopeSql string `json:"dataScopeSql" form:"dataScopeSql" gorm:"-"`
}

// TableName Role 表名
func (Role) TableName() string {
	return "sys_role"
}
