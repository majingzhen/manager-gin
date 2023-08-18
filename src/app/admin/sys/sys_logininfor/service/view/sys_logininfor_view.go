// Package view 自动生成模板 SysLogininfor
// @description <TODO description class purpose>
// @author
// @File: sys_logininfor
// @version 1.0.0
// @create 2023-08-18 13:41:26
package view

// SysLogininforView 结构体

type SysLogininforView struct {
	Id string `json:"id"`

	UserName string `json:"userName"`

	Ipaddr string `json:"ipaddr"`

	LoginLocation string `json:"loginLocation"`

	Browser string `json:"browser"`

	Os string `json:"os"`

	Status string `json:"status"`

	Msg string `json:"msg"`

	LoginTime string `json:"loginTime"`
}
