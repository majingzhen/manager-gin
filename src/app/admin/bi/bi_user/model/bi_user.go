// Package model 自动生成模板 BiUser
// @description <TODO description class purpose>
// @author Administrator
// @File: bi_user
// @version 1.0.0
// @create 2023-08-16 14:53:36
package model

import (
	"time"
)

// BiUser 结构体

type BiUser struct {
	Id string `json:"id" form:"id" gorm:"column:id;comment:;size:32;"`

	TestName string `json:"testName" form:"testName" gorm:"column:test_name;comment:;size:128;"`

	Gender int `json:"gender" form:"gender" gorm:"column:gender;comment:;size:8;"`

	Remark string `json:"remark" form:"remark" gorm:"column:remark;comment:;size:256;"`

	Birth time.Time `json:"birth" form:"birth" gorm:"column:birth;comment:;"`

	CreateTime time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:;"`

	UpdateTime time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:;"`
}

// TableName BiUser 表名
func (BiUser) TableName() string {
	return "BI_USER"
}
