// Package view 自动生成模板 SysOperLog
// @description <TODO description class purpose>
// @author
// @File: sys_oper_log
// @version 1.0.0
// @create 2023-08-18 13:41:26
package view

// SysOperLogView 结构体

type SysOperLogView struct {
	Id string `json:"id"`

	Title string `json:"title"`

	BusinessType int `json:"businessType"`

	Method string `json:"method"`

	RequestMethod string `json:"requestMethod"`

	OperatorType int `json:"operatorType"`

	OperName string `json:"operName"`

	DeptName string `json:"deptName"`

	OperUrl string `json:"operUrl"`

	OperIp string `json:"operIp"`

	OperLocation string `json:"operLocation"`

	OperParam string `json:"operParam"`

	JsonResult string `json:"jsonResult"`

	Status int `json:"status"`

	ErrorMsg string `json:"errorMsg"`

	OperTime string `json:"operTime"`

	CostTime int `json:"costTime"`
}
