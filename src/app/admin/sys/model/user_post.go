// Package model 自动生成模板 UserPost
// @description <TODO description class purpose>
// @author
// @File: user_post
// @version 1.0.0
// @create 2023-08-21 17:37:55
package model

// UserPost 结构体

type UserPost struct {
	UserId string `json:"userId" form:"userId" gorm:"column:user_id;comment:用户ID;"`

	PostId string `json:"postId" form:"postId" gorm:"column:post_id;comment:岗位ID;"`
}

// TableName UserPost 表名
func (UserPost) TableName() string {
	return "sys_user_post"
}
