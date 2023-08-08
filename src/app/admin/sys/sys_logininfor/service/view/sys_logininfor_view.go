// Package view 自动生成模板 SysLogininfor
// @description <TODO description class purpose>
// @author
// @File: sys_logininfor
// @version 1.0.0
// @create 2023-08-08 10:06:19
package view

// SysLogininforView 结构体

type SysLogininforView struct {
	Browser string `json:"browser"`

	Id int `json:"id"`

	Ipaddr string `json:"ipaddr"`

	LoginLocation string `json:"loginLocation"`

	LoginTime string `json:"loginTime"`

	Msg string `json:"msg"`

	Os string `json:"os"`

	Status interface{} `json:"status"`

	UserName string `json:"userName"`
}
