// Package view 自动生成模板 SysNotice
// @description <TODO description class purpose>
// @author
// @File: sys_notice
// @version 1.0.0
// @create 2023-08-08 10:06:19
package view

// SysNoticeView 结构体

type SysNoticeView struct {
	CreateBy string `json:"createBy"`

	CreateTime string `json:"createTime"`

	Id int `json:"id"`

	NoticeContent interface{} `json:"noticeContent"`

	NoticeTitle string `json:"noticeTitle"`

	NoticeType interface{} `json:"noticeType"`

	Remark string `json:"remark"`

	Status interface{} `json:"status"`

	UpdateBy string `json:"updateBy"`

	UpdateTime string `json:"updateTime"`
}
