// Package view 自动生成模板 TableColumn
// @description <TODO description class purpose>
// @author
// @File: table_column
// @version 1.0.0
// @create 2023-08-31 09:09:53
package view

import "manager-gin/src/common"

// TableColumnView 结构体

type TableColumnView struct {
	ColumnComment string `json:"columnComment" form:"columnComment"`
	ColumnName    string `json:"columnName" form:"columnName"`
	ColumnType    string `json:"columnType" form:"columnType"`
	CreateBy      string `json:"createBy" form:"createBy"`
	CreateTime    string `json:"createTime" form:"createTime"`
	DictType      string `json:"dictType" form:"dictType"`
	GoField       string `json:"goField" form:"goField"`
	GoType        string `json:"goType" form:"goType"`
	HtmlType      string `json:"htmlType" form:"htmlType"`
	Id            string `json:"id" form:"id"`
	IsEdit        string `json:"isEdit" form:"isEdit"`
	IsIncrement   string `json:"isIncrement" form:"isIncrement"`
	IsInsert      string `json:"isInsert" form:"isInsert"`
	IsList        string `json:"isList" form:"isList"`
	IsPk          string `json:"isPk" form:"isPk"`
	IsQuery       string `json:"isQuery" form:"isQuery"`
	IsRequired    string `json:"isRequired" form:"isRequired"`
	QueryType     string `json:"queryType" form:"queryType"`
	Sort          int    `json:"sort" form:"sort"`
	TableId       string `json:"tableId" form:"tableId"`
	UpdateBy      string `json:"updateBy" form:"updateBy"`
	UpdateTime    string `json:"updateTime" form:"updateTime"`
}

type TableColumnPageView struct {
	common.PageView
	// TODO 按需修改
	ColumnComment string `json:"columnComment" form:"columnComment"`
	ColumnName    string `json:"columnName" form:"columnName"`
	ColumnType    string `json:"columnType" form:"columnType"`
	CreateBy      string `json:"createBy" form:"createBy"`
	CreateTime    string `json:"createTime" form:"createTime"`
	DictType      string `json:"dictType" form:"dictType"`
	GoField       string `json:"goField" form:"goField"`
	GoType        string `json:"goType" form:"goType"`
	HtmlType      string `json:"htmlType" form:"htmlType"`
	Id            string `json:"id" form:"id"`
	IsEdit        string `json:"isEdit" form:"isEdit"`
	IsIncrement   string `json:"isIncrement" form:"isIncrement"`
	IsInsert      string `json:"isInsert" form:"isInsert"`
	IsList        string `json:"isList" form:"isList"`
	IsPk          string `json:"isPk" form:"isPk"`
	IsQuery       string `json:"isQuery" form:"isQuery"`
	IsRequired    string `json:"isRequired" form:"isRequired"`
	QueryType     string `json:"queryType" form:"queryType"`
	Sort          int    `json:"sort" form:"sort"`
	TableId       string `json:"tableId" form:"tableId"`
	UpdateBy      string `json:"updateBy" form:"updateBy"`
	UpdateTime    string `json:"updateTime" form:"updateTime"`
}

type TableColumnQueryView struct {
	// TODO 按需修改
	ColumnComment string `json:"columnComment" form:"columnComment"`
	ColumnName    string `json:"columnName" form:"columnName"`
	ColumnType    string `json:"columnType" form:"columnType"`
	CreateBy      string `json:"createBy" form:"createBy"`
	CreateTime    string `json:"createTime" form:"createTime"`
	DictType      string `json:"dictType" form:"dictType"`
	GoField       string `json:"goField" form:"goField"`
	GoType        string `json:"goType" form:"goType"`
	HtmlType      string `json:"htmlType" form:"htmlType"`
	Id            string `json:"id" form:"id"`
	IsEdit        string `json:"isEdit" form:"isEdit"`
	IsIncrement   string `json:"isIncrement" form:"isIncrement"`
	IsInsert      string `json:"isInsert" form:"isInsert"`
	IsList        string `json:"isList" form:"isList"`
	IsPk          string `json:"isPk" form:"isPk"`
	IsQuery       string `json:"isQuery" form:"isQuery"`
	IsRequired    string `json:"isRequired" form:"isRequired"`
	QueryType     string `json:"queryType" form:"queryType"`
	Sort          int    `json:"sort" form:"sort"`
	TableId       string `json:"tableId" form:"tableId"`
	UpdateBy      string `json:"updateBy" form:"updateBy"`
	UpdateTime    string `json:"updateTime" form:"updateTime"`
}
