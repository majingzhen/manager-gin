// Package model 自动生成模板 Table
// @description <TODO description class purpose>
// @author
// @File: table
// @version 1.0.0
// @create 2023-08-31 09:09:53
package model

import (
	"time"
)

// Table 结构体
type Table struct {
	BusinessName string `json:"businessName" form:"businessName" gorm:"column:business_name;comment:生成业务名;"`

	CreateBy string `json:"createBy" form:"createBy" gorm:"column:create_by;comment:创建者;"`

	CreateTime *time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;"`

	FunctionAuthor string `json:"functionAuthor" form:"functionAuthor" gorm:"column:function_author;comment:生成功能作者;"`

	FunctionName string `json:"functionName" form:"functionName" gorm:"column:function_name;comment:生成功能名;"`

	GenPath string `json:"genPath" form:"genPath" gorm:"column:gen_path;comment:生成路径（不填默认项目路径）;"`

	GenType string `json:"genType" form:"genType" gorm:"column:gen_type;comment:生成代码方式（0zip压缩包 1自定义路径）;"`

	Id string `json:"id" form:"id" gorm:"column:id;comment:编号;"`

	ModuleName string `json:"moduleName" form:"moduleName" gorm:"column:module_name;comment:生成模块名;"`

	Options string `json:"options" form:"options" gorm:"column:options;comment:其它生成选项;"`

	PackageName string `json:"packageName" form:"packageName" gorm:"column:package_name;comment:生成包路径;"`

	Remark string `json:"remark" form:"remark" gorm:"column:remark;comment:备注;"`

	StructName string `json:"structName" form:"structName" gorm:"column:struct_name;comment:结构体名称;"`

	SubTableFkName string `json:"subTableFkName" form:"subTableFkName" gorm:"column:sub_table_fk_name;comment:子表关联的外键名;"`

	SubTableName string `json:"subTableName" form:"subTableName" gorm:"column:sub_table_name;comment:关联子表的表名;"`

	TableComment string `json:"tableComment" form:"tableComment" gorm:"column:table_comment;comment:表描述;"`

	Name string `json:"tableName" form:"tableName" gorm:"column:table_name;comment:表名称;"`

	TplCategory string `json:"tplCategory" form:"tplCategory" gorm:"column:tpl_category;comment:使用的模板（crud单表操作 tree树表操作）;"`

	UpdateBy string `json:"updateBy" form:"updateBy" gorm:"column:update_by;comment:更新者;"`

	UpdateTime *time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:更新时间;"`
}

// TableName Table 表名
func (t Table) TableName() string {
	return "gen_table"
}
