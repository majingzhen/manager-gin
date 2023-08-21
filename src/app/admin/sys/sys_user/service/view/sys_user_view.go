// Package view 自动生成模板 SysUser
// @description <TODO description class purpose>
// @author
// @File: sys_user
// @version 1.0.0
// @create 2023-08-21 14:20:37
package view

import "manager-gin/src/app/admin/sys/sys_role/service/view"

// SysUserView 结构体

type SysUserView struct {
	Id          string              `json:"id" form:"id"`
	DeptId      string              `json:"deptId" form:"deptId"`
	UserName    string              `json:"userName" form:"userName"`
	NickName    string              `json:"nickName" form:"nickName"`
	UserType    string              `json:"userType" form:"userType"`
	Email       string              `json:"email" form:"email"`
	PhoneNumber string              `json:"phoneNumber" form:"phoneNumber"`
	Sex         string              `json:"sex" form:"sex"`
	Avatar      string              `json:"avatar" form:"avatar"`
	Password    string              `json:"password" form:"password"`
	Salt        string              `json:"salt" form:"salt"`
	Status      string              `json:"status" form:"status"`
	DeletedAt   string              `json:"deletedAt" form:"deletedAt"`
	LoginIp     string              `json:"loginIp" form:"loginIp"`
	LoginDate   string              `json:"loginDate" form:"loginDate"`
	CreateBy    string              `json:"createBy" form:"createBy"`
	CreateTime  string              `json:"createTime" form:"createTime"`
	UpdateBy    string              `json:"updateBy" form:"updateBy"`
	UpdateTime  string              `json:"updateTime" form:"updateTime"`
	Remark      string              `json:"remark" form:"remark"`
	Roles       *[]view.SysRoleView `json:"roles"`
}

type SysUserPageView struct {
	// TODO 按需修改
	Id          string `json:"id" form:"id"`
	DeptId      string `json:"deptId" form:"deptId"`
	UserName    string `json:"userName" form:"userName"`
	NickName    string `json:"nickName" form:"nickName"`
	UserType    string `json:"userType" form:"userType"`
	Email       string `json:"email" form:"email"`
	PhoneNumber string `json:"phoneNumber" form:"phoneNumber"`
	Sex         string `json:"sex" form:"sex"`
	Avatar      string `json:"avatar" form:"avatar"`
	Password    string `json:"password" form:"password"`
	Salt        string `json:"salt" form:"salt"`
	Status      string `json:"status" form:"status"`
	DeletedAt   string `json:"deletedAt" form:"deletedAt"`
	LoginIp     string `json:"loginIp" form:"loginIp"`
	LoginDate   string `json:"loginDate" form:"loginDate"`
	CreateBy    string `json:"createBy" form:"createBy"`
	CreateTime  string `json:"createTime" form:"createTime"`
	UpdateBy    string `json:"updateBy" form:"updateBy"`
	UpdateTime  string `json:"updateTime" form:"updateTime"`
	Remark      string `json:"remark" form:"remark"`

	OrderByColumn string `json:"orderByColumn" form:"orderByColumn"` //排序字段
	IsAsc         string `json:"isAsc" form:"isAsc"`                 //排序方式
	PageNum       int    `json:"pageNum" form:"pageNum"`             //当前页码
	PageSize      int    `json:"pageSize" form:"pageSize"`           //每页数
}
