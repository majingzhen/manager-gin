// Package model 自动生成模板 SysLogininfor
// @description <TODO description class purpose>
// @author
// @File: sys_logininfor
// @version 1.0.0
// @create 2023-08-18 13:41:26
package model

import (
	"time"
)

// SysLogininfor 结构体

type SysLogininfor struct {
	Id string `json:"id" form:"id" gorm:"column:id;comment:访问ID;"`

	UserName string `json:"userName" form:"userName" gorm:"column:user_name;comment:用户账号;"`

	Ipaddr string `json:"ipaddr" form:"ipaddr" gorm:"column:ipaddr;comment:登录IP地址;"`

	LoginLocation string `json:"loginLocation" form:"loginLocation" gorm:"column:login_location;comment:登录地点;"`

	Browser string `json:"browser" form:"browser" gorm:"column:browser;comment:浏览器类型;"`

	Os string `json:"os" form:"os" gorm:"column:os;comment:操作系统;"`

	Status string `json:"status" form:"status" gorm:"column:status;comment:登录状态（0成功 1失败）;"`

	Msg string `json:"msg" form:"msg" gorm:"column:msg;comment:提示消息;"`

	LoginTime time.Time `json:"loginTime" form:"loginTime" gorm:"column:login_time;comment:访问时间;"`
}

// TableName SysLogininfor 表名
func (SysLogininfor) TableName() string {
	return "sys_logininfor"
}
