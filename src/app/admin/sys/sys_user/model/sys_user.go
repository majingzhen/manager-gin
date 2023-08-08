// Package model 自动生成模板 SysUser
// @description <TODO description class purpose>
// @author
// @File: sys_user
// @version 1.0.0
// @create 2023-08-08 10:06:19
package model

import (
	"time"
)

// SysUser 结构体

type SysUser struct {
	Avatar string `json:"avatar" form:"avatar" gorm:"column:avatar;comment:;"`

	CreateBy string `json:"createBy" form:"createBy" gorm:"column:create_by;comment:;"`

	CreateTime time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:;"`

	DelFlag interface{} `json:"delFlag" form:"delFlag" gorm:"column:del_flag;comment:;"`

	DeptId int `json:"deptId" form:"deptId" gorm:"column:dept_id;comment:;"`

	Email string `json:"email" form:"email" gorm:"column:email;comment:;"`

	Id int `json:"id" form:"id" gorm:"column:id;comment:;"`

	LoginDate time.Time `json:"loginDate" form:"loginDate" gorm:"column:login_date;comment:;"`

	LoginIp string `json:"loginIp" form:"loginIp" gorm:"column:login_ip;comment:;"`

	NickName string `json:"nickName" form:"nickName" gorm:"column:nick_name;comment:;"`

	Password string `json:"password" form:"password" gorm:"column:password;comment:;"`

	Phonenumber string `json:"phonenumber" form:"phonenumber" gorm:"column:phonenumber;comment:;"`

	Remark string `json:"remark" form:"remark" gorm:"column:remark;comment:;"`

	Sex interface{} `json:"sex" form:"sex" gorm:"column:sex;comment:;"`

	Status interface{} `json:"status" form:"status" gorm:"column:status;comment:;"`

	UpdateBy string `json:"updateBy" form:"updateBy" gorm:"column:update_by;comment:;"`

	UpdateTime time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:;"`

	UserName string `json:"userName" form:"userName" gorm:"column:user_name;comment:;"`

	UserType string `json:"userType" form:"userType" gorm:"column:user_type;comment:;"`
}

// TableName SysUser 表名
func (SysUser) TableName() string {
	return "sys_user"
}
