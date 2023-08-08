// Package view 自动生成模板 SysJobLog
// @description <TODO description class purpose>
// @author
// @File: sys_job_log
// @version 1.0.0
// @create 2023-08-08 10:06:19
package view

// SysJobLogView 结构体

type SysJobLogView struct {
	CreateTime string `json:"createTime"`

	ExceptionInfo string `json:"exceptionInfo"`

	Id int `json:"id"`

	InvokeTarget string `json:"invokeTarget"`

	JobGroup string `json:"jobGroup"`

	JobMessage string `json:"jobMessage"`

	JobName string `json:"jobName"`

	Status interface{} `json:"status"`
}
