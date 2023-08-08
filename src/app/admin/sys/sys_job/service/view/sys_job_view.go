// Package view 自动生成模板 SysJob
// @description <TODO description class purpose>
// @author
// @File: sys_job
// @version 1.0.0
// @create 2023-08-08 10:06:19
package view

// SysJobView 结构体

type SysJobView struct {
	Concurrent interface{} `json:"concurrent"`

	CreateBy string `json:"createBy"`

	CreateTime string `json:"createTime"`

	CronExpression string `json:"cronExpression"`

	Id int `json:"id"`

	InvokeTarget string `json:"invokeTarget"`

	JobGroup string `json:"jobGroup"`

	JobName string `json:"jobName"`

	MisfirePolicy string `json:"misfirePolicy"`

	Remark string `json:"remark"`

	Status interface{} `json:"status"`

	UpdateBy string `json:"updateBy"`

	UpdateTime string `json:"updateTime"`
}
