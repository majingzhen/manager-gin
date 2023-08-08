// Package view 自动生成模板 SysConfig
// @description <TODO description class purpose>
// @author
// @File: sys_config
// @version 1.0.0
// @create 2023-08-08 10:06:19
package view

// SysConfigView 结构体

type SysConfigView struct {
	ConfigKey string `json:"configKey"`

	ConfigName string `json:"configName"`

	ConfigType interface{} `json:"configType"`

	ConfigValue string `json:"configValue"`

	CreateBy string `json:"createBy"`

	CreateTime string `json:"createTime"`

	Id int `json:"id"`

	Remark string `json:"remark"`

	UpdateBy string `json:"updateBy"`

	UpdateTime string `json:"updateTime"`
}
