// Package model 自动生成模板 User
// @description <TODO description class purpose>
// @author
// @File: user
// @version 1.0.0
// @create 2023-08-18 14:02:24
package model

import (
	"manager-gin/src/global"
	"time"
)

// User 结构体

type User struct {
	global.GvaModel
	DeptId string `json:"orgId" form:"orgId" gorm:"column:dept_id;comment:部门ID;"`

	UserName string `json:"userName" form:"userName" gorm:"column:user_name;comment:用户账号;"`

	NickName string `json:"nickName" form:"nickName" gorm:"column:nick_name;comment:用户昵称;"`

	UserType string `json:"userType" form:"userType" gorm:"column:user_type;comment:用户类型（00系统用户）;"`

	Email string `json:"email" form:"email" gorm:"column:email;comment:用户邮箱;"`

	PhoneNumber string `json:"phoneNumber" form:"phoneNumber" gorm:"column:phone_number;comment:手机号码;"`

	Sex string `json:"sex" form:"sex" gorm:"column:sex;comment:用户性别（0男 1女 2未知）;"`

	Avatar string `json:"avatar" form:"avatar" gorm:"column:avatar;comment:头像地址;"`

	Password string `json:"password" form:"password" gorm:"column:password;comment:密码;"`

	Salt string `json:"salt" form:"salt" gorm:"column:salt;comment:盐值;;"`

	Status string `json:"status" form:"status" gorm:"column:status;comment:帐号状态（0正常 1停用）;"`

	LoginIp string `json:"loginIp" form:"loginIp" gorm:"column:login_ip;comment:最后登录IP;"`

	LoginDate *time.Time `json:"loginDate" form:"loginDate" gorm:"column:login_date;comment:最后登录时间"`

	Remark string `json:"remark" form:"remark" gorm:"column:remark;comment:备注;"`

	DataScopeSql string `json:"dataScopeSql" form:"dataScopeSql" gorm:"-"`
}

// TableName User 表名
func (User) TableName() string {
	return "sys_user"
}
