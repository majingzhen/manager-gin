// Package model 自动生成模板 SysOperLog
// @description <TODO description class purpose>
// @author
// @File: sys_oper_log
// @version 1.0.0
// @create 2023-08-08 10:06:19
package model

import (
	"time"
)

// SysOperLog 结构体

type SysOperLog struct {
	BusinessType int `json:"businessType" form:"businessType" gorm:"column:business_type;comment:;"`

	CostTime int `json:"costTime" form:"costTime" gorm:"column:cost_time;comment:;"`

	DeptName string `json:"deptName" form:"deptName" gorm:"column:dept_name;comment:;"`

	ErrorMsg string `json:"errorMsg" form:"errorMsg" gorm:"column:error_msg;comment:;"`

	Id int `json:"id" form:"id" gorm:"column:id;comment:;"`

	JsonResult string `json:"jsonResult" form:"jsonResult" gorm:"column:json_result;comment:;"`

	Method string `json:"method" form:"method" gorm:"column:method;comment:;"`

	OperIp string `json:"operIp" form:"operIp" gorm:"column:oper_ip;comment:;"`

	OperLocation string `json:"operLocation" form:"operLocation" gorm:"column:oper_location;comment:;"`

	OperName string `json:"operName" form:"operName" gorm:"column:oper_name;comment:;"`

	OperParam string `json:"operParam" form:"operParam" gorm:"column:oper_param;comment:;"`

	OperTime time.Time `json:"operTime" form:"operTime" gorm:"column:oper_time;comment:;"`

	OperUrl string `json:"operUrl" form:"operUrl" gorm:"column:oper_url;comment:;"`

	OperatorType int `json:"operatorType" form:"operatorType" gorm:"column:operator_type;comment:;"`

	RequestMethod string `json:"requestMethod" form:"requestMethod" gorm:"column:request_method;comment:;"`

	Status int `json:"status" form:"status" gorm:"column:status;comment:;"`

	Title string `json:"title" form:"title" gorm:"column:title;comment:;"`
}

// TableName SysOperLog 表名
func (SysOperLog) TableName() string {
	return "sys_oper_log"
}
