// Package model 自动生成模板 SysJobLog
// @description <TODO description class purpose>
// @author
// @File: sys_job_log
// @version 1.0.0
// @create 2023-08-18 13:41:26
package model

import (
	"time"
)

// SysJobLog 结构体

type SysJobLog struct {
	Id string `json:"id" form:"id" gorm:"column:id;comment:任务日志ID;"`

	JobName string `json:"jobName" form:"jobName" gorm:"column:job_name;comment:任务名称;"`

	JobGroup string `json:"jobGroup" form:"jobGroup" gorm:"column:job_group;comment:任务组名;"`

	InvokeTarget string `json:"invokeTarget" form:"invokeTarget" gorm:"column:invoke_target;comment:调用目标字符串;"`

	JobMessage string `json:"jobMessage" form:"jobMessage" gorm:"column:job_message;comment:日志信息;"`

	Status string `json:"status" form:"status" gorm:"column:status;comment:执行状态（0正常 1失败）;"`

	ExceptionInfo string `json:"exceptionInfo" form:"exceptionInfo" gorm:"column:exception_info;comment:异常信息;"`

	CreateTime time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;"`
}

// TableName SysJobLog 表名
func (SysJobLog) TableName() string {
	return "sys_job_log"
}
