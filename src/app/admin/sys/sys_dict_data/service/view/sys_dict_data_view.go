// Package view 自动生成模板 SysDictData
// @description <TODO description class purpose>
// @author
// @File: sys_dict_data
// @version 1.0.0
// @create 2023-08-18 13:41:26
package view

// SysDictDataView 结构体

type SysDictDataView struct {
	Id string `json:"id"`

	DictSort int `json:"dictSort"`

	DictLabel string `json:"dictLabel"`

	DictValue string `json:"dictValue"`

	DictType string `json:"dictType"`

	CssClass string `json:"cssClass"`

	ListClass string `json:"listClass"`

	IsDefault string `json:"isDefault"`

	Status string `json:"status"`

	CreateBy string `json:"createBy"`

	CreateTime string `json:"createTime"`

	UpdateBy string `json:"updateBy"`

	UpdateTime string `json:"updateTime"`

	Remark string `json:"remark"`
}
