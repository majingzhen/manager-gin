// Package model 自动生成模板 SysNotice
// @description <TODO description class purpose>
// @author
// @File: sys_notice
// @version 1.0.0
// @create 2023-08-18 13:41:26
package model

import (
	"time"
)

// SysNotice 结构体

type SysNotice struct {
	Id string `json:"id" form:"id" gorm:"column:id;comment:公告ID;"`

	NoticeTitle string `json:"noticeTitle" form:"noticeTitle" gorm:"column:notice_title;comment:公告标题;"`

	NoticeType string `json:"noticeType" form:"noticeType" gorm:"column:notice_type;comment:公告类型（1通知 2公告）;"`

	NoticeContent interface{} `json:"noticeContent" form:"noticeContent" gorm:"column:notice_content;comment:公告内容;"`

	Status string `json:"status" form:"status" gorm:"column:status;comment:公告状态（0正常 1关闭）;"`

	CreateBy string `json:"createBy" form:"createBy" gorm:"column:create_by;comment:创建者;"`

	CreateTime time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;"`

	UpdateBy string `json:"updateBy" form:"updateBy" gorm:"column:update_by;comment:更新者;"`

	UpdateTime time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:更新时间;"`

	Remark string `json:"remark" form:"remark" gorm:"column:remark;comment:备注;"`
}

// TableName SysNotice 表名
func (SysNotice) TableName() string {
	return "sys_notice"
}
