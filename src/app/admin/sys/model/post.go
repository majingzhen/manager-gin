// Package model 自动生成模板 Post
// @description <TODO description class purpose>
// @author
// @File: post
// @version 1.0.0
// @create 2023-08-21 17:37:55
package model

import (
	"time"
)

// Post 结构体

type Post struct {
	Id string `json:"id" form:"id" gorm:"column:id;comment:岗位ID;"`

	PostCode string `json:"postCode" form:"postCode" gorm:"column:post_code;comment:岗位编码;"`

	PostName string `json:"postName" form:"postName" gorm:"column:post_name;comment:岗位名称;"`

	PostSort int `json:"postSort" form:"postSort" gorm:"column:post_sort;comment:显示顺序;"`

	Status string `json:"status" form:"status" gorm:"column:status;comment:状态（0正常 1停用）;"`

	CreateBy string `json:"createBy" form:"createBy" gorm:"column:create_by;comment:创建者;"`

	CreateTime *time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;"`

	UpdateBy string `json:"updateBy" form:"updateBy" gorm:"column:update_by;comment:更新者;"`

	UpdateTime *time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:更新时间;"`

	Remark string `json:"remark" form:"remark" gorm:"column:remark;comment:备注;"`
}

// TableName Post 表名
func (Post) TableName() string {
	return "sys_post"
}
