// Package model 自动生成模板 SysPost
// @description <TODO description class purpose>
// @author
// @File: sys_post
// @version 1.0.0
// @create 2023-08-08 10:06:19
package model

import (
	"time"
)

// SysPost 结构体

type SysPost struct {
	CreateBy string `json:"createBy" form:"createBy" gorm:"column:create_by;comment:;"`

	CreateTime time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:;"`

	Id int `json:"id" form:"id" gorm:"column:id;comment:;"`

	PostCode string `json:"postCode" form:"postCode" gorm:"column:post_code;comment:;"`

	PostName string `json:"postName" form:"postName" gorm:"column:post_name;comment:;"`

	PostSort int `json:"postSort" form:"postSort" gorm:"column:post_sort;comment:;"`

	Remark string `json:"remark" form:"remark" gorm:"column:remark;comment:;"`

	Status interface{} `json:"status" form:"status" gorm:"column:status;comment:;"`

	UpdateBy string `json:"updateBy" form:"updateBy" gorm:"column:update_by;comment:;"`

	UpdateTime time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:;"`
}

// TableName SysPost 表名
func (SysPost) TableName() string {
	return "sys_post"
}
