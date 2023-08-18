// Package view 自动生成模板 SysConfig
// @description <TODO description class purpose>
// @author
// @File: sys_config
// @version 1.0.0
// @create 2023-08-18 13:41:26
package view

// SysConfigView 结构体

type SysConfigView struct {
	Id string `json:"id"`

	ConfigName string `json:"configName"`

	ConfigKey string `json:"configKey"`

	ConfigValue string `json:"configValue"`

	ConfigType string `json:"configType"`

	CreateBy string `json:"createBy"`

	CreateTime string `json:"createTime"`

	UpdateBy string `json:"updateBy"`

	UpdateTime string `json:"updateTime"`

	Remark string `json:"remark"`
}
