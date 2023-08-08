// Package model 自动生成模板 SysUserPost
// @description <TODO description class purpose>
// @author
// @File: sys_user_post
// @version 1.0.0
// @create 2023-08-08 10:06:19
package model

// SysUserPost 结构体

type SysUserPost struct {
	PostId int `json:"postId" form:"postId" gorm:"column:post_id;comment:;"`

	UserId int `json:"userId" form:"userId" gorm:"column:user_id;comment:;"`
}

// TableName SysUserPost 表名
func (SysUserPost) TableName() string {
	return "sys_user_post"
}
