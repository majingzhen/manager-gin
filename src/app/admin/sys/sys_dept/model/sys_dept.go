// Package model 自动生成模板 SysDept
// @description <TODO description class purpose>
// @author
// @File: sys_dept
// @version 1.0.0
// @create 2023-08-08 10:06:18
package model

import (
	"time"
)

// SysDept 结构体

type SysDept struct {
	Ancestors string `json:"ancestors" form:"ancestors" gorm:"column:ancestors;comment:;"`

	CreateBy string `json:"createBy" form:"createBy" gorm:"column:create_by;comment:;"`

	CreateTime time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:;"`

	DelFlag interface{} `json:"delFlag" form:"delFlag" gorm:"column:del_flag;comment:;"`

	DeptName string `json:"deptName" form:"deptName" gorm:"column:dept_name;comment:;"`

	Email string `json:"email" form:"email" gorm:"column:email;comment:;"`

	Id int `json:"id" form:"id" gorm:"column:id;comment:;"`

	Leader string `json:"leader" form:"leader" gorm:"column:leader;comment:;"`

	OrderNum int `json:"orderNum" form:"orderNum" gorm:"column:order_num;comment:;"`

	ParentId int `json:"parentId" form:"parentId" gorm:"column:parent_id;comment:;"`

	Phone string `json:"phone" form:"phone" gorm:"column:phone;comment:;"`

	Status interface{} `json:"status" form:"status" gorm:"column:status;comment:;"`

	UpdateBy string `json:"updateBy" form:"updateBy" gorm:"column:update_by;comment:;"`

	UpdateTime time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:;"`
}

// TableName SysDept 表名
func (SysDept) TableName() string {
	return "sys_dept"
}
