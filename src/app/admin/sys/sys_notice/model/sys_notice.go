// Package model 自动生成模板 SysNotice
// @description <TODO description class purpose>
// @author
// @File: sys_notice
// @version 1.0.0
// @create 2023-08-08 10:06:19
package model

import (
	"time"
)

// SysNotice 结构体

type SysNotice struct {
	CreateBy string `json:"createBy" form:"createBy" gorm:"column:create_by;comment:;"`

	CreateTime time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:;"`

	Id int `json:"id" form:"id" gorm:"column:id;comment:;"`

	NoticeContent interface{} `json:"noticeContent" form:"noticeContent" gorm:"column:notice_content;comment:;"`

	NoticeTitle string `json:"noticeTitle" form:"noticeTitle" gorm:"column:notice_title;comment:;"`

	NoticeType interface{} `json:"noticeType" form:"noticeType" gorm:"column:notice_type;comment:;"`

	Remark string `json:"remark" form:"remark" gorm:"column:remark;comment:;"`

	Status interface{} `json:"status" form:"status" gorm:"column:status;comment:;"`

	UpdateBy string `json:"updateBy" form:"updateBy" gorm:"column:update_by;comment:;"`

	UpdateTime time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:;"`
}

// TableName SysNotice 表名
func (SysNotice) TableName() string {
	return "sys_notice"
}
