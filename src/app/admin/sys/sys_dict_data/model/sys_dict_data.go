// Package model 自动生成模板 SysDictData
// @description <TODO description class purpose>
// @author
// @File: sys_dict_data
// @version 1.0.0
// @create 2023-08-18 13:41:26
package model

import (
	"time"
)

// SysDictData 结构体

type SysDictData struct {
	Id string `json:"id" form:"id" gorm:"column:id;comment:字典编码;"`

	DictSort int `json:"dictSort" form:"dictSort" gorm:"column:dict_sort;comment:字典排序;"`

	DictLabel string `json:"dictLabel" form:"dictLabel" gorm:"column:dict_label;comment:字典标签;"`

	DictValue string `json:"dictValue" form:"dictValue" gorm:"column:dict_value;comment:字典键值;"`

	DictType string `json:"dictType" form:"dictType" gorm:"column:dict_type;comment:字典类型;"`

	CssClass string `json:"cssClass" form:"cssClass" gorm:"column:css_class;comment:样式属性（其他样式扩展）;"`

	ListClass string `json:"listClass" form:"listClass" gorm:"column:list_class;comment:表格回显样式;"`

	IsDefault string `json:"isDefault" form:"isDefault" gorm:"column:is_default;comment:是否默认（Y是 N否）;"`

	Status string `json:"status" form:"status" gorm:"column:status;comment:状态（0正常 1停用）;"`

	CreateBy string `json:"createBy" form:"createBy" gorm:"column:create_by;comment:创建者;"`

	CreateTime time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;"`

	UpdateBy string `json:"updateBy" form:"updateBy" gorm:"column:update_by;comment:更新者;"`

	UpdateTime time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:更新时间;"`

	Remark string `json:"remark" form:"remark" gorm:"column:remark;comment:备注;"`
}

// TableName SysDictData 表名
func (SysDictData) TableName() string {
	return "sys_dict_data"
}
