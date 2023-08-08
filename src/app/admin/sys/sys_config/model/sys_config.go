// Package model 自动生成模板 SysConfig
// @description <TODO description class purpose>
// @author
// @File: sys_config
// @version 1.0.0
// @create 2023-08-08 10:06:18
package model

import (
	"time"
)

// SysConfig 结构体

type SysConfig struct {
	ConfigKey string `json:"configKey" form:"configKey" gorm:"column:config_key;comment:;"`

	ConfigName string `json:"configName" form:"configName" gorm:"column:config_name;comment:;"`

	ConfigType interface{} `json:"configType" form:"configType" gorm:"column:config_type;comment:;"`

	ConfigValue string `json:"configValue" form:"configValue" gorm:"column:config_value;comment:;"`

	CreateBy string `json:"createBy" form:"createBy" gorm:"column:create_by;comment:;"`

	CreateTime time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:;"`

	Id int `json:"id" form:"id" gorm:"column:id;comment:;"`

	Remark string `json:"remark" form:"remark" gorm:"column:remark;comment:;"`

	UpdateBy string `json:"updateBy" form:"updateBy" gorm:"column:update_by;comment:;"`

	UpdateTime time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:;"`
}

// TableName SysConfig 表名
func (SysConfig) TableName() string {
	return "sys_config"
}
