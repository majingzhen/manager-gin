// Package model 自动生成模板 SysJob
// @description <TODO description class purpose>
// @author
// @File: sys_job
// @version 1.0.0
// @create 2023-08-08 10:06:19
package model

import (
	"time"
)

// SysJob 结构体

type SysJob struct {
	Concurrent interface{} `json:"concurrent" form:"concurrent" gorm:"column:concurrent;comment:;"`

	CreateBy string `json:"createBy" form:"createBy" gorm:"column:create_by;comment:;"`

	CreateTime time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:;"`

	CronExpression string `json:"cronExpression" form:"cronExpression" gorm:"column:cron_expression;comment:;"`

	Id int `json:"id" form:"id" gorm:"column:id;comment:;"`

	InvokeTarget string `json:"invokeTarget" form:"invokeTarget" gorm:"column:invoke_target;comment:;"`

	JobGroup string `json:"jobGroup" form:"jobGroup" gorm:"column:job_group;comment:;"`

	JobName string `json:"jobName" form:"jobName" gorm:"column:job_name;comment:;"`

	MisfirePolicy string `json:"misfirePolicy" form:"misfirePolicy" gorm:"column:misfire_policy;comment:;"`

	Remark string `json:"remark" form:"remark" gorm:"column:remark;comment:;"`

	Status interface{} `json:"status" form:"status" gorm:"column:status;comment:;"`

	UpdateBy string `json:"updateBy" form:"updateBy" gorm:"column:update_by;comment:;"`

	UpdateTime time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:;"`
}

// TableName SysJob 表名
func (SysJob) TableName() string {
	return "sys_job"
}
