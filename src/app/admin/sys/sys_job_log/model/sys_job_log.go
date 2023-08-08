// Package model 自动生成模板 SysJobLog
// @description <TODO description class purpose>
// @author
// @File: sys_job_log
// @version 1.0.0
// @create 2023-08-08 10:06:19
package model

import (
	"time"
)

// SysJobLog 结构体

type SysJobLog struct {
	CreateTime time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:;"`

	ExceptionInfo string `json:"exceptionInfo" form:"exceptionInfo" gorm:"column:exception_info;comment:;"`

	Id int `json:"id" form:"id" gorm:"column:id;comment:;"`

	InvokeTarget string `json:"invokeTarget" form:"invokeTarget" gorm:"column:invoke_target;comment:;"`

	JobGroup string `json:"jobGroup" form:"jobGroup" gorm:"column:job_group;comment:;"`

	JobMessage string `json:"jobMessage" form:"jobMessage" gorm:"column:job_message;comment:;"`

	JobName string `json:"jobName" form:"jobName" gorm:"column:job_name;comment:;"`

	Status interface{} `json:"status" form:"status" gorm:"column:status;comment:;"`
}

// TableName SysJobLog 表名
func (SysJobLog) TableName() string {
	return "sys_job_log"
}
