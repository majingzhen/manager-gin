// Package view 自动生成模板 SysUser
// @description <TODO description class purpose>
// @author
// @File: sys_user
// @version 1.0.0
// @create 2023-08-18 14:02:24
package view

import "manager-gin/src/app/admin/sys/sys_role/service/view"

// SysUserView 结构体

type SysUserView struct {
	Id          string              `json:"id"`
	DeptId      string              `json:"deptId"`
	UserName    string              `json:"userName"`
	NickName    string              `json:"nickName"`
	UserType    string              `json:"userType"`
	Email       string              `json:"email"`
	PhoneNumber string              `json:"phoneNumber"`
	Sex         string              `json:"sex"`
	Avatar      string              `json:"avatar"`
	Password    string              `json:"password"`
	Salt        string              `json:"salt"`
	Status      string              `json:"status"`
	DeletedAt   string              `json:"deletedAt"`
	LoginIp     string              `json:"loginIp"`
	LoginDate   string              `json:"loginDate"`
	CreateBy    string              `json:"createBy"`
	CreateTime  string              `json:"createTime"`
	UpdateBy    string              `json:"updateBy"`
	UpdateTime  string              `json:"updateTime"`
	Remark      string              `json:"remark"`
	Roles       *[]view.SysRoleView `json:"roles"`
}
