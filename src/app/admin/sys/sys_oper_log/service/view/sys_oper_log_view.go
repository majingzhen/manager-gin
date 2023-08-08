// Package view 自动生成模板 SysOperLog
// @description <TODO description class purpose>
// @author
// @File: sys_oper_log
// @version 1.0.0
// @create 2023-08-08 10:06:19
package view

// SysOperLogView 结构体

type SysOperLogView struct {
	BusinessType int `json:"businessType"`

	CostTime int `json:"costTime"`

	DeptName string `json:"deptName"`

	ErrorMsg string `json:"errorMsg"`

	Id int `json:"id"`

	JsonResult string `json:"jsonResult"`

	Method string `json:"method"`

	OperIp string `json:"operIp"`

	OperLocation string `json:"operLocation"`

	OperName string `json:"operName"`

	OperParam string `json:"operParam"`

	OperTime string `json:"operTime"`

	OperUrl string `json:"operUrl"`

	OperatorType int `json:"operatorType"`

	RequestMethod string `json:"requestMethod"`

	Status int `json:"status"`

	Title string `json:"title"`
}
