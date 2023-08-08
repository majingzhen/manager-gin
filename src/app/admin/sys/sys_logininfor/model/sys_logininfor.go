// Package model 自动生成模板 SysLogininfor
// @description <TODO description class purpose>
// @author
// @File: sys_logininfor
// @version 1.0.0
// @create 2023-08-08 10:06:19
package model

import (
	"time"
)

// SysLogininfor 结构体

type SysLogininfor struct {
	Browser string `json:"browser" form:"browser" gorm:"column:browser;comment:;"`

	Id int `json:"id" form:"id" gorm:"column:id;comment:;"`

	Ipaddr string `json:"ipaddr" form:"ipaddr" gorm:"column:ipaddr;comment:;"`

	LoginLocation string `json:"loginLocation" form:"loginLocation" gorm:"column:login_location;comment:;"`

	LoginTime time.Time `json:"loginTime" form:"loginTime" gorm:"column:login_time;comment:;"`

	Msg string `json:"msg" form:"msg" gorm:"column:msg;comment:;"`

	Os string `json:"os" form:"os" gorm:"column:os;comment:;"`

	Status interface{} `json:"status" form:"status" gorm:"column:status;comment:;"`

	UserName string `json:"userName" form:"userName" gorm:"column:user_name;comment:;"`
}

// TableName SysLogininfor 表名
func (SysLogininfor) TableName() string {
	return "sys_logininfor"
}
