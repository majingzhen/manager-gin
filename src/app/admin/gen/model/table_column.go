// Package model 自动生成模板 TableColumn
// @description <TODO description class purpose>
// @author
// @File: table_column
// @version 1.0.0
// @create 2023-08-31 09:09:53
package model

import (
	"time"
)

// TableColumn 结构体
type TableColumn struct {
	ColumnComment string `json:"columnComment" form:"columnComment" gorm:"column:column_comment;comment:列描述;"`

	ColumnName string `json:"columnName" form:"columnName" gorm:"column:column_name;comment:列名称;"`

	ColumnType string `json:"columnType" form:"columnType" gorm:"column:column_type;comment:列类型;"`

	DataType string `json:"dataType" form:"dataType" gorm:"column:data_type;comment:数据类型;"`

	ColumnLength int `json:"columnLength" form:"columnLength" gorm:"column:column_length;comment:列长度;"`

	CreateBy string `json:"createBy" form:"createBy" gorm:"column:create_by;comment:创建者;"`

	JsonField string `json:"jsonField" form:"jsonField" gorm:"column:json_field;comment:json字段名;"`

	CreateTime *time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;"`

	DictType string `json:"dictType" form:"dictType" gorm:"column:dict_type;comment:字典类型;"`

	GoField string `json:"goField" form:"goField" gorm:"column:go_field;comment:字段名;"`

	GoType string `json:"goType" form:"goType" gorm:"column:go_type;comment:类型;"`

	DefaultValue string `json:"defaultValue" form:"defaultValue" gorm:"column:default_value;comment:默认值;"`

	HtmlType string `json:"htmlType" form:"htmlType" gorm:"column:html_type;comment:显示类型（文本框、文本域、下拉框、复选框、单选框、日期控件）;"`

	Id string `json:"id" form:"id" gorm:"column:id;comment:编号;"`

	IsEdit string `json:"isEdit" form:"isEdit" gorm:"column:is_edit;comment:是否编辑字段（1是）;"`

	IsIncrement string `json:"isIncrement" form:"isIncrement" gorm:"column:is_increment;comment:是否自增（1是）;"`

	IsInsert string `json:"isInsert" form:"isInsert" gorm:"column:is_insert;comment:是否为插入字段（1是）;"`

	IsList string `json:"isList" form:"isList" gorm:"column:is_list;comment:是否列表字段（1是）;"`

	IsPk string `json:"isPk" form:"isPk" gorm:"column:is_pk;comment:是否主键（1是）;"`

	IsQuery string `json:"isQuery" form:"isQuery" gorm:"column:is_query;comment:是否查询字段（1是）;"`

	IsRequired string `json:"isRequired" form:"isRequired" gorm:"column:is_required;comment:是否必填（1是）;"`

	QueryType string `json:"queryType" form:"queryType" gorm:"column:query_type;comment:查询方式（等于、不等于、大于、小于、范围）;"`

	Sort int `json:"sort" form:"sort" gorm:"column:sort;comment:排序;"`

	TableId string `json:"tableId" form:"tableId" gorm:"column:table_id;comment:归属表编号;"`

	UpdateBy string `json:"updateBy" form:"updateBy" gorm:"column:update_by;comment:更新者;"`

	UpdateTime *time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:更新时间;"`
}

// TableName TableColumn 表名
func (TableColumn) TableName() string {
	return "gen_table_column"
}
