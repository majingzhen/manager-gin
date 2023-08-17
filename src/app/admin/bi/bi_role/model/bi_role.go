// Package model 自动生成模板 BiRole
// @description <TODO description class purpose>
// @author Majz
// @File: bi_role
// @version 1.0.0
// @create 2023-08-16 14:53:36
package model

import (
	"time"
)

// BiRole 结构体

type BiRole struct {
	Id string `json:"id" form:"id" gorm:"column:ID;comment:;size:32;"`

	TestName string `json:"testName" form:"testName" gorm:"column:TEST_NAME;comment:;size:128;"`

	Gender int `json:"gender" form:"gender" gorm:"column:GENDER;comment:;size:8;"`

	Remark string `json:"remark" form:"remark" gorm:"column:REMARK;comment:;size:256;"`

	Birth time.Time `json:"birth" form:"birth" gorm:"column:BIRTH;comment:;"`

	CreateTime time.Time `json:"createTime" form:"createTime" gorm:"column:CREATE_TIME;comment:;"`

	UpdateTime time.Time `json:"updateTime" form:"updateTime" gorm:"column:UPDATE_TIME;comment:;"`
}

// TableName BiRole 表名
func (BiRole) TableName() string {
	return "BI_ROLE"
}
