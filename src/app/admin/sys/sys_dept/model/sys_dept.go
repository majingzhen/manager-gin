// Package model 自动生成模板 SysDept
// @description <TODO description class purpose>
// @author
// @File: sys_dept
// @version 1.0.0
// @create 2023-08-21 10:27:01
package model

import (
	"time"
)

// SysDept 结构体

type SysDept struct {
	Id string `json:"id" form:"id" gorm:"column:id;comment:组织id;"`

	ParentId string `json:"parentId" form:"parentId" gorm:"column:parent_id;comment:父组织id;"`

	Ancestors string `json:"ancestors" form:"ancestors" gorm:"column:ancestors;comment:祖级列表;"`

	DeptName string `json:"deptName" form:"deptName" gorm:"column:dept_name;comment:组织名称;"`

	OrderNum int `json:"orderNum" form:"orderNum" gorm:"column:order_num;comment:显示顺序;"`

	Leader string `json:"leader" form:"leader" gorm:"column:leader;comment:负责人;"`

	Phone string `json:"phone" form:"phone" gorm:"column:phone;comment:联系电话;"`

	Email string `json:"email" form:"email" gorm:"column:email;comment:邮箱;"`

	Status string `json:"status" form:"status" gorm:"column:status;comment:部门状态（0正常 1停用）;"`

	DeletedAt time.Time `json:"deletedAt" form:"deletedAt" gorm:"column:deleted_at;comment:删除标志;"`

	CreateBy string `json:"createBy" form:"createBy" gorm:"column:create_by;comment:创建者;"`

	CreateTime time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;"`

	UpdateBy string `json:"updateBy" form:"updateBy" gorm:"column:update_by;comment:更新者;"`

	UpdateTime time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:更新时间;"`

	DataScopeSql string `json:"dataScopeSql" form:"dataScopeSql" gorm:"-"`
}

// TableName SysDept 表名
func (SysDept) TableName() string {
	return "sys_dept"
}
