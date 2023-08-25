// Package model 自动生成模板 SysDictType
// @description <TODO description class purpose>
// @author
// @File: sys_dict_type
// @version 1.0.0
// @create 2023-08-18 13:41:26
package model

import (
	"time"
)

// SysDictType 结构体

type SysDictType struct {
	Id string `json:"id" form:"id" gorm:"column:id;comment:字典主键;"`

	DictName string `json:"dictName" form:"dictName" gorm:"column:dict_name;comment:字典名称;"`

	DictType string `json:"dictType" form:"dictType" gorm:"column:dict_type;comment:字典类型;"`

	Status string `json:"status" form:"status" gorm:"column:status;comment:状态（0正常 1停用）;"`

	CreateBy string `json:"createBy" form:"createBy" gorm:"column:create_by;comment:创建者;"`

	CreateTime *time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;"`

	UpdateBy string `json:"updateBy" form:"updateBy" gorm:"column:update_by;comment:更新者;"`

	UpdateTime *time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:更新时间;"`

	Remark string `json:"remark" form:"remark" gorm:"column:remark;comment:备注;"`
}

// TableName SysDictType 表名
func (SysDictType) TableName() string {
	return "sys_dict_type"
}
