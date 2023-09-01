// Package view 自动生成模板 Table
// @description <TODO description class purpose>
// @author
// @File: table
// @version 1.0.0
// @create 2023-08-31 09:09:53
package view

import "manager-gin/src/common"

// TableView 结构体

type TableView struct {
	BusinessName   string `json:"businessName" form:"businessName"`
	CreateBy       string `json:"createBy" form:"createBy"`
	CreateTime     string `json:"createTime" form:"createTime"`
	FunctionAuthor string `json:"functionAuthor" form:"functionAuthor"`
	FunctionName   string `json:"functionName" form:"functionName"`
	GenPath        string `json:"genPath" form:"genPath"`
	GenType        string `json:"genType" form:"genType"`
	Id             string `json:"id" form:"id"`
	ModuleName     string `json:"moduleName" form:"moduleName"`
	Options        string `json:"options" form:"options"`
	PackageName    string `json:"packageName" form:"packageName"`
	Remark         string `json:"remark" form:"remark"`
	StructName     string `json:"structName" form:"structName"`
	SubTableFkName string `json:"subTableFkName" form:"subTableFkName"`
	SubTableName   string `json:"subTableName" form:"subTableName"`
	TableComment   string `json:"tableComment" form:"tableComment"`
	TableName      string `json:"tableName" form:"tableName"`
	TplCategory    string `json:"tplCategory" form:"tplCategory"`
	UpdateBy       string `json:"updateBy" form:"updateBy"`
	UpdateTime     string `json:"updateTime" form:"updateTime"`
}

type TablePageView struct {
	common.PageView
	// TODO 按需修改
	BusinessName   string `json:"businessName" form:"businessName"`
	CreateBy       string `json:"createBy" form:"createBy"`
	CreateTime     string `json:"createTime" form:"createTime"`
	FunctionAuthor string `json:"functionAuthor" form:"functionAuthor"`
	FunctionName   string `json:"functionName" form:"functionName"`
	Path           string `json:"genPath" form:"genPath"`
	Type           string `json:"genType" form:"genType"`
	Id             string `json:"id" form:"id"`
	ModuleName     string `json:"moduleName" form:"moduleName"`
	Options        string `json:"options" form:"options"`
	PackageName    string `json:"packageName" form:"packageName"`
	Remark         string `json:"remark" form:"remark"`
	StructName     string `json:"structName" form:"structName"`
	SubTableFkName string `json:"subTableFkName" form:"subTableFkName"`
	SubTableName   string `json:"subTableName" form:"subTableName"`
	TableComment   string `json:"tableComment" form:"tableComment"`
	TableName      string `json:"tableName" form:"tableName"`
	TplCategory    string `json:"tplCategory" form:"tplCategory"`
	UpdateBy       string `json:"updateBy" form:"updateBy"`
	UpdateTime     string `json:"updateTime" form:"updateTime"`
}

type TableQueryView struct {
	// TODO 按需修改
	BusinessName   string `json:"businessName" form:"businessName"`
	CreateBy       string `json:"createBy" form:"createBy"`
	CreateTime     string `json:"createTime" form:"createTime"`
	FunctionAuthor string `json:"functionAuthor" form:"functionAuthor"`
	FunctionName   string `json:"functionName" form:"functionName"`
	Path           string `json:"genPath" form:"genPath"`
	Type           string `json:"genType" form:"genType"`
	Id             string `json:"id" form:"id"`
	ModuleName     string `json:"moduleName" form:"moduleName"`
	Options        string `json:"options" form:"options"`
	PackageName    string `json:"packageName" form:"packageName"`
	Remark         string `json:"remark" form:"remark"`
	StructName     string `json:"structName" form:"structName"`
	SubTableFkName string `json:"subTableFkName" form:"subTableFkName"`
	SubTableName   string `json:"subTableName" form:"subTableName"`
	TableComment   string `json:"tableComment" form:"tableComment"`
	TableName      string `json:"tableName" form:"tableName"`
	TplCategory    string `json:"tplCategory" form:"tplCategory"`
	UpdateBy       string `json:"updateBy" form:"updateBy"`
	UpdateTime     string `json:"updateTime" form:"updateTime"`
}
