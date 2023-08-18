// Package view 自动生成模板 SysJob
// @description <TODO description class purpose>
// @author
// @File: sys_job
// @version 1.0.0
// @create 2023-08-18 13:41:26
package view

// SysJobView 结构体

type SysJobView struct {
	Id string `json:"id"`

	JobName string `json:"jobName"`

	JobGroup string `json:"jobGroup"`

	InvokeTarget string `json:"invokeTarget"`

	CronExpression string `json:"cronExpression"`

	MisfirePolicy string `json:"misfirePolicy"`

	Concurrent string `json:"concurrent"`

	Status string `json:"status"`

	CreateBy string `json:"createBy"`

	CreateTime string `json:"createTime"`

	UpdateBy string `json:"updateBy"`

	UpdateTime string `json:"updateTime"`

	Remark string `json:"remark"`
}
