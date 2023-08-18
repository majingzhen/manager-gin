// Package view 自动生成模板 SysJobLog
// @description <TODO description class purpose>
// @author
// @File: sys_job_log
// @version 1.0.0
// @create 2023-08-18 13:41:26
package view

// SysJobLogView 结构体

type SysJobLogView struct {
	Id string `json:"id"`

	JobName string `json:"jobName"`

	JobGroup string `json:"jobGroup"`

	InvokeTarget string `json:"invokeTarget"`

	JobMessage string `json:"jobMessage"`

	Status string `json:"status"`

	ExceptionInfo string `json:"exceptionInfo"`

	CreateTime string `json:"createTime"`
}
