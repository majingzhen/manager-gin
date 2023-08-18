// Package view 自动生成模板 SysNotice
// @description <TODO description class purpose>
// @author
// @File: sys_notice
// @version 1.0.0
// @create 2023-08-18 13:41:26
package view

// SysNoticeView 结构体

type SysNoticeView struct {
	Id string `json:"id"`

	NoticeTitle string `json:"noticeTitle"`

	NoticeType string `json:"noticeType"`

	NoticeContent interface{} `json:"noticeContent"`

	Status string `json:"status"`

	CreateBy string `json:"createBy"`

	CreateTime string `json:"createTime"`

	UpdateBy string `json:"updateBy"`

	UpdateTime string `json:"updateTime"`

	Remark string `json:"remark"`
}
