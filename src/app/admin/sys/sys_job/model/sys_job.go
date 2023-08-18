// Package model 自动生成模板 SysJob
// @description <TODO description class purpose>
// @author
// @File: sys_job
// @version 1.0.0
// @create 2023-08-18 13:41:26
package model

import (
	"time"
)

// SysJob 结构体

type SysJob struct {
	Id string `json:"id" form:"id" gorm:"column:id;comment:任务ID;"`

	JobName string `json:"jobName" form:"jobName" gorm:"column:job_name;comment:任务名称;"`

	JobGroup string `json:"jobGroup" form:"jobGroup" gorm:"column:job_group;comment:任务组名;"`

	InvokeTarget string `json:"invokeTarget" form:"invokeTarget" gorm:"column:invoke_target;comment:调用目标字符串;"`

	CronExpression string `json:"cronExpression" form:"cronExpression" gorm:"column:cron_expression;comment:cron执行表达式;"`

	MisfirePolicy string `json:"misfirePolicy" form:"misfirePolicy" gorm:"column:misfire_policy;comment:计划执行错误策略（1立即执行 2执行一次 3放弃执行）;"`

	Concurrent string `json:"concurrent" form:"concurrent" gorm:"column:concurrent;comment:是否并发执行（0允许 1禁止）;"`

	Status string `json:"status" form:"status" gorm:"column:status;comment:状态（0正常 1暂停）;"`

	CreateBy string `json:"createBy" form:"createBy" gorm:"column:create_by;comment:创建者;"`

	CreateTime time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;"`

	UpdateBy string `json:"updateBy" form:"updateBy" gorm:"column:update_by;comment:更新者;"`

	UpdateTime time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:更新时间;"`

	Remark string `json:"remark" form:"remark" gorm:"column:remark;comment:备注信息;"`
}

// TableName SysJob 表名
func (SysJob) TableName() string {
	return "sys_job"
}
