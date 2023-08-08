// Package view 自动生成模板 SysUser
// @description <TODO description class purpose>
// @author
// @File: sys_user
// @version 1.0.0
// @create 2023-08-08 10:06:19
package view

// SysUserView 结构体

type SysUserView struct {
	Avatar string `json:"avatar"`

	CreateBy string `json:"createBy"`

	CreateTime string `json:"createTime"`

	DelFlag interface{} `json:"delFlag"`

	DeptId int `json:"deptId"`

	Email string `json:"email"`

	Id int `json:"id"`

	LoginDate string `json:"loginDate"`

	LoginIp string `json:"loginIp"`

	NickName string `json:"nickName"`

	Password string `json:"password"`

	Phonenumber string `json:"phonenumber"`

	Remark string `json:"remark"`

	Sex interface{} `json:"sex"`

	Status interface{} `json:"status"`

	UpdateBy string `json:"updateBy"`

	UpdateTime string `json:"updateTime"`

	UserName string `json:"userName"`

	UserType string `json:"userType"`
}
