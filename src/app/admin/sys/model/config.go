// Package model 自动生成模板 Config
// @description <TODO description class purpose>
// @author
// @File: config
// @version 1.0.0
// @create 2023-08-21 14:20:37
package model

import (
	"time"
)

// Config 结构体

type Config struct {
	Id string `json:"id" form:"id" gorm:"column:id;comment:参数主键;"`

	ConfigName string `json:"configName" form:"configName" gorm:"column:config_name;comment:参数名称;"`

	ConfigKey string `json:"configKey" form:"configKey" gorm:"column:config_key;comment:参数键名;"`

	ConfigValue string `json:"configValue" form:"configValue" gorm:"column:config_value;comment:参数键值;"`

	ConfigType string `json:"configType" form:"configType" gorm:"column:config_type;comment:系统内置（Y是 N否）;"`

	CreateBy string `json:"createBy" form:"createBy" gorm:"column:create_by;comment:创建者;"`

	CreateTime *time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;"`

	UpdateBy string `json:"updateBy" form:"updateBy" gorm:"column:update_by;comment:更新者;"`

	UpdateTime *time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:更新时间;"`

	Remark string `json:"remark" form:"remark" gorm:"column:remark;comment:备注;"`
}

// TableName Config 表名
func (Config) TableName() string {
	return "sys_config"
}
