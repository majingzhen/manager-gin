// Package view 自动生成模板 SysDictData
// @description <TODO description class purpose>
// @author
// @File: sys_dict_data
// @version 1.0.0
// @create 2023-08-20 19:08:42
package view

// SysDictDataView 结构体

type SysDictDataView struct {
	Id         string `json:"id" form:"id"`
	DictSort   int    `json:"dictSort" form:"dictSort"`
	DictLabel  string `json:"dictLabel" form:"dictLabel"`
	DictValue  string `json:"dictValue" form:"dictValue"`
	DictType   string `json:"dictType" form:"dictType"`
	CssClass   string `json:"cssClass" form:"cssClass"`
	ListClass  string `json:"listClass" form:"listClass"`
	IsDefault  string `json:"isDefault" form:"isDefault"`
	Status     string `json:"status" form:"status"`
	CreateBy   string `json:"createBy" form:"createBy"`
	CreateTime string `json:"createTime" form:"createTime"`
	UpdateBy   string `json:"updateBy" form:"updateBy"`
	UpdateTime string `json:"updateTime" form:"updateTime"`
	Remark     string `json:"remark" form:"remark"`
}

type SysDictDataPageView struct {
	// TODO 按需修改
	Id         string `json:"id" form:"id"`
	DictSort   int    `json:"dictSort" form:"dictSort"`
	DictLabel  string `json:"dictLabel" form:"dictLabel"`
	DictValue  string `json:"dictValue" form:"dictValue"`
	DictType   string `json:"dictType" form:"dictType"`
	CssClass   string `json:"cssClass" form:"cssClass"`
	ListClass  string `json:"listClass" form:"listClass"`
	IsDefault  string `json:"isDefault" form:"isDefault"`
	Status     string `json:"status" form:"status"`
	CreateBy   string `json:"createBy" form:"createBy"`
	CreateTime string `json:"createTime" form:"createTime"`
	UpdateBy   string `json:"updateBy" form:"updateBy"`
	UpdateTime string `json:"updateTime" form:"updateTime"`
	Remark     string `json:"remark" form:"remark"`

	OrderByColumn string `json:"orderByColumn" form:"orderByColumn"` //排序字段
	IsAsc         string `json:"isAsc" form:"isAsc"`                 //排序方式
	PageNum       int    `json:"pageNum" form:"pageNum"`             //当前页码
	PageSize      int    `json:"pageSize" form:"pageSize"`           //每页数
}
