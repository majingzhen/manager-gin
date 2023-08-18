// Package model 自动生成模板 SysOperLog
// @description <TODO description class purpose>
// @author
// @File: sys_oper_log
// @version 1.0.0
// @create 2023-08-18 13:41:26
package model

import (
	"time"
)

// SysOperLog 结构体

type SysOperLog struct {
	Id string `json:"id" form:"id" gorm:"column:id;comment:日志主键;"`

	Title string `json:"title" form:"title" gorm:"column:title;comment:模块标题;"`

	BusinessType int `json:"businessType" form:"businessType" gorm:"column:business_type;comment:业务类型（0其它 1新增 2修改 3删除）;"`

	Method string `json:"method" form:"method" gorm:"column:method;comment:方法名称;"`

	RequestMethod string `json:"requestMethod" form:"requestMethod" gorm:"column:request_method;comment:请求方式;"`

	OperatorType int `json:"operatorType" form:"operatorType" gorm:"column:operator_type;comment:操作类别（0其它 1后台用户 2手机端用户）;"`

	OperName string `json:"operName" form:"operName" gorm:"column:oper_name;comment:操作人员;"`

	DeptName string `json:"deptName" form:"deptName" gorm:"column:dept_name;comment:部门名称;"`

	OperUrl string `json:"operUrl" form:"operUrl" gorm:"column:oper_url;comment:请求URL;"`

	OperIp string `json:"operIp" form:"operIp" gorm:"column:oper_ip;comment:主机地址;"`

	OperLocation string `json:"operLocation" form:"operLocation" gorm:"column:oper_location;comment:操作地点;"`

	OperParam string `json:"operParam" form:"operParam" gorm:"column:oper_param;comment:请求参数;"`

	JsonResult string `json:"jsonResult" form:"jsonResult" gorm:"column:json_result;comment:返回参数;"`

	Status int `json:"status" form:"status" gorm:"column:status;comment:操作状态（0正常 1异常）;"`

	ErrorMsg string `json:"errorMsg" form:"errorMsg" gorm:"column:error_msg;comment:错误消息;"`

	OperTime time.Time `json:"operTime" form:"operTime" gorm:"column:oper_time;comment:操作时间;"`

	CostTime int `json:"costTime" form:"costTime" gorm:"column:cost_time;comment:消耗时间;"`
}

// TableName SysOperLog 表名
func (SysOperLog) TableName() string {
	return "sys_oper_log"
}
