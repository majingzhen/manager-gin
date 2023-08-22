// Package model 自动生成模板 SysUserPost
// @description <TODO description class purpose>
// @author
// @File: sys_user_post
// @version 1.0.0
// @create 2023-08-21 17:37:55
package model

// SysUserPost 结构体

type SysUserPost struct {
	UserId string `json:"userId" form:"userId" gorm:"column:user_id;comment:用户ID;"`

	PostId string `json:"postId" form:"postId" gorm:"column:post_id;comment:岗位ID;"`
}

// TableName SysUserPost 表名
func (SysUserPost) TableName() string {
	return "sys_user_post"
}
