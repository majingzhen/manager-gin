// Package model 自动生成模板 Notice
// @description <TODO description class purpose>
// @author matuto
// @version 1.0.0
// @create 2023-09-11 17:58:01
package model

import (
	"time"
)

// Notice 结构体
type Notice struct {
	Id string `gorm:"column:id;comment:公告ID;"`

	NoticeTitle string `gorm:"column:notice_title;comment:公告标题;"`

	NoticeType string `gorm:"column:notice_type;comment:公告类型（1通知 2公告）;"`

	NoticeContent []byte `gorm:"column:notice_content;comment:公告内容;"`

	Status string `gorm:"column:status;comment:公告状态（0正常 1关闭）;"`

	CreateBy string `gorm:"column:create_by;comment:创建者;"`

	CreateTime *time.Time `gorm:"column:create_time;comment:创建时间;"`

	UpdateBy string `gorm:"column:update_by;comment:更新者;"`

	UpdateTime *time.Time `gorm:"column:update_time;comment:更新时间;"`

	Remark string `gorm:"column:remark;comment:备注;"`
}

// TableName Notice 表名
func (Notice) TableName() string {
	return "sys_notice"
}
